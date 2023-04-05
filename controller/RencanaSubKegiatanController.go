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

type rencanaSubKegiatanController struct {
	rencanaSubKegiatanService service.RencanaSubKegiatanService
}

func NewRencanaSubKegiatanController(rencanaSubKegiatanService service.RencanaSubKegiatanService) *rencanaSubKegiatanController {
	return &rencanaSubKegiatanController{rencanaSubKegiatanService}
}

func (c *rencanaSubKegiatanController) GetRencanaSubKegiatans(cntx *gin.Context) {
	var rencanaSubKegiatans, err = c.rencanaSubKegiatanService.FindAll()
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
	}

	var rencanaSubKegiatansResponse []responses.RencanaSubKegiatanResponse

	for _, rencanaSubKegiatan := range rencanaSubKegiatans {
		var rencanaSubKegiatanResponse = helper.ConvertToRencanaSubKegiatanResponse(rencanaSubKegiatan)
		rencanaSubKegiatansResponse = append(rencanaSubKegiatansResponse, rencanaSubKegiatanResponse)
	}

	cntx.JSON(http.StatusOK, rencanaSubKegiatansResponse)
}

func (c *rencanaSubKegiatanController) GetRencanaSubKegiatan(cntx *gin.Context) {
	var idString = cntx.Param("id")
	var id, _ = strconv.Atoi(idString)

	var rencanaSubKegiatan, err = c.rencanaSubKegiatanService.FindById(id)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": "Data tidak ditemukan",
		})
		return
	}

	var rencanaSubKegiatanResponse = helper.ConvertToRencanaSubKegiatanResponse(rencanaSubKegiatan)

	cntx.JSON(http.StatusOK, rencanaSubKegiatanResponse)
}

func (c *rencanaSubKegiatanController) GetRencanaSubKegiatansBySearch(cntx *gin.Context) {
	var whereClauseString = cntx.Request.URL.Query()
	var whereClauseInterface = map[string]interface{}{}

	for k, v := range whereClauseString {
		interfaceKey := k
		interfaceVal := v

		whereClauseInterface[interfaceKey] = interfaceVal
	}

	var rencanaSubKegiatans, err = c.rencanaSubKegiatanService.FindBySearch(whereClauseInterface)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
	}

	var rencanaSubKegiatansResponse []responses.RencanaSubKegiatanResponse

	for _, rencanaSubKegiatan := range rencanaSubKegiatans {
		var rencanaSubKegiatanResponse = helper.ConvertToRencanaSubKegiatanResponse(rencanaSubKegiatan)
		rencanaSubKegiatansResponse = append(rencanaSubKegiatansResponse, rencanaSubKegiatanResponse)
	}

	cntx.JSON(http.StatusOK, rencanaSubKegiatansResponse)
}

func (c *rencanaSubKegiatanController) CreateRencanaSubKegiatan(cntx *gin.Context) {
	var rencanaSubKegiatanRequest request.CreateRencanaSubKegiatanRequest

	var err = cntx.ShouldBindJSON(&rencanaSubKegiatanRequest)
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

	rencanaSubKegiatan, err := c.rencanaSubKegiatanService.Create(rencanaSubKegiatanRequest)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, cntx.Error(err))
		return
	}

	var rencanaSubKegiatanResponse = helper.ConvertToRencanaSubKegiatanResponse(rencanaSubKegiatan)

	cntx.JSON(http.StatusCreated, rencanaSubKegiatanResponse)
}

func (c *rencanaSubKegiatanController) UpdateRencanaSubKegiatan(cntx *gin.Context) {
	var rencanaSubKegiatanRequest request.UpdateRencanaSubKegiatanRequest

	var err = cntx.ShouldBindJSON(&rencanaSubKegiatanRequest)
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

	rencanaSubKegiatan, err := c.rencanaSubKegiatanService.Update(id, rencanaSubKegiatanRequest)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, cntx.Error(err))
		return
	}

	var rencanaSubKegiatanResponse = helper.ConvertToRencanaSubKegiatanResponse(rencanaSubKegiatan)

	cntx.JSON(http.StatusOK, rencanaSubKegiatanResponse)
}

func (c *rencanaSubKegiatanController) DeleteRencanaSubkKegiatan(cntx *gin.Context) {
	var idString = cntx.Param("id")
	var id, _ = strconv.Atoi(idString)

	_, err := c.rencanaSubKegiatanService.Delete(id)
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
