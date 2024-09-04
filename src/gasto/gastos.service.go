package gasto

import (
	"context"
	"log"

	database "baptiste.com/database"
	gastos "baptiste.com/src/gasto/models"
)

type GastoService struct {
	database database.Database
}

func ConstructorGastoService(database database.Database) *GastoService {
	return &GastoService{
		database: database,
	}
}

func (gs *GastoService) CreateGasto(ctx context.Context, gasto *gastos.Gastos) error {
	result := gs.database.Db.Create(&gasto)
	if result.Error != nil {
		log.Fatal("failed to insert record:", result.Error)
		return result.Error
	}

	return nil
}

func (gs *GastoService) Hello(ctx context.Context) string {
	return "Hola mi primer endpoint"
}
