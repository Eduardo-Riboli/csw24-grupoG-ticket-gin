package entities

type TenantUpRequest struct {
    Name                   string `json:"name" validate:"max=100"`
    ContactInfo            string `json:"contact_info" validate:"max=255"`
    SpecificConfigurations string `json:"specific_configurations" validate:"max=255"`
}