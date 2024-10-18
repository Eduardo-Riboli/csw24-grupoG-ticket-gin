package repositories

import (
    "github.com/grupoG/csw24-grupoG-ticket-gin/models"
    "gorm.io/gorm"
)

type EventRepository struct {
    DB *gorm.DB
}

func NewEventRepository(db *gorm.DB) *EventRepository {
    return &EventRepository{DB: db}
}

func (r *EventRepository) GetAll() ([]models.Event, error) {
    var events []models.Event
    if err := r.DB.Find(&events).Error; err != nil {
        return nil, err
    }
    return events, nil
}

func (r *EventRepository) GetByID(id uint) (models.Event, error) {
    var event models.Event
    if err := r.DB.First(&event, id).Error; err != nil {
        return models.Event{}, err
    }
    return event, nil
}

func (r *EventRepository) Create(event models.Event) (models.Event, error) {
    if err := r.DB.Create(&event).Error; err != nil {
        return models.Event{}, err
    }
    return event, nil
}

func (r *EventRepository) Update(event models.Event) (models.Event, error) {
    if err := r.DB.Save(&event).Error; err != nil {
        return models.Event{}, err
    }
    return event, nil
}

func (r *EventRepository) Delete(id uint) error {
    if err := r.DB.Delete(&models.Event{}, id).Error; err != nil {
        return err
    }
    return nil
}