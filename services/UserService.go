package services

import (
	"log"

	"github.com/docentre/docentre/initializers"
	"github.com/docentre/docentre/models"
)

func UserCreate(name string, email string, password string) (models.User, error) {
	user := models.User{Username: name, Email: email, Password: password, Identity: "user"}
	// user := models.User{Name: "Erin", Email: "aaa.com", Password: "123456"}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		return models.User{}, result.Error
	}

	return user, nil
}

func UserLogin(name string, password string) (models.User, error) {
	var user models.User
	result := initializers.DB.Where(&models.User{Username: name, Password: password}).First(&user)

	if result.Error != nil {
		return models.User{}, result.Error
	}
	user.Password = ""
	return user, nil
}

func GetUsersByUsername(name string) ([]models.User, error) {
	var users []models.User
	// perform like query
	result := initializers.DB.Where("username LIKE ?", "%"+name+"%").Find(&users)

	log.Default().Println(users)
	if result.Error != nil {
		return []models.User{}, result.Error
	}
	return users, nil
}
