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

type instansiOnProgramController struct {
	instansiOnProgramService service.InstansiOnProgramService
}

func NewInstansiOnProgramController(instansiOnProgramService service.InstansiOnProgramService) *instansiOnProgramController {
	return &instansiOnProgramController{instansiOnProgramService}
}

func (c *instansiOnProgramController) GetInstansiOnPrograms(cntx *gin.Context) {
	var instansiOnPrograms, err = c.instansiOnProgramService.FindAll()
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	var instansiOnProgramsResponse []responses.InstansiOnProgramResponse

	for _, instansiOnProgram := range instansiOnPrograms {
		var instansiOnProgramResponse = helper.ConvertToInstansiOnProgramResponse(instansiOnProgram)
		instansiOnProgramsResponse = append(instansiOnProgramsResponse, instansiOnProgramResponse)
	}

	cntx.JSON(http.StatusOK, instansiOnProgramsResponse)
}

func (c *instansiOnProgramController) GetInstansiOnProgram(cntx *gin.Context) {
	var idString = cntx.Param("id")
	var id, _ = strconv.Atoi(idString)

	var instansiOnProgram, err = c.instansiOnProgramService.FindById(id)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": "Data tidak ditemukan",
		})
		return
	}

	var instansiOnProgramResponse = helper.ConvertToInstansiOnProgramResponse(instansiOnProgram)

	cntx.JSON(http.StatusOK, instansiOnProgramResponse)
}

func (c *instansiOnProgramController) CreateInstansiOnProgram(cntx *gin.Context) {
	var instansiOnProgramRequest request.CreateInstansiOnProgramRequest

	var err = cntx.ShouldBindJSON(&instansiOnProgramRequest)
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

	instansiOnProgram, err := c.instansiOnProgramService.Create(instansiOnProgramRequest)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	var instansiOnProgramResponse = helper.ConvertToInstansiOnProgramResponse(instansiOnProgram)

	cntx.JSON(http.StatusOK, instansiOnProgramResponse)
}

func (c *instansiOnProgramController) UpdateInstansiOnProgram(cntx *gin.Context) {
	var instansiOnProgramRequest request.UpdateInstansiOnProgramRequest

	var err = cntx.ShouldBindJSON(&instansiOnProgramRequest)
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

	instansiOnProgram, err := c.instansiOnProgramService.Update(id, instansiOnProgramRequest)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	var instansiOnProgramResponse = helper.ConvertToInstansiOnProgramResponse(instansiOnProgram)

	cntx.JSON(http.StatusOK, instansiOnProgramResponse)
}

func (c *instansiOnProgramController) DeleteInstansiOnProgram(cntx *gin.Context) {
	var idString = cntx.Param("id")
	var id, _ = strconv.Atoi(idString)

	_, err := c.instansiOnProgramService.Delete(id)
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
