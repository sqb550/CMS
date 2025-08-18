package userServices

import (
	"CMS/app/models"
	"CMS/config/database"
)

func GetUserByUsername(Username string) (*models.User, error) {
	var user models.User
	result := database.DB.Where("username=?", Username).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func Register(user *models.User) error {
	result := database.DB.Create(user)
	return result.Error
}
