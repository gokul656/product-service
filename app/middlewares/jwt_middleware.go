package middlewares

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	h "github.com/gokul656/product-service/app/utils"
)

func ValidateRequest() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := validateToken(ctx)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, h.CreateResponse(http.StatusUnauthorized, err.Error(), ""))
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}

func validateToken(c *gin.Context) error {
	return errors.New("unauthorized request")
}
