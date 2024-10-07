package utils

import (
	"github.com/in-mai-space/disco/entities/models"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(u *models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}
