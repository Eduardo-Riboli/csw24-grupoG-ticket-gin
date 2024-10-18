package repositories

import (
    "github.com/grupoG/csw24-grupoG-ticket-gin/models"
    "gorm.io/gorm"
)

type TicketRepository struct {
    DB *gorm.DB
}

func NewTicketRepository(db *gorm.DB) *TicketRepository {
    return &TicketRepository{DB: db}
}

func (r *TicketRepository) GetAll() ([]models.Ticket, error) {
    var tickets []models.Ticket
    if err := r.DB.Find(&tickets).Error; err != nil {
        return nil, err
    }
    return tickets, nil
}

func (r *TicketRepository) GetByID(id uint) (models.Ticket, error) {
    var ticket models.Ticket
    if err := r.DB.First(&ticket, id).Error; err != nil {
        return models.Ticket{}, err
    }
    return ticket, nil
}

func (r *TicketRepository) Create(ticket models.Ticket) (models.Ticket, error) {
    if err := r.DB.Create(&ticket).Error; err != nil {
        return models.Ticket{}, err
    }
    return ticket, nil
}

func (r *TicketRepository) Update(ticket models.Ticket) (models.Ticket, error) {
    if err := r.DB.Save(&ticket).Error; err != nil {
        return models.Ticket{}, err
    }
    return ticket, nil
}

func (r *TicketRepository) Delete(id uint) error {
    if err := r.DB.Delete(&models.Ticket{}, id).Error; err != nil {
        return err
    }
    return nil
}
