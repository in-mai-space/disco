package base

import (
	"github.com/google/uuid"
	"github.com/in-mai-space/disco/entities/models"
	"gorm.io/gorm"
)

type UserServiceInterface interface {
	CreateUser(user *models.User) (*models.User, error)
	GetUserByID(id uuid.UUID) (*models.User, error)
	UpdateUser(user *models.User, id uuid.UUID) (*models.User, error)
	DeleteUser(id uuid.UUID) error
}

type userService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) UserServiceInterface {
	return &userService{db: db}
}

func (u *userService) CreateUser(user *models.User) (*models.User, error) {
	return CreateUser(user, u.db)
}

func (u *userService) GetUserByID(id uuid.UUID) (*models.User, error) {
	return GetUserByID(id, u.db)
}

func (u *userService) UpdateUser(user *models.User, id uuid.UUID) (*models.User, error) {
	return UpdateUser(user, id, u.db)
}

func (u *userService) DeleteUser(id uuid.UUID) error {
	return DeleteUser(id, u.db)
}
