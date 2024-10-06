package base

import (
	"github.com/google/uuid"
	"github.com/in-mai-space/disco/entities/models"
	"gorm.io/gorm"
)

func FindByEmail(email string, db *gorm.DB) (*models.User, error) {
	var user models.User

	result := db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func CreateUser(user *models.User, db *gorm.DB) (*models.User, error) {
	err := db.Create(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func GetUserByID(id uuid.UUID, db *gorm.DB) (*models.User, error) {
	var user models.User

	result := db.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func UpdateUser(user *models.User, db *gorm.DB) (*models.User, error) {
	err := db.Save(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func DeleteUser(id uuid.UUID, db *gorm.DB) error {
	err := db.Delete(&models.User{}, id).Error
	return err
}
