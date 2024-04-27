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
// @Router /user [post]
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

// @Summary Login a user
// @Description Login a user with username and password.
// @Tags User
// @Accept json
// @Produce json
// @Param body body controllers.UserLogin.requestBody true " "
// @Success 200 {object} controllers.UserLogin.successResponseBody
// @Failure 404 {object} controllers.UserLogin.userNotFoundResponseBody
// @Router /login [post]
func UserLogin(c *gin.Context) {
	type requestBody struct {
		Username string `json:"username" binding:"required" example:"username"`
		Password string `json:"password" binding:"required" example:"password"`
	}
	type userNotFoundResponseBody struct {
		// Should always be nil.
		// XXX: Consider removing the field; also swaggo fails to generate example with null value.
		User *UserDto `json:"user"`
		Msg  string   `json:"msg" example:"User not found"`
	}
	type successResponseBody struct {
		User UserDto `json:"user"`
	}

	var body requestBody
	c.Bind(&body)

	user, err := services.UserLogin(body.Username, body.Password)
	if err != nil {
		log.Default().Println(err)
		c.JSON(404, userNotFoundResponseBody{
			User: nil,
			Msg:  "User not found",
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

// @Summary Get users by username
// @Description Get users with the same given username.
// @Tags User
// @Accept json
// @Produce json
// @Param body body controllers.GetUsersByUsername.requestBody true " "
// @Success 200 {object} controllers.GetUsersByUsername.successResponseBody
// @Failure 200 {object} controllers.GetUsersByUsername.usersNotFoundResponseBody
// @Router /users [post]
func GetUsersByUsername(c *gin.Context) {
	type requestBody struct {
		Username string `json:"username" binding:"required" example:"username"`
	}
	type usersNotFoundResponseBody struct {
		// Should always be empty.
		// XXX: Consider removing the field.
		Users []UserDto `json:"users"`
		Msg   string    `json:"msg" example:"Users not found"`
	}
	type successResponseBody struct {
		Users []UserDto `json:"users"`
	}

	var body requestBody
	c.Bind(&body)

	users, err := services.GetUsersByUsername(body.Username)
	if err != nil {
		log.Default().Println(err)
		c.JSON(200, usersNotFoundResponseBody{
			Users: []UserDto{},
			Msg:   "Users not found",
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

	c.JSON(200, successResponseBody{
		Users: usersDto,
	})
}
