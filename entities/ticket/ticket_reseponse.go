package entities

type TicketResponse struct {
    ID               uint    `json:"id"`
    EventID          uint    `json:"event_id"`
    TenantID         uint    `json:"tenant_id"`
    OriginalPrice    float64 `json:"original_price"`
    SellerID         uint    `json:"seller_id"`
    VerificationCode string  `json:"verification_code"`
    Status           string  `json:"status"`
}