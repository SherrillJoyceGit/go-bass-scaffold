package main

import (
	"github.com/SherrillJoyceGit/go-bass-scaffold/api"
	"github.com/SherrillJoyceGit/go-bass-scaffold/bootstrap"
	"github.com/SherrillJoyceGit/go-bass-scaffold/config"
	"github.com/SherrillJoyceGit/go-bass-scaffold/controller"
	"github.com/SherrillJoyceGit/go-bass-scaffold/controller/dao"
	"github.com/SherrillJoyceGit/go-bass-scaffold/db"
	"go.uber.org/dig"
	"log"
)

func main() {

	// construct dig container
	container, err := BuildContainer()

	// construct database connection
	err = container.Invoke(db.NewDbAccess)

	// construct controller
	err = container.Invoke(controller.NewFishController)

	// start HTTP server gracefully
	err = container.Invoke(bootstrap.BeGraceful)
	if err != nil {
		log.Panicln(err)
	}
}

// BuildContainer 创建上下文容器
func BuildContainer() (*dig.Container, error) {
	container := dig.New()
	var err error

	// 提供-http服务构建
	err = container.Provide(api.InitRestApi)

	// 提供-Graceful Shutdown
	err = container.Provide(bootstrap.BeGraceful)

	// 提供-内部依赖组件
	err = container.Provide(dao.NewFishDao)

	// 提供-配置
	err = container.Provide(config.NewViper)
	err = container.Provide(db.NewFileConfig)

	return container, err
}
