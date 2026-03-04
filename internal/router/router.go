package router

import (
	"github.com/gin-gonic/gin"
	"github.com/yeliixhh/linkapi/internal/handler"
	"github.com/yeliixhh/linkapi/internal/middleware"
	"github.com/yeliixhh/linkapi/internal/types/interfaces"
	"go.uber.org/dig"
)

type RouteParams struct {
	dig.In

	AuthHandler *handler.AuthHandler

	UserServer interfaces.UserService
}

// 创建服务实例
func NewRouter(params RouteParams) (*gin.Engine, error) {

	r := gin.New()

	// 检测健康状态
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
			"code":   200,
		})
	})

	// 不需要权限控制
	v1 := r.Group("/api/v1")

	// 需要权限控制
	v1Auth := r.Group("/api/v1")

	// 使用授权中间件
	v1Auth.Use(middleware.Auth(params.UserServer))

	RegisterAuthRouter(v1, params.AuthHandler)

	return r, nil
}

func RegisterAuthRouter(r *gin.RouterGroup, h *handler.AuthHandler) {
	auth := r.Group("/auth")
	{
		auth.POST("/register", h.Register)
		auth.POST("/login", h.Login)
	}
}
