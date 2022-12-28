package controller

import (
	"kemiskinan/helper"
	"kemiskinan/responses"
	"kemiskinan/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type kabupatenKotaController struct {
	kabupatenKotaService service.KabupatenKotaService
}

func NewKabupatenKotaController(kabupatenKotaService service.KabupatenKotaService) *kabupatenKotaController {
	return &kabupatenKotaController{kabupatenKotaService}
}

func (c *kabupatenKotaController) GetKabupatenKotas(cntx *gin.Context) {
	var kabupatenKotas, err = c.kabupatenKotaService.FindAll()
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
	}

	var kabupatenKotasResponse []responses.Kabupaten_KotaResponse

	for _, kabupatenKota := range kabupatenKotas {
		var kabupatenKotaResponse = helper.ConvertToKabupatenKotaResponse(kabupatenKota)
		kabupatenKotasResponse = append(kabupatenKotasResponse, kabupatenKotaResponse)
	}

	cntx.JSON(http.StatusOK, kabupatenKotasResponse)
}

func (c *kabupatenKotaController) GetKabupatenKota(cntx *gin.Context) {
	var id = cntx.Param("id")

	var kabupatenKota, err = c.kabupatenKotaService.FindById(id)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": "Data tidak ditemukan",
		})
		return
	}

	var kabupatenKotaResponse = helper.ConvertToKabupatenKotaResponse(kabupatenKota)

	cntx.JSON(http.StatusOK, kabupatenKotaResponse)
}
