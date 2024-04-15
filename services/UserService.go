package services

import (
	"github.com/docentre/docentre/initializers"
	"github.com/docentre/docentre/models"
)

func UserCreate(name string, email string, password string) (models.User, error) {
	user := models.User{Name: name, Email: email, Password: password}
	// user := models.User{Name: "Erin", Email: "aaa.com", Password: "123456"}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		return models.User{}, result.Error
	}

	return user, nil
}

func UserLogin(name string, password string) (models.User, error) {
	var user models.User
	result := initializers.DB.Where(&models.User{Name: name, Password: password}).First(&user)

	if result.Error != nil {
		return models.User{}, result.Error
	}
	user.Password = ""
	return user, nil
}
