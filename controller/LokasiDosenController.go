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

type lokasiDosenController struct {
	lokasiDosenService service.LokasiDosenService
}

func NewLokasiDosenController(lokasiDosenService service.LokasiDosenService) *lokasiDosenController {
	return &lokasiDosenController{lokasiDosenService}
}

func (c *lokasiDosenController) GetLokasiDosens(cntx *gin.Context) {
	var lokasiDosens, err = c.lokasiDosenService.FindAll()
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	var lokasiDosensResponse []responses.LokasiDosenResponse

	for _, lokasiDosen := range lokasiDosens {
		var lokasiDosenResponse = helper.ConvertToLokasiDosenResponse(lokasiDosen)
		lokasiDosensResponse = append(lokasiDosensResponse, lokasiDosenResponse)
	}

	cntx.JSON(http.StatusOK, lokasiDosensResponse)
}

func (c *lokasiDosenController) GetLokasiDosen(cntx *gin.Context) {
	var idString = cntx.Param("id")
	var id, _ = strconv.Atoi(idString)

	var lokasiDosen, err = c.lokasiDosenService.FindById(id)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": "Data tidak ditemukan",
		})
		return
	}

	var lokasiDosenResponse = helper.ConvertToLokasiDosenResponse(lokasiDosen)

	cntx.JSON(http.StatusOK, lokasiDosenResponse)
}

func (c *lokasiDosenController) GetLokasiDosenWithRelation(cntx *gin.Context) {
	var lokasiDosens, err = c.lokasiDosenService.FindAllRelation()
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	var lokasiDosensResponse []responses.LokasiDosenWithRelationResponse

	for _, lokasiDosen := range lokasiDosens {
		var lokasiDosenResponse = helper.ConvertToLokasiDosenWithRelationResponse(lokasiDosen)
		lokasiDosensResponse = append(lokasiDosensResponse, lokasiDosenResponse)
	}

	cntx.JSON(http.StatusOK, lokasiDosensResponse)
}

func (c *lokasiDosenController) GetLokasiDosenRelationByDosenId(cntx *gin.Context) {
	var dosenIdString = cntx.Param("dosenid")
	var dosenId, _ = strconv.Atoi(dosenIdString)

	var lokasiDosens, err = c.lokasiDosenService.FindRelationByDosenId(dosenId)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	var lokasiDosensResponse []responses.LokasiDosenWithRelationResponse

	for _, lokasiDosen := range lokasiDosens {
		var lokasiDosenResponse = helper.ConvertToLokasiDosenWithRelationResponse(lokasiDosen)
		lokasiDosensResponse = append(lokasiDosensResponse, lokasiDosenResponse)
	}

	cntx.JSON(http.StatusOK, lokasiDosensResponse)

}

func (c *lokasiDosenController) CreateLokasiDosen(cntx *gin.Context) {
	var lokasiDosenRequest request.CreateLokasiDosenRequest

	var err = cntx.ShouldBindJSON(&lokasiDosenRequest)
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

	lokasiDosen, err := c.lokasiDosenService.Create(lokasiDosenRequest)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	var lokasiDosenResponse = helper.ConvertToLokasiDosenResponse(lokasiDosen)

	cntx.JSON(http.StatusCreated, lokasiDosenResponse)
}

func (c *lokasiDosenController) UpdateLokasiDosen(cntx *gin.Context) {
	var lokasiDosenRequest request.UpdateLokasiDosenRequest

	var err = cntx.ShouldBindJSON(&lokasiDosenRequest)
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

	lokasiDosen, err := c.lokasiDosenService.Update(id, lokasiDosenRequest)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"errors": cntx.Error(err),
		})
		return
	}

	var lokasiDosenResponse = helper.ConvertToLokasiDosenResponse(lokasiDosen)

	cntx.JSON(http.StatusOK, lokasiDosenResponse)
}

func (c *lokasiDosenController) DeleteLokasiDosen(cntx *gin.Context) {
	var idString = cntx.Param("id")
	var id, _ = strconv.Atoi(idString)

	_, err := c.lokasiDosenService.Delete(id)
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
