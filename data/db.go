package data

import (
	"log"

	"PNC-Parcial-2/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(dsn string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error al conectar con base de datos: ", err)
	}

	err = db.AutoMigrate(&entities.Book{})
	if err != nil {
		log.Fatal("Error al migrar tablas: ", err)
	}

	DB = db
	return db
}
