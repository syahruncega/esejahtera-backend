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

type detailSubKegiatanController struct {
	detailSubKegiatanService service.DetailSubKegiatanService
}

func NewDetailSubKegiatanController(detailSubKegiatanService service.DetailSubKegiatanService) *detailSubKegiatanController {
	return &detailSubKegiatanController{detailSubKegiatanService}
}

func (c *detailSubKegiatanController) GetDetailSubKegiatans(cntx *gin.Context) {
	var detailSubKegiatans, err = c.detailSubKegiatanService.FindAll()
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
	}

	var detailSubKegiatansResponse []responses.Detail_Sub_KegiatanResponse

	for _, detailSubKegiatan := range detailSubKegiatans {
		var detailSubKegiatanResponse = helper.ConvertToDetailSubKegiatanResponse(detailSubKegiatan)
		detailSubKegiatansResponse = append(detailSubKegiatansResponse, detailSubKegiatanResponse)
	}

	cntx.JSON(http.StatusOK, detailSubKegiatansResponse)
}

func (c *detailSubKegiatanController) GetDetailSubKegiatan(cntx *gin.Context) {
	var idString = cntx.Param("id")
	var id, _ = strconv.Atoi(idString)

	var detailSubKegiatan, err = c.detailSubKegiatanService.FindById(id)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": "Data tidak ditemukan",
		})
		return
	}

	var detailSubKegiatanResponse = helper.ConvertToDetailSubKegiatanResponse(detailSubKegiatan)

	cntx.JSON(http.StatusOK, detailSubKegiatanResponse)
}

func (c *detailSubKegiatanController) GetDetailSubKegiatanWithRelation(cntx *gin.Context) {
	var detailSubKegiatanRelations, err = c.detailSubKegiatanService.FindAllRelation()
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
	}

	var detailSubKegiatansResponse []responses.Detail_Sub_KegiatanWithSub_KegiatanResponse

	for _, detailSubKegiatanRelation := range detailSubKegiatanRelations {
		var detailSubKegiatanRelationResponse = helper.ConvertToDetailSubKegiatanWithSubKegiatanResponse(detailSubKegiatanRelation)
		detailSubKegiatansResponse = append(detailSubKegiatansResponse, detailSubKegiatanRelationResponse)
	}

	cntx.JSON(http.StatusOK, detailSubKegiatansResponse)

}

func (c *detailSubKegiatanController) CreateDetailSubKegiatan(cntx *gin.Context) {
	var detailSubKegiatanRequest request.CreateDetail_Sub_KegiatanRequest

	var err = cntx.ShouldBindJSON(&detailSubKegiatanRequest)
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

	detailSubKegiatan, err := c.detailSubKegiatanService.Create(detailSubKegiatanRequest)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	var detailSubKegiatanResponse = helper.ConvertToDetailSubKegiatanResponse(detailSubKegiatan)

	cntx.JSON(http.StatusCreated, gin.H{
		"data": detailSubKegiatanResponse,
	})
}

func (c *detailSubKegiatanController) UpdateDetailSubKegiatan(cntx *gin.Context) {
	var detailSubKegiatanRequest request.UpdateDetail_Sub_KegiatanRequest

	var err = cntx.ShouldBindJSON(&detailSubKegiatanRequest)
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

	detailSubKegiatan, err := c.detailSubKegiatanService.Update(id, detailSubKegiatanRequest)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"errors": cntx.Error(err),
		})
		return
	}

	var detailSubKegiatanResponse = helper.ConvertToDetailSubKegiatanResponse(detailSubKegiatan)

	cntx.JSON(http.StatusOK, gin.H{
		"data": detailSubKegiatanResponse,
	})
}

func (c *detailSubKegiatanController) DeleteDetailSubKegitan(cntx *gin.Context) {
	var idString = cntx.Param("id")
	var id, _ = strconv.Atoi(idString)

	_, err := c.detailSubKegiatanService.Delete(id)
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
