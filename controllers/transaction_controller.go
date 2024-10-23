package controllers

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    entities "github.com/grupoG/csw24-grupoG-ticket-gin/entities/transaction"
    "github.com/grupoG/csw24-grupoG-ticket-gin/services"
)

type TransactionController struct {
    Service *services.TransactionService
}

func NewTransactionController(service *services.TransactionService) *TransactionController {
    return &TransactionController{Service: service}
}

// GetAllTransactions godoc
// @Summary Get all transactions
// @Description Get a list of all transactions
// @Tags transactions
// @Produce json
// @Success 200 {array} entities.TransactionResponse
// @Failure 500 {object} map[string]string
// @Router /transactions [get]
func (ctrl *TransactionController) GetAllTransactions(c *gin.Context) {
    transactions, err := ctrl.Service.GetAllTransactions()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, transactions)
}

// GetTransactionByID godoc
// @Summary Get transaction by ID
// @Description Get details of a specific transaction by ID
// @Tags transactions
// @Produce json
// @Param id path int true "Transaction ID"
// @Success 200 {object} entities.TransactionResponse
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /transactions/{id} [get]
func (ctrl *TransactionController) GetTransactionByID(c *gin.Context) {
    idParam := c.Param("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    transaction, err := ctrl.Service.GetTransactionByID(uint(id))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Transaction not found"})
        return
    }

    c.JSON(http.StatusOK, transaction)
}

// CreateTransaction godoc
// @Summary Create a new transaction
// @Description Create a new transaction with the given details
// @Tags transactions
// @Accept json
// @Produce json
// @Param transaction body entities.TransactionCrRequest true "Transaction request body"
// @Success 201 {object} entities.TransactionResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /transactions [post]
func (ctrl *TransactionController) CreateTransaction(c *gin.Context) {
    var transactionRequest entities.TransactionCrRequest
    if err := c.ShouldBindJSON(&transactionRequest); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    newTransaction, err := ctrl.Service.CreateTransaction(transactionRequest)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, newTransaction)
}

// UpdateTransaction godoc
// @Summary Update a transaction
// @Description Update details of an existing transaction by ID
// @Tags transactions
// @Accept json
// @Produce json
// @Param id path int true "Transaction ID"
// @Param transaction body entities.TransactionUpRequest true "Transaction request body"
// @Success 200 {object} entities.TransactionResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /transactions/{id} [put]
func (ctrl *TransactionController) UpdateTransaction(c *gin.Context) {
    idParam := c.Param("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    var transactionRequest entities.TransactionUpRequest
    if err := c.ShouldBindJSON(&transactionRequest); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    updatedTransaction, err := ctrl.Service.UpdateTransaction(uint(id), transactionRequest)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, updatedTransaction)
}

// DeleteTransaction godoc
// @Summary Delete a transaction
// @Description Delete a transaction by ID
// @Tags transactions
// @Produce json
// @Param id path int true "Transaction ID"
// @Success 204 {object} nil
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /transactions/{id} [delete]
func (ctrl *TransactionController) DeleteTransaction(c *gin.Context) {
    idParam := c.Param("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    if err := ctrl.Service.DeleteTransaction(uint(id)); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusNoContent, nil)
}