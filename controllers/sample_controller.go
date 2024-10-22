package controllers

import (
    "github.com/grupoG/csw24-grupoG-ticket-gin/models"
    "github.com/grupoG/csw24-grupoG-ticket-gin/services"
    "github.com/grupoG/csw24-grupoG-ticket-gin/utils"
    "net/http"
    "github.com/gin-gonic/gin"
)

type SampleController struct {
    Service *services.SampleService
}

func NewSampleController(service *services.SampleService) *SampleController {
    return &SampleController{Service: service}
}

// GetAllSamples godoc
// @Summary Get all samples
// @Description Get a list of all samples
// @Tags samples
// @Produce json
// @Success 200 {array} models.Sample
// @Failure 500 {object} utils.ErrorResponse
// @Router /samples [get]
func (ctrl *SampleController) GetAllSamples(c *gin.Context) {
    samples, err := ctrl.Service.GetAllSamples()
    if err != nil {
        utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
        return
    }
    c.JSON(http.StatusOK, samples)
}

// CreateSample godoc
// @Summary Create a new sample
// @Description Create a new sample with the given details
// @Tags samples
// @Accept json
// @Produce json
// @Param sample body models.Sample true "Sample request body"
// @Success 201 {object} models.Sample
// @Failure 400 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /samples [post]
func (ctrl *SampleController) CreateSample(c *gin.Context) {
    var sample models.Sample
    if err := c.ShouldBindJSON(&sample); err != nil {
        utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
        return
    }
    newSample, err := ctrl.Service.CreateSample(sample)
    if err != nil {
        utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
        return
    }
    c.JSON(http.StatusCreated, newSample)
}