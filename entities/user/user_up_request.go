package entities

type UserUpRequest struct {
    Name  string `json:"name" validate:"max=100"`
    Email string `json:"email" validate:"email,max=100"`
}