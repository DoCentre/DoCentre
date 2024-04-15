package controllers

import (
	"log"

	"github.com/docentre/docentre/services"
	"github.com/gin-gonic/gin"
)

type UserDto struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Identity string `json:"identity"`
}

func UserCreate(c *gin.Context) {

	var body struct {
		Username string `json:"username" binding:"required"`
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	c.Bind(&body)

	user, err := services.UserCreate(body.Username, body.Email, body.Password)

	if err != nil {
		log.Default().Println(err)
		c.Status(400)
		c.JSON(400, gin.H{
			"msg": "User/Email already exists",
		})
		return
	}

	userDto := UserDto{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Identity: user.Identity,
	}

	c.JSON(200, gin.H{
		"user": userDto,
	})
}

func UserLogin(c *gin.Context) {
	var body struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	c.Bind(&body)

	user, err := services.UserLogin(body.Username, body.Password)

	if err != nil {
		log.Default().Println(err)
		c.JSON(404, gin.H{
			"user": nil,
			"msg":  "User not found",
		})
		return
	}

	userDto := UserDto{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Identity: user.Identity,
	}

	c.JSON(200, gin.H{
		"user": userDto,
	})
}

func GetUsersByUsername(c *gin.Context) {
	var body struct {
		Username string `json:"username" binding:"required"`
	}

	c.Bind(&body)

	users, err := services.GetUsersByUsername(body.Username)

	if err != nil {
		log.Default().Println(err)
		c.JSON(200, gin.H{
			"users": []UserDto{},
			"msg":   "Users not found",
		})
		return
	}

	var usersDto []UserDto

	for _, user := range users {
		userDto := UserDto{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
			Identity: user.Identity,
		}
		usersDto = append(usersDto, userDto)
	}

	c.JSON(200, gin.H{
		"users": usersDto,
	})
}
