package controller

import (
	"kemiskinan/helper"
	"kemiskinan/responses"
	"kemiskinan/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type kecamatanController struct {
	kecamatanService service.KecamatanService
}

func NewKecamatanController(kecamatanService service.KecamatanService) *kecamatanController {
	return &kecamatanController{kecamatanService}
}

func (c *kecamatanController) GetKecamatans(cntx *gin.Context) {
	var kabupatenKotaId = cntx.Query("kabupatenkotaid")

	var kecamatans, err = c.kecamatanService.FindByKabupatenKota(kabupatenKotaId)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
	}

	var kecamatansResponse []responses.KecamatanResponse

	for _, kecamatan := range kecamatans {
		var kecamatanResponse = helper.ConvertToKecamatanResponse(kecamatan)
		kecamatansResponse = append(kecamatansResponse, kecamatanResponse)
	}

	cntx.JSON(http.StatusOK, kecamatansResponse)
}

func (c *kecamatanController) GetKecamatan(cntx *gin.Context) {
	var id = cntx.Param("id")

	var kecamatan, err = c.kecamatanService.FindById(id)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": "Data tidak ditemukan",
		})
		return
	}

	var kecamatanResponse = helper.ConvertToKecamatanResponse(kecamatan)

	cntx.JSON(http.StatusOK, kecamatanResponse)
}
