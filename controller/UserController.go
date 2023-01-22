package controller

import (
	"fmt"
	"kemiskinan/helper"
	"kemiskinan/request"
	"kemiskinan/responses"
	"kemiskinan/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type userController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) *userController {
	return &userController{userService}
}

func (c *userController) GetUsers(cntx *gin.Context) {
	var users, err = c.userService.FindAll()
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
	}

	var usersResponse []responses.UserResponse

	for _, user := range users {
		var userResponse = helper.ConvertToUserResponse(user)
		usersResponse = append(usersResponse, userResponse)
	}

	cntx.JSON(http.StatusOK, usersResponse)
}

func (c *userController) GetUser(cntx *gin.Context) {
	var idString = cntx.Param("id")
	var id, _ = strconv.Atoi(idString)

	var user, err = c.userService.FindById(id)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": "Data tidak ditemukan",
		})
	}

	var userResponse = helper.ConvertToUserResponse(user)

	cntx.JSON(http.StatusOK, userResponse)
}

func (c *userController) CreateUser(cntx *gin.Context) {
	var userRequest request.CreateUserRequest

	var err = cntx.ShouldBindJSON(&userRequest)
	if err != nil {
		var errorMessages = []string{}

		for _, e := range err.(validator.ValidationErrors) {
			var errorMessage = fmt.Sprintf("Error on field %s, condition : %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": errorMessages,
		})
		return
	}

	user, err := c.userService.Create(userRequest)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	var userResponse = helper.ConvertToUserResponse(user)

	cntx.JSON(http.StatusCreated, gin.H{
		"data": userResponse,
	})
}

func (c *userController) UpdateUser(cntx *gin.Context) {
	var userRequest request.UpdateUserRequest

	var err = cntx.ShouldBindJSON(&userRequest)
	if err != nil {
		var errorMessages = []string{}

		for _, e := range err.(validator.ValidationErrors) {
			var errorMessage = fmt.Sprintf("Error on field %s, condition : %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": errorMessages,
		})
		return
	}

	var idString = cntx.Param("id")
	var id, _ = strconv.Atoi(idString)

	user, err := c.userService.Update(id, userRequest)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"errors": cntx.Error(err),
		})
		return
	}

	var userResponse = helper.ConvertToUserResponse(user)

	cntx.JSON(http.StatusOK, gin.H{
		"data": userResponse,
	})
}

func (c *userController) DeleteUser(cntx *gin.Context) {
	var idString = cntx.Param("id")
	var id, _ = strconv.Atoi(idString)

	_, err := c.userService.Delete(id)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": "Data gagal dihapus",
		})
		return
	}

	cntx.JSON(http.StatusOK, gin.H{
		"status": "Data berhasil dihapus",
	})
}
