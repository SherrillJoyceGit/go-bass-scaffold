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
	app := api.InitRestApi()
	bootstrap.BeGraceful(app)
}
