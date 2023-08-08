package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HomeResponse struct {
	Message string `json:"message"`
	Status  bool   `json:"status"`
}

type MessageError struct {
	Message string `json:"message"`
	Url     string `json:"url"`
}

func HomeHandler(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, HomeResponse{
		Message: "Welcome to Baptiste API REST☻.",
		Status:  true,
	})
}

func HelloHandler(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, HomeResponse{
		Message: "Hello Word",
		Status:  true,
	})
}
