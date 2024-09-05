package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mohamadhasan1992/go-rest-api.git/models"
	"github.com/mohamadhasan1992/go-rest-api.git/utils"
)

func signup(context *gin.Context) {
	var err error
	var user models.User
	err = context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request body!"})
		return
	}
	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save!"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "user created successfully!"})
}

func login(context *gin.Context) {
	var err error
	var user models.User
	err = context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request body!"})
		return
	}
	err = user.ValidateCredentials()
	if err != nil {
		context.JSON(http.StatusUnauthorized, err)
		return
	}
	token, err := utils.GenerateToken(user.Email, user.Id)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "could not authenticate user!"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "successfully loged in.", "token": token})
}
