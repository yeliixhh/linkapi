package interfaces

import (
	"context"

	"github.com/yeliixhh/linkapi/internal/types"
)

// 用户业务层
type UserService interface {
	// 注册用户
	RegisterUser(ctx context.Context, username, password, nickname string) error

	// 登录
	Login(ctx context.Context, req *types.AuthLoginRequest) (*types.AuthLoginResponse, error)

	// 生成token
	GenerateToken(ctx context.Context, user *types.SysUser) (accessToken string, refreshToken string, err error)

	// 校验token
	ValidateToken(ctx context.Context, accessToken string) (*types.SysUser, error)

	// 查询用户信息
	QueryUserInfo(ctx context.Context) (*types.UserInfo, error)
}

// 用户存储层
type UserRepository interface {
	// 查询用户通过用户名称
	QueryUserByUserName(ctx context.Context, username string) *types.SysUser

	// 通过用户id查询
	QueryUserById(ctx context.Context, Id string) (*types.SysUser, error)

	// 创建用户
	CreateUser(ctx context.Context, user *types.SysUser) error
}
