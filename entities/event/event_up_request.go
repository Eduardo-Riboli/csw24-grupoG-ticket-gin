package entities

import (
	"time"
)

type EventUpRequest struct {
    TenantID uint      `json:"tenant_id"`
    Name     string    `json:"name" validate:"max=100"`
    Location string    `json:"location" validate:"max=255"`
    Date     time.Time `json:"date"`
}