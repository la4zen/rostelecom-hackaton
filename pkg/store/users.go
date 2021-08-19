package store

import (
	"database/sql"
	"errors"

	"github.com/la4zen/rostelecom-hackaton/pkg/models"
)

func (s *Store) GetUser(user *models.User) error {
	result := s.DB.First(&user, "password=? and login=?", user.Password, user.Login)
	if result.Error != nil {
		if result.Error == sql.ErrNoRows {
			return errors.New("login or password is invalid")
		}
	}
	return nil
}

func (s *Store) CreateUser(user *models.User) error {
	result := s.DB.First(&user, "login=?", user.Login)
	if result.RowsAffected > 0 {
		return errors.New("user with this login already exists")
	}
	s.DB.Create(&user)
	return nil
}
