package entities

type UserCrRequest struct {
    TenantID uint   `json:"tenant_id" validate:"required"`
    Name     string `json:"name" validate:"required,max=100"`
    Email    string `json:"email" validate:"required,email,max=100"`
}