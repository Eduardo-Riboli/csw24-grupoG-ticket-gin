package entities

type NotificationPreferencesCrRequest struct {
    UserID        uint `json:"user_id" validate:"required"`
    ReceiveEmails bool `json:"receive_emails" validate:"required"`
}