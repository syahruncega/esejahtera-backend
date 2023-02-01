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

type indikatorKegiatanController struct {
	indikatorKegiatanService service.IndikatorKegiatanService
}

func NewIndikatorKegiatanController(indikatorKegiatanService service.IndikatorKegiatanService) *indikatorKegiatanController {
	return &indikatorKegiatanController{indikatorKegiatanService}
}

func (c *indikatorKegiatanController) GetIndikatorKegiatans(cntx *gin.Context) {
	var indikatorKegiatans, err = c.indikatorKegiatanService.FindAll()
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
	}

	var indikatorKegiatansResponse []responses.IndikatorKegiatanResponse

	for _, indikatorKegiatan := range indikatorKegiatans {
		var indikatorKegiatanResponse = helper.ConvertToIndikatorKegiatanResponse(indikatorKegiatan)
		indikatorKegiatansResponse = append(indikatorKegiatansResponse, indikatorKegiatanResponse)
	}

	cntx.JSON(http.StatusOK, indikatorKegiatansResponse)
}

func (c *indikatorKegiatanController) GetIndikatorKegiatan(cntx *gin.Context) {
	var idString = cntx.Param("id")
	var id, _ = strconv.Atoi(idString)

	var indikatorKegiatan, err = c.indikatorKegiatanService.FindById(id)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": "Data tidak ditemukan",
		})
		return
	}

	var indikatorKegiatanResponse = helper.ConvertToIndikatorKegiatanResponse(indikatorKegiatan)

	cntx.JSON(http.StatusOK, indikatorKegiatanResponse)
}

func (c *indikatorKegiatanController) CreateIndikatorKegiatan(cntx *gin.Context) {
	var indikatorKegiatanRequest request.CreateIndikatorKegiatanRequest

	var err = cntx.ShouldBindJSON(&indikatorKegiatanRequest)
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

	indikatorKegiatan, err := c.indikatorKegiatanService.Create(indikatorKegiatanRequest)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	var indikatorKegiatanResponse = helper.ConvertToIndikatorKegiatanResponse(indikatorKegiatan)

	cntx.JSON(http.StatusCreated, indikatorKegiatanResponse)
}

func (c *indikatorKegiatanController) UpdateIndikatorKegiatan(cntx *gin.Context) {
	var indikatorKegiatanRequest request.UpdateIndikatorKegiatanRequest

	var err = cntx.ShouldBindJSON(&indikatorKegiatanRequest)
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

	indikatorKegiatan, err := c.indikatorKegiatanService.Update(id, indikatorKegiatanRequest)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"errors": cntx.Error(err),
		})
		return
	}

	var indikatorKegiatanResponse = helper.ConvertToIndikatorKegiatanResponse(indikatorKegiatan)

	cntx.JSON(http.StatusOK, indikatorKegiatanResponse)
}

func (c *indikatorKegiatanController) DeleteIndikatorKegiatan(cntx *gin.Context) {
	var idString = cntx.Param("id")
	var id, _ = strconv.Atoi(idString)

	_, err := c.indikatorKegiatanService.Delete(id)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": "Data gagal dihapus",
		})
		return
	}

	cntx.JSON(http.StatusOK, gin.H{
		"status": "Data Berhasil Dihapus",
	})
}
