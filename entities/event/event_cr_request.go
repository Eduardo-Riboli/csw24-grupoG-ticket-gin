package entities

import (
	"time"
)

type EventCrRequest struct {
    TenantID uint      `json:"tenant_id" validate:"required"`
    Name     string    `json:"name" validate:"required,max=100"`
    Location string    `json:"location" validate:"required,max=255"`
    Date     time.Time `json:"date" validate:"required"`
}