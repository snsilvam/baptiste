package server

import (
	"context"
	"errors"
	"log"
	"net/http"

	"baptiste.com/database"
	"baptiste.com/repository"
	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
	"github.com/unrolled/secure"
)

type Config struct {
	Port          string
	ProjectID     string
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

	corsMiddleware := cors.New(b.config.CorsOptions)
	routerWithCors := corsMiddleware.Handler(b.router)

	secureMiddleware := secure.New(b.config.SecureOptions)
	finalHandler := secureMiddleware.Handler(routerWithCors)
	server := &http.Server{
		Addr:    b.config.Port,
		Handler: finalHandler,
	}

	log.Println("Starting server on port", server.Addr)
	if err := http.ListenAndServe(b.config.Port, b.router); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
