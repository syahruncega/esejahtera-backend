package controller

import (
	"kemiskinan/helper"
	"kemiskinan/responses"
	"kemiskinan/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type keluargaSigiController struct {
	keluargaSigiService service.KeluargaSigiService
}

func NewKeluargaSigiController(keluargaSigiService service.KeluargaSigiService) *keluargaSigiController {
	return &keluargaSigiController{keluargaSigiService}
}

func (c *keluargaSigiController) GetKeluargas(cntx *gin.Context) {

	var keluargaSigis, err = c.keluargaSigiService.FindAll()
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
	}

	var keluargaSigisResponse []responses.KeluargaResponse

	for _, keluargaSigi := range keluargaSigis {
		var keluargaResponse = helper.ConvertToKeluargaSigiResponse(keluargaSigi)
		keluargaSigisResponse = append(keluargaSigisResponse, keluargaResponse)
	}

	cntx.JSON(http.StatusOK, keluargaSigisResponse)
}

func (c *keluargaSigiController) GetKeluargaById(cntx *gin.Context) {
	var idString = cntx.Param("id")
	var id, _ = strconv.Atoi(idString)

	var keluargaSigi, err = c.keluargaSigiService.FindById(id)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": "Data tidak ditemukan",
		})
		return
	}

	var keluargaSigiResponse = helper.ConvertToKeluargaSigiResponse(keluargaSigi)

	cntx.JSON(http.StatusOK, keluargaSigiResponse)
}

func (c *keluargaSigiController) GetKeluargaByIdKeluarga(cntx *gin.Context) {
	var idKeluarga = cntx.Param("idkeluarga")

	var keluargaSigi, err = c.keluargaSigiService.FindByIdKeluarga(idKeluarga)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": "Data tidak ditemukan",
		})
		return
	}

	var keluargaSigiResponse = helper.ConvertToKeluargaSigiResponse(keluargaSigi)

	cntx.JSON(http.StatusOK, keluargaSigiResponse)

}
