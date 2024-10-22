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

// GetAllEvents godoc
// @Summary Get all events
// @Description Get a list of all events
// @Tags events
// @Produce json
// @Success 200 {array} entities.Event
// @Failure 500 {object} utils.ErrorResponse
// @Router /events [get]
func (ctrl *EventController) GetAllEvents(c *gin.Context) {
    events, err := ctrl.Service.GetAllEvents()
    if (err != nil) {
        utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
        return
    }
    c.JSON(http.StatusOK, events)
}

// GetEventByID godoc
// @Summary Get event by ID
// @Description Get details of a specific event by ID
// @Tags events
// @Produce json
// @Param id path int true "Event ID"
// @Success 200 {object} entities.Event
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Router /events/{id} [get]
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

// CreateEvent godoc
// @Summary Create a new event
// @Description Create a new event with the given details
// @Tags events
// @Accept json
// @Produce json
// @Param event body entities.EventCrRequest true "Event request body"
// @Success 201 {object} entities.Event
// @Failure 400 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /events [post]
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

// UpdateEvent godoc
// @Summary Update an event
// @Description Update details of an existing event by ID
// @Tags events
// @Accept json
// @Produce json
// @Param id path int true "Event ID"
// @Param event body entities.EventUpRequest true "Event request body"
// @Success 200 {object} entities.Event
// @Failure 400 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /events/{id} [put]
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

// DeleteEvent godoc
// @Summary Delete an event
// @Description Delete an event by ID
// @Tags events
// @Produce json
// @Param id path int true "Event ID"
// @Success 204 {object} nil
// @Failure 400 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /events/{id} [delete]
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