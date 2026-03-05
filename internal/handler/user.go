package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/yeliixhh/linkapi/internal/logger"
	"github.com/yeliixhh/linkapi/internal/types"
	"github.com/yeliixhh/linkapi/internal/types/interfaces"
	"github.com/yeliixhh/linkapi/internal/utils"
)

type UserHandler struct {
	UserServer interfaces.UserService
}

// 创建user handler
func NewUserHandler(userServer interfaces.UserService) (*UserHandler, error) {
	handler := &UserHandler{
		UserServer: userServer,
	}

	return handler, nil
}

// 查询用户信息
func (h *UserHandler) QueryUserInfo(ctx *gin.Context) {

	info, err := h.UserServer.QueryUserInfo(ctx)

	if err != nil {
		logger.Info("获取用户信息失败: %s", err.Error())

		utils.HttpResponse(ctx, types.FailResult("获取用户信息失败"))

		return
	}

	utils.HttpResponse(ctx, info)
}
