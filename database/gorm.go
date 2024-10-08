package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	gastos "baptiste.com/src/gasto/models"
)

type Database struct {
	Db *gorm.DB
}

func ConstructorDatabase(dsn string) (*Database, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error conectando a la base de datos: ", err)
		return nil, err
	}
	db.Exec("SET search_path TO baptiste")

	err = db.AutoMigrate(&gastos.Gastos{})
	if err != nil {
		return nil, err
	}

	return &Database{
		Db: db,
	}, nil
}
