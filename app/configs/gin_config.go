package configs

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gokul656/product-service/app/utils"
)

func ErrorHandler(c *gin.Context, err interface{}) {
	response := utils.CreateResponse(http.StatusInternalServerError, fmt.Sprintf("%s", err), "")
	c.AbortWithStatusJSON(http.StatusInternalServerError, response)
}
