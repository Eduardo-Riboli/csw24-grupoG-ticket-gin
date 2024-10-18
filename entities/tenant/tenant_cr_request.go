package entities

type TenantCrRequest struct {
	Name                   string `json:"name" validate:"required,max=100"`
	ContactInfo            string `json:"contact_info" validate:"max=255"`
	SpecificConfigurations string `json:"specific_configurations" validate:"max=255"`
}
