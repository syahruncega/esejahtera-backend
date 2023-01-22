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

type analisController struct {
	analisService service.AnalisService
}

func NewAnalisController(analisService service.AnalisService) *analisController {
	return &analisController{analisService}
}

func (c *analisController) GetAnaliss(cntx *gin.Context) {
	var analiss, err = c.analisService.FindAll()
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	var analissResponse []responses.AnalisResponse

	for _, analis := range analiss {
		var analisResponse = helper.ConvertToAnalisResponse(analis)
		analissResponse = append(analissResponse, analisResponse)
	}

	cntx.JSON(http.StatusOK, analissResponse)
}

func (c *analisController) GetAnalis(cntx *gin.Context) {
	var idString = cntx.Param("id")
	var id, _ = strconv.Atoi(idString)

	var analis, err = c.analisService.FindById(id)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": "Data tidak ditemukan",
		})
		return
	}

	var analisResponse = helper.ConvertToAnalisResponse(analis)

	cntx.JSON(http.StatusOK, analisResponse)
}

func (c *analisController) GetAnalisWithRelation(cntx *gin.Context) {
	var analiss, err = c.analisService.FindAllRelation()
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	var analissResponse []responses.AnalisWithRelationResponse

	for _, analis := range analiss {
		var analisResponse = helper.ConvertToAnalisWithRelationResponse(analis)
		analissResponse = append(analissResponse, analisResponse)
	}

	cntx.JSON(http.StatusOK, analissResponse)
}

func (c *analisController) CreateAnalis(cntx *gin.Context) {
	var analisRequest request.CreateAnalisRequest

	var err = cntx.ShouldBindJSON(&analisRequest)
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

	analis, err := c.analisService.Create(analisRequest)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	var analisResponse = helper.ConvertToAnalisResponse(analis)

	cntx.JSON(http.StatusCreated, gin.H{
		"data": analisResponse,
	})
}

func (c *analisController) UpdateAnalis(cntx *gin.Context) {
	var analisRequest request.UpdateAnalisRequest

	var err = cntx.ShouldBindJSON(&analisRequest)
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

	analis, err := c.analisService.Update(id, analisRequest)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"errors": cntx.Error(err),
		})
		return
	}

	var analisResponse = helper.ConvertToAnalisResponse(analis)

	cntx.JSON(http.StatusOK, gin.H{
		"data": analisResponse,
	})
}

func (c *analisController) DeleteAnalis(cntx *gin.Context) {
	var idString = cntx.Param("id")
	var id, _ = strconv.Atoi(idString)

	_, err := c.analisService.Delete(id)
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
