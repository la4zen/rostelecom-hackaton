package util

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/la4zen/rostelecom-hackaton/pkg/models"
)

var Secret = []byte("SecretKEYYERTYjmazerngwaqenrigsdefg")

func GenerateToken(user *models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &models.Claims{
		ID: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour * 7).Unix(),
		},
	})
	t, err := token.SignedString(Secret)
	if err != nil {
		return "", err
	}
	return t, nil
}
