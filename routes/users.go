package routes

import (
	"event_booking/models"
	"event_booking/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func ComparePassword(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func ValidateUser(email_id string, password string) bool {
	user, err := models.GetUserByEmail(email_id) // Get user from DB by email
	if err != nil {
		fmt.Println("failed to get email")
		return false // User not found or other DB error
	}

	// Compare the provided password with the stored password hash
	return ComparePassword(password, user.Password)
}
func create_user(ctx *gin.Context) {
	var user models.User
	err := ctx.BindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	hashedPassword, _ := HashPassword(user.Password)
	user.Password = hashedPassword
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
func login_user(ctx *gin.Context) {
	var user models.User
	err := ctx.BindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	fmt.Println("user", user)
	isValid := ValidateUser(user.Email, user.Password)
	if !isValid {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid email or password",
		})
		return
	}
	token, err := utils.GenerateTokenJwt(user.Email, user.Id)
	fmt.Print("error", err)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Server Error unable to login",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   token,
	})

}

func UserRouter(server *gin.Engine) {
	server.POST("/signup", create_user)
	server.POST("/login", login_user)
}
