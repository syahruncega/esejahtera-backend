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

type indikatorSasaranController struct {
	indikatorSasaranService service.IndikatorSasaranService
}

func NewIndikatorSasaranController(indikatorSasaranService service.IndikatorSasaranService) *indikatorSasaranController {
	return &indikatorSasaranController{indikatorSasaranService}
}

func (c *indikatorSasaranController) GetIndikatorSasarans(cntx *gin.Context) {
	var indikatorSasarans, err = c.indikatorSasaranService.FindAll()
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
	}

	var indikatorSasaransResponse []responses.IndikatorSasaranResponse

	for _, indikatorSasaran := range indikatorSasarans {
		var indikatorSasaranResponse = helper.ConvertToIndikatorSasaranResponse(indikatorSasaran)
		indikatorSasaransResponse = append(indikatorSasaransResponse, indikatorSasaranResponse)
	}

	cntx.JSON(http.StatusOK, indikatorSasaransResponse)
}

func (c *indikatorSasaranController) GetIndikatorSasaran(cntx *gin.Context) {
	var idString = cntx.Param("id")
	var id, _ = strconv.Atoi(idString)

	var indikatorSasaran, err = c.indikatorSasaranService.FindById(id)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": "Data tidak ditemukan",
		})
		return
	}

	var indikatorSasaranResponse = helper.ConvertToIndikatorSasaranResponse(indikatorSasaran)

	cntx.JSON(http.StatusOK, indikatorSasaranResponse)
}

func (c *indikatorSasaranController) CreateIndikatorSasaran(cntx *gin.Context) {
	var indikatorSasaranRequest request.CreateIndikatorSasaranRequest

	var err = cntx.ShouldBindJSON(&indikatorSasaranRequest)
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

	indikatorSasaran, err := c.indikatorSasaranService.Create(indikatorSasaranRequest)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	var indikatorSasaranResponse = helper.ConvertToIndikatorSasaranResponse(indikatorSasaran)

	cntx.JSON(http.StatusOK, indikatorSasaranResponse)
}

func (c *indikatorSasaranController) UpdateIndikatorSasaran(cntx *gin.Context) {
	var indikatorSasaranRequest request.UpdateIndikatorSasaranRequest

	var err = cntx.ShouldBindJSON(&indikatorSasaranRequest)
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

	indikatorSasaran, err := c.indikatorSasaranService.Update(id, indikatorSasaranRequest)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	var indikatorSasaranResponse = helper.ConvertToIndikatorSasaranResponse(indikatorSasaran)

	cntx.JSON(http.StatusOK, indikatorSasaranResponse)
}

func (c *indikatorSasaranController) DeleteIndikatorSasaran(cntx *gin.Context) {
	var idString = cntx.Param("id")
	var id, _ = strconv.Atoi(idString)

	_, err := c.indikatorSasaranService.Delete(id)
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
