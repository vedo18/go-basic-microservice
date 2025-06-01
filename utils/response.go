package utils

import (
	"github.com/gin-gonic/gin"
)

type APIResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func SendSuccess(c *gin.Context, code int, data interface{}) {
	c.JSON(code, APIResponse{
		Status: "success",
		Data:   data,
	})
}

func SendError(c *gin.Context, code int, message string) {
	c.JSON(code, APIResponse{
		Status:  "error",
		Message: message,
	})
}
