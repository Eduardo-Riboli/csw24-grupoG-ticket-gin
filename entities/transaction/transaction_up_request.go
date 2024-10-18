package entities

import "time"

type TransactionUpRequest struct {
    SalePrice         float64   `json:"sale_price"`
    TransactionDate   time.Time `json:"transaction_date"`
    TransactionStatus string    `json:"transaction_status" validate:"max=50"`
}