package handler

import (
	"absensi/auth"
	"absensi/helper"
	"absensi/user"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
	authService auth.Service
}

func NewUserHandler(userService user.Service, authService auth.Service) *userHandler {
	return &userHandler{userService, authService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errorMessage := gin.H{"errors": helper.FormatValidatorError(err)}
		response := helper.APIResponse("Register account failed", http.StatusBadRequest, "failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	newUser, err := h.userService.RegisterUser(input)
	if err != nil {

		errorMessage := gin.H{"errors": helper.FormatValidatorError(err)}
		response := helper.APIResponse("Register account failed", http.StatusBadRequest, "failed", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	token, err := h.authService.GenerateToken(newUser.ID)
	if err != nil {

		errorMessage := gin.H{"errors": helper.FormatValidatorError(err)}
		response := helper.APIResponse("Register account failed", http.StatusBadRequest, "failed", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	formatter := user.FormatUser(newUser, token)
	response := helper.APIResponse("Account has been registered", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}
func (h *userHandler) LoginUser(c *gin.Context) {
	var input user.LoginUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errorMessage := gin.H{"errors": helper.FormatValidatorError(err)}
		response := helper.APIResponse("Login failed", http.StatusBadRequest, "failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	userLogin, err := h.userService.LoginUser(input)

	if err != nil {

		errorMessage := gin.H{"errors": helper.FormatValidatorError(err)}
		response := helper.APIResponse("Login failed", http.StatusBadRequest, "failed", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	token, err := h.authService.GenerateToken(userLogin.ID)
	if err != nil {

		errorMessage := gin.H{"errors": helper.FormatValidatorError(err)}
		response := helper.APIResponse("Register account failed", http.StatusBadRequest, "failed", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatUser(userLogin, token)
	response := helper.APIResponse("Account has been registered", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) UploadAvatar(c *gin.Context) {
	file, err := c.FormFile("avatar")
	if err != nil {

		data := gin.H{"is_upload": false}
		response := helper.APIResponse("Failed to upload avatar image", http.StatusBadRequest, "failed", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	userID := "cd3c5838-5acb-43b9-bb21-de5ee511a3d1"

	path := fmt.Sprintf("images/%s-%s", userID, file.Filename)
	err = c.SaveUploadedFile(file, path)
	if err != nil {

		data := gin.H{"is_upload": false}
		response := helper.APIResponse("Failed to upload avatar image", http.StatusBadRequest, "failed", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.userService.SaveAvatar(userID, path)
	if err != nil {

		data := gin.H{"is_upload": false}
		response := helper.APIResponse("Failed to upload avatar image", http.StatusBadRequest, "failed", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{"is_upload": true}
	response := helper.APIResponse("Success to upload avatar image", http.StatusOK, "success", data)
	c.JSON(http.StatusBadRequest, response)

}
