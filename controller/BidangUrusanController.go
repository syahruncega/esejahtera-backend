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

type bidangUrusanController struct {
	bidangUrusanService service.BidangUrusanService
}

func NewBidangUrusanController(bidangUrusanService service.BidangUrusanService) *bidangUrusanController {
	return &bidangUrusanController{bidangUrusanService}
}

func (c *bidangUrusanController) GetBidangUrusans(cntx *gin.Context) {
	var bidangUrusans, err = c.bidangUrusanService.FindAll()
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
	}

	var bidangUrusansResponse []responses.BidangUrusanResponse

	for _, bidangUrusan := range bidangUrusans {
		var bidangUrusanResponse = helper.ConvertToBidangUrusanResponse(bidangUrusan)
		bidangUrusansResponse = append(bidangUrusansResponse, bidangUrusanResponse)
	}

	cntx.JSON(http.StatusOK, bidangUrusansResponse)
}

func (c *bidangUrusanController) GetBidangUrusan(cntx *gin.Context) {
	var idString = cntx.Param("id")
	var id, _ = strconv.Atoi(idString)

	var bidangUrusan, err = c.bidangUrusanService.FindById(id)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": "Data tidak ditemukan",
		})
		return
	}

	var bidangUrusanResponse = helper.ConvertToBidangUrusanResponse(bidangUrusan)

	cntx.JSON(http.StatusOK, bidangUrusanResponse)
}

func (c *bidangUrusanController) CreateBidangUrusan(cntx *gin.Context) {
	var bidangUrusanRequest request.CreateBidangUrusan

	var err = cntx.ShouldBindJSON(&bidangUrusanRequest)
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

	bidangUrusan, err := c.bidangUrusanService.Create(bidangUrusanRequest)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	var bidangUrusanResponse = helper.ConvertToBidangUrusanResponse(bidangUrusan)

	cntx.JSON(http.StatusCreated, gin.H{
		"data": bidangUrusanResponse,
	})
}

func (c *bidangUrusanController) UpdateBidangUrusan(cntx *gin.Context) {
	var bidangUrusanRequest request.UpdateBidangUrusan

	var err = cntx.ShouldBindJSON(&bidangUrusanRequest)
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

	bidangUrusan, err := c.bidangUrusanService.Update(id, bidangUrusanRequest)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"errors": cntx.Error(err),
		})
		return
	}

	var bidangUrusanResponse = helper.ConvertToBidangUrusanResponse(bidangUrusan)

	cntx.JSON(http.StatusOK, gin.H{
		"data": bidangUrusanResponse,
	})
}

func (c *bidangUrusanController) DeleteBidangUrusan(cntx *gin.Context) {
	var idString = cntx.Param("id")
	var id, _ = strconv.Atoi(idString)

	_, err := c.bidangUrusanService.Delete(id)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	cntx.JSON(http.StatusOK, gin.H{
		"status": "Data berhasil dihapus",
	})
}
