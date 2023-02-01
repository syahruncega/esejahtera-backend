package controller

import (
	"fmt"
	"kemiskinan/helper"
	"kemiskinan/request"
	"kemiskinan/responses"
	"kemiskinan/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type indikatorProgramController struct {
	indikatorProgramService service.IndikatorProgramService
}

func NewIndikatorProgramController(indikatorProgramService service.IndikatorProgramService) *indikatorProgramController {
	return &indikatorProgramController{indikatorProgramService}
}

func (c *indikatorProgramController) GetIndikatorPrograms(cntx *gin.Context) {
	var indikatorPrograms, err = c.indikatorProgramService.FindAll()
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
	}

	var indikatorProgramsResponse []responses.IndikatorProgramResponse

	for _, indikatorProgram := range indikatorPrograms {
		var indikatorProgramResponse = helper.ConvertToIndikatorProgramResponse(indikatorProgram)
		indikatorProgramsResponse = append(indikatorProgramsResponse, indikatorProgramResponse)
	}

	cntx.JSON(http.StatusOK, indikatorProgramsResponse)
}

func (c *indikatorProgramController) GetIndikatorProgram(cntx *gin.Context) {
	var idString = cntx.Param("id")
	var id, _ = strconv.Atoi(idString)

	var indikatorProgram, err = c.indikatorProgramService.FindById(id)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": "Data tidak ditemukan",
		})
		return
	}

	var indikatorProgramResponse = helper.ConvertToIndikatorProgramResponse(indikatorProgram)

	cntx.JSON(http.StatusOK, indikatorProgramResponse)
}

func (c *indikatorProgramController) CreateIndikatorProgram(cntx *gin.Context) {
	var indikatorProgramRequest request.CreateIndikatorProgramRequest

	var err = cntx.ShouldBindJSON(&indikatorProgramRequest)
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

	indikatorProgram, err := c.indikatorProgramService.Create(indikatorProgramRequest)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	var indikatorProgramResponse = helper.ConvertToIndikatorProgramResponse(indikatorProgram)

	cntx.JSON(http.StatusCreated, gin.H{
		"data": indikatorProgramResponse,
	})
}

func (c *indikatorProgramController) UpdateIndikatorProgram(cntx *gin.Context) {
	var indikatorProgramRequest request.UpdateIndikatorProgramRequest

	var err = cntx.ShouldBindJSON(&indikatorProgramRequest)
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

	indikatorProgram, err := c.indikatorProgramService.Update(id, indikatorProgramRequest)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"errors": cntx.Error(err),
		})
		return
	}

	var indikatorProgramResponse = helper.ConvertToIndikatorProgramResponse(indikatorProgram)

	cntx.JSON(http.StatusOK, indikatorProgramResponse)
}

func (c *indikatorProgramController) DeleteIndikatorProgram(cntx *gin.Context) {
	var idString = cntx.Param("id")
	var id, _ = strconv.Atoi(idString)

	_, err := c.indikatorProgramService.Delete(id)
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
