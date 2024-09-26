package gasto

import "time"

type Gastos struct {
	ID              uint `gorm:"primaryKey"`
	Nombre          string
	CostoDelGasto   int
	FechaDelGasto   time.Time `json:"fecha_del_gasto"`
	FechaFormateada string    `json:"fecha_formateada,omitempty"`
	Categoria       string
	Lugar           string
}
