package controllers

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    entities "github.com/grupoG/csw24-grupoG-ticket-gin/entities/tenant"
    "github.com/grupoG/csw24-grupoG-ticket-gin/services"
)

type TenantController struct {
    Service *services.TenantService
}

func NewTenantController(service *services.TenantService) *TenantController {
    return &TenantController{Service: service}
}

// GetAllTenants godoc
// @Summary Get all tenants
// @Description Get a list of all tenants
// @Tags tenants
// @Produce json
// @Success 200 {array} entities.TenantResponse
// @Failure 500 {object} map[string]string
// @Router /tenants [get]
func (ctrl *TenantController) GetAllTenants(c *gin.Context) {
    tenants, err := ctrl.Service.GetAllTenants()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, tenants)
}

// GetTenantByID godoc
// @Summary Get tenant by ID
// @Description Get details of a specific tenant by ID
// @Tags tenants
// @Produce json
// @Param id path int true "Tenant ID"
// @Success 200 {object} entities.TenantResponse
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /tenants/{id} [get]
func (controller *TenantController) GetTenantByID(c *gin.Context) {
    idParam := c.Param("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    tenant, err := controller.Service.GetTenantByID(uint(id))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Tenant not found"})
        return
    }

    c.JSON(http.StatusOK, tenant)
}

// CreateTenant godoc
// @Summary Create a new tenant
// @Description Create a new tenant with the given details
// @Tags tenants
// @Accept json
// @Produce json
// @Param tenant body entities.TenantCrRequest true "Tenant request body"
// @Success 201 {object} entities.TenantResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /tenants [post]
func (ctrl *TenantController) CreateTenant(c *gin.Context) {
    var tenantRequest entities.TenantCrRequest
    if err := c.ShouldBindJSON(&tenantRequest); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    newTenant, err := ctrl.Service.CreateTenant(tenantRequest)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, newTenant)
}

// UpdateTenant godoc
// @Summary Update a tenant
// @Description Update details of an existing tenant by ID
// @Tags tenants
// @Accept json
// @Produce json
// @Param id path int true "Tenant ID"
// @Param tenant body entities.TenantUpRequest true "Tenant request body"
// @Success 200 {object} entities.TenantResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /tenants/{id} [put]
func (ctrl *TenantController) UpdateTenant(c *gin.Context) {
    idParam := c.Param("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    var tenantRequest entities.TenantUpRequest
    if err := c.ShouldBindJSON(&tenantRequest); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    updatedTenant, err := ctrl.Service.UpdateTenant(uint(id), tenantRequest)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, updatedTenant)
}

// DeleteTenant godoc
// @Summary Delete a tenant
// @Description Delete a tenant by ID
// @Tags tenants
// @Produce json
// @Param id path int true "Tenant ID"
// @Success 204 {object} nil
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /tenants/{id} [delete]
func (ctrl *TenantController) DeleteTenant(c *gin.Context) {
    idParam := c.Param("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    if err := ctrl.Service.DeleteTenant(uint(id)); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusNoContent, nil)
}