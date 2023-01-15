package controller

import (
	"kemiskinan/helper"
	"kemiskinan/responses"
	"kemiskinan/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type keluargaDonggalaController struct {
	keluargaDonggalaService service.KeluargaDonggalaService
}

func NewKeluargaDonggalaController(keluargaDonggalaService service.KeluargaDonggalaService) *keluargaDonggalaController {
	return &keluargaDonggalaController{keluargaDonggalaService}
}

func (c *keluargaDonggalaController) GetKeluargas(cntx *gin.Context) {

	var keluargaDonggalas, err = c.keluargaDonggalaService.FindAll()
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
	}

	var keluargaDonggalasResponse []responses.KeluargaResponse

	for _, keluarga := range keluargaDonggalas {
		var keluargaResponse = helper.ConvertToKeluargaDonggalaResponse(keluarga)
		keluargaDonggalasResponse = append(keluargaDonggalasResponse, keluargaResponse)
	}

	cntx.JSON(http.StatusOK, keluargaDonggalasResponse)
}

func (c *keluargaDonggalaController) GetKeluargaById(cntx *gin.Context) {
	var idString = cntx.Param("id")
	var id, _ = strconv.Atoi(idString)

	var keluargaDonggala, err = c.keluargaDonggalaService.FindById(id)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": "Data tidak ditemukan",
		})
		return
	}

	var keluargaDonggalaResponse = helper.ConvertToKeluargaDonggalaResponse(keluargaDonggala)

	cntx.JSON(http.StatusOK, keluargaDonggalaResponse)
}

func (c *keluargaDonggalaController) GetKeluargaByIdKeluarga(cntx *gin.Context) {
	var idKeluarga = cntx.Param("idkeluarga")

	var keluargaDonggala, err = c.keluargaDonggalaService.FindByIdKeluarga(idKeluarga)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": "Data tidak ditemukan",
		})
		return
	}

	var keluargaDonggalaResponse = helper.ConvertToKeluargaDonggalaResponse(keluargaDonggala)

	cntx.JSON(http.StatusOK, keluargaDonggalaResponse)

}
