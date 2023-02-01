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

type detailInstansiController struct {
	detailInstansiService service.DetailInstansiService
}

func NewDetailInstansiController(detailInstansiService service.DetailInstansiService) *detailInstansiController {
	return &detailInstansiController{detailInstansiService}
}

func (c *detailInstansiController) GetDetailInstansis(cntx *gin.Context) {
	var instansiId = cntx.Query("instansiid")

	var detailInstansis []model.DetailInstansi
	var err error

	var detailInstansisResponse []responses.DetailInstansiResponse

	if instansiId != "" {
		detailInstansis, err = c.detailInstansiService.FindByInstansiId(instansiId)
	} else {
		detailInstansis, err = c.detailInstansiService.FindAll()
	}

	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
	}

	for _, detailInstansi := range detailInstansis {
		var detailInstansiResponse = helper.ConvertToDetailInstansiResponse(detailInstansi)
		detailInstansisResponse = append(detailInstansisResponse, detailInstansiResponse)
	}

	cntx.JSON(http.StatusOK, detailInstansisResponse)
}

func (c *detailInstansiController) GetDetailInstansi(cntx *gin.Context) {
	var idString = cntx.Param("id")
	var id, _ = strconv.Atoi(idString)

	var detailInstansi, err = c.detailInstansiService.FindById(id)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": "Data tidak ditemukan",
		})
		return
	}

	var detailInstansiResponse = helper.ConvertToDetailInstansiResponse(detailInstansi)

	cntx.JSON(http.StatusOK, detailInstansiResponse)
}

func (c *detailInstansiController) CreateDetailInstansi(cntx *gin.Context) {
	var detailInstansiRequest request.CreateDetailInstansiRequest

	var err = cntx.ShouldBindJSON(&detailInstansiRequest)
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

	detailInstansi, err := c.detailInstansiService.Create(detailInstansiRequest)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	var detailInstansiResponse = helper.ConvertToDetailInstansiResponse(detailInstansi)

	cntx.JSON(http.StatusCreated, detailInstansiResponse)
}

func (c *detailInstansiController) UpdateDetailInstansi(cntx *gin.Context) {
	var detailInstansiRequest request.UpdateDetailInstansiRequest

	var err = cntx.ShouldBindJSON(&detailInstansiRequest)
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

	detailInstansi, err := c.detailInstansiService.Update(id, detailInstansiRequest)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"errors": cntx.Error(err),
		})
		return
	}

	var detailInstansiResponse = helper.ConvertToDetailInstansiResponse(detailInstansi)

	cntx.JSON(http.StatusOK, detailInstansiResponse)
}

func (c *detailInstansiController) DeleteDetailInstansi(cntx *gin.Context) {
	var idString = cntx.Param("id")
	var id, _ = strconv.Atoi(idString)

	_, err := c.detailInstansiService.Delete(id)
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
