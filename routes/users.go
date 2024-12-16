package routes

import (
	"event_booking/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func create_user(ctx *gin.Context) {
	var user models.User
	err := ctx.BindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = user.Save()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"user": user,
	})

}

func UserRouter(server *gin.Engine) {
	server.POST("/signup", create_user)
}
