package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	entities "github.com/grupoG/csw24-grupoG-ticket-gin/entities/tenant"
	"github.com/grupoG/csw24-grupoG-ticket-gin/services"
	"github.com/grupoG/csw24-grupoG-ticket-gin/utils"
)

type TenantController struct {
	Service *services.TenantService
}

func NewTenantController(service *services.TenantService) *TenantController {
	return &TenantController{Service: service}
}

func (ctrl *TenantController) GetAllTenants(c *gin.Context) {
	tenants, err := ctrl.Service.GetAllTenants()
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, tenants)
}

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

func (ctrl *TenantController) CreateTenant(c *gin.Context) {
	var tenantRequest entities.TenantRequest
	if err := c.ShouldBindJSON(&tenantRequest); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	newTenant, err := ctrl.Service.CreateTenant(tenantRequest)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusCreated, newTenant)
}
