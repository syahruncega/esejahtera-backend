package controller

import (
	"kemiskinan/helper"
	"kemiskinan/responses"
	"kemiskinan/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type kelurahanController struct {
	kelurahanService service.KelurahanService
}

func NewKelurahanController(kelurahanService service.KelurahanService) *kelurahanController {
	return &kelurahanController{kelurahanService}
}

func (c *kelurahanController) GetKelurahans(cntx *gin.Context) {
	var kecamatanId = cntx.Query("kecamatanid")

	var kelurahans, err = c.kelurahanService.FindByKecamatanId(kecamatanId)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
	}

	var kelurahansResponse []responses.KelurahanResponse

	for _, kelurahan := range kelurahans {
		var kelurahanResponse = helper.ConvertToKelurahanResponse(kelurahan)
		kelurahansResponse = append(kelurahansResponse, kelurahanResponse)
	}

	cntx.JSON(http.StatusOK, kelurahansResponse)
}

func (c *kelurahanController) GetKelurahan(cntx *gin.Context) {
	var id = cntx.Param("id")

	var kelurahan, err = c.kelurahanService.FindById(id)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": "Data tidak ditemukan",
		})
		return
	}

	var kelurahanResponse = helper.ConvertToKelurahanResponse(kelurahan)

	cntx.JSON(http.StatusOK, kelurahanResponse)
}
