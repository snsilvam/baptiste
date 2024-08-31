package main

import (
	"log"
	"os"

	"baptiste.com/config"
	"baptiste.com/src/gasto"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PORT := os.Getenv("PORT")
	DNS := os.Getenv("DNS")
	server, err := config.ConstructorServer(PORT, DNS)
	if err != nil {
		log.Fatal("error inicializando el servidor ", err)
	}

	gastoModule := gasto.ConstructorGastoModule(*server.Database)
	gastoModule.RegisterGastoRoutes(server.Router)

	server.StartServer()
}
