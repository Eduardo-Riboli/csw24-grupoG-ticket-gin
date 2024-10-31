package entities

type TicketAuthRequest struct {
    VerificationCode string `json:"verification_code" validate:"required"`
}