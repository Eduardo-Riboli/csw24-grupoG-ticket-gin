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

func (ctrl *TicketController) GetAllTickets(c *gin.Context) {
    tickets, err := ctrl.Service.GetAllTickets()
    if err != nil {
        utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
        return
    }
    c.JSON(http.StatusOK, tickets)
}

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