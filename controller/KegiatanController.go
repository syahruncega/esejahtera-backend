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

type kegiatanController struct {
	kegiatanService service.KegiatanService
}

func NewKegiatanController(kegiatanService service.KegiatanService) *kegiatanController {
	return &kegiatanController{kegiatanService}
}

func (c *kegiatanController) GetKegiatans(cntx *gin.Context) {
	var kegiatans, err = c.kegiatanService.FindAll()
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
	}

	var kegiatansResponse []responses.KegiatanResponse

	for _, kegiatan := range kegiatans {
		var kegiatanResponse = helper.ConvertToKegiatanResponse(kegiatan)
		kegiatansResponse = append(kegiatansResponse, kegiatanResponse)
	}

	cntx.JSON(http.StatusOK, kegiatansResponse)
}

func (c *kegiatanController) GetKegiatan(cntx *gin.Context) {
	var idString = cntx.Param("id")
	var id, _ = strconv.Atoi(idString)

	var kegiatan, err = c.kegiatanService.FindById(id)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": "Data tidak ditemukan",
		})
		return
	}

	var kegiatanResponse = helper.ConvertToKegiatanResponse(kegiatan)

	cntx.JSON(http.StatusOK, kegiatanResponse)
}

func (c *kegiatanController) GetKegiatanWithRelation(cntx *gin.Context) {
	var kegiatanRelations, err = c.kegiatanService.FindAllRelation()
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
	}

	var kegiatansRelationResponse []responses.KegiatanWithProgramResponse

	for _, kegiatanRelation := range kegiatanRelations {
		var kegiatanRelationResponse = helper.ConvertToKegiatanWithProgramResponse(kegiatanRelation)
		kegiatansRelationResponse = append(kegiatansRelationResponse, kegiatanRelationResponse)
	}

	cntx.JSON(http.StatusOK, kegiatansRelationResponse)
}

func (c *kegiatanController) CreateKegiatan(cntx *gin.Context) {
	var kegiatanRequest request.CreateKegiatanRequest

	var err = cntx.ShouldBindJSON(&kegiatanRequest)
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

	kegiatan, err := c.kegiatanService.Create(kegiatanRequest)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	var kegiatanResponse = helper.ConvertToKegiatanResponse(kegiatan)

	cntx.JSON(http.StatusCreated, gin.H{
		"data": kegiatanResponse,
	})
}

func (c *kegiatanController) UpdateKegiatan(cntx *gin.Context) {
	var kegiatanRequest request.UpdateKegiatanRequest

	var err = cntx.ShouldBindJSON(&kegiatanRequest)
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

	kegiatan, err := c.kegiatanService.Update(id, kegiatanRequest)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"errors": cntx.Error(err),
		})
		return
	}

	var kegiatanResponse = helper.ConvertToKegiatanResponse(kegiatan)

	cntx.JSON(http.StatusOK, gin.H{
		"data": kegiatanResponse,
	})
}

func (c *kegiatanController) DeleteKegiatan(cntx *gin.Context) {
	var idString = cntx.Param("id")
	var id, _ = strconv.Atoi(idString)

	_, err := c.kegiatanService.Delete(id)
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
