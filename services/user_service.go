package services

import (
	"log"

	"github.com/docentre/docentre/models"
	"github.com/docentre/docentre/repositories"
)

func UserCreate(name string, email string, password string) (models.User, error) {
	user := models.User{Username: name, Email: email, Password: password, Identity: "user"}

	result := repositories.DB.Create(&user)

	if result.Error != nil {
		return models.User{}, result.Error
	}

	return user, nil
}

func UserLogin(name string, password string) (models.User, error) {
	var user models.User
	result := repositories.DB.Where(&models.User{Username: name, Password: password}).First(&user)

	if result.Error != nil {
		return models.User{}, result.Error
	}
	user.Password = ""
	return user, nil
}

func GetUsersByUsername(name string) ([]models.User, error) {
	var users []models.User
	// perform like query
	result := repositories.DB.Where("username LIKE ?", "%"+name+"%").Find(&users)

	log.Println(users)
	if result.Error != nil {
		return []models.User{}, result.Error
	}
	return users, nil
}
