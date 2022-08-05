package controller

import (
	"github.com/SherrillJoyceGit/go-bass-scaffold/controller/dao"
	"github.com/SherrillJoyceGit/go-bass-scaffold/model"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/dig"
)

type FishController struct {
	dig.In
	fishDao *dao.FishDao
	app     *fiber.App
}

func (b *FishController) Ping() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		return ctx.SendString("Get a fish !~")
	}
}

func NewFishController(fishDao *dao.FishDao, app *fiber.App) {
	ctrl := &FishController{
		fishDao: fishDao,
		//app: app,
	}
	app.Post("/helloFish", ctrl.HelloPostFish())
}

func (ctrl *FishController) HelloPostFish() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		p := new(model.CastBaitParam)
		if err := ctx.BodyParser(p); err != nil {
			return ctx.JSON(fiber.Map{
				"code":    "10001",
				"message": "Wrong fish!",
			})
		}
		if fish, err := ctrl.fishDao.HelloPostFish(p); err == nil {
			return ctx.JSON(fiber.Map{
				"code":    "10000",
				"message": "(" + fish.FishName + ")" + fish.Feeling,
			})
		} else {
			return ctx.JSON(fiber.Map{
				"code":    "10001",
				"message": "Wrong fish!",
			})
		}

	}
}
