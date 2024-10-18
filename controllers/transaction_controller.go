package controllers

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    entities "github.com/grupoG/csw24-grupoG-ticket-gin/entities/transaction"
    "github.com/grupoG/csw24-grupoG-ticket-gin/services"
    "github.com/grupoG/csw24-grupoG-ticket-gin/utils"
)

type TransactionController struct {
    Service *services.TransactionService
}

func NewTransactionController(service *services.TransactionService) *TransactionController {
    return &TransactionController{Service: service}
}

func (ctrl *TransactionController) GetAllTransactions(c *gin.Context) {
    transactions, err := ctrl.Service.GetAllTransactions()
    if err != nil {
        utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
        return
    }
    c.JSON(http.StatusOK, transactions)
}

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

func (ctrl *TransactionController) CreateTransaction(c *gin.Context) {
    var transactionRequest entities.TransactionCrRequest
    if err := c.ShouldBindJSON(&transactionRequest); err != nil {
        utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
        return
    }
    newTransaction, err := ctrl.Service.CreateTransaction(transactionRequest)
    if err != nil {
        utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
        return
    }
    c.JSON(http.StatusCreated, newTransaction)
}

func (ctrl *TransactionController) UpdateTransaction(c *gin.Context) {
    idParam := c.Param("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    var transactionRequest entities.TransactionUpRequest
    if err := c.ShouldBindJSON(&transactionRequest); err != nil {
        utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
        return
    }

    updatedTransaction, err := ctrl.Service.UpdateTransaction(uint(id), transactionRequest)
    if err != nil {
        utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
        return
    }
    c.JSON(http.StatusOK, updatedTransaction)
}

func (ctrl *TransactionController) DeleteTransaction(c *gin.Context) {
    idParam := c.Param("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    if err := ctrl.Service.DeleteTransaction(uint(id)); err != nil {
        utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
        return
    }
    c.JSON(http.StatusNoContent, nil)
}