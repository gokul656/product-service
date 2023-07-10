package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	helper "github.com/gokul656/product-service/app/utils"
)

func AuthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, helper.CreateResponse(http.StatusOK, "success", fmt.Sprintf("Hello %s!", c.Param("user"))))
}
