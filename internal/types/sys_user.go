package types

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SysUser struct {
	// 用户id
	Id string `json:"id" gorm:"primaryKey"`
	// 登录名
	Username string `json:"username"`
	// 密码
	Password string `json:"password"`
	// 用户昵称
	Nickname string `json:"nickname"`
	// 是否为管理员
	IsAdmin int `json:"isAdmin"`
	// 创建者
	CreatedAt time.Time `json:"createAt"`
	// 修改时间
	UpdatedAt time.Time `json:"updateAt"`
}

func (u *SysUser) BeforeCreate(tx *gorm.DB) (err error) {

	u.Id = uuid.New().String()

	return
}

type UserInfo struct {
	// 用户id
	Id string `json:"id"`

	// 登录名
	Username string `json:"username"`

	// 用户昵称
	Nickname string `json:"nickname"`

	// 是否为管理员
	IsAdmin int `json:"isAdmin"`

	// 创建者
	CreatedAt time.Time `json:"createAt"`

	// 修改时间
	UpdatedAt time.Time `json:"updateAt"`
}

func (u *UserInfo) ValueOf(user *SysUser) (err error) {
	if user == nil {
		return errors.New("用户不存在")
	}

	u.Id = user.Id
	u.Username = user.Username
	u.Nickname = user.Nickname
	u.IsAdmin = user.IsAdmin
	u.CreatedAt = user.CreatedAt
	u.UpdatedAt = user.UpdatedAt

	return nil
}
