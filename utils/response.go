package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GenerateErrorResponse(message string, c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{
		"error": message,
	})
}
