package entities

type TicketRefundRequest struct {
    TicketID uint `json:"ticket_id" validate:"required"`
}