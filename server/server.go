package server

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Config struct {
	Port        string
	DatabaseUrl string
}

type Server interface {
	Config() *Config
}

type Broker struct {
	config *Config
	router *gin.
}

func (b *Broker) Config() *Config {
	return b.config
}

func NewServer(ctx context.Context, config *Config) (*Broker, error) {
	if config.Port == "" {
		return nil, errors.New("port is required")
	}

	if config.DatabaseUrl == "" {
		return nil, errors.New("url of  database is requiered")
	}

	broker := &Broker{
		config: config,
		router: gin.New(),
	}

	return broker, nil
}

func (b *Broker) Start(binder func(s Server, r *gin.IRouter)) {
	//b.router = new router
	//binder(b, b.router)
	log.Println("Starting server on port", b.Config().Port)
	if err := http.ListenAndServe(b.config.Port, b.router); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
