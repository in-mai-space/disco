package utils

import (
	"github.com/in-mai-space/disco/entities/models"
	"github.com/in-mai-space/disco/models"
	"golang.org/x/crypto/bcrypt"
)

func (u *models.User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}
