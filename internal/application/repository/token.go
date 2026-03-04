package repository

import (
	"context"
	"errors"

	"github.com/yeliixhh/linkapi/internal/types"
	"github.com/yeliixhh/linkapi/internal/types/interfaces"
	"gorm.io/gorm"
)

type SysTokenRepositoryImpl struct {
	db *gorm.DB
}

// 创建SysToken实例
func NewSysTokenRepository(db *gorm.DB) (interfaces.SysTokenRepository, error) {

	repo := &SysTokenRepositoryImpl{
		db: db,
	}

	return repo, nil

}

// 创建Token
func (repo *SysTokenRepositoryImpl) CreateSysToken(ctx context.Context, token *types.SysToken) error {

	return repo.db.WithContext(ctx).Create(token).Error

}

func (repo *SysTokenRepositoryImpl) QuerySysTokenByToken(ctx context.Context, token string) (*types.SysToken, error) {
	sysToken := &types.SysToken{}
	if repo.db.WithContext(ctx).Where("token = ?", token).First(sysToken).Error != nil {
		return nil, errors.New("token不存在")
	}

	return sysToken, nil
}
