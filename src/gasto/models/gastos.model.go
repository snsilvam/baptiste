package gasto

import "time"

type Gastos struct {
	ID            uint `gorm:"primaryKey"`
	Nombre        string
	CostoDelGasto int
	FechaDelGasto time.Time
	Categoria     string
	Lugar         string
}
