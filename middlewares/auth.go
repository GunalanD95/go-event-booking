package middlewares

import (
	"event_booking/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authenticate(ctx *gin.Context) {
	auth_token := ctx.Request.Header.Get("Authorization")

	if auth_token == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "invalid token",
		})
		return
	}
	_, userId, err := utils.VerifyToken(auth_token)
	fmt.Println("user_id", userId, "err", err)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "not logged in",
		})
		return
	}
	ctx.Set("user-id", userId)
	ctx.Next()
}
