package controller

import (
	"github.com/gofiber/fiber/v2"
	"reflect"
	"strings"
)

type Controller struct {
	ControllerName string
}

type ControllerInterface interface {
	initRoute(*fiber.App, ControllerInterface) *fiber.App
}

func (c *Controller) initRoute(app *fiber.App, ci ControllerInterface) *fiber.App {
	r_con_v := reflect.ValueOf(ci)
	r_con_t := reflect.TypeOf(ci)
	typ := r_con_v.Type()
	//若没有制定 controllerName，使用默认 {path}controller
	if c.ControllerName == "" {
		conName := strings.ToLower(r_con_t.Elem().Name())
		if strings.HasSuffix(conName, "controller") {
			c.ControllerName = conName[:len(conName)-len("controller")]
		} else {
			c.ControllerName = conName
		}
	}
	// 按照方法名前缀获取action，无法获取默认 get
	for i := 0; i < typ.NumMethod(); i++ {
		method := r_con_v.MethodByName(typ.Method(i).Name)
		methodName := strings.ToLower(typ.Method(i).Name)
		urlPath := "/" + c.ControllerName + "/" + methodName
		if strings.HasPrefix(methodName, "get") {
			app.Add(fiber.MethodGet, strings.ToLower(urlPath), (method.Call(nil)[0].Interface()).(fiber.Handler))
		} else if strings.HasPrefix(methodName, "post") {
			app.Add(fiber.MethodPost, strings.ToLower(urlPath), (method.Call(nil)[0].Interface()).(fiber.Handler))
		} else {
			app.Add(fiber.MethodGet, strings.ToLower(urlPath), (method.Call(nil)[0].Interface()).(fiber.Handler))
		}

	}
	return app
}
func InitController(app *fiber.App, ci ...ControllerInterface) *fiber.App {
	for _, c := range ci {
		app = c.initRoute(app, c)
	}
	return app
}
