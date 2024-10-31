package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	entities "github.com/grupoG/csw24-grupoG-ticket-gin/entities/ticket"
	"github.com/grupoG/csw24-grupoG-ticket-gin/services"
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
// @Success 200 {array} entities.TicketResponse
// @Failure 500 {object} map[string]string
// @Router /tickets [get]
func (ctrl *TicketController) GetAllTickets(c *gin.Context) {
    tickets, err := ctrl.Service.GetAllTickets()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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
// @Success 200 {object} entities.TicketResponse
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
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
// @Success 201 {object} entities.TicketResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /tickets [post]
func (ctrl *TicketController) CreateTicket(c *gin.Context) {
    var ticketRequest entities.TicketCrRequest
    if err := c.ShouldBindJSON(&ticketRequest); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    newTicket, err := ctrl.Service.CreateTicket(ticketRequest)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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
// @Success 200 {object} entities.TicketResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
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
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    updatedTicket, err := ctrl.Service.UpdateTicket(uint(id), ticketRequest)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /tickets/{id} [delete]
func (ctrl *TicketController) DeleteTicket(c *gin.Context) {
    idParam := c.Param("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    if err := ctrl.Service.DeleteTicket(uint(id)); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusNoContent, nil)
}

// PurchaseTicket godoc
// @Summary Purchase a ticket
// @Description Purchase a ticket by providing necessary details
// @Tags tickets
// @Accept json
// @Produce json
// @Param ticket body entities.TicketPurchaseRequest true "Ticket purchase request body"
// @Success 201 {object} entities.TicketResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /tickets/purchase [post]
func (ctrl *TicketController) PurchaseTicket(c *gin.Context) {
    var ticketRequest entities.TicketPurchaseRequest
    if err := c.ShouldBindJSON(&ticketRequest); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    newTicket, err := ctrl.Service.PurchaseTicket(ticketRequest)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, newTicket)
}

// SellTicket godoc
// @Summary Sell a ticket
// @Description List a ticket for sale on the platform
// @Tags tickets
// @Accept json
// @Produce json
// @Param ticket body entities.TicketSellRequest true "Ticket sell request body"
// @Success 201 {object} entities.TicketResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /tickets/sell [post]
func (ctrl *TicketController) SellTicket(c *gin.Context) {
    var ticketRequest entities.TicketSellRequest
    if err := c.ShouldBindJSON(&ticketRequest); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    newTicket, err := ctrl.Service.SellTicket(ticketRequest)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, newTicket)
}

// AuthenticateTicket godoc
// @Summary Authenticate a ticket
// @Description Authenticate a ticket by scanning its verification code
// @Tags tickets
// @Accept json
// @Produce json
// @Param ticket body entities.TicketAuthRequest true "Ticket authentication request body"
// @Success 200 {object} entities.TicketResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /tickets/authenticate [post]
func (ctrl *TicketController) AuthenticateTicket(c *gin.Context) {
    var ticketRequest entities.TicketAuthRequest
    if err := c.ShouldBindJSON(&ticketRequest); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    authenticatedTicket, err := ctrl.Service.AuthenticateTicket(ticketRequest)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, authenticatedTicket)
}

// RefundTicket godoc
// @Summary Refund a ticket
// @Description Request a refund for a purchased ticket
// @Tags tickets
// @Accept json
// @Produce json
// @Param ticket body entities.TicketRefundRequest true "Ticket refund request body"
// @Success 200 {object} entities.TicketResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /tickets/refund [post]
func (ctrl *TicketController) RefundTicket(c *gin.Context) {
    var ticketRequest entities.TicketRefundRequest
    if err := c.ShouldBindJSON(&ticketRequest); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    refundedTicket, err := ctrl.Service.RefundTicket(ticketRequest)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, refundedTicket)
}