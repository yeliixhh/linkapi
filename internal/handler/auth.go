package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/yeliixhh/linkapi/internal/logger"
	"github.com/yeliixhh/linkapi/internal/types"
	"github.com/yeliixhh/linkapi/internal/types/interfaces"
	"github.com/yeliixhh/linkapi/internal/utils"
)

type AuthHandler struct {
	authService interfaces.UserService
}

func NewAuthHandler(authService interfaces.UserService) (*AuthHandler, error) {

	return &AuthHandler{authService}, nil
}

// 注册
func (h *AuthHandler) Register(c *gin.Context) {
	req := &types.AuthRegisterRequest{}

	if c.ShouldBindJSON(req) != nil {
		utils.HttpResponse(c, types.FailResult("参数异常"))
	}

	err := req.ValidParam()

	if err != nil {
		logger.Log.Info(fmt.Sprintf("注册参数异常: %s", err.Error()))
		utils.HttpResponse(c, types.FailResultError(err))
		return
	}

	err = h.authService.RegisterUser(c, req.Username, req.Password, req.NickName)

	if err != nil {
		utils.HttpResponse(c, types.FailResultError(err))
		return
	}

	utils.HttpResponse(c, types.SuccessResult("ok"))
}

// 登录
func (h *AuthHandler) Login(c *gin.Context) {
	req := &types.AuthLoginRequest{}

	if c.ShouldBindJSON(req) != nil {
		utils.HttpResponse(c, types.FailResult("参数异常"))
	}

	if err := req.ValidParam(); err != nil {
		utils.HttpResponse(c, types.FailResultError(err))
	}

	login, err := h.authService.Login(c, req)

	if err != nil {
		utils.HttpResponse(c, types.FailResultError(err))
		return
	}

	utils.HttpResponse(c, types.SuccessResult(login))
}
