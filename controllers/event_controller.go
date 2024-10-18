package controllers

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    entities "github.com/grupoG/csw24-grupoG-ticket-gin/entities/event"
    "github.com/grupoG/csw24-grupoG-ticket-gin/services"
    "github.com/grupoG/csw24-grupoG-ticket-gin/utils"
)

type EventController struct {
    Service *services.EventService
}

func NewEventController(service *services.EventService) *EventController {
    return &EventController{Service: service}
}

func (ctrl *EventController) GetAllEvents(c *gin.Context) {
    events, err := ctrl.Service.GetAllEvents()
    if err != nil {
        utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
        return
    }
    c.JSON(http.StatusOK, events)
}

func (ctrl *EventController) GetEventByID(c *gin.Context) {
    idParam := c.Param("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    event, err := ctrl.Service.GetEventByID(uint(id))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
        return
    }

    c.JSON(http.StatusOK, event)
}

func (ctrl *EventController) CreateEvent(c *gin.Context) {
    var eventRequest entities.EventCrRequest
    if err := c.ShouldBindJSON(&eventRequest); err != nil {
        utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
        return
    }
    newEvent, err := ctrl.Service.CreateEvent(eventRequest)
    if err != nil {
        utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
        return
    }
    c.JSON(http.StatusCreated, newEvent)
}

func (ctrl *EventController) UpdateEvent(c *gin.Context) {
    idParam := c.Param("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    var eventRequest entities.EventUpRequest
    if err := c.ShouldBindJSON(&eventRequest); err != nil {
        utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
        return
    }

    updatedEvent, err := ctrl.Service.UpdateEvent(uint(id), eventRequest)
    if err != nil {
        utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
        return
    }
    c.JSON(http.StatusOK, updatedEvent)
}

func (ctrl *EventController) DeleteEvent(c *gin.Context) {
    idParam := c.Param("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    if err := ctrl.Service.DeleteEvent(uint(id)); err != nil {
        utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
        return
    }
    c.JSON(http.StatusNoContent, nil)
}