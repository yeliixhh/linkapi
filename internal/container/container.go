package container

import (
	"github.com/yeliixhh/linkapi/internal/application/repository"
	"github.com/yeliixhh/linkapi/internal/application/service"
	"github.com/yeliixhh/linkapi/internal/config"
	"github.com/yeliixhh/linkapi/internal/database"
	"github.com/yeliixhh/linkapi/internal/handler"
	"github.com/yeliixhh/linkapi/internal/logger"
	"github.com/yeliixhh/linkapi/internal/router"
	"go.uber.org/dig"
)

// 创建容器
func NewApplication() *dig.Container {
	// 初始化日志
	logger.InitLogger()

	c := dig.New()

	// 加载配置文件
	must(c.Provide(config.NewConfig))

	// 初始化数据库
	must(c.Provide(database.NewDB))

	// repository
	must(c.Provide(repository.NewUserRepository))
	must(c.Provide(repository.NewSysTokenRepository))

	// service
	must(c.Provide(service.NewAuthService))

	// handler
	must(c.Provide(handler.NewAuthHandler))

	// 将服务实例导入容器中
	must(c.Provide(router.NewRouter))

	return c
}

// 如果有报错抛出异常
func must(err error) {
	if err != nil {
		panic(err)
	}
}
