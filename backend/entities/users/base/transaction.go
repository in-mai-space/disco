package base

import (
	"errors"

	"github.com/google/uuid"
	"github.com/in-mai-space/disco/entities/models"
	utils "github.com/in-mai-space/disco/errors"
	"gorm.io/gorm"
)

func FindByEmail(email string, db *gorm.DB) (*models.User, error) {
	var user models.User
	result := db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, utils.NotFoundError("user")
		}
		return nil, utils.InternalServerError(result.Error.Error())
	}
	return &user, nil
}

func CreateUser(user *models.User, db *gorm.DB) (*models.User, error) {
	if err := db.Create(user).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, utils.AlreadyExistsError("user")
		}
		return nil, utils.CreateFailedError("user")
	}
	return user, nil
}

func GetUserByID(id uuid.UUID, db *gorm.DB) (*models.User, error) {
	var user models.User
	result := db.First(&user, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, utils.NotFoundError("user")
		}
		return nil, utils.InternalServerError(result.Error.Error())
	}
	return &user, nil
}

func UpdateUser(user *models.User, id uuid.UUID, db *gorm.DB) (*models.User, error) {
	var existingUser models.User
	result := db.First(&existingUser, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, utils.NotFoundError("user")
		}
		return nil, utils.InternalServerError(result.Error.Error())
	}

	user.ID = id
	result = db.Save(user)
	if result.Error != nil {
		return nil, utils.UpdateFailedError("user")
	}
	return user, nil
}

func DeleteUser(id uuid.UUID, db *gorm.DB) error {
	result := db.Delete(&models.User{}, id)
	if result.Error != nil {
		return utils.DeleteFailedError("user")
	}
	if result.RowsAffected == 0 {
		return utils.NotFoundError("user")
	}
	return nil
}
