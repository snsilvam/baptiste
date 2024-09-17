package main

import (
	"log"
	"net/http"
	"os"

	"baptiste.com/config"
	"baptiste.com/database"
	"baptiste.com/src/gasto"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PORT := os.Getenv("PORT")
	DNS := os.Getenv("DNS")
	server, err := config.ConstructorServer(PORT, DNS, database.ConstructorDatabase)
	if err != nil {
		log.Fatal("error inicializando el servidor ", err)
	}
	server.Router.Use(cors.Default())

	gastoModule := gasto.ConstructorGastoModule(*server.Database)
	gastoModule.RegisterGastoRoutes(server.Router)

	// Crear una ruta GET en el endpoint "/hello"
	server.Router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, World!. Welcome to baptiste api♥",
		})
	})
	server.StartServer()
}
