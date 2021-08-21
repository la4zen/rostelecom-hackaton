package store

import (
	"log"

	"github.com/la4zen/rostelecom-hackaton/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Store struct {
	DB *gorm.DB
}

func New() *Store {
	_db, err := gorm.Open(postgres.Open("host=database user=postgres password=12345 sslmode=disable database=postgres"))
	if err != nil {
		log.Fatal(err)
	}
	_db.AutoMigrate(&models.User{}, &models.Room{})
	return &Store{
		DB: _db,
	}
}
