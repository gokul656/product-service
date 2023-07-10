package utils

import (
	t "github.com/gokul656/product-service/app/types"
)

func CreateResponse(statuscode int, message, data string) t.APIResponse {
	return t.APIResponse{
		StatusCode: statuscode,
		Message:    message,
		Data:       data,
	}
}
