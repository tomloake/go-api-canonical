package main

import (
	// "net/http"

	// "github.com/labstack/echo/v4"
	"go-api-canonical/src/internal/router"

	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Go API Canonical
// @version 1.0
// @description API Example Service
// @title Go API Canonical

// @host 127.0.0.1:8585
// @BasePath /api

// @schemes http https
// @produce	application/json
// @consumes application/json

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	r := router.New()

	r.GET("/swagger/*", echoSwagger.WrapHandler)
	r.Logger.Fatal(r.Start("127.0.0.1:8585"))

	// r.GET("/swagger/*", echoSwagger.WrapHandler)
}