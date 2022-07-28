package controller

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"reflect"
	"strings"
)

type Controller struct {
	controllerName string
}
type ControllerInterface interface {
	initRoute(*fiber.App, ControllerInterface) *fiber.App
}

func (c *Controller) initRoute(app *fiber.App, ci ControllerInterface) *fiber.App {
	r_con_v := reflect.ValueOf(ci)
	r_con_t := reflect.TypeOf(ci)
	typ := r_con_v.Type()
	//若没有制定 controllerName，使用默认 {path}controller
	if c.controllerName == "" {
		conName := strings.ToLower(r_con_t.Elem().Name())
		if strings.HasSuffix(conName, "controller") {
			c.controllerName = conName[:len(conName)-len("controller")]
		} else {
			c.controllerName = conName
		}
	}
	fmt.Println("controllerName:" + c.controllerName + "  " + r_con_t.String() + " ")
	// 按照方法名前缀获取action，无法获取默认 get
	for i := 0; i < r_con_v.NumMethod(); i++ {
		//fmt.Println(fmt.Sprintf("method[%d]:%s type is %s", i, typ.Method(i).Name, typ.Method(i).Type))
		methodName := strings.ToLower(typ.Method(i).Name)
		urlPath := "/" + c.controllerName + "/" + methodName
		if strings.HasPrefix(methodName, "get") {
			fmt.Println("action get,urlPath:" + strings.ToLower(urlPath))
			//app.Add(fiber.MethodGet, strings.ToLower(urlPath), reflect.New(typ.Elem()))
		} else if strings.HasPrefix(methodName, "post") {
			fmt.Println("action post,urlPath:" + strings.ToLower(urlPath))
			//app.Add(fiber.MethodGet, strings.ToLower(urlPath), reflect.New(typ.Elem()))
		} else {
			fmt.Println("action default get,urlPath:" + strings.ToLower(urlPath))
			//app.Add(fiber.MethodGet, strings.ToLower(urlPath), reflect.New(typ.Elem()))
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
