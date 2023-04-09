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

type bidangUrusanOnInstansiController struct {
	bidangUrusanOnInstansiService service.BidangUrusanOnInstansiService
}

func NewBidangUrusanOnInstansiController(bidangUrusanOnInstansiService service.BidangUrusanOnInstansiService) *bidangUrusanOnInstansiController {
	return &bidangUrusanOnInstansiController{bidangUrusanOnInstansiService}
}

func (c *bidangUrusanOnInstansiController) GetBidangUrusanOnInstansis(cntx *gin.Context) {
	var bidangUrusanOnInstansis []model.BidangUrusanOnInstansi
	var err error

	var instansiIdString = cntx.Query("instansiid")
	var instansiId, _ = strconv.Atoi(instansiIdString)

	if instansiIdString != "" {
		bidangUrusanOnInstansis, err = c.bidangUrusanOnInstansiService.FindByInstansiId(instansiId)
	} else {
		bidangUrusanOnInstansis, err = c.bidangUrusanOnInstansiService.FindAll()
	}

	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
	}

	var bidangUrusanOnInstansisResponse []responses.BidangUrusanOnInstansiResponse

	for _, bidangUrusanOnInstansi := range bidangUrusanOnInstansis {
		var bidangUrusanOnInstansiResponse = helper.ConvertToBidangUrusanOnInstansiResponse(bidangUrusanOnInstansi)
		bidangUrusanOnInstansisResponse = append(bidangUrusanOnInstansisResponse, bidangUrusanOnInstansiResponse)
	}

	cntx.JSON(http.StatusOK, bidangUrusanOnInstansisResponse)
}

func (c *bidangUrusanOnInstansiController) GetBidangUrusanOnInstansi(cntx *gin.Context) {
	var idString = cntx.Param("id")
	var id, _ = strconv.Atoi(idString)

	var bidangUrusanOnInstansi, err = c.bidangUrusanOnInstansiService.FindById(id)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": "Data tidak ditemukan",
		})
		return
	}

	var bidangUrusanOnInstansiResponse = helper.ConvertToBidangUrusanOnInstansiResponse(bidangUrusanOnInstansi)

	cntx.JSON(http.StatusOK, bidangUrusanOnInstansiResponse)
}

func (c *bidangUrusanOnInstansiController) GetBidangUrusanOnInstansiBySearch(cntx *gin.Context) {
	var whereClauseString = cntx.Request.URL.Query()
	var whereClauseInterface = make(map[string]interface{})

	for k, v := range whereClauseString {
		interfaceKey := k
		interfaceVal := v

		whereClauseInterface[interfaceKey] = interfaceVal
	}

	var bidangUrusanOnInstansis, err = c.bidangUrusanOnInstansiService.FindBySearch(whereClauseInterface)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
	}

	var bidangUrusanOnInstansisResponse []responses.BidangUrusanOnInstansiResponse

	for _, bidangUrusanOnInstansi := range bidangUrusanOnInstansis {
		var bidangUrusanOnInstansiResponse = helper.ConvertToBidangUrusanOnInstansiResponse(bidangUrusanOnInstansi)
		bidangUrusanOnInstansisResponse = append(bidangUrusanOnInstansisResponse, bidangUrusanOnInstansiResponse)
	}

	cntx.JSON(http.StatusOK, bidangUrusanOnInstansisResponse)
}

func (c *bidangUrusanOnInstansiController) CreateBidangUrusanOnInstansi(cntx *gin.Context) {
	var bidangUrusanOnInstansiRequest request.CreateBidangUrusanOnInstansiRequest

	var err = cntx.ShouldBindJSON(&bidangUrusanOnInstansiRequest)
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

	bidangUrusanOnInstansi, err := c.bidangUrusanOnInstansiService.Create(bidangUrusanOnInstansiRequest)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	var bidangUrusanOnInstansiResponse = helper.ConvertToBidangUrusanOnInstansiResponse(bidangUrusanOnInstansi)

	cntx.JSON(http.StatusOK, bidangUrusanOnInstansiResponse)
}

func (c *bidangUrusanOnInstansiController) UpdateBidangUrusanOnInstansi(cntx *gin.Context) {
	var bidangUrusanOnInstansiRequest request.UpdateBidangUrusanOnInstansiRequest

	var err = cntx.ShouldBindJSON(&bidangUrusanOnInstansiRequest)
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

	bidangUrusanOnInstansi, err := c.bidangUrusanOnInstansiService.Update(id, bidangUrusanOnInstansiRequest)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	var bidangUrusanOnInstansiResponse = helper.ConvertToBidangUrusanOnInstansiResponse(bidangUrusanOnInstansi)

	cntx.JSON(http.StatusOK, bidangUrusanOnInstansiResponse)
}

func (c *bidangUrusanOnInstansiController) DeleteBidangUrusanOnInstansi(cntx *gin.Context) {
	var idString = cntx.Param("id")
	var id, _ = strconv.Atoi(idString)

	_, err := c.bidangUrusanOnInstansiService.Delete(id)
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
