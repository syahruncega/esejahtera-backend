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

type detailSubKegiatanController struct {
	detailSubKegiatanService service.DetailSubKegiatanService
}

func NewDetailSubKegiatanController(detailSubKegiatanService service.DetailSubKegiatanService) *detailSubKegiatanController {
	return &detailSubKegiatanController{detailSubKegiatanService}
}

func (c *detailSubKegiatanController) GetDetailSubKegiatans(cntx *gin.Context) {
	var detailSubKegiatans, err = c.detailSubKegiatanService.FindAll()

	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
	}

	var detailSubKegiatansResponse []responses.DetailSubKegiatanResponse

	for _, detailSubKegiatan := range detailSubKegiatans {
		var detailSubKegiatanResponse = helper.ConvertToDetailSubKegiatanResponse(detailSubKegiatan)
		detailSubKegiatansResponse = append(detailSubKegiatansResponse, detailSubKegiatanResponse)
	}

	cntx.JSON(http.StatusOK, detailSubKegiatansResponse)
}

func (c *detailSubKegiatanController) GetDetailSubKegiatan(cntx *gin.Context) {
	var idString = cntx.Param("id")
	var id, _ = strconv.Atoi(idString)

	var detailSubKegiatan, err = c.detailSubKegiatanService.FindById(id)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": "Data tidak ditemukan",
		})
		return
	}

	var detailSubKegiatanResponse = helper.ConvertToDetailSubKegiatanResponse(detailSubKegiatan)

	cntx.JSON(http.StatusOK, detailSubKegiatanResponse)
}

func (c *detailSubKegiatanController) CreateDetailSubKegiatan(cntx *gin.Context) {
	var detailSubKegiatanRequest request.CreateDetailSubKegiatanRequest

	var err = cntx.ShouldBindJSON(&detailSubKegiatanRequest)
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

	detailSubKegiatan, err := c.detailSubKegiatanService.Create(detailSubKegiatanRequest)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	var detailSubKegiatanResponse = helper.ConvertToDetailSubKegiatanResponse(detailSubKegiatan)

	cntx.JSON(http.StatusCreated, detailSubKegiatanResponse)
}

func (c *detailSubKegiatanController) UpdateDetailSubKegiatan(cntx *gin.Context) {
	var detailSubKegiatanRequest request.UpdateDetailSubKegiatanRequest

	var err = cntx.ShouldBindJSON(&detailSubKegiatanRequest)
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

	detailSubKegiatan, err := c.detailSubKegiatanService.Update(id, detailSubKegiatanRequest)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	var detailSubKegiatanResponse = helper.ConvertToDetailSubKegiatanResponse(detailSubKegiatan)

	cntx.JSON(http.StatusOK, detailSubKegiatanResponse)
}

func (c *detailSubKegiatanController) DeleteDetailSubKegiatan(cntx *gin.Context) {
	var idString = cntx.Param("id")
	var id, _ = strconv.Atoi(idString)

	_, err := c.detailSubKegiatanService.Delete(id)
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
