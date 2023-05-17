package server

import (
	"context"
	"errors"
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
	//router *mux.Router
}

func (b *Broker) Config() *Config {
	return b.config
}

func NewServer(ctx context.Context, config *Config) (*Broker, error) {
	if config.Port == "" {
		return nil, errors.New("Port is required")
	}

	if config.DatabaseUrl == "" {
		return nil, errors.New("Url of database is required")
	}

	broker := &Broker{
		config: config,
		//router:
	}
	return broker, nil
}
