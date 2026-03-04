package types

import "errors"

// 注册DTO
type AuthRegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	NickName string `json:"nickName"`
}

func (a *AuthRegisterRequest) ValidParam() error {
	if a.Username == "" {
		return errors.New("用户名不能为空")
	}

	if a.NickName == "" {
		return errors.New("用户昵称不能为空")
	}

	if a.Password == "" {
		return errors.New("密码不能为空")
	}

	return nil
}

// 登录请求体
type AuthLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (a *AuthLoginRequest) ValidParam() error {
	if a.Username == "" {
		return errors.New("登录名称不能为空")
	}
	if a.Password == "" {
		return errors.New("密码不能为空")
	}

	return nil
}

type AuthLoginResponse struct {
	Id           string `json:"id"`
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}
