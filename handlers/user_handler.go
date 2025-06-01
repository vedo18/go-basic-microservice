package handlers

import (
	"basicMicroservice/models"
	"basicMicroservice/services"
	"basicMicroservice/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func GetAllUsers(c *gin.Context) {
	users, err := services.GetAllUsers()
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, "Failed to fetch users")
		return
	}

	utils.SendSuccess(c, http.StatusOK, users)
}

func GetUserById(c *gin.Context) {
	id := c.Param("id")

	user, err := services.GetUserById(id)
	if err != nil {
		utils.SendError(c, http.StatusNotFound, err.Error())
		return
	}

	utils.SendSuccess(c, http.StatusOK, user)
}

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.SendError(c, http.StatusBadRequest, "Invalid JSON")
		return
	}

	if err := validate.Struct(user); err != nil {
		utils.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	createdUser, err := services.CreateUser(user)
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, "Failed to create user in database")
		return
	}

	utils.SendSuccess(c, http.StatusCreated, createdUser)
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")

	// Check if user exists
	existingUser, err := services.GetUserById(id)
	if err != nil {
		utils.SendError(c, http.StatusNotFound, err.Error())
		return
	}

	var updateData models.User
	if err := c.ShouldBindJSON(&updateData); err != nil {
		utils.SendError(c, http.StatusBadRequest, "Invalid JSON")
		return
	}

	// Merge existing with new data
	if updateData.Name == "" {
		updateData.Name = existingUser.Name
	}
	if updateData.Email == "" {
		updateData.Email = existingUser.Email
	}

	updatedUser, err := services.UpdateUser(id, updateData)
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, "Failed to update user")
		return
	}

	utils.SendSuccess(c, http.StatusOK, updatedUser)
}
