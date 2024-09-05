package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mohamadhasan1992/go-rest-api.git/utils"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")
	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Login first!"})
		return
	}
	userId, err := utils.VerifyToken(token)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "you are not loggedin!"})
		return
	}
	context.Set("userId", userId)
	context.Next()
}
