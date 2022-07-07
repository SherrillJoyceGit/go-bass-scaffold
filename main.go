package main

import (
	"fmt"
	"github.com/SherrillJoyceGit/go-bass-scaffold/api"
	"github.com/SherrillJoyceGit/go-bass-scaffold/bootstrap"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	app := api.InitRestApi()
	// Listen from a different goroutine
	go func() {
		if err := app.Listen(":8080"); err != nil {
			log.Panic(err)
		}
	}()
	bootstrap.Init()
	c := make(chan os.Signal, 1)                    // Create channel to signify a signal being sent
	signal.Notify(c, os.Interrupt, syscall.SIGTERM) // When an interrupt or termination signal is sent, notify the channel

	_ = <-c // This blocks the main thread until an interrupt is received
	fmt.Println("Gracefully shutting down...")
	_ = app.Shutdown()

	fmt.Println("Running cleanup tasks...")

	// Your cleanup tasks go here
	// db.Close()
	// redisConn.Close()
	fmt.Println("App was successful shutdown.")
}
