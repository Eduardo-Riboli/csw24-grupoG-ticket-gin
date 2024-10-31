package services

import (
	"errors"
	"log"

	entities "github.com/grupoG/csw24-grupoG-ticket-gin/entities/ticket"
	"github.com/grupoG/csw24-grupoG-ticket-gin/models"
	"github.com/grupoG/csw24-grupoG-ticket-gin/repositories"
)

type TicketService struct {
    Repository *repositories.TicketRepository
}

func NewTicketService(repo *repositories.TicketRepository) *TicketService {
    return &TicketService{Repository: repo}
}

func (ticketService *TicketService) GetAllTickets() ([]entities.TicketResponse, error) {
    tickets, err := ticketService.Repository.GetAll()
    if err != nil {
        return nil, err
    }

    var ticketResponses []entities.TicketResponse
    for _, ticket := range tickets {
        ticketResponses = append(ticketResponses, entities.TicketResponse{
            ID:               ticket.ID,
            EventID:          ticket.EventID,
            TenantID:         ticket.TenantID,
            OriginalPrice:    ticket.OriginalPrice,
            SellerID:         ticket.SellerID,
            VerificationCode: ticket.VerificationCode,
            Status:           ticket.Status,
        })
    }

    return ticketResponses, nil
}

func (ticketService *TicketService) GetTicketByID(id uint) (entities.TicketResponse, error) {
    ticket, err := ticketService.Repository.GetByID(id)
    if err != nil {
        return entities.TicketResponse{}, err
    }

    return entities.TicketResponse{
        ID:               ticket.ID,
        EventID:          ticket.EventID,
        TenantID:         ticket.TenantID,
        OriginalPrice:    ticket.OriginalPrice,
        SellerID:         ticket.SellerID,
        VerificationCode: ticket.VerificationCode,
        Status:           ticket.Status,
    }, nil
}

func (ticketService *TicketService) CreateTicket(ticketRequest entities.TicketCrRequest) (entities.TicketResponse, error) {
    ticket := models.Ticket{
        EventID:          ticketRequest.EventID,
        TenantID:         ticketRequest.TenantID,
        OriginalPrice:    ticketRequest.OriginalPrice,
        SellerID:         ticketRequest.SellerID,
        VerificationCode: ticketRequest.VerificationCode,
        Status:           ticketRequest.Status,
    }

    createdTicket, err := ticketService.Repository.Create(ticket)
    if err != nil {
        return entities.TicketResponse{}, err
    }

    return entities.TicketResponse{
        ID:               createdTicket.ID,
        EventID:          createdTicket.EventID,
        TenantID:         createdTicket.TenantID,
        OriginalPrice:    createdTicket.OriginalPrice,
        SellerID:         createdTicket.SellerID,
        VerificationCode: createdTicket.VerificationCode,
        Status:           createdTicket.Status,
    }, nil
}

func (ticketService *TicketService) UpdateTicket(id uint, ticketRequest entities.TicketUpRequest) (entities.TicketResponse, error) {
    ticket, err := ticketService.Repository.GetByID(id)
    if err != nil {
        return entities.TicketResponse{}, err
    }

    if ticketRequest.OriginalPrice != 0 {
        ticket.OriginalPrice = ticketRequest.OriginalPrice
    }
    if ticketRequest.VerificationCode != "" {
        ticket.VerificationCode = ticketRequest.VerificationCode
    }
    if ticketRequest.Status != "" {
        ticket.Status = ticketRequest.Status
    }

    updatedTicket, err := ticketService.Repository.Update(ticket)
    if err != nil {
        return entities.TicketResponse{}, err
    }

    return entities.TicketResponse{
        ID:               updatedTicket.ID,
        EventID:          updatedTicket.EventID,
        TenantID:         updatedTicket.TenantID,
        OriginalPrice:    updatedTicket.OriginalPrice,
        SellerID:         updatedTicket.SellerID,
        VerificationCode: updatedTicket.VerificationCode,
        Status:           updatedTicket.Status,
    }, nil
}

func (ticketService *TicketService) DeleteTicket(id uint) error {
    return ticketService.Repository.Delete(id)
}

func (ticketService *TicketService) NotifyUser(userID uint, message string) error {
    log.Printf("Notifying user %d: %s", userID, message)
    return nil
}

func (ticketService *TicketService) PurchaseTicket(ticketRequest entities.TicketPurchaseRequest) (entities.TicketResponse, error) {
    ticket := models.Ticket{
        EventID:          ticketRequest.EventID,
        TenantID:         ticketRequest.TenantID,
        OriginalPrice:    ticketRequest.OriginalPrice,
        SellerID:         ticketRequest.SellerID,
        VerificationCode: ticketRequest.VerificationCode,
        Status:           "Purchased",
    }

    createdTicket, err := ticketService.Repository.Create(ticket)
    if err != nil {
        return entities.TicketResponse{}, err
    }

        // Send confirmation email or push notification here

    return entities.TicketResponse{
        ID:               createdTicket.ID,
        EventID:          createdTicket.EventID,
        TenantID:         createdTicket.TenantID,
        OriginalPrice:    createdTicket.OriginalPrice,
        SellerID:         createdTicket.SellerID,
        VerificationCode: createdTicket.VerificationCode,
        Status:           createdTicket.Status,
    }, nil
}

func (ticketService *TicketService) SellTicket(ticketRequest entities.TicketSellRequest) (entities.TicketResponse, error) {
    ticket := models.Ticket{
        EventID:          ticketRequest.EventID,
        TenantID:         ticketRequest.TenantID,
        OriginalPrice:    ticketRequest.OriginalPrice,
        SellerID:         ticketRequest.SellerID,
        VerificationCode: ticketRequest.VerificationCode,
        Status:           "For Sale",
    }

    createdTicket, err := ticketService.Repository.Create(ticket)
    if err != nil {
        return entities.TicketResponse{}, err
    }

    // Send notification to the seller here

    return entities.TicketResponse{
        ID:               createdTicket.ID,
        EventID:          createdTicket.EventID,
        TenantID:         createdTicket.TenantID,
        OriginalPrice:    createdTicket.OriginalPrice,
        SellerID:         createdTicket.SellerID,
        VerificationCode: createdTicket.VerificationCode,
        Status:           createdTicket.Status,
    }, nil
}

func (ticketService *TicketService) AuthenticateTicket(ticketRequest entities.TicketAuthRequest) (entities.TicketResponse, error) {
    ticket, err := ticketService.Repository.GetByVerificationCode(ticketRequest.VerificationCode)
    if err != nil {
        return entities.TicketResponse{}, err
    }

    if ticket.Status != "Purchased" {
        return entities.TicketResponse{}, errors.New("ticket is not valid for entry")
    }

    ticket.Status = "Used"
    updatedTicket, err := ticketService.Repository.Update(ticket)
    if err != nil {
        return entities.TicketResponse{}, err
    }

    return entities.TicketResponse{
        ID:               updatedTicket.ID,
        EventID:          updatedTicket.EventID,
        TenantID:         updatedTicket.TenantID,
        OriginalPrice:    updatedTicket.OriginalPrice,
        SellerID:         updatedTicket.SellerID,
        VerificationCode: updatedTicket.VerificationCode,
        Status:           updatedTicket.Status,
    }, nil
}

func (ticketService *TicketService) RefundTicket(ticketRequest entities.TicketRefundRequest) (entities.TicketResponse, error) {
    ticket, err := ticketService.Repository.GetByID(ticketRequest.TicketID)
    if err != nil {
        return entities.TicketResponse{}, err
    }

    // Check refund policies here
    if ticket.Status != "Purchased" {
        return entities.TicketResponse{}, errors.New("ticket is not eligible for refund")
    }

    ticket.Status = "Refunded"
    updatedTicket, err := ticketService.Repository.Update(ticket)
    if err != nil {
        return entities.TicketResponse{}, err
    }

    transaction, err := ticketService.Repository.GetTransactionByTicketID(ticket.ID)
    if err != nil {
        return entities.TicketResponse{}, err
    }
    transaction.TransactionStatus = "Refunded"
    _, err = ticketService.Repository.UpdateTransaction(transaction)
    if err != nil {
        return entities.TicketResponse{}, err
    }

    err = ticketService.NotifyUser(ticket.SellerID, "Your ticket has been refunded.")
    if err != nil {
        return entities.TicketResponse{}, err
    }

    return entities.TicketResponse{
        ID:               updatedTicket.ID,
        EventID:          updatedTicket.EventID,
        TenantID:         updatedTicket.TenantID,
        OriginalPrice:    updatedTicket.OriginalPrice,
        SellerID:         updatedTicket.SellerID,
        VerificationCode: updatedTicket.VerificationCode,
        Status:           updatedTicket.Status,
    }, nil
}
