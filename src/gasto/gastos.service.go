package gasto

import (
	"context"
	"errors"
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
	if gasto.Nombre == "" {
		return errors.New("nombre required")
	}
	if gasto.CostoDelGasto == 0 {
		return errors.New("costo required")
	}
	if gasto.Categoria == "" {
		gasto.Categoria = "No registrado."
	}
	if gasto.Lugar == "" {
		gasto.Lugar = "No registrado."
	}

	result := gs.database.Db.Create(&gasto)
	if result.Error != nil {
		log.Fatal("failed to insert record:", result.Error)
		return result.Error
	}

	return nil
}

func (gs *GastoService) GetAllGastos(ctx context.Context) ([]*gastos.Gastos, error) {
	var allGastos []*gastos.Gastos

	// Buscar todos los registros de gastos en la base de datos
	result := gs.database.Db.Find(&allGastos)
	if result.Error != nil {
		return nil, result.Error
	}

	// Formato de fecha amigable
	const layout = "02-01-2006" // Formato: dd-mm-yyyy

	// Iterar sobre los resultados y aplicar formato a la fecha
	for _, gasto := range allGastos {
		// Asignar la fecha formateada a un campo nuevo o reemplazar
		gasto.FechaFormateada = gasto.FechaDelGasto.Format(layout)
	}

	return allGastos, nil
}

func (gs *GastoService) UpdateGasto(ctx context.Context, gasto *gastos.Gastos) error {
	if gasto.ID == 0 {
		return errors.New("id required")
	}
	result := gs.database.Db.Model(&gastos.Gastos{}).Where("id = ?", gasto.ID).Updates(gasto)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
