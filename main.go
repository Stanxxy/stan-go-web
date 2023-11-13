package main

import (
	"log"

	"github.com/Stanxxy/stan-go-web/config"
	"github.com/Stanxxy/stan-go-web/internal/controller"
	"github.com/Stanxxy/stan-go-web/internal/core"

	// "github.com/Stanxxy/stan-go-web/internal/models"
	echo "github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	config, err := config.NewConfig()
	if err != nil {
		log.Fatalf("%+v\n", err)
	}
	// create server
	server := core.NewServer(config)
	// serve files for dev
	server.ServeStaticFiles()

	healthCtrl := &controller.Healthcheck{}

	// api endpoints
	controller.RegisterAuthRoutes(server)
	controller.RegisterUserRoutes(server)
	controller.RegisterBusinessRoutes(server)

	// pages
	// u := server.Echo.Group("/users")
	// u.GET("", userListCtrl.GetUsers)
	// u.GET("/:id", userCtrl.GetUser)

	// metric / health endpoint according to RFC 5785
	server.Echo.GET("/.well-known/health-check", healthCtrl.GetHealthcheck)
	server.Echo.GET("/.well-known/metrics", echo.WrapHandler(promhttp.Handler()))

	// we will do migrate here

	server.InitDB()

	// Start server
	go func() {
		if err := server.Start(config.Address); err != nil {
			server.Echo.Logger.Info("shutting down the server")
		}
	}()

	server.GracefulShutdown()
}
