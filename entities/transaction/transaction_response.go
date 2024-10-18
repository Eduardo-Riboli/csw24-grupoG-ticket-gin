package entities

import "time"

type TransactionResponse struct {
    ID                uint      `json:"id"`
    TenantID          uint      `json:"tenant_id"`
    BuyerID           uint      `json:"buyer_id"`
    TicketID          uint      `json:"ticket_id"`
    SalePrice         float64   `json:"sale_price"`
    TransactionDate   time.Time `json:"transaction_date"`
    TransactionStatus string    `json:"transaction_status"`
}