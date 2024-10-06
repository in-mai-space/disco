package base

import (
	"github.com/google/uuid"
	"github.com/in-mai-space/disco/entities/models"
	"gorm.io/gorm"
)

type UserService interface {
	CreateUser(user *models.User) (*models.User, error)
	GetUserByID(id uuid.UUID) (*models.User, error)
	UpdateUser(user *models.User) (*models.User, error)
	DeleteUser(id uuid.UUID) error
}

type userService struct {
	db *gorm.DB
}