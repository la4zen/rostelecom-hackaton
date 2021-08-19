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
	_db, err := gorm.Open(postgres.Open("host=localhost user=postgres password=897+897 sslmode=disable database=rostelecom"))
	if err != nil {
		log.Fatal(err)
	}
	_db.AutoMigrate(&models.User{})
	return &Store{
		DB: _db,
	}
}
