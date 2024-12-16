package main

import (
	"event_booking/db"
	"fmt"

	"event_booking/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Starting the application")
	err := db.InitDB()
	if err != nil {
		fmt.Println("Unable to start the application. Error:", err)
		return
	}
	server := gin.Default()
	routes.EventRouter(server)
	routes.UserRouter(server)
	server.Run(":8080")
}
