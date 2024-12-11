package middlewares

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/restAPI/utils"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")
	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Authentication token not provided"})
		return
	}

	userId, err := utils.VerifyToken(token)
	fmt.Println("userId in createEvent===> ", userId)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not Authorized", "error": err})
		return
	}

	context.Set("userId", userId)
	// this will make sure that the next request handler runs in line
	context.Next()
}
