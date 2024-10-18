package controllers

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    entities "github.com/grupoG/csw24-grupoG-ticket-gin/entities/user"
    "github.com/grupoG/csw24-grupoG-ticket-gin/services"
    "github.com/grupoG/csw24-grupoG-ticket-gin/utils"
)

type UserController struct {
    Service *services.UserService
}

func NewUserController(service *services.UserService) *UserController {
    return &UserController{Service: service}
}

func (ctrl *UserController) GetAllUsers(c *gin.Context) {
    users, err := ctrl.Service.GetAllUsers()
    if err != nil {
        utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
        return
    }
    c.JSON(http.StatusOK, users)
}

func (ctrl *UserController) GetUserByID(c *gin.Context) {
    idParam := c.Param("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    user, err := ctrl.Service.GetUserByID(uint(id))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }

    c.JSON(http.StatusOK, user)
}

func (ctrl *UserController) CreateUser(c *gin.Context) {
    var userRequest entities.UserCrRequest
    if err := c.ShouldBindJSON(&userRequest); err != nil {
        utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
        return
    }
    newUser, err := ctrl.Service.CreateUser(userRequest)
    if err != nil {
        utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
        return
    }
    c.JSON(http.StatusCreated, newUser)
}

func (ctrl *UserController) UpdateUser(c *gin.Context) {
    idParam := c.Param("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    var userRequest entities.UserUpRequest
    if err := c.ShouldBindJSON(&userRequest); err != nil {
        utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
        return
    }

    updatedUser, err := ctrl.Service.UpdateUser(uint(id), userRequest)
    if err != nil {
        utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
        return
    }
    c.JSON(http.StatusOK, updatedUser)
}

func (ctrl *UserController) DeleteUser(c *gin.Context) {
    idParam := c.Param("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    if err := ctrl.Service.DeleteUser(uint(id)); err != nil {
        utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
        return
    }
    c.JSON(http.StatusNoContent, nil)
}