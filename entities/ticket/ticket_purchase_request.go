package entities

type TicketPurchaseRequest struct {
    EventID          uint    `json:"event_id" validate:"required"`
    TenantID         uint    `json:"tenant_id" validate:"required"`
    OriginalPrice    float64 `json:"original_price" validate:"required"`
    SellerID         uint    `json:"seller_id" validate:"required"`
    VerificationCode string  `json:"verification_code" validate:"required,max=100"`
}
