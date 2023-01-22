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

type pusbangController struct {
	pusbangService service.PusbangService
}

func NewPusbangController(pusbangService service.PusbangService) *pusbangController {
	return &pusbangController{pusbangService}
}

func (c *pusbangController) GetPusbangs(cntx *gin.Context) {
	var pusbangs, err = c.pusbangService.FindAll()
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	var pusbangsResponse []responses.PusbangResponse

	for _, pusbang := range pusbangs {
		var pusbangResponse = helper.ConvertToPusbangResponse(pusbang)
		pusbangsResponse = append(pusbangsResponse, pusbangResponse)
	}

	cntx.JSON(http.StatusOK, pusbangsResponse)
}

func (c *pusbangController) GetPusbang(cntx *gin.Context) {
	var idString = cntx.Param("id")
	var id, _ = strconv.Atoi(idString)

	var pusbang, err = c.pusbangService.FindById(id)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": "Data tidak ditemukan",
		})
		return
	}

	var pusbangResponse = helper.ConvertToPusbangResponse(pusbang)

	cntx.JSON(http.StatusOK, pusbangResponse)
}

func (c *pusbangController) GetPusbangWithRelation(cntx *gin.Context) {
	var pusbangs, err = c.pusbangService.FindAllRelation()
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	var pusbangsResponse []responses.PusbangWithRelationResponse

	for _, pusbang := range pusbangs {
		var pusbangResponse = helper.ConvertToPusbangWithRelationResponse(pusbang)
		pusbangsResponse = append(pusbangsResponse, pusbangResponse)
	}

	cntx.JSON(http.StatusOK, pusbangsResponse)
}

func (c *pusbangController) CreatePusbang(cntx *gin.Context) {
	var pusbangRequest request.CreatePusbangRequest

	var err = cntx.ShouldBindJSON(&pusbangRequest)
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

	pusbang, err := c.pusbangService.Create(pusbangRequest)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	var pusbangResponse = helper.ConvertToPusbangResponse(pusbang)

	cntx.JSON(http.StatusCreated, pusbangResponse)
}

func (c *pusbangController) UpdatePusbang(cntx *gin.Context) {
	var pusbangRequest request.UpdatePusbangRequest

	var err = cntx.ShouldBindJSON(&pusbangRequest)
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

	pusbang, err := c.pusbangService.Update(id, pusbangRequest)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"errors": cntx.Error(err),
		})
		return
	}

	var pusbangResponse = helper.ConvertToPusbangResponse(pusbang)

	cntx.JSON(http.StatusOK, pusbangResponse)
}

func (c *pusbangController) DeletePusbang(cntx *gin.Context) {
	var idString = cntx.Param("id")
	var id, _ = strconv.Atoi(idString)

	_, err := c.pusbangService.Delete(id)
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
