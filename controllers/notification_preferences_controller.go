package controllers

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    entities "github.com/grupoG/csw24-grupoG-ticket-gin/entities/notification_preferences"
    "github.com/grupoG/csw24-grupoG-ticket-gin/services"
    "github.com/grupoG/csw24-grupoG-ticket-gin/utils"
)

type NotificationPreferencesController struct {
    Service *services.NotificationPreferencesService
}

func NewNotificationPreferencesController(service *services.NotificationPreferencesService) *NotificationPreferencesController {
    return &NotificationPreferencesController{Service: service}
}

func (ctrl *NotificationPreferencesController) GetAllPreferences(c *gin.Context) {
    preferences, err := ctrl.Service.GetAllPreferences()
    if err != nil {
        utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
        return
    }
    c.JSON(http.StatusOK, preferences)
}

func (ctrl *NotificationPreferencesController) GetPreferenceByID(c *gin.Context) {
    idParam := c.Param("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    preference, err := ctrl.Service.GetPreferenceByID(uint(id))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Preference not found"})
        return
    }

    c.JSON(http.StatusOK, preference)
}

func (ctrl *NotificationPreferencesController) CreatePreference(c *gin.Context) {
    var request entities.NotificationPreferencesCrRequest
    if err := c.ShouldBindJSON(&request); err != nil {
        utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
        return
    }
    newPreference, err := ctrl.Service.CreatePreference(request)
    if err != nil {
        utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
        return
    }
    c.JSON(http.StatusCreated, newPreference)
}

func (ctrl *NotificationPreferencesController) UpdatePreference(c *gin.Context) {
    idParam := c.Param("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    var request entities.NotificationPreferencesUpRequest
    if err := c.ShouldBindJSON(&request); err != nil {
        utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
        return
    }

    updatedPreference, err := ctrl.Service.UpdatePreference(uint(id), request)
    if err != nil {
        utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
        return
    }
    c.JSON(http.StatusOK, updatedPreference)
}

func (ctrl *NotificationPreferencesController) DeletePreference(c *gin.Context) {
    idParam := c.Param("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    if err := ctrl.Service.DeletePreference(uint(id)); err != nil {
        utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
        return
    }
    c.JSON(http.StatusNoContent, nil)
}