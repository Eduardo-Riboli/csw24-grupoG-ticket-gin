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

func (repo *TicketRepository) GetByVerificationCode(code string) (models.Ticket, error) {
    var ticket models.Ticket
    if err := repo.DB.Where("verification_code = ?", code).First(&ticket).Error; err != nil {
        return ticket, err
    }
    return ticket, nil
}

func (r *TicketRepository) GetTransactionByTicketID(ticketID uint) (models.Transaction, error) {
    var transaction models.Transaction
    if err := r.DB.Where("ticket_id = ?", ticketID).First(&transaction).Error; err != nil {
        return transaction, err
    }
    return transaction, nil
}

func (r *TicketRepository) UpdateTransaction(transaction models.Transaction) (models.Transaction, error) {
    if err := r.DB.Save(&transaction).Error; err != nil {
        return transaction, err
    }
    return transaction, nil
}
