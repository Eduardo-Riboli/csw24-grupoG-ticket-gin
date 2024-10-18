package entities

type NotificationPreferencesResponse struct {
    ID            uint `json:"id"`
    UserID        uint `json:"user_id"`
    ReceiveEmails bool `json:"receive_emails"`
}