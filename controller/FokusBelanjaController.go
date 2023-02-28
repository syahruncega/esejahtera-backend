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

type fokusBelanjaController struct {
	fokusBelanjaService service.FokusBelanjaService
}

func NewFokusBelanjaController(fokusBelanjaService service.FokusBelanjaService) *fokusBelanjaController {
	return &fokusBelanjaController{fokusBelanjaService}
}

func (c *fokusBelanjaController) GetFokusBelanjas(cntx *gin.Context) {
	var subKegiatanIdString = cntx.Query("subkegiatanid")

	var fokusBelanjas []model.FokusBelanja
	var err error

	if subKegiatanIdString != "" {
		var subKegiatanId, _ = strconv.Atoi(subKegiatanIdString)
		fokusBelanjas, err = c.fokusBelanjaService.FindBySubKegiatanId(subKegiatanId)
	} else {
		fokusBelanjas, err = c.fokusBelanjaService.FindAll()
	}

	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
	}

	var fokusBelanjasResponse []responses.FokusBelanjaResponse

	for _, fokusBelanja := range fokusBelanjas {
		var fokusBelanjaResponse = helper.ConvertToFokusBelanjaResponse(fokusBelanja)
		fokusBelanjasResponse = append(fokusBelanjasResponse, fokusBelanjaResponse)
	}

	cntx.JSON(http.StatusOK, fokusBelanjasResponse)
}

func (c *fokusBelanjaController) GetFokusBelanja(cntx *gin.Context) {
	var idString = cntx.Param("id")
	var id, _ = strconv.Atoi(idString)

	var fokusBelanja, err = c.fokusBelanjaService.FindById(id)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": "Data tidak ditemukan",
		})
		return
	}

	var fokusBelanjaResponse = helper.ConvertToFokusBelanjaResponse(fokusBelanja)

	cntx.JSON(http.StatusOK, fokusBelanjaResponse)
}

func (c *fokusBelanjaController) CreateFokusBelanja(cntx *gin.Context) {
	var fokusBelanjaRequest request.CreateFokusBelanjaRequest

	var err = cntx.ShouldBindJSON(&fokusBelanjaRequest)
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

	fokusBelanja, err := c.fokusBelanjaService.Create(fokusBelanjaRequest)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	var fokusBelanjaResponse = helper.ConvertToFokusBelanjaResponse(fokusBelanja)

	cntx.JSON(http.StatusCreated, fokusBelanjaResponse)
}

func (c *fokusBelanjaController) UpdateFokusBelanja(cntx *gin.Context) {
	var fokusBelanjaRequest request.UpdateFokusBelanjaRequest

	var err = cntx.ShouldBindJSON(&fokusBelanjaRequest)
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

	fokusBelanja, err := c.fokusBelanjaService.Update(id, fokusBelanjaRequest)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"errors": cntx.Error(err),
		})
		return
	}

	var fokusBelanjaResponse = helper.ConvertToFokusBelanjaResponse(fokusBelanja)

	cntx.JSON(http.StatusOK, fokusBelanjaResponse)
}

func (c *fokusBelanjaController) DeleteFokusBelanja(cntx *gin.Context) {
	var idString = cntx.Param("id")
	var id, _ = strconv.Atoi(idString)

	_, err := c.fokusBelanjaService.Delete(id)
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
