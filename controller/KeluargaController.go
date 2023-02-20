package controller

import (
	"fmt"
	"kemiskinan/helper"
	"kemiskinan/request"
	"kemiskinan/responses"
	"kemiskinan/service"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type keluargaController struct {
	keluargaService service.KeluargaService
	individuService service.IndividuService
}

func NewKeluargaController(keluargaService service.KeluargaService, individuService service.IndividuService) *keluargaController {
	return &keluargaController{
		keluargaService,
		individuService,
	}
}

func (c *keluargaController) GetKeluargas(cntx *gin.Context) {
	var keluargas, err = c.keluargaService.FindAll()
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
	}

	var keluargasResponse []responses.KeluargaResponse

	for _, keluarga := range keluargas {
		var keluargaResponse = helper.ConvertToKeluargaResponse(keluarga)
		keluargasResponse = append(keluargasResponse, keluargaResponse)
	}

	cntx.JSON(http.StatusOK, keluargasResponse)
}

func (c *keluargaController) GetKeluargaById(cntx *gin.Context) {
	var idString = cntx.Param("id")
	var id, _ = strconv.Atoi(idString)

	var keluarga, err = c.keluargaService.FindById(id)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": "Data tidak ditemukan",
		})
		return
	}

	var keluargaResponse = helper.ConvertToKeluargaResponse(keluarga)

	cntx.JSON(http.StatusOK, keluargaResponse)
}

func (c *keluargaController) GetKeluargaByIdKeluarga(cntx *gin.Context) {
	var idKeluarga = cntx.Param("idkeluarga")

	var keluarga, err = c.keluargaService.FindByIdKeluarga(idKeluarga)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": "Data tidak ditemukan",
		})
		return
	}

	var keluargaResponse = helper.ConvertToKeluargaResponse(keluarga)

	cntx.JSON(http.StatusOK, keluargaResponse)
}

func (c *keluargaController) GetKeluargaBySearch(cntx *gin.Context) {
	var pageRowString = cntx.Query("pagerow")
	var limit, _ = strconv.Atoi(pageRowString)
	var pageRowFloat, _ = strconv.ParseFloat(pageRowString, 64)

	var halamanString = cntx.Query("halaman")
	var halaman, _ = strconv.Atoi(halamanString)

	var whereClauseString = cntx.Request.URL.Query()
	var whereClauseInterface = make(map[string]interface{})

	for k, v := range whereClauseString {
		if k == "halaman" || k == "pagerow" {
			continue
		}

		interfaceKey := k
		interfaceVal := v

		whereClauseInterface[interfaceKey] = interfaceVal
	}

	var totalData, err = c.keluargaService.CountData(whereClauseInterface)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
	}

	var offset = limit * (halaman - 1)

	var hasilBagi = float64(totalData) / float64(pageRowFloat)
	var jumlahHalaman = int(math.Ceil(hasilBagi))

	keluargas, err := c.keluargaService.FindBySearch(whereClauseInterface, limit, offset)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
	}

	var keluargasResponse []responses.KeluargaResponse

	for _, keluarga := range keluargas {
		var keluargaResponse = helper.ConvertToKeluargaResponse(keluarga)
		keluargasResponse = append(keluargasResponse, keluargaResponse)
	}

	cntx.JSON(http.StatusOK, gin.H{
		"data":          keluargasResponse,
		"totalData":     totalData,
		"jumlahHalaman": jumlahHalaman,
	})
}

func (c *keluargaController) UpdateKeluarga(cntx *gin.Context) {
	var keluargaRequest request.UpdateKeluargaRequest

	var err = cntx.ShouldBindJSON(&keluargaRequest)
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

	keluarga, err := c.keluargaService.Update(id, keluargaRequest)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"errors": cntx.Error(err),
		})
		return
	}

	var keluargaResponse = helper.ConvertToKeluargaResponse(keluarga)

	cntx.JSON(http.StatusOK, keluargaResponse)
}

// func (c *keluargaController) TestDistinct(cntx *gin.Context) {
// 	var distinctsKelurahan, err = c.keluargaService.DistinctKelurahan("7205140")
// 	if err != nil {
// 		cntx.JSON(http.StatusBadRequest, gin.H{
// 			"error": cntx.Error(err),
// 		})
// 		return
// 	}

// 	var distinctsKelurahanResponse []responses.DistinctKelurahan

// 	for _, distinctKelurahan := range distinctsKelurahan {
// 		var distinctKelurahanResponse = helper.ConvertToDistinctKelurahanResponse(distinctKelurahan)
// 		distinctsKelurahanResponse = append(distinctsKelurahanResponse, distinctKelurahanResponse)
// 	}

// 	cntx.JSON(http.StatusOK, distinctsKelurahanResponse)

// }
