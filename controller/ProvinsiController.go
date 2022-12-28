package controller

import (
	"kemiskinan/helper"
	"kemiskinan/responses"
	"kemiskinan/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type provinsiController struct {
	provinsiService service.ProvinsiService
}

func NewProvinsiController(provinsiService service.ProvinsiService) *provinsiController {
	return &provinsiController{provinsiService}
}

func (c *provinsiController) GetProvinsis(cntx *gin.Context) {
	var provinsis, err = c.provinsiService.FindAll()
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
	}

	var provinsisResponse []responses.ProvinsiResponse

	for _, provinsi := range provinsis {
		var provinsiResponse = helper.ConvertToProvinsiResponse(provinsi)
		provinsisResponse = append(provinsisResponse, provinsiResponse)
	}

	cntx.JSON(http.StatusOK, provinsisResponse)
}

func (c *provinsiController) GetProvinsi(cntx *gin.Context) {
	var id = cntx.Param("id")

	var provinsi, err = c.provinsiService.FindById(id)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": "Data tidak ditemukan",
		})
		return
	}

	var provinsiResponse = helper.ConvertToProvinsiResponse(provinsi)

	cntx.JSON(http.StatusOK, provinsiResponse)
}
