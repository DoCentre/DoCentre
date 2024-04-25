package controllers

import (
	"log"

	"github.com/docentre/docentre/services"
	"github.com/gin-gonic/gin"
)

type UserDto struct {
	ID       uint   `json:"id" example:"1"`
	Username string `json:"username" example:"username"`
	Email    string `json:"email" example:"email@mail.com"`
	Identity string `json:"identity" example:"user"`
}

// @Summary Create a user
// @Description Create a new user; the user will be created with the identity "user".
// @Tags User
// @Accept json
// @Produce json
// @Param body body controllers.UserCreate.requestBody true " "
// @Success 200 {object} controllers.UserCreate.successResponseBody
// @Failure 400 {object} controllers.UserCreate.existedResponseBody
// @Router /user/create [post]
func UserCreate(c *gin.Context) {
	type requestBody struct {
		Username string `json:"username" binding:"required" example:"username"`
		Email    string `json:"email" binding:"required" example:"email@mail.com"`
		Password string `json:"password" binding:"required" example:"password"`
	}
	type existedResponseBody struct {
		Msg string `json:"msg" example:"User/Email already exists"`
	}
	type successResponseBody struct {
		User UserDto `json:"user"`
	}

	var body requestBody
	c.Bind(&body)

	user, err := services.UserCreate(body.Username, body.Email, body.Password)
	if err != nil {
		log.Default().Println(err)
		c.Status(400)
		c.JSON(400, existedResponseBody{
			Msg: "User/Email already exists",
		})
		return
	}

	userDto := UserDto{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Identity: user.Identity,
	}

	c.JSON(200, successResponseBody{
		User: userDto,
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
