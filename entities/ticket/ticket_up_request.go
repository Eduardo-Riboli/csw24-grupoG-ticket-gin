package entities

type TicketUpRequest struct {
    OriginalPrice    float64 `json:"original_price"`
    VerificationCode string  `json:"verification_code" validate:"max=100"`
    Status           string  `json:"status" validate:"max=50"`
}