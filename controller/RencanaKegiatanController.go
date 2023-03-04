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

type rencanaKegiatanController struct {
	rencanaKegiatanService service.RencanaKegiatanService
}

func NewRencanaKegiatanController(rencanaKegiatanService service.RencanaKegiatanService) *rencanaKegiatanController {
	return &rencanaKegiatanController{rencanaKegiatanService}
}

func (c *rencanaKegiatanController) GetRencanaKegiatans(cntx *gin.Context) {
	var rencanaKegiatans, err = c.rencanaKegiatanService.FindAll()
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
	}

	var rencanaKegiatansResponse []responses.RencanaKegiatanResponse

	for _, rencanaKegiatan := range rencanaKegiatans {
		var rencanaKegiatanResponse = helper.ConvertToRencanaKegiatanResponse(rencanaKegiatan)
		rencanaKegiatansResponse = append(rencanaKegiatansResponse, rencanaKegiatanResponse)
	}

	cntx.JSON(http.StatusOK, rencanaKegiatansResponse)
}

func (c *rencanaKegiatanController) GetRencanaKegiatan(cntx *gin.Context) {
	var idString = cntx.Param("id")
	var id, _ = strconv.Atoi(idString)

	var rencanaKegiatan, err = c.rencanaKegiatanService.FindById(id)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": "Data tidak ditemukan",
		})
		return
	}

	var rencanaKegiatanResponse = helper.ConvertToRencanaKegiatanResponse(rencanaKegiatan)

	cntx.JSON(http.StatusOK, rencanaKegiatanResponse)
}

func (c *rencanaKegiatanController) CreateRencanaKegiatan(cntx *gin.Context) {
	var rencanaKegiatanRequest request.CreateRencanaKegiatanRequest

	var err = cntx.ShouldBindJSON(&rencanaKegiatanRequest)
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

	rencanaKegiatan, err := c.rencanaKegiatanService.Create(rencanaKegiatanRequest)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	var rencanaKegiatanResponse = helper.ConvertToRencanaKegiatanResponse(rencanaKegiatan)

	cntx.JSON(http.StatusCreated, rencanaKegiatanResponse)
}

func (c *rencanaKegiatanController) UpdateRencanaKegiatan(cntx *gin.Context) {
	var rencanaKegiatanRequest request.UpdateRencanaKegiatanRequest

	var err = cntx.ShouldBindJSON(&rencanaKegiatanRequest)
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

	rencanaKegiatan, err := c.rencanaKegiatanService.Update(id, rencanaKegiatanRequest)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	var rencanaKegiatanResponse = helper.ConvertToRencanaKegiatanResponse(rencanaKegiatan)

	cntx.JSON(http.StatusOK, rencanaKegiatanResponse)
}

func (c *rencanaKegiatanController) DeleteRencanaKegiatan(cntx *gin.Context) {
	var idString = cntx.Param("id")
	var id, _ = strconv.Atoi(idString)

	_, err := c.rencanaKegiatanService.Delete(id)
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
