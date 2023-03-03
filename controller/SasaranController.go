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

type sasaranController struct {
	sasaranService service.SasaranService
}

func NewSasaranController(sasaranService service.SasaranService) *sasaranController {
	return &sasaranController{sasaranService}
}

func (c *sasaranController) GetSasarans(cntx *gin.Context) {
	var sasarans, err = c.sasaranService.FindAll()
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
	}

	var sasaransResponse []responses.SasaranResponse

	for _, sasaran := range sasarans {
		var sasaranResponse = helper.ConvertToSasaranResponse(sasaran)
		sasaransResponse = append(sasaransResponse, sasaranResponse)
	}

	cntx.JSON(http.StatusOK, sasaransResponse)
}

func (c *sasaranController) GetSasaran(cntx *gin.Context) {
	var idString = cntx.Param("id")
	var id, _ = strconv.Atoi(idString)

	var sasaran, err = c.sasaranService.FindById(id)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": "Data tidak ditemukan",
		})
		return
	}

	var sasaranResponse = helper.ConvertToSasaranResponse(sasaran)

	cntx.JSON(http.StatusOK, sasaranResponse)
}

func (c *sasaranController) CreateSasaran(cntx *gin.Context) {
	var sasaranRequest request.CreateSasaranRequest

	var err = cntx.ShouldBindJSON(&sasaranRequest)
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

	sasaran, err := c.sasaranService.Create(sasaranRequest)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	var sasaranResponse = helper.ConvertToSasaranResponse(sasaran)

	cntx.JSON(http.StatusOK, sasaranResponse)
}

func (c *sasaranController) UpdateSasaran(cntx *gin.Context) {
	var sasaranRequest request.UpdateSasaranRequest

	var err = cntx.ShouldBindJSON(&sasaranRequest)
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

	sasaran, err := c.sasaranService.Update(id, sasaranRequest)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	var sasaranResponse = helper.ConvertToSasaranResponse(sasaran)

	cntx.JSON(http.StatusOK, sasaranResponse)
}

func (c *sasaranController) DeleteSasaran(cntx *gin.Context) {
	var idString = cntx.Param("id")
	var id, _ = strconv.Atoi(idString)

	_, err := c.sasaranService.Delete(id)
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
