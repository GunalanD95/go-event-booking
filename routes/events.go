package routes

import (
	"event_booking/models"
	"event_booking/utils"
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
	auth_token := ctx.Request.Header.Get("Authorization")

	if auth_token == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid token",
		})
		return
	}
	ok, _ := utils.VerifyToken(auth_token)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "not logged in",
		})
		return
	}

	var event models.Event
	err := ctx.BindJSON(&event)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
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
	// Use the event_id to fetch the event
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
	_, err = models.GetEventById(eventID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Could not find event with this event_id"})
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
	_, err = models.GetEventById(eventID)
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

func EventRouter(server *gin.Engine) {
	server.GET("/events", get_events)
	server.POST("/create_event", create_event)
	server.GET("/event/:id", get_event)
	server.PUT("/event/:id", update_event)
	server.DELETE("/event/:id", delete_event)
}
