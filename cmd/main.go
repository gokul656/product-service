package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gokul656/product-service/app/bootstraps"
	env "github.com/gokul656/product-service/app/configs"
	"github.com/gokul656/product-service/app/controllers"
	"github.com/gokul656/product-service/app/middlewares"
)

func main() {
	log.Println("Starting Product Service")

	cfg := env.GetEnv()

	r := gin.Default()
	public := r.Group("dev")
	public.GET("/ping", pingHandler)

	protected := r.Group("prod")
	protected.Use(middlewares.ValidateRequest())
	protected.GET("/user/:user", controllers.AuthHandler)

	bootstraps.Setup()

	r.Run(fmt.Sprintf(":%s", cfg.Port))
}

func pingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
