package entities

type UserResponse struct {
    ID       uint   `json:"id"`
    TenantID uint   `json:"tenant_id"`
    Name     string `json:"name"`
    Email    string `json:"email"`
}