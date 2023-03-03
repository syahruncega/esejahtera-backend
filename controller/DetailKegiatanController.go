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

type detailKegiatanController struct {
	detailKegiatanService service.DetailKegiatanService
}

func NewDetailKegiatanController(detailKegiatanService service.DetailKegiatanService) *detailKegiatanController {
	return &detailKegiatanController{detailKegiatanService}
}

func (c *detailKegiatanController) GetDetailKegiatans(cntx *gin.Context) {
	var detailKegiatans, err = c.detailKegiatanService.FindAll()

	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
	}

	var detailKegiatansResponse []responses.DetailKegiatanResponse

	for _, detailKegiatan := range detailKegiatans {
		var detailKegiatanResponse = helper.ConvertToDetailKegiatanResponse(detailKegiatan)
		detailKegiatansResponse = append(detailKegiatansResponse, detailKegiatanResponse)
	}

	cntx.JSON(http.StatusOK, detailKegiatansResponse)
}

func (c *detailKegiatanController) GetDetailKegiatan(cntx *gin.Context) {
	var idString = cntx.Param("id")
	var id, _ = strconv.Atoi(idString)

	var detailKegiatan, err = c.detailKegiatanService.FindById(id)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": "Data tidak ditemukan",
		})
		return
	}

	var detailKegiatanResponse = helper.ConvertToDetailKegiatanResponse(detailKegiatan)

	cntx.JSON(http.StatusOK, detailKegiatanResponse)
}

func (c *detailKegiatanController) CreateDetailKegiatan(cntx *gin.Context) {
	var detailKegiatanRequest request.CreateDetailKegiatanRequest

	var err = cntx.ShouldBindJSON(&detailKegiatanRequest)
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

	detailKegiatan, err := c.detailKegiatanService.Create(detailKegiatanRequest)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	var detailKegiatanResponse = helper.ConvertToDetailKegiatanResponse(detailKegiatan)

	cntx.JSON(http.StatusCreated, detailKegiatanResponse)
}

func (c *detailKegiatanController) UpdateDetailKegiatan(cntx *gin.Context) {
	var detailKegiatanRequest request.UpdateDetailKegiatanRequest

	var err = cntx.ShouldBindJSON(&detailKegiatanRequest)
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

	detailKegiatan, err := c.detailKegiatanService.Update(id, detailKegiatanRequest)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	var detailKegiatanResponse = helper.ConvertToDetailKegiatanResponse(detailKegiatan)

	cntx.JSON(http.StatusOK, detailKegiatanResponse)
}

func (c *detailKegiatanController) DeleteDetailKegiatan(cntx *gin.Context) {
	var idString = cntx.Param("id")
	var id, _ = strconv.Atoi(idString)

	_, err := c.detailKegiatanService.Delete(id)
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
