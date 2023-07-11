package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gokul656/product-service/app/bootstraps"
	"github.com/gokul656/product-service/app/configs"
	"github.com/gokul656/product-service/app/handlers"
)

func main() {
	log.Println("Starting Product Service")

	cfg := configs.GetEnv()

	r := gin.Default()
	r.Use(gin.CustomRecovery(configs.ErrorHandler))

	public := r.Group("dev")
	public.GET("/ping", pingHandler)
	public.POST("/user", handlers.RegisterUser)
	public.GET("/user/:user", handlers.AuthHandler)
	public.DELETE("/user/:user", handlers.DeleteUser)
	public.GET("/user", handlers.GetAllUsers)

	// protected := r.Group("prod")
	// protected.Use(middlewares.ValidateRequest())

	bootstraps.Setup()

	r.Run(fmt.Sprintf(":%s", cfg.Port))
}

func pingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
