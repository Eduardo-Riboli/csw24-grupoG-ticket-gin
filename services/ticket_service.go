package services

import (
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
