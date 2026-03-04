package repository

import (
	"context"
	"errors"

	"github.com/yeliixhh/linkapi/internal/types"
	"github.com/yeliixhh/linkapi/internal/types/interfaces"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) (interfaces.UserRepository, error) {
	authRepositoryImpl := &UserRepositoryImpl{db}
	return authRepositoryImpl, nil
}

// 通过用户名搜索用户
func (a UserRepositoryImpl) QueryUserByUserName(ctx context.Context, username string) *types.SysUser {
	user := &types.SysUser{}

	err := a.db.WithContext(ctx).Where("username = ?", username).First(user).Error

	if err != nil {
		return nil
	}

	return user
}

// 创建用户
func (a UserRepositoryImpl) CreateUser(ctx context.Context, user *types.SysUser) error {
	return a.db.WithContext(ctx).Create(user).Error
}

func (a UserRepositoryImpl) QueryUserById(ctx context.Context, Id string) (*types.SysUser, error) {
	user := &types.SysUser{}

	err := a.db.WithContext(ctx).Where("id = ?", Id).First(user).Error

	if err != nil {
		return nil, errors.New("用户不存在")
	}

	return user, nil
}
