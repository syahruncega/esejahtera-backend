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

type indikatorSubKegiatanController struct {
	indikatorSubKegiatanService service.IndikatorSubKegiatanService
}

func NewIndikatorSubKegiatanController(indikatorSubKegiatanService service.IndikatorSubKegiatanService) *indikatorSubKegiatanController {
	return &indikatorSubKegiatanController{indikatorSubKegiatanService}
}

func (c *indikatorSubKegiatanController) GetIndikatorSubKegiatans(cntx *gin.Context) {
	var indikatorSubKegiatans, err = c.indikatorSubKegiatanService.FindAll()
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
	}

	var indikatorSubKegiatansResponse []responses.IndikatorSubKegiatanResponse

	for _, indikatorSubKegiatan := range indikatorSubKegiatans {
		var indikatorSubKegiatanResponse = helper.ConvertToIndikatorSubKegiatanResponse(indikatorSubKegiatan)
		indikatorSubKegiatansResponse = append(indikatorSubKegiatansResponse, indikatorSubKegiatanResponse)
	}

	cntx.JSON(http.StatusOK, indikatorSubKegiatansResponse)
}

func (c *indikatorSubKegiatanController) GetIndikatorSubKegiatan(cntx *gin.Context) {
	var idString = cntx.Param("id")
	var id, _ = strconv.Atoi(idString)

	var indikatorSubKegiatan, err = c.indikatorSubKegiatanService.FindById(id)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": "Data tidak ditemukan",
		})
		return
	}

	var indikatorSubKegiatanResponse = helper.ConvertToIndikatorSubKegiatanResponse(indikatorSubKegiatan)

	cntx.JSON(http.StatusOK, indikatorSubKegiatanResponse)
}

func (c *indikatorSubKegiatanController) CreateIndikatorSubKegiatan(cntx *gin.Context) {
	var indikatorSubKegiatanRequest request.UpdateIndikatorSubKegiatanRequest

	var err = cntx.ShouldBindJSON(&indikatorSubKegiatanRequest)
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

	indikatorSubKegiatan, err := c.indikatorSubKegiatanService.Create(request.CreateIndikatorSubKegiatanRequest(indikatorSubKegiatanRequest))
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	var indikatorSubKegiatanResponse = helper.ConvertToIndikatorSubKegiatanResponse(indikatorSubKegiatan)

	cntx.JSON(http.StatusCreated, indikatorSubKegiatanResponse)
}

func (c *indikatorSubKegiatanController) UpdateIndikatorSubKegiatan(cntx *gin.Context) {
	var indikatorSubKegiatanRequest request.UpdateIndikatorSubKegiatanRequest

	var err = cntx.ShouldBindJSON(&indikatorSubKegiatanRequest)
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

	indikatorSubKegiatan, err := c.indikatorSubKegiatanService.Update(id, indikatorSubKegiatanRequest)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"errors": cntx.Error(err),
		})
		return
	}

	var indikatorSubKegiatanResponse = helper.ConvertToIndikatorSubKegiatanResponse(indikatorSubKegiatan)

	cntx.JSON(http.StatusOK, indikatorSubKegiatanResponse)
}

func (c *indikatorSubKegiatanController) DeleteIndikatorSubKegiatan(cntx *gin.Context) {
	var idString = cntx.Param("id")
	var id, _ = strconv.Atoi(idString)

	_, err := c.indikatorSubKegiatanService.Delete(id)
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
