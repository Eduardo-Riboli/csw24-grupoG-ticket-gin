package repositories

import (
	"github.com/grupoG/csw24-grupoG-ticket-gin/models"
	"gorm.io/gorm"
)

type TenantRepository struct {
	DB *gorm.DB
}

func NewTenantRepository(db *gorm.DB) *TenantRepository {
	return &TenantRepository{DB: db}
}

func (r *TenantRepository) GetAll() ([]models.Tenant, error) {
	var tenants []models.Tenant
	if err := r.DB.Find(&tenants).Error; err != nil {
		return nil, err
	}
	return tenants, nil
}

func (r *TenantRepository) GetByID(id uint) (models.Tenant, error) {
	var tenant models.Tenant
	if err := r.DB.First(&tenant, id).Error; err != nil {
		return models.Tenant{}, err
	}
	return tenant, nil
}

func (r *TenantRepository) Create(tenant models.Tenant) (models.Tenant, error) {
	if err := r.DB.Create(&tenant).Error; err != nil {
		return models.Tenant{}, err
	}
	return tenant, nil
}
