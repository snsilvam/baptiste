package handlers

import (
	"encoding/json"
	"net/http"

	"baptiste.com/server"
	"github.com/gin-gonic/gin"
)

type HomeResponse struct {
	Message string `json:"message"`
	Status  bool   `json:"status"`
}

func HomeHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(HomeResponse{
			Message: "Welcome to Baptiste API REST☻.",
			Status:  true,
		})
	}
}

func GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, HomeResponse{
		Message: "Welcome to Baptiste API REST☻.",
		Status:  true,
	})
}
