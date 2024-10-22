package controllers

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    entities "github.com/grupoG/csw24-grupoG-ticket-gin/entities/ticket"
    "github.com/grupoG/csw24-grupoG-ticket-gin/services"
    "github.com/grupoG/csw24-grupoG-ticket-gin/utils"
)

type TicketController struct {
    Service *services.TicketService
}

func NewTicketController(service *services.TicketService) *TicketController {
    return &TicketController{Service: service}
}

// GetAllTickets godoc
// @Summary Get all tickets
// @Description Get a list of all tickets
// @Tags tickets
// @Produce json
// @Success 200 {array} entities.Ticket
// @Failure 500 {object} utils.ErrorResponse
// @Router /tickets [get]
func (ctrl *TicketController) GetAllTickets(c *gin.Context) {
    tickets, err := ctrl.Service.GetAllTickets()
    if err != nil {
        utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
        return
    }
    c.JSON(http.StatusOK, tickets)
}

// GetTicketByID godoc
// @Summary Get ticket by ID
// @Description Get details of a specific ticket by ID
// @Tags tickets
// @Produce json
// @Param id path int true "Ticket ID"
// @Success 200 {object} entities.Ticket
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Router /tickets/{id} [get]
func (ctrl *TicketController) GetTicketByID(c *gin.Context) {
    idParam := c.Param("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    ticket, err := ctrl.Service.GetTicketByID(uint(id))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Ticket not found"})
        return
    }

    c.JSON(http.StatusOK, ticket)
}

// CreateTicket godoc
// @Summary Create a new ticket
// @Description Create a new ticket with the given details
// @Tags tickets
// @Accept json
// @Produce json
// @Param ticket body entities.TicketCrRequest true "Ticket request body"
// @Success 201 {object} entities.Ticket
// @Failure 400 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /tickets [post]
func (ctrl *TicketController) CreateTicket(c *gin.Context) {
    var ticketRequest entities.TicketCrRequest
    if err := c.ShouldBindJSON(&ticketRequest); err != nil {
        utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
        return
    }
    newTicket, err := ctrl.Service.CreateTicket(ticketRequest)
    if err != nil {
        utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
        return
    }
    c.JSON(http.StatusCreated, newTicket)
}

// UpdateTicket godoc
// @Summary Update a ticket
// @Description Update details of an existing ticket by ID
// @Tags tickets
// @Accept json
// @Produce json
// @Param id path int true "Ticket ID"
// @Param ticket body entities.TicketUpRequest true "Ticket request body"
// @Success 200 {object} entities.Ticket
// @Failure 400 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /tickets/{id} [put]
func (ctrl *TicketController) UpdateTicket(c *gin.Context) {
    idParam := c.Param("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    var ticketRequest entities.TicketUpRequest
    if err := c.ShouldBindJSON(&ticketRequest); err != nil {
        utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
        return
    }

    updatedTicket, err := ctrl.Service.UpdateTicket(uint(id), ticketRequest)
    if err != nil {
        utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
        return
    }
    c.JSON(http.StatusOK, updatedTicket)
}

// DeleteTicket godoc
// @Summary Delete a ticket
// @Description Delete a ticket by ID
// @Tags tickets
// @Produce json
// @Param id path int true "Ticket ID"
// @Success 204 {object} nil
// @Failure 400 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /tickets/{id} [delete]
func (ctrl *TicketController) DeleteTicket(c *gin.Context) {
    idParam := c.Param("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    if err := ctrl.Service.DeleteTicket(uint(id)); err != nil {
        utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
        return
    }
    c.JSON(http.StatusNoContent, nil)
}