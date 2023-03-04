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

type programOnKegiatanController struct {
	programOnKegiatanService service.ProgramOnKegiatanService
}

func NewProgramOnKegiatanController(programOnKegiatanService service.ProgramOnKegiatanService) *programOnKegiatanController {
	return &programOnKegiatanController{programOnKegiatanService}
}

func (c *programOnKegiatanController) GetProgramOnKegiatans(cntx *gin.Context) {
	var programOnKegiatans, err = c.programOnKegiatanService.FindAll()
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
	}

	var programOnKegiatansResponse []responses.ProgramOnKegiatanResponse

	for _, programOnKegiatan := range programOnKegiatans {
		var programOnKegiatanResponse = helper.ConvertToProgramOnKegiatanResponse(programOnKegiatan)
		programOnKegiatansResponse = append(programOnKegiatansResponse, programOnKegiatanResponse)
	}

	cntx.JSON(http.StatusOK, programOnKegiatansResponse)
}

func (c *programOnKegiatanController) GetProgramOnKegiatan(cntx *gin.Context) {
	var idString = cntx.Param("id")
	var id, _ = strconv.Atoi(idString)

	var programOnKegiatan, err = c.programOnKegiatanService.FindById(id)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": "Data tidak ditemukan",
		})
		return
	}

	var programOnKegiatanResponse = helper.ConvertToProgramOnKegiatanResponse(programOnKegiatan)

	cntx.JSON(http.StatusOK, programOnKegiatanResponse)
}

func (c *programOnKegiatanController) CreateProgramOnKegiatan(cntx *gin.Context) {
	var programOnKegiatanRequest request.CreateProgramOnKegiatanRequest

	var err = cntx.ShouldBindJSON(&programOnKegiatanRequest)
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

	programOnKegiatan, err := c.programOnKegiatanService.Create(programOnKegiatanRequest)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	var programOnKegiatanResponse = helper.ConvertToProgramOnKegiatanResponse(programOnKegiatan)

	cntx.JSON(http.StatusOK, programOnKegiatanResponse)
}

func (c *programOnKegiatanController) UpdateProgramOnKegiatan(cntx *gin.Context) {
	var programOnKegiatanRequest request.UpdateProgramOnKegiatanRequest

	var err = cntx.ShouldBindJSON(programOnKegiatanRequest)
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

	programOnKegiatan, err := c.programOnKegiatanService.Update(id, programOnKegiatanRequest)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	var programOnKegiatanResponse = helper.ConvertToProgramOnKegiatanResponse(programOnKegiatan)

	cntx.JSON(http.StatusOK, programOnKegiatanResponse)
}

func (c *programOnKegiatanController) DeleteProgramOnKegiatan(cntx *gin.Context) {
	var idString = cntx.Param("id")
	var id, _ = strconv.Atoi(idString)

	_, err := c.programOnKegiatanService.Delete(id)
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
