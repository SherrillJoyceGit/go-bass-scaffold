package controller

import (
	"fmt"
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
	fmt.Println("controllerName:" + c.ControllerName + "  " + r_con_t.String() + " ")
	// 按照方法名前缀获取action，无法获取默认 get
	for i := 0; i < typ.NumMethod(); i++ {
		//fmt.Println(fmt.Sprintf("method[%d]:%s type is %s", i, typ.Method(i).Name, typ.Method(i).Type))
		methodName := strings.ToLower(typ.Method(i).Name)
		urlPath := "/" + c.ControllerName + "/" + methodName
		if strings.HasPrefix(methodName, "get") {
			fmt.Println("action get,urlPath:" + strings.ToLower(urlPath))
			fmt.Println(typ.Method(i).Func.Type().String())
			method := r_con_v.MethodByName(typ.Method(i).Name)
			h, ok := (makeHandelFunction(fiber.Handler(nil), method)).(fiber.Handler)
			//h, ok := (method.Interface()).(fiber.Handler)
			if !ok {
				fmt.Println(ok)
				//panic("error")
			}
			app.Add(fiber.MethodGet, strings.ToLower(urlPath), h)
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
func makeHandelFunction(f interface{}, vf reflect.Value) interface{} {
	rf := reflect.TypeOf(f)
	if rf.Kind() != reflect.Func {
		panic("expects a function")
	}
	//vtf := reflect.ValueOf(f)
	wrapperF := reflect.MakeFunc(rf, func(in []reflect.Value) []reflect.Value {
		fmt.Println("before")
		out := vf.Call(in)
		fmt.Println("after")
		return out
	})
	return wrapperF.Interface()
}
