package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gokul656/product-service/app/dao"
	"github.com/gokul656/product-service/app/services"
	"github.com/gokul656/product-service/app/utils"
)

func AuthHandler(c *gin.Context) {
	response, err := services.GetUser(c.Param("user"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.CreateResponse(http.StatusBadRequest, err.Error(), ""))
		return
	}
	c.JSON(http.StatusOK, response)
}

func RegisterUser(c *gin.Context) {
	var request dao.UserDTO
	c.ShouldBindJSON(&request)

	response, err := services.CreateUser(&request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.CreateResponse(http.StatusBadRequest, err.Error(), ""))
		return
	}
	c.JSON(http.StatusOK, response)
}

func GetAllUsers(c *gin.Context) {
	response := services.GetAllUsers()
	c.JSON(http.StatusOK, response)
}

func DeleteUser(c *gin.Context) {
	_ = services.DeleteUserByName(c.Param("user"))
	c.JSON(http.StatusOK, utils.CreateResponse(http.StatusOK, "user deleted successfully", ""))
}
