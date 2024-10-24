package controllers

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    entities "github.com/grupoG/csw24-grupoG-ticket-gin/entities/notification_preferences"
    "github.com/grupoG/csw24-grupoG-ticket-gin/services"
)

type NotificationPreferencesController struct {
    Service *services.NotificationPreferencesService
}

func NewNotificationPreferencesController(service *services.NotificationPreferencesService) *NotificationPreferencesController {
    return &NotificationPreferencesController{Service: service}
}

// GetAllPreferences godoc
// @Summary Get all notification preferences
// @Description Get a list of all notification preferences
// @Tags notification_preferences
// @Produce json
// @Success 200 {array} entities.NotificationPreferencesResponse
// @Failure 500 {object} map[string]string
// @Router /preferences [get]
func (ctrl *NotificationPreferencesController) GetAllPreferences(c *gin.Context) {
    preferences, err := ctrl.Service.GetAllPreferences()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, preferences)
}

// GetPreferenceByID godoc
// @Summary Get notification preference by ID
// @Description Get details of a specific notification preference by ID
// @Tags notification_preferences
// @Produce json
// @Param id path int true "Preference ID"
// @Success 200 {object} entities.NotificationPreferencesResponse
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /preferences/{id} [get]
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

// CreatePreference godoc
// @Summary Create a new notification preference
// @Description Create a new notification preference with the given details
// @Tags notification_preferences
// @Accept json
// @Produce json
// @Param preference body entities.NotificationPreferencesCrRequest true "Preference request body"
// @Success 201 {object} entities.NotificationPreferencesResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /preferences [post]
func (ctrl *NotificationPreferencesController) CreatePreference(c *gin.Context) {
    var request entities.NotificationPreferencesCrRequest
    if err := c.ShouldBindJSON(&request); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    newPreference, err := ctrl.Service.CreatePreference(request)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, newPreference)
}

// UpdatePreference godoc
// @Summary Update a notification preference
// @Description Update details of an existing notification preference by ID
// @Tags notification_preferences
// @Accept json
// @Produce json
// @Param id path int true "Preference ID"
// @Param preference body entities.NotificationPreferencesUpRequest true "Preference request body"
// @Success 200 {object} entities.NotificationPreferencesResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /preferences/{id} [put]
func (ctrl *NotificationPreferencesController) UpdatePreference(c *gin.Context) {
    idParam := c.Param("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    var request entities.NotificationPreferencesUpRequest
    if err := c.ShouldBindJSON(&request); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    updatedPreference, err := ctrl.Service.UpdatePreference(uint(id), request)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, updatedPreference)
}

// DeletePreference godoc
// @Summary Delete a notification preference
// @Description Delete a notification preference by ID
// @Tags notification_preferences
// @Produce json
// @Param id path int true "Preference ID"
// @Success 204 {object} nil
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /preferences/{id} [delete]
func (ctrl *NotificationPreferencesController) DeletePreference(c *gin.Context) {
    idParam := c.Param("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    if err := ctrl.Service.DeletePreference(uint(id)); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusNoContent, nil)
}