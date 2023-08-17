package handlers

import (
	"github.com/gin-gonic/gin"
)

type ErrorMessage struct {
	Message string `json:"message"`
}

func NotFoundHandler(c *gin.Context) {
	errorMessage := ErrorMessage{Message: "Baptiste no encontro la ruta solicitada. ☻"}
	c.JSON(404, errorMessage)
}
