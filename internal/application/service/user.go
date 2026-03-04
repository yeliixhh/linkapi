package service

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/yeliixhh/linkapi/internal/logger"
	"github.com/yeliixhh/linkapi/internal/types"
	"github.com/yeliixhh/linkapi/internal/types/interfaces"
	"github.com/yeliixhh/linkapi/internal/utils"
)

var (
	REGISTER_ERROR      = errors.New("注册失败")
	USERNAME_EXISTS     = errors.New("用户名已存在")
	USERNAME_NOT_EXISTS = errors.New("用户名已不存在")
	PASSWORD_ERROR      = errors.New("用户名或密码错误")
)

var (
	secretOnce sync.Once
	secretKey  string
)

type UserServiceImpl struct {
	userRepo  interfaces.UserRepository
	tokenRepo interfaces.SysTokenRepository
}

func NewAuthService(authRepository interfaces.UserRepository, tokenRepo interfaces.SysTokenRepository) (interfaces.UserService, error) {

	authServiceImpl := &UserServiceImpl{
		userRepo:  authRepository,
		tokenRepo: tokenRepo,
	}

	return authServiceImpl, nil
}

// 注册用户
func (s *UserServiceImpl) RegisterUser(ctx context.Context, username, password, nickname string) error {
	user := s.userRepo.QueryUserByUserName(ctx, username)
	if user != nil {
		return USERNAME_EXISTS
	}

	hash, err := utils.GeneratePasswordHash(password)
	if err != nil {
		logger.Log.Error(fmt.Sprintf("生成密码hash异常: %s", password))
		return REGISTER_ERROR
	}

	user = &types.SysUser{
		Username: username,
		Password: hash,
		Nickname: nickname,
		IsAdmin:  0,
	}

	err = s.userRepo.CreateUser(ctx, user)

	if err != nil {
		logger.Log.Error(fmt.Sprintf("创建用户异常: %s", err.Error()))
		return REGISTER_ERROR
	}

	return nil
}

// 登录
func (s *UserServiceImpl) Login(ctx context.Context, req *types.AuthLoginRequest) (*types.AuthLoginResponse, error) {

	user := s.userRepo.QueryUserByUserName(ctx, req.Username)

	if user == nil {
		return nil, USERNAME_NOT_EXISTS
	}

	if !utils.ComparePassword(req.Password, user.Password) {
		return nil, PASSWORD_ERROR
	}

	accessToken, refreshToken, err := s.GenerateToken(ctx, user)
	if err != nil {
		return nil, errors.New("token生成失败")
	}

	response := &types.AuthLoginResponse{
		Id:           user.Id,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return response, nil
}

// 生成token
func (s *UserServiceImpl) GenerateToken(ctx context.Context, user *types.SysUser) (accessToken string, refreshToken string, err error) {

	if user == nil {
		return "", "", USERNAME_NOT_EXISTS
	}

	now := time.Now()
	accessTokenExpiresAt := now.Add(24 * time.Hour)
	refreshTokenExpiresAt := now.Add(24 * 7 * time.Hour)
	secretBytes := []byte(getJwtSecretKey())
	accessTokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.Id,
		"exp":     accessTokenExpiresAt.Unix(),
		"iat":     now.Unix(),
		"type":    "access",
	})
	accessToken, err1 := accessTokenClaims.SignedString(secretBytes)
	if err1 != nil {
		logger.Log.Info(fmt.Sprintf("token生成失败: %s", err1.Error()))
		return "", "", errors.New("token生成失败")
	}
	refreshTokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.Id,
		"exp":     refreshTokenExpiresAt.Unix(),
		"iat":     now.Unix(),
		"type":    "refresh",
	})
	refreshToken, err1 = refreshTokenClaims.SignedString(secretBytes)
	if err1 != nil {
		logger.Log.Info(fmt.Sprintf("token生成失败: %s", err1.Error()))
		return "", "", errors.New("token生成失败")
	}

	access := &types.SysToken{
		UserId:    user.Id,
		TokenType: "access",
		IsRevoked: false,
		Token:     accessToken,
		ExpiresAt: accessTokenExpiresAt,
	}

	refresh := &types.SysToken{
		UserId:    user.Id,
		TokenType: "refresh",
		IsRevoked: false,
		Token:     refreshToken,
		ExpiresAt: refreshTokenExpiresAt,
	}

	if s.tokenRepo.CreateSysToken(ctx, access) != nil {

		return "", "", errors.New("token生成失败")
	}
	if s.tokenRepo.CreateSysToken(ctx, refresh) != nil {
		return "", "", errors.New("token生成失败")
	}

	return accessToken, refreshToken, nil
}

// 获取jwt密钥
func getJwtSecretKey() string {

	secretOnce.Do(func() {
		secret := os.Getenv("JWT_SECRET")

		if secret == "" {
			randomBytes := make([]byte, 32) // 32 bytes = 256 bits
			_, err := rand.Read(randomBytes)
			if err != nil {
				panic("JWT密钥生成失败")
			}
			// 转为十六进制字符串（便于存储或打印）
			secret = base64.StdEncoding.EncodeToString(randomBytes)
		}
		secretKey = secret
	})

	return secretKey
}

// 校验token是否合法
func (s *UserServiceImpl) ValidateToken(ctx context.Context, accessToken string) (*types.SysUser, error) {
	parse, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("token异常")
		}
		return []byte(getJwtSecretKey()), nil
	})

	if err != nil {
		return nil, err
	}

	if parse.Valid != true {
		return nil, errors.New("token异常")
	}

	token, err := s.tokenRepo.QuerySysTokenByToken(ctx, accessToken)

	if err != nil || token.IsRevoked || token.TokenType != "access" {
		return nil, errors.New("token异常")
	}

	user, err := s.userRepo.QueryUserById(ctx, token.UserId)

	if err != nil {
		return nil, errors.New("token异常")
	}

	return user, nil
}
