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

func (ctrl *SampleController) GetAllSamples(c *gin.Context) {
    samples, err := ctrl.Service.GetAllSamples()
    if err != nil {
        utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
        return
    }
    c.JSON(http.StatusOK, samples)
}

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