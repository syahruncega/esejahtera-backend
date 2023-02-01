package controller

import (
	"fmt"
	"kemiskinan/helper"
	"kemiskinan/model"
	"kemiskinan/request"
	"kemiskinan/responses"
	"kemiskinan/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type detailProgramController struct {
	detailProgramService service.DetailProgramService
}

func NewDetailProgramController(detailProgramService service.DetailProgramService) *detailProgramController {
	return &detailProgramController{detailProgramService}
}

func (c *detailProgramController) GetDetailPrograms(cntx *gin.Context) {
	var instansiId = cntx.Query("instansiid")

	var detailPrograms []model.DetailProgram
	var err error

	if instansiId != "" {
		detailPrograms, err = c.detailProgramService.FindByInstansiId(instansiId)
	} else {
		detailPrograms, err = c.detailProgramService.FindAll()
	}

	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
	}

	var detailProgramsResponse []responses.DetailProgramResponse

	for _, detailProgram := range detailPrograms {
		var detailProgramResponse = helper.ConvertToDetailProgramResponse(detailProgram)
		detailProgramsResponse = append(detailProgramsResponse, detailProgramResponse)
	}

	cntx.JSON(http.StatusOK, detailProgramsResponse)
}

func (c *detailProgramController) GetDetailProgram(cntx *gin.Context) {
	var idString = cntx.Param("id")
	var id, _ = strconv.Atoi(idString)

	var detailProgram, err = c.detailProgramService.FindById(id)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": "Data tidak ditemukan",
		})
		return
	}

	var detailProgramResponse = helper.ConvertToDetailProgramResponse(detailProgram)

	cntx.JSON(http.StatusOK, detailProgramResponse)
}

func (c *detailProgramController) CreateDetailProgram(cntx *gin.Context) {
	var detailProgramRequest request.CreateDetailProgramRequest

	var err = cntx.ShouldBindJSON(&detailProgramRequest)
	if err != nil {
		var errorMessages = []string{}

		for _, e := range err.(validator.ValidationErrors) {
			var errorMessage = fmt.Sprintf("Error on field %s, condition : %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": errorMessages,
		})
		return
	}

	detailProgram, err := c.detailProgramService.Create(detailProgramRequest)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	var detailProgramResponse = helper.ConvertToDetailProgramResponse(detailProgram)

	cntx.JSON(http.StatusCreated, detailProgramResponse)
}

func (c *detailProgramController) UpdateDetailProgram(cntx *gin.Context) {
	var detailProgramRequest request.UpdateDetailProgramRequest

	var err = cntx.ShouldBindJSON(&detailProgramRequest)
	if err != nil {
		var errorMessages = []string{}

		for _, e := range err.(validator.ValidationErrors) {
			var errorMessage = fmt.Sprintf("Error on field %s, condition : %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": errorMessages,
		})
		return
	}

	var idString = cntx.Param("id")
	var id, _ = strconv.Atoi(idString)

	detailProgram, err := c.detailProgramService.Update(id, detailProgramRequest)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	var detailProgramResponse = helper.ConvertToDetailProgramResponse(detailProgram)

	cntx.JSON(http.StatusOK, detailProgramResponse)
}

func (c *detailProgramController) DeleteDetailProgram(cntx *gin.Context) {
	var idString = cntx.Param("id")
	var id, _ = strconv.Atoi(idString)

	_, err := c.detailProgramService.Delete(id)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": "Data gagal dihapus",
		})
		return
	}

	cntx.JSON(http.StatusOK, gin.H{
		"status": "Data berhasil dihapus",
	})
}
