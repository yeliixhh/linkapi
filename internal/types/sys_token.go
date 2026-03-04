package types

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// 用户token记录表
type SysToken struct {
	Id     string `json:"id"`
	UserId string `json:"userId"`
	// token类型 access为验证token refresh为刷新token
	TokenType string    `json:"tokenType"`
	Token     string    `json:"token"`
	IsRevoked bool      `json:"isRevoked"`
	ExpiresAt time.Time `json:"expiresAt"`
	// 创建者
	CreatedAt time.Time `json:"createAt"`
	// 修改时间
	UpdatedAt time.Time `json:"updateAt"`
}

func (u *SysToken) BeforeCreate(tx *gorm.DB) (err error) {

	u.Id = uuid.New().String()

	return
}
