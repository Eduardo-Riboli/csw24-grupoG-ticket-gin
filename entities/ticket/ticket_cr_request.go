package entities

type TicketCrRequest struct {
    EventID          uint    `json:"event_id" validate:"required"`
    TenantID         uint    `json:"tenant_id" validate:"required"`
    OriginalPrice    float64 `json:"original_price" validate:"required"`
    SellerID         uint    `json:"seller_id" validate:"required"`
    VerificationCode string  `json:"verification_code" validate:"required,max=100"`
    Status           string  `json:"status" validate:"required,max=50"`
}