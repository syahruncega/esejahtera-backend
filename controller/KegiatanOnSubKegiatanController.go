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

type kegiatanOnSubKegiatanController struct {
	kegiatanOnSubKegiatanService service.KegiatanOnSubKegiatanService
}

func NewKegiatanOnSubKegiatanController(kegiatanOnSubKegiatanService service.KegiatanOnSubKegiatanService) *kegiatanOnSubKegiatanController {
	return &kegiatanOnSubKegiatanController{kegiatanOnSubKegiatanService}
}

func (c *kegiatanOnSubKegiatanController) GetKegiatanOnSubKegiatans(cntx *gin.Context) {
	var kegiatanOnSubKegiatans []model.KegiatanOnSubKegiatan
	var err error

	var kegiatanIdString = cntx.Query("kegiatanid")
	var kegiatanId, _ = strconv.Atoi(kegiatanIdString)

	if kegiatanIdString != "" {
		kegiatanOnSubKegiatans, err = c.kegiatanOnSubKegiatanService.FindByKegiatanId(kegiatanId)
	} else {
		kegiatanOnSubKegiatans, err = c.kegiatanOnSubKegiatanService.FindAll()
	}

	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	var kegiatanOnSubKegiatansResponse []responses.KegiatanOnSubKegiatanResponse

	for _, kegiatanOnSubKegiatan := range kegiatanOnSubKegiatans {
		var kegiatanOnSubKegiatanResponse = helper.ConvertToKegiatanOnSubKegiatanResponse(kegiatanOnSubKegiatan)
		kegiatanOnSubKegiatansResponse = append(kegiatanOnSubKegiatansResponse, kegiatanOnSubKegiatanResponse)
	}

	cntx.JSON(http.StatusOK, kegiatanOnSubKegiatansResponse)
}

func (c *kegiatanOnSubKegiatanController) GetKegiatanOnSubKegiatan(cntx *gin.Context) {
	var idString = cntx.Param("id")
	var id, _ = strconv.Atoi(idString)

	var kegiatanOnSubKegiatan, err = c.kegiatanOnSubKegiatanService.FindById(id)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": "Data tidak ditemukan",
		})
		return
	}

	var kegiatanOnSubKegiatanResponse = helper.ConvertToKegiatanOnSubKegiatanResponse(kegiatanOnSubKegiatan)

	cntx.JSON(http.StatusOK, kegiatanOnSubKegiatanResponse)
}

func (c *kegiatanOnSubKegiatanController) CreateKegiatanOnSubKegiatan(cntx *gin.Context) {
	var kegiatanOnSubKegiatanRequest request.CreateKegiatanOnSubKegiatanRequest

	var err = cntx.ShouldBindJSON(&kegiatanOnSubKegiatanRequest)
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

	kegiatanOnSubKegiatan, err := c.kegiatanOnSubKegiatanService.Create(kegiatanOnSubKegiatanRequest)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	var kegiatanOnSubKegiatanResponse = helper.ConvertToKegiatanOnSubKegiatanResponse(kegiatanOnSubKegiatan)

	cntx.JSON(http.StatusOK, kegiatanOnSubKegiatanResponse)
}

func (c *kegiatanOnSubKegiatanController) UpdateKegiatanOnSubKegiatan(cntx *gin.Context) {
	var kegiatanOnSubKegiatanRequest request.UpdateKegiatanOnSubKegiatanRequest

	var err = cntx.ShouldBindJSON(&kegiatanOnSubKegiatanRequest)
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

	kegiatanOnSubKegiatan, err := c.kegiatanOnSubKegiatanService.Update(id, kegiatanOnSubKegiatanRequest)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	var kegiatanOnSubKegiatanResponse = helper.ConvertToKegiatanOnSubKegiatanResponse(kegiatanOnSubKegiatan)

	cntx.JSON(http.StatusOK, kegiatanOnSubKegiatanResponse)
}

func (c *kegiatanOnSubKegiatanController) DeleteKegiatanOnSubKegiatan(cntx *gin.Context) {
	var idString = cntx.Param("id")
	var id, _ = strconv.Atoi(idString)

	_, err := c.kegiatanOnSubKegiatanService.Delete(id)
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
