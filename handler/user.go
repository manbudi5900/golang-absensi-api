package handler

import (
	"absensi/helper"
	"absensi/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
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
	formatter := user.FormatUser(newUser, "12121121212")
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

	formatter := user.FormatUser(userLogin, "12121121212")
	response := helper.APIResponse("Account has been registered", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}
