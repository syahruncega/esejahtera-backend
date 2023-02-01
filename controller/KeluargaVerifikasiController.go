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

type keluargaVerifikasiController struct {
	keluargaVerifikasiService service.KeluargaVerifikasiService
}

func NewKeluargaVerifikasiController(keluargaVerifikasiService service.KeluargaVerifikasiService) *keluargaVerifikasiController {
	return &keluargaVerifikasiController{keluargaVerifikasiService}
}

func (c *keluargaVerifikasiController) GetKeluargaVerifikasis(cntx *gin.Context) {
	var keluargaVerifikasis, err = c.keluargaVerifikasiService.FindAll()
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
	}

	var keluargaVerifikasisResponse []responses.KeluargaVerifikasiResponse

	for _, keluargaVerifikasi := range keluargaVerifikasis {
		var keluargaVerifikasiResponse = helper.ConvertToKeluargaVerifikasiResponse(keluargaVerifikasi)
		keluargaVerifikasisResponse = append(keluargaVerifikasisResponse, keluargaVerifikasiResponse)
	}

	cntx.JSON(http.StatusOK, keluargaVerifikasisResponse)
}

func (c *keluargaVerifikasiController) GetKeluargaVerifikasi(cntx *gin.Context) {
	var idString = cntx.Param("id")
	var id, _ = strconv.Atoi(idString)

	var keluargaVerifikasi, err = c.keluargaVerifikasiService.FindById(id)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": "Data tidak ditemukan",
		})
		return
	}

	var keluargaVerifikasiResponse = helper.ConvertToKeluargaVerifikasiResponse(keluargaVerifikasi)

	cntx.JSON(http.StatusOK, keluargaVerifikasiResponse)
}

func (c *keluargaVerifikasiController) GetKeluargaVerifikasiByIdKeluarga(cntx *gin.Context) {
	var idKeluarga = cntx.Param("idkeluarga")

	var keluargaVerifikasi, err = c.keluargaVerifikasiService.FindByIdKeluarga(idKeluarga)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": "Data tidak ditemukan",
		})
		return
	}

	var keluargaVerifikasiResponse = helper.ConvertToKeluargaVerifikasiResponse(keluargaVerifikasi)

	cntx.JSON(http.StatusOK, keluargaVerifikasiResponse)
}

func (c *keluargaVerifikasiController) GetKeluargaVerifikasiBySearch(cntx *gin.Context) {
	var whereClauseString = cntx.Request.URL.Query()
	var whereClauseInterface = make(map[string]interface{})

	for k, v := range whereClauseString {
		interfaceKey := k
		interfaceValue := v

		whereClauseInterface[interfaceKey] = interfaceValue
	}

	var keluargaVerifikasis, err = c.keluargaVerifikasiService.FindBySearch(whereClauseInterface)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
	}

	var keluargaVerifikasisResponse []responses.KeluargaVerifikasiResponse

	for _, keluargaVerifikasi := range keluargaVerifikasis {
		var keluargaVerifikasiResponse = helper.ConvertToKeluargaVerifikasiResponse(keluargaVerifikasi)
		keluargaVerifikasisResponse = append(keluargaVerifikasisResponse, keluargaVerifikasiResponse)
	}

	cntx.JSON(http.StatusOK, keluargaVerifikasisResponse)

}

func (c *keluargaVerifikasiController) CreateKeluargaVerifikasi(cntx *gin.Context) {
	var keluargaVerifikasiRequest request.CreateKeluargaVerifikasiRequest

	var err = cntx.ShouldBindJSON(&keluargaVerifikasiRequest)
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

	keluargaVerifikasi, err := c.keluargaVerifikasiService.Create(keluargaVerifikasiRequest)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"errors": cntx.Error(err),
		})
		return
	}

	var keluargaVerifikasiResponse = helper.ConvertToKeluargaVerifikasiResponse(keluargaVerifikasi)

	cntx.JSON(http.StatusCreated, keluargaVerifikasiResponse)
}

func (c *keluargaVerifikasiController) UpdateKeluargaVerifikasi(cntx *gin.Context) {
	var keluargaVerifikasiRequest request.UpdateKeluargaVerifikasiRequest

	var err = cntx.ShouldBindJSON(&keluargaVerifikasiRequest)
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

	keluargaVerifikasi, err := c.keluargaVerifikasiService.Update(id, keluargaVerifikasiRequest)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"errors": cntx.Error(err),
		})
		return
	}

	var keluargaVerifikasiResponse = helper.ConvertToKeluargaVerifikasiResponse(keluargaVerifikasi)

	cntx.JSON(http.StatusOK, keluargaVerifikasiResponse)
}

func (c *keluargaVerifikasiController) DeleteKeluargaVerifikasi(cntx *gin.Context) {
	var idString = cntx.Param("id")
	var id, _ = strconv.Atoi(idString)

	_, err := c.keluargaVerifikasiService.Delete(id)
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
