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

type individuVerifikasiController struct {
	individuVerifikasiService service.IndividuVerifikasiService
}

func NewIndividuVerifikasiController(individuVerifikasiService service.IndividuVerifikasiService) *individuVerifikasiController {
	return &individuVerifikasiController{individuVerifikasiService}
}

func (c *individuVerifikasiController) GetIndividuVerifikasis(cntx *gin.Context) {
	var individuVerifikasis, err = c.individuVerifikasiService.FindAll()
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
	}

	var individuVerifikasisResponse []responses.IndividuVerifikasiResponse

	for _, individuVerifikasi := range individuVerifikasis {
		var individuVerifikasiResponse = helper.ConvertToIndividuVerifikasiResponse(individuVerifikasi)
		individuVerifikasisResponse = append(individuVerifikasisResponse, individuVerifikasiResponse)
	}

	cntx.JSON(http.StatusOK, individuVerifikasisResponse)
}

func (c *individuVerifikasiController) GetIndividuVerifikasi(cntx *gin.Context) {
	var idString = cntx.Param("id")
	var id, _ = strconv.Atoi(idString)

	var individuVerifikasi, err = c.individuVerifikasiService.FindById(id)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": "Data tidak ditemukan",
		})
		return
	}

	var individuVerifikasiResponse = helper.ConvertToIndividuVerifikasiResponse(individuVerifikasi)

	cntx.JSON(http.StatusOK, individuVerifikasiResponse)
}

func (c *individuVerifikasiController) GetIndividuVerifikasiByIdKeluarga(cntx *gin.Context) {
	var idKeluarga = cntx.Param("idkeluarga")

	var individuVerifikasis, err = c.individuVerifikasiService.FindByIdKeluarga(idKeluarga)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
	}

	var individuVerifikasisResponse []responses.IndividuVerifikasiResponse

	for _, individuVerifikasi := range individuVerifikasis {
		var individuVerifikasiResponse = helper.ConvertToIndividuVerifikasiResponse(individuVerifikasi)
		individuVerifikasisResponse = append(individuVerifikasisResponse, individuVerifikasiResponse)
	}

	cntx.JSON(http.StatusOK, individuVerifikasisResponse)
}

func (c *individuVerifikasiController) GetIndividuVerifikasiBySearch(cntx *gin.Context) {
	var whereClauseString = cntx.Request.URL.Query()
	var whereClauseInterface = make(map[string]interface{})

	for k, v := range whereClauseString {
		interfaceKey := k
		interfaceValue := v

		whereClauseInterface[interfaceKey] = interfaceValue
	}

	var individuVerifikasis, err = c.individuVerifikasiService.FindBySearch(whereClauseInterface)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
	}

	var individuVerifikasisResponse []responses.IndividuVerifikasiResponse

	for _, individuVerifikasi := range individuVerifikasis {
		var individuVerifikasiResponse = helper.ConvertToIndividuVerifikasiResponse(individuVerifikasi)
		individuVerifikasisResponse = append(individuVerifikasisResponse, individuVerifikasiResponse)
	}

	cntx.JSON(http.StatusOK, individuVerifikasisResponse)
}

func (c *individuVerifikasiController) CreateIndividuVerifikasi(cntx *gin.Context) {
	var individuVerifikasiRequest request.CreateIndividuVerifikasiRequest

	var err = cntx.ShouldBindJSON(&individuVerifikasiRequest)
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

	individuVerifikasi, err := c.individuVerifikasiService.Create(individuVerifikasiRequest)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"errors": cntx.Error(err),
		})
		return
	}

	var individuVerifikasiResponse = helper.ConvertToIndividuVerifikasiResponse(individuVerifikasi)

	cntx.JSON(http.StatusCreated, individuVerifikasiResponse)
}

func (c *individuVerifikasiController) UpdateIndividuVerifikasi(cntx *gin.Context) {
	var individuVerifikasiRequest request.UpdateIndividuVerifikasiRequest

	var err = cntx.ShouldBindJSON(&individuVerifikasiRequest)
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

	individuVerifikasi, err := c.individuVerifikasiService.Update(id, individuVerifikasiRequest)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"errors": cntx.Error(err),
		})
		return
	}

	var individuVerifikasiResponse = helper.ConvertToIndividuVerifikasiResponse(individuVerifikasi)

	cntx.JSON(http.StatusOK, individuVerifikasiResponse)
}

func (c *individuVerifikasiController) DeleteIndividuVerifikasi(cntx *gin.Context) {
	var idString = cntx.Param("id")
	var id, _ = strconv.Atoi(idString)

	_, err := c.individuVerifikasiService.Delete(id)
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
