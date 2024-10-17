package entities

type TenantResponse struct {
	ID 					   uint   `json:"id"`
	Name                   string `json:"name"`
	ContactInfo            string `json:"contact_info"`
	SpecificConfigurations string `json:"specific_configurations"`
}