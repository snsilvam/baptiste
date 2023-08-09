package server

import (
	"context"
	"errors"
	"log"
	"net/http"

	"baptiste.com/database"
	"baptiste.com/repository"
	"github.com/gin-gonic/gin"
)

type Config struct {
	Port      string
	ProjectID string
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

	if config.ProjectID == "" {
		return nil, errors.New("projectID of  database is required")
	}

	broker := &Broker{
		config: config,
		//create a new routeador
		router: gin.New(),
	}

	return broker, nil
}

func (b *Broker) Start(binder func(s Server, r *gin.Engine)) {
	ctx := context.Background()
	repo, err := database.NewClientFirestore(ctx, b.config.ProjectID)

	if err != nil {
		log.Fatal("error creating a new client of firestore: ", err)
	}

	repository.SetRepository(repo)
	repository.SetUsersRepository(repo)
	repository.SetTrackingMonthlyFixedExpensesRepository(repo)

	b.router = gin.New()

	binder(b, b.router)

	log.Println("Starting server on port", b.Config().Port)

	if err := http.ListenAndServe(b.config.Port, b.router); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
