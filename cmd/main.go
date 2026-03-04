package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yeliixhh/linkapi/internal/container"
	"github.com/yeliixhh/linkapi/internal/logger"
)

// 应用启动主入口
func main() {
	c := container.NewApplication()

	logger.Log.Info("Router is ready")

	// 启动容器
	c.Invoke(func(g *gin.Engine) {
		logger.Log.Info("Router is ready")

		g.Run()
	})
}
