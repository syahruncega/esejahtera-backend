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

type detailLokasiController struct {
	detailLokasiService service.DetailLokasiService
}

func NewDetailLokasiController(detailLokasiService service.DetailLokasiService) *detailLokasiController {
	return &detailLokasiController{detailLokasiService}
}

func (c *detailLokasiController) GetDetailLokasis(cntx *gin.Context) {
	var detailLokasis, err = c.detailLokasiService.FindAll()
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
	}

	var detailLokasisResponse []responses.Detail_LokasiResponse

	for _, detailLokasi := range detailLokasis {
		var detailLokasiResponse = helper.ConvertToDetailLokasiResponse(detailLokasi)
		detailLokasisResponse = append(detailLokasisResponse, detailLokasiResponse)
	}

	cntx.JSON(http.StatusOK, detailLokasisResponse)
}

func (c *detailLokasiController) GetDetailLokasi(cntx *gin.Context) {
	var idString = cntx.Param("id")
	var id, _ = strconv.Atoi(idString)

	var detailLokasi, err = c.detailLokasiService.FindById(id)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": "Data tidak ditemukan",
		})
		return
	}

	var detailLokasiResponse = helper.ConvertToDetailLokasiResponse(detailLokasi)

	cntx.JSON(http.StatusOK, detailLokasiResponse)
}

func (c *detailLokasiController) CreateDetailLokasi(cntx *gin.Context) {
	var detailLokasiRequest request.CreateDetail_LokasiRequest

	var err = cntx.ShouldBindJSON(&detailLokasiRequest)
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

	detailLokasi, err := c.detailLokasiService.Create(detailLokasiRequest)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	var detailLokasiResponse = helper.ConvertToDetailLokasiResponse(detailLokasi)

	cntx.JSON(http.StatusCreated, gin.H{
		"data": detailLokasiResponse,
	})
}

func (c *detailLokasiController) UpdateDetailLokasi(cntx *gin.Context) {
	var detailLokasiRequest request.UpdateDetail_LokasiRequest

	var err = cntx.ShouldBindJSON(&detailLokasiRequest)
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

	detailLokasi, err := c.detailLokasiService.Update(id, detailLokasiRequest)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"errors": cntx.Error(err),
		})
		return
	}

	var detailLokasiResponse = helper.ConvertToDetailLokasiResponse(detailLokasi)

	cntx.JSON(http.StatusOK, gin.H{
		"data": detailLokasiResponse,
	})
}

func (c *detailLokasiController) DeleteDetailLokasi(cntx *gin.Context) {
	var idString = cntx.Param("id")
	var id, _ = strconv.Atoi(idString)

	_, err := c.detailLokasiService.Delete(id)
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
