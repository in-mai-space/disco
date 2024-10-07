package base

import (
	"github.com/go-playground/validator"
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
	db       *gorm.DB
	validate *validator.Validate
}

func NewUserService(db *gorm.DB) UserServiceInterface {
	return &userService{
		db:       db,
		validate: validator.New(),
	}
}

func (u *userService) CreateUser(user *models.User) (*models.User, error) {
	if err := u.validate.Struct(user); err != nil {
		return nil, err
	}
	return CreateUser(user, u.db)
}

func (u *userService) GetUserByID(id uuid.UUID) (*models.User, error) {
	return GetUserByID(id, u.db)
}

func (u *userService) UpdateUser(user *models.User, id uuid.UUID) (*models.User, error) {
	if err := u.validate.Struct(user); err != nil {
		return nil, err
	}
	return UpdateUser(user, id, u.db)
}

func (u *userService) DeleteUser(id uuid.UUID) error {
	return DeleteUser(id, u.db)
}
