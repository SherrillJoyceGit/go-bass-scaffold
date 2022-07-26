package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func InitRestApi() *fiber.App {
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})

	// 跨域配置
	app.Use(cors.New())

	// 绑定日志中间件
	//app.Use(middle.LoggerToLogstash())

	app.Use(recover.New())

	return app
}
