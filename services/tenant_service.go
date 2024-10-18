package services

import (
    entities "github.com/grupoG/csw24-grupoG-ticket-gin/entities/tenant"
    "github.com/grupoG/csw24-grupoG-ticket-gin/models"
    "github.com/grupoG/csw24-grupoG-ticket-gin/repositories"
)

type TenantService struct {
    Repository *repositories.TenantRepository
}

func NewTenantService(repo *repositories.TenantRepository) *TenantService {
    return &TenantService{Repository: repo}
}

func (tenantService *TenantService) GetAllTenants() ([]entities.TenantResponse, error) {
    tenants, err := tenantService.Repository.GetAll()
    if err != nil {
        return nil, err
    }

    var tenantResponses []entities.TenantResponse
    for _, tenant := range tenants {
        tenantResponses = append(tenantResponses, entities.TenantResponse{
            ID:                     tenant.ID,
            Name:                   tenant.Name,
            ContactInfo:            tenant.ContactInfo,
            SpecificConfigurations: tenant.SpecificConfigurations,
        })
    }

    return tenantResponses, nil
}

func (tenantService *TenantService) GetTenantByID(id uint) (entities.TenantResponse, error) {
    tenant, err := tenantService.Repository.GetByID(id)
    if err != nil {
        return entities.TenantResponse{}, err
    }

    return entities.TenantResponse{
        ID:                     tenant.ID,
        Name:                   tenant.Name,
        ContactInfo:            tenant.ContactInfo,
        SpecificConfigurations: tenant.SpecificConfigurations,
    }, nil
}

func (tenantService *TenantService) CreateTenant(tenantRequest entities.TenantCrRequest) (entities.TenantResponse, error) {
    tenant := models.Tenant{
        Name:                   tenantRequest.Name,
        ContactInfo:            tenantRequest.ContactInfo,
        SpecificConfigurations: tenantRequest.SpecificConfigurations,
    }

    createdTenant, err := tenantService.Repository.Create(tenant)
    if err != nil {
        return entities.TenantResponse{}, err
    }

    return entities.TenantResponse{
        ID:                     createdTenant.ID,
        Name:                   createdTenant.Name,
        ContactInfo:            createdTenant.ContactInfo,
        SpecificConfigurations: createdTenant.SpecificConfigurations,
    }, nil
}

func (tenantService *TenantService) UpdateTenant(id uint, tenantRequest entities.TenantUpRequest) (entities.TenantResponse, error) {
    tenant, err := tenantService.Repository.GetByID(id)

    if err != nil {
        return entities.TenantResponse{}, err
    }

    if tenantRequest.Name != "" {
        tenant.Name = tenantRequest.Name
    }
    if tenantRequest.ContactInfo != "" {
        tenant.ContactInfo = tenantRequest.ContactInfo
    }
    if tenantRequest.SpecificConfigurations != "" {
        tenant.SpecificConfigurations = tenantRequest.SpecificConfigurations
    }

    updatedTenant, err := tenantService.Repository.Update(tenant)
    if err != nil {
        return entities.TenantResponse{}, err
    }

    return entities.TenantResponse{
        ID:                     updatedTenant.ID,
        Name:                   updatedTenant.Name,
        ContactInfo:            updatedTenant.ContactInfo,
        SpecificConfigurations: updatedTenant.SpecificConfigurations,
    }, nil
}

func (tenantService *TenantService) DeleteTenant(id uint) error {
    return tenantService.Repository.Delete(id)
}