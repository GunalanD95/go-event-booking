package routes

import (
	"event_booking/middlewares"
	"event_booking/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func get_events(ctx *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	fmt.Println("Events: ", events)
	ctx.JSON(http.StatusOK, gin.H{
		"events": events,
	})
}

func create_event(ctx *gin.Context) {
	var event models.Event
	err := ctx.BindJSON(&event)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	userIdRaw, exists := ctx.Get("user-id")
	if !exists {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "User ID not found"})
		return
	}
	event.UserId = userIdRaw.(int) // Type assertion with a check
	err = event.Save()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"event": event,
	})
}

func get_event(ctx *gin.Context) {
	idParam := ctx.Param("id")
	eventID, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	event, err := models.GetEventById(eventID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"event": event})
}

func update_event(ctx *gin.Context) {
	idParam := ctx.Param("id")
	eventID, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	event, err := models.GetEventById(eventID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Could not find event with this event_id"})
		return
	}
	userIdRaw, exists := ctx.Get("user-id")
	if !exists {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "User ID not found"})
		return
	}

	if event.UserId != userIdRaw.(int) {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "you don't have permission to delete this event",
		})
		return
	}
	var updatedEvent models.Event
	if err := ctx.BindJSON(&updatedEvent); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = models.UpdateEventById(eventID, updatedEvent)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Could not find event with this event_id"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"event": updatedEvent})

}

func delete_event(ctx *gin.Context) {
	idParam := ctx.Param("id")
	eventID, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	event, err := models.GetEventById(eventID)
	userIdRaw, exists := ctx.Get("user-id")
	if !exists {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "User ID not found"})
		return
	}

	if event.UserId != userIdRaw.(int) {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "you don't have permission to delete this event",
		})
		return
	}
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Could not find event with this event_id"})
		return
	}
	err = models.DeleteEventById(eventID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Could not delete event with this event_id"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"sucess": "event delete sucessfully"})
}

func register_event(ctx *gin.Context) {
	// Retrieve userId from the middleware context
	userIdRaw, exists := ctx.Get("user-id")
	if !exists {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "User ID not found in context"})
		return
	}

	userId, ok := userIdRaw.(int)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID format"})
		return
	}

	// Parse JSON input for the event registration
	var event models.EventRegistration
	if err := ctx.ShouldBindJSON(&event); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	event.UserId = userId
	err := event.Save()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "interval server error"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message":  "Event registered successfully",
		"user-id":  userId,
		"name":     event.Name,
		"event_id": event.EventId,
	})
}

func EventRouter(server *gin.Engine) {
	server.GET("/events", get_events)
	server.GET("/event/:id", get_event)

	auth_routes := server.Group("/")
	auth_routes.POST("/create_event", middlewares.Authenticate, create_event)
	auth_routes.PUT("/event/:id", middlewares.Authenticate, update_event)
	auth_routes.DELETE("/event/:id", middlewares.Authenticate, delete_event)

	auth_routes.POST("/register_event", middlewares.Authenticate, register_event)
}
