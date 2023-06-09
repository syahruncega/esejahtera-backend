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

type tagKeluargaController struct {
	tagKeluargaService service.TagKeluargaService
}

func NewTagKeluargaController(tagKeluargaService service.TagKeluargaService) *tagKeluargaController {
	return &tagKeluargaController{tagKeluargaService}
}

func (c *tagKeluargaController) GetTagKeluargas(cntx *gin.Context) {
	var tagKeluargas, err = c.tagKeluargaService.FindAll()
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	var tagKeluargasResponse []responses.TagKeluargaResponse

	for _, tagKeluarga := range tagKeluargas {
		var tagKeluargaResponse = helper.ConvertToTagKeluargaResponse(tagKeluarga)
		tagKeluargasResponse = append(tagKeluargasResponse, tagKeluargaResponse)
	}

	cntx.JSON(http.StatusOK, tagKeluargasResponse)
}

func (c *tagKeluargaController) GetTagKeluarga(cntx *gin.Context) {
	var idString = cntx.Param("id")
	var id, _ = strconv.Atoi(idString)

	var tagKeluarga, err = c.tagKeluargaService.FindById(id)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": "Data tidak ditemukan",
		})
		return
	}

	var tagKeluargaResponse = helper.ConvertToTagKeluargaResponse(tagKeluarga)

	cntx.JSON(http.StatusOK, tagKeluargaResponse)
}

func (c *tagKeluargaController) GetTagKeluargaBySearch(cntx *gin.Context) {
	var whereClauseString = cntx.Request.URL.Query()
	var whereClauseInterface = make(map[string]interface{})

	for k, v := range whereClauseString {
		interfaceKey := k
		interfaceVal := v

		whereClauseInterface[interfaceKey] = interfaceVal
	}

	var tagKeluargas, err = c.tagKeluargaService.FindBySearch(whereClauseInterface)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
	}

	var tagKeluargasResponse []responses.TagKeluargaResponse

	for _, tagKeluarga := range tagKeluargas {
		var tagKeluargaResponse = helper.ConvertToTagKeluargaResponse(tagKeluarga)
		tagKeluargasResponse = append(tagKeluargasResponse, tagKeluargaResponse)
	}

	cntx.JSON(http.StatusOK, tagKeluargasResponse)
}

func (c *tagKeluargaController) CreateTagKeluarga(cntx *gin.Context) {
	var tagKeluargaRequest request.CreateTagKeluargaRequest

	var err = cntx.ShouldBindJSON(&tagKeluargaRequest)
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

	tagKeluarga, err := c.tagKeluargaService.Create(tagKeluargaRequest)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	var tagKeluargaResponse = helper.ConvertToTagKeluargaResponse(tagKeluarga)

	cntx.JSON(http.StatusOK, tagKeluargaResponse)
}

func (c *tagKeluargaController) UpdateTagKeluarga(cntx *gin.Context) {
	var tagKeluargaRequest request.UpdateTagKeluargaRequest

	var err = cntx.ShouldBindJSON(&tagKeluargaRequest)
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

	tagKeluarga, err := c.tagKeluargaService.Update(id, tagKeluargaRequest)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	var tagKeluargaResponse = helper.ConvertToTagKeluargaResponse(tagKeluarga)

	cntx.JSON(http.StatusOK, tagKeluargaResponse)
}

func (c *tagKeluargaController) DeleteTagKeluarga(cntx *gin.Context) {
	var idString = cntx.Param("idString")
	var id, _ = strconv.Atoi(idString)

	_, err := c.tagKeluargaService.Delete(id)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": "Data gagal dihapus",
		})
		return
	}

	cntx.JSON(http.StatusOK, "Data berhasil dihapus")
}
