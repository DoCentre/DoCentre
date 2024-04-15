package controllers

import (
	"log"

	"github.com/docentre/docentre/services"
	"github.com/gin-gonic/gin"
)

type UserDto struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func UserCreate(c *gin.Context) {

	var body struct {
		Name     string `json:"name" binding:"required"`
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	c.Bind(&body)

	user, err := services.UserCreate(body.Name, body.Email, body.Password)

	if err != nil {
		log.Default().Println(err)
		c.Status(400)
		c.JSON(400, gin.H{
			"msg": "User/Email already exists",
		})
		return
	}

	userDto := UserDto{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}

	c.JSON(200, gin.H{
		"user": userDto,
	})
}

func UserLogin(c *gin.Context) {
	var body struct {
		Name     string `json:"name" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	c.Bind(&body)

	user, err := services.UserLogin(body.Name, body.Password)

	if err != nil {
		log.Default().Println(err)
		c.JSON(404, gin.H{
			"user": nil,
			"msg":  "User not found",
		})
		return
	}

	userDto := UserDto{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}

	c.JSON(200, gin.H{
		"user": userDto,
	})
}
