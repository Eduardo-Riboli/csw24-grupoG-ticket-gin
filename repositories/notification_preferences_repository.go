package repositories

import (
    "github.com/grupoG/csw24-grupoG-ticket-gin/models"
    "gorm.io/gorm"
)

type NotificationPreferencesRepository struct {
    DB *gorm.DB
}

func NewNotificationPreferencesRepository(db *gorm.DB) *NotificationPreferencesRepository {
    return &NotificationPreferencesRepository{DB: db}
}

func (r *NotificationPreferencesRepository) GetAll() ([]models.NotificationPreferences, error) {
    var preferences []models.NotificationPreferences
    if err := r.DB.Find(&preferences).Error; err != nil {
        return nil, err
    }
    return preferences, nil
}

func (r *NotificationPreferencesRepository) GetByID(id uint) (models.NotificationPreferences, error) {
    var preference models.NotificationPreferences
    if err := r.DB.First(&preference, id).Error; err != nil {
        return models.NotificationPreferences{}, err
    }
    return preference, nil
}

func (r *NotificationPreferencesRepository) Create(preference models.NotificationPreferences) (models.NotificationPreferences, error) {
    if err := r.DB.Create(&preference).Error; err != nil {
        return models.NotificationPreferences{}, err
    }
    return preference, nil
}

func (r *NotificationPreferencesRepository) Update(preference models.NotificationPreferences) (models.NotificationPreferences, error) {
    if err := r.DB.Save(&preference).Error; err != nil {
        return models.NotificationPreferences{}, err
    }
    return preference, nil
}

func (r *NotificationPreferencesRepository) Delete(id uint) error {
    if err := r.DB.Delete(&models.NotificationPreferences{}, id).Error; err != nil {
        return err
    }
    return nil
}