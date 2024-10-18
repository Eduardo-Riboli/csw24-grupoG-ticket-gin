package services

import (
    entities "github.com/grupoG/csw24-grupoG-ticket-gin/entities/transaction"
    "github.com/grupoG/csw24-grupoG-ticket-gin/models"
    "github.com/grupoG/csw24-grupoG-ticket-gin/repositories"
)

type TransactionService struct {
    Repository *repositories.TransactionRepository
}

func NewTransactionService(repo *repositories.TransactionRepository) *TransactionService {
    return &TransactionService{Repository: repo}
}

func (transactionService *TransactionService) GetAllTransactions() ([]entities.TransactionResponse, error) {
    transactions, err := transactionService.Repository.GetAll()
    if err != nil {
        return nil, err
    }

    var transactionResponses []entities.TransactionResponse
    for _, transaction := range transactions {
        transactionResponses = append(transactionResponses, entities.TransactionResponse{
            ID:                transaction.ID,
            TenantID:          transaction.TenantID,
            BuyerID:           transaction.BuyerID,
            TicketID:          transaction.TicketID,
            SalePrice:         transaction.SalePrice,
            TransactionDate:   transaction.TransactionDate,
            TransactionStatus: transaction.TransactionStatus,
        })
    }

    return transactionResponses, nil
}

func (transactionService *TransactionService) GetTransactionByID(id uint) (entities.TransactionResponse, error) {
    transaction, err := transactionService.Repository.GetByID(id)
    if err != nil {
        return entities.TransactionResponse{}, err
    }

    return entities.TransactionResponse{
        ID:                transaction.ID,
        TenantID:          transaction.TenantID,
        BuyerID:           transaction.BuyerID,
        TicketID:          transaction.TicketID,
        SalePrice:         transaction.SalePrice,
        TransactionDate:   transaction.TransactionDate,
        TransactionStatus: transaction.TransactionStatus,
    }, nil
}

func (transactionService *TransactionService) CreateTransaction(transactionRequest entities.TransactionCrRequest) (entities.TransactionResponse, error) {
    transaction := models.Transaction{
        TenantID:          transactionRequest.TenantID,
        BuyerID:           transactionRequest.BuyerID,
        TicketID:          transactionRequest.TicketID,
        SalePrice:         transactionRequest.SalePrice,
        TransactionDate:   transactionRequest.TransactionDate,
        TransactionStatus: transactionRequest.TransactionStatus,
    }

    createdTransaction, err := transactionService.Repository.Create(transaction)
    if err != nil {
        return entities.TransactionResponse{}, err
    }

    return entities.TransactionResponse{
        ID:                createdTransaction.ID,
        TenantID:          createdTransaction.TenantID,
        BuyerID:           createdTransaction.BuyerID,
        TicketID:          createdTransaction.TicketID,
        SalePrice:         createdTransaction.SalePrice,
        TransactionDate:   createdTransaction.TransactionDate,
        TransactionStatus: createdTransaction.TransactionStatus,
    }, nil
}

func (transactionService *TransactionService) UpdateTransaction(id uint, transactionRequest entities.TransactionUpRequest) (entities.TransactionResponse, error) {
    transaction, err := transactionService.Repository.GetByID(id)
    if err != nil {
        return entities.TransactionResponse{}, err
    }

    if transactionRequest.SalePrice != 0 {
        transaction.SalePrice = transactionRequest.SalePrice
    }
    if !transactionRequest.TransactionDate.IsZero() {
        transaction.TransactionDate = transactionRequest.TransactionDate
    }
    if transactionRequest.TransactionStatus != "" {
        transaction.TransactionStatus = transactionRequest.TransactionStatus
    }

    updatedTransaction, err := transactionService.Repository.Update(transaction)
    if err != nil {
        return entities.TransactionResponse{}, err
    }

    return entities.TransactionResponse{
        ID:                updatedTransaction.ID,
        TenantID:          updatedTransaction.TenantID,
        BuyerID:           updatedTransaction.BuyerID,
        TicketID:          updatedTransaction.TicketID,
        SalePrice:         updatedTransaction.SalePrice,
        TransactionDate:   updatedTransaction.TransactionDate,
        TransactionStatus: updatedTransaction.TransactionStatus,
    }, nil
}

func (transactionService *TransactionService) DeleteTransaction(id uint) error {
    return transactionService.Repository.Delete(id)
}