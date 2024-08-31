package config

import (
	"errors"
	"fmt"

	"baptiste.com/database"
	"github.com/gin-gonic/gin"
)

type Server struct {
	Port     string
	Database *database.DatabasePostgres
	Router   *gin.Engine
}

func ConstructorServer(port string, dns string) (*Server, error) {
	if port == "" {
		return nil, errors.New("el servidor necesita un puerto para inicializar la applicacion")
	}
	if dns == "" {
		return nil, errors.New("el servidor necesita un dns para inicializar la base de datos")
	}

	router := gin.Default()
	database, err := database.ConstructorDatabasePostgres(dns)
	if err != nil {
		return nil, err
	}

	return &Server{
		Port:     port,
		Database: database,
		Router:   router,
	}, nil
}

func (s *Server) StartServer() {
	fmt.Println("Servidor inicializado en el puerto: ", s.Port)
	s.Router.Run(s.Port)
}
