package server

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
	"github.com/unrolled/secure"
)

type Config struct {
	Port          string
	SecureOptions secure.Options
	CorsOptions   cors.Options
	Audience      string
	Domain        string
}

type Server interface {
	Config() *Config
}

type Broker struct {
	config *Config

	router *gin.Engine
}

func (b *Broker) Config() *Config {
	return b.config
}

func NewServer(ctx context.Context, config *Config) (*Broker, error) {
	if config.Port == "" {
		return nil, errors.New("port is required")
	}

	broker := &Broker{
		config: config,
		//create a new routeador
		router: gin.New(),
	}

	return broker, nil
}

func (b *Broker) Start(binder func(s Server, r *gin.Engine)) {
	b.router = gin.New()

	binder(b, b.router)

	server := &http.Server{
		Addr: b.config.Port,
	}

	log.Println("Starting server on port", server.Addr)
	if err := http.ListenAndServe(b.config.Port, b.router); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
