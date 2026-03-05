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

	// 启动容器
	c.Invoke(func(g *gin.Engine, config *config.Config) {
		addr := fmt.Sprintf("%s:%s", config.ServerConf.Host, config.ServerConf.Port)

		logger.Info("服务启动: %s", addr)

		g.Run(addr)

	})
}
