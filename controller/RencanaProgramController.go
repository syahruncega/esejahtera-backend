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

type rencanaProgramController struct {
	rencanaProgramService service.RencanaProgramService
}

func NewRencanaProgramController(rencanaProgramService service.RencanaProgramService) *rencanaProgramController {
	return &rencanaProgramController{rencanaProgramService}
}

func (c *rencanaProgramController) GetRencanaPrograms(cntx *gin.Context) {
	var rencanaPrograms, err = c.rencanaProgramService.FindAll()

	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
	}

	var rencanaProgramsResponse []responses.RencanaProgramResponse

	for _, rencanaProgram := range rencanaPrograms {
		var rencanaProgramResponse = helper.ConvertToRencanaProgramResponse(rencanaProgram)
		rencanaProgramsResponse = append(rencanaProgramsResponse, rencanaProgramResponse)
	}

	cntx.JSON(http.StatusOK, rencanaProgramsResponse)
}

func (c *rencanaProgramController) GetRencanaProgram(cntx *gin.Context) {
	var idString = cntx.Param("id")
	var id, _ = strconv.Atoi(idString)

	var rencanaProgram, err = c.rencanaProgramService.FindById(id)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": "Data tidak ditemukan",
		})
		return
	}

	var rencanaProgramResponse = helper.ConvertToRencanaProgramResponse(rencanaProgram)

	cntx.JSON(http.StatusOK, rencanaProgramResponse)
}

func (c *rencanaProgramController) CreateRencanaProgram(cntx *gin.Context) {
	var rencanaProgramRequest request.CreateRencanaProgramRequest

	var err = cntx.ShouldBindJSON(&rencanaProgramRequest)
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

	rencanaProgram, err := c.rencanaProgramService.Create(rencanaProgramRequest)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	var rencanaProgramResponse = helper.ConvertToRencanaProgramResponse(rencanaProgram)

	cntx.JSON(http.StatusCreated, rencanaProgramResponse)
}

func (c *rencanaProgramController) UpdateRencanaProgram(cntx *gin.Context) {
	var rencanaProgramRequest request.UpdateRencanaProgramRequest

	var err = cntx.ShouldBindJSON(&rencanaProgramRequest)
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

	rencanaProgram, err := c.rencanaProgramService.Update(id, rencanaProgramRequest)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	var rencanaProgramResponse = helper.ConvertToRencanaProgramResponse(rencanaProgram)

	cntx.JSON(http.StatusOK, rencanaProgramResponse)
}

func (c *rencanaProgramController) DeleteRencanaProgram(cntx *gin.Context) {
	var idString = cntx.Param("id")
	var id, _ = strconv.Atoi(idString)

	_, err := c.rencanaProgramService.Delete(id)
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
