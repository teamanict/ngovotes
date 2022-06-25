package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/teamanict/ngovotes/config"
	"github.com/teamanict/ngovotes/routes"
	"github.com/teamanict/ngovotes/server"
)

func init() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("\n\u001b[96msee you again👋\u001b[0m")
		os.Exit(1)
	}()
}

func main() {
	config, err := config.ParseConfig("config")
	if err != nil {
		panic(err)
	}

	// Setup webserver
	ws, err := server.SetupWebServer(config)
	if err != nil {
		panic(err)
	}

	// Register routes
	routes.RegisterAPIRoutes(ws.App)

	// Run webserver
	ws.ListenWebServer()
}
