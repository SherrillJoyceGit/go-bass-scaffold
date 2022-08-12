package main

import (
	"github.com/SherrillJoyceGit/go-bass-scaffold/api"
	"github.com/SherrillJoyceGit/go-bass-scaffold/bootstrap"
)

func main() {
	app := api.InitRestApi()
	bootstrap.BeGraceful(app)
}
