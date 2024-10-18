package entities

import "time"

type TransactionCrRequest struct {
    TenantID          uint      `json:"tenant_id" validate:"required"`
    BuyerID           uint      `json:"buyer_id" validate:"required"`
    TicketID          uint      `json:"ticket_id" validate:"required"`
    SalePrice         float64   `json:"sale_price" validate:"required"`
    TransactionDate   time.Time `json:"transaction_date" validate:"required"`
    TransactionStatus string    `json:"transaction_status" validate:"required,max=50"`
}