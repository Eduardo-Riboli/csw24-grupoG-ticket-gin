package repositories

import (
    "github.com/grupoG/csw24-grupoG-ticket-gin/models"
    "gorm.io/gorm"
)

type TransactionRepository struct {
    DB *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *TransactionRepository {
    return &TransactionRepository{DB: db}
}

func (r *TransactionRepository) GetAll() ([]models.Transaction, error) {
    var transactions []models.Transaction
    if err := r.DB.Find(&transactions).Error; err != nil {
        return nil, err
    }
    return transactions, nil
}

func (r *TransactionRepository) GetByID(id uint) (models.Transaction, error) {
    var transaction models.Transaction
    if err := r.DB.First(&transaction, id).Error; err != nil {
        return models.Transaction{}, err
    }
    return transaction, nil
}

func (r *TransactionRepository) Create(transaction models.Transaction) (models.Transaction, error) {
    if err := r.DB.Create(&transaction).Error; err != nil {
        return models.Transaction{}, err
    }
    return transaction, nil
}

func (r *TransactionRepository) Update(transaction models.Transaction) (models.Transaction, error) {
    if err := r.DB.Save(&transaction).Error; err != nil {
        return models.Transaction{}, err
    }
    return transaction, nil
}

func (r *TransactionRepository) Delete(id uint) error {
    if err := r.DB.Delete(&models.Transaction{}, id).Error; err != nil {
        return err
    }
    return nil
}