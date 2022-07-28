package controller

import "github.com/gofiber/fiber/v2"

type FishController struct {
	Controller
}

func (b *FishController) Ping() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		return ctx.SendString("ping a fish !~")
	}
}
func (b *FishController) GetPing() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		return ctx.SendString("GetPing a fish !~")
	}
}
func (b *FishController) PostPing() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		return ctx.SendString("PostPing a fish !~")
	}
}
