package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/yeliixhh/linkapi/internal/config"
	"github.com/yeliixhh/linkapi/internal/container"
	"github.com/yeliixhh/linkapi/internal/logger"
)

// 应用启动主入口
func main() {
	c := container.NewApplication()

	logger.Log.Info("Router is ready")

	// 启动容器
	c.Invoke(func(g *gin.Engine, config *config.Config) {
		logger.Log.Info("Router is ready")

		g.Run(fmt.Sprintf("%s:%s", config.ServerConf.Host, config.ServerConf.Port))
	})
}
