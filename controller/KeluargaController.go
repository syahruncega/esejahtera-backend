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

type keluargaController struct {
	keluargaService service.KeluargaService
}

func NewKeluargaController(keluargaService service.KeluargaService) *keluargaController {
	return &keluargaController{keluargaService}
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
	var whereClauseString = cntx.Request.URL.Query()
	var whereClauseInterface = make(map[string]interface{})

	for k, v := range whereClauseString {
		interfaceKey := k
		interfaceVal := v

		whereClauseInterface[interfaceKey] = interfaceVal
	}

	var keluargas, err = c.keluargaService.FindBySearch(whereClauseInterface)
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

// func (c *keluargaController) GetKeluargaBySearch(cntx *gin.Context) {
// 	var whereClause = cntx.Request.URL.Query()

// 	value := whereClause["test"]
// 	var testint []int
// 	for _, val := range value {
// 		var testtis, _ = strconv.Atoi(val)
// 		testint = append(testint, testtis)
// 	}

// 	// var keluargas, err = c.keluargaService.FindBySearch(whereClause)
// 	cntx.JSON(http.StatusOK, testint)
// }

func (c *keluargaController) CountPenerimaByKabupatenKota(cntx *gin.Context) {
	var kabupatenKotaId = cntx.Param("kabupatenkotaid")
	var penerimaParameter = cntx.Param("penerimaparameter")
	var nilai = cntx.Param("nilai")

	var jumlah, err = c.keluargaService.CountPenerimaByKabupatenKota(kabupatenKotaId, penerimaParameter, nilai)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	var jumlahResponse = helper.ConvertToJumlahResponse(jumlah)

	cntx.JSON(http.StatusOK, jumlahResponse)
}

func (c *keluargaController) CountPenerimaByKecamatan(cntx *gin.Context) {
	var kecamatanId = cntx.Param("kecamatanid")
	var penerimaParameter = cntx.Param("penerimaparameter")
	var nilai = cntx.Param("nilai")

	var jumlah, err = c.keluargaService.CountPenerimaByKecamatan(kecamatanId, penerimaParameter, nilai)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	var jumlahResponse = helper.ConvertToJumlahResponse(jumlah)

	cntx.JSON(http.StatusOK, jumlahResponse)

}

func (c *keluargaController) CountPenerimaByKelurahan(cntx *gin.Context) {
	var kelurahanId = cntx.Param("kelurahanid")
	var penerimaParameter = cntx.Param("penerimaparameter")
	var nilai = cntx.Param("nilai")

	var jumlah, err = c.keluargaService.CountPenerimaByKelurahan(kelurahanId, penerimaParameter, nilai)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	var jumlahResponse = helper.ConvertToJumlahResponse(jumlah)

	cntx.JSON(http.StatusOK, jumlahResponse)
}

func (c *keluargaController) CountDesilByKabupatenKota(cntx *gin.Context) {
	var kabupatenKotaId = cntx.Param("kabupatenkotaid")
	var nilaiDesil = cntx.Param("nilaidesil")

	var jumlah, err = c.keluargaService.CountDesilByKabupatenKota(kabupatenKotaId, nilaiDesil)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	var jumlahResponse = helper.ConvertToJumlahResponse(jumlah)

	cntx.JSON(http.StatusOK, jumlahResponse)
}

func (c *keluargaController) CountDesilByKecamatan(cntx *gin.Context) {
	var kecamatanId = cntx.Param("kecamatanid")
	var nilaiDesil = cntx.Param("nilaidesil")

	var jumlah, err = c.keluargaService.CountDesilByKecamatan(kecamatanId, nilaiDesil)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	var jumlahResponse = helper.ConvertToJumlahResponse(jumlah)

	cntx.JSON(http.StatusOK, jumlahResponse)
}

func (c *keluargaController) CountDesilByKelurahan(cntx *gin.Context) {
	var kelurahanId = cntx.Param("kelurahanid")
	var nilaiDesil = cntx.Param("nilaidesil")

	var jumlah, err = c.keluargaService.CountDesilByKelurahan(kelurahanId, nilaiDesil)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	var jumlahResponse = helper.ConvertToJumlahResponse(jumlah)

	cntx.JSON(http.StatusOK, jumlahResponse)
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
