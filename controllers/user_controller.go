package controllers

import (
	"log"
	"net/http"

	"github.com/docentre/docentre/services"
	"github.com/gin-gonic/gin"
)

type UserDto struct {
	ID       uint   `json:"id" example:"1"`
	Username string `json:"username" example:"username"`
	Email    string `json:"email" example:"email@mail.com"`
	Identity string `json:"identity" example:"user"`
}

// invalidResponseBody is the response body for request with invalid request body.
type invalidResponseBody struct {
	Msg string `json:"msg" example:"Invalid request body"`
}

// @Summary Create a user
// @Description Create a new user; the user will be created with the identity "user".
// @Tags User
// @Accept json
// @Produce json
// @Param body body controllers.UserCreate.requestBody true " "
// @Success 200 {object} controllers.UserCreate.successResponseBody
// @Failure 400 {object} invalidResponseBody
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
	err := c.Bind(&body)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, invalidResponseBody{
			Msg: "Invalid request body",
		})
		return
	}

	user, err := services.UserCreate(body.Username, body.Email, body.Password)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, existedResponseBody{
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
	c.JSON(http.StatusOK, successResponseBody{
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
// @Failure 400 {object} invalidResponseBody
// @Failure 404 {object} controllers.UserLogin.userNotFoundResponseBody
// @Router /login [post]
func UserLogin(c *gin.Context) {
	type requestBody struct {
		Username string `json:"username" binding:"required" example:"username"`
		Password string `json:"password" binding:"required" example:"password"`
	}
	type userNotFoundResponseBody struct {
		Msg string `json:"msg" example:"User not found"`
	}
	type successResponseBody struct {
		User UserDto `json:"user"`
	}

	var body requestBody
	err := c.Bind(&body)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, invalidResponseBody{
			Msg: "Invalid request body",
		})
		return
	}

	user, err := services.UserLogin(body.Username, body.Password)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusNotFound, userNotFoundResponseBody{
			Msg: "User not found",
		})
		return
	}

	userDto := UserDto{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Identity: user.Identity,
	}
	c.JSON(http.StatusOK, successResponseBody{
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
// @Failure 400 {object} invalidResponseBody
// @Failure 500 {object} controllers.GetUsersByUsername.unexpectedErrorResponseBody
// @Router /users [post]
func GetUsersByUsername(c *gin.Context) {
	type requestBody struct {
		Username string `json:"username" binding:"required" example:"username"`
	}
	// errorResponseBody the request failed due to unexpected error.
	type unexpectedErrorResponseBody struct {
		Msg string `json:"msg" example:"Unexpected error"`
	}
	type successResponseBody struct {
		Users []UserDto `json:"users"`
	}

	var body requestBody
	err := c.Bind(&body)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, invalidResponseBody{
			Msg: "Invalid request body",
		})
		return
	}

	// NOTE: A non-exist username does not return an error. It returns an empty array.
	users, err := services.GetUsersByUsername(body.Username)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, unexpectedErrorResponseBody{
			Msg: "Unexpected error",
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
	c.JSON(http.StatusOK, successResponseBody{
		Users: usersDto,
	})
}

// @Summary Get all users
// @Description Get all users.
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object} controllers.GetUsers.successResponseBody
// @Failure 500 {object} controllers.GetUsers.unexpectedErrorResponseBody
// @Router /users [get]
func GetUsers(c *gin.Context) {
	// errorResponseBody the request failed due to unexpected error.
	type unexpectedErrorResponseBody struct {
		Msg string `json:"msg" example:"Unexpected error"`
	}
	type successResponseBody struct {
		Users []UserDto `json:"users"`
	}

	users, err := services.GetUsers()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, unexpectedErrorResponseBody{
			Msg: "Unexpected error",
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
	c.JSON(http.StatusOK, successResponseBody{
		Users: usersDto,
	})
}
