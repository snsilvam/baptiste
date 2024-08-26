package database

import (
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Start(dsn string) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error conectando a la base de datos: ", err)
	}

	db.AutoMigrate(&Gastos{})
}

type Gastos struct {
	ID            uint `gorm:"primaryKey"`
	Nombre        string
	CostoDelGasto int
	FechaDelGasto time.Time
	Categoria     string
	Lugar         string
}
