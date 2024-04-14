package controllers

import (
	"log"

	"github.com/docentre/docentre/initializers"
	"github.com/docentre/docentre/models"
	"github.com/gin-gonic/gin"
)

func UserCreate(c *gin.Context) {
	// get data off request body
	var body struct {
		Name     string `json:"name" binding:"required"`
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	c.Bind(&body)

	// create user
	user := models.User{Name: body.Name, Email: body.Email, Password: body.Password}
	// user := models.User{Name: "Erin", Email: "aaa.com", Password: "123456"}

	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.Status(400)
		return
	}

	// user.ID             // returns inserted data's primary key
	// result.Error        // returns error
	// result.RowsAffected // returns inserted records count

	c.JSON(200, gin.H{
		"user": user,
	})
}

func UserLogin(c *gin.Context) {
	// get data off request body
	var body struct {
		Name     string `json:"name" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	c.Bind(&body)

	var user models.User
	result := initializers.DB.Where(&models.User{Name: body.Name, Password: body.Password}).First(&user)

	if result.Error != nil {
		log.Default().Println(result.Error)
		c.JSON(200, gin.H{
			"user": nil,
		})
		return
	}

	c.JSON(200, gin.H{
		"user": user,
	})
}
