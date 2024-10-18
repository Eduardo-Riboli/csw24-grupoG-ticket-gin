package entities

import (
	"time"
)

type EventResponse struct {
    ID       uint      `json:"id"`
    TenantID uint      `json:"tenant_id"`
    Name     string    `json:"name"`
    Location string    `json:"location"`
    Date     time.Time `json:"date"`
}