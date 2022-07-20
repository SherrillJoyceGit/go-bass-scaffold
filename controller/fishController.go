package controller

import "github.com/gofiber/fiber/v2"

type FishController struct {
}

func (b *FishController) Ping() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		return ctx.SendString("Get a fish !~")
	}
}
