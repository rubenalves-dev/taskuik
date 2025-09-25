package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type JsonResponse struct {
	Data    any    `json:"data,omitempty"`
	Status  int    `json:"status"`
	Message string `json:"message,omitempty"`
}

func OkJson(c *gin.Context, message string, data any) {
	Json(c, http.StatusOK, message, data)
}

func ErrJson(c *gin.Context, status int, err error) {
	Json(c, status, err.Error(), nil)
}

func Json(c *gin.Context, status int, message string, data any) {
	response := JsonResponse{
		Data:    data,
		Message: message,
		Status:  status,
	}
	c.JSON(status, response)
}
