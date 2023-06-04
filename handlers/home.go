package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HomeResponse struct {
	Message string `json:"message"`
	Status  bool   `json:"status"`
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
