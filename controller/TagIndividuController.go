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

type tagIndividuController struct {
	tagIndividuService service.TagIndividuService
}

func NewTagIndividuController(tagIndividuService service.TagIndividuService) *tagIndividuController {
	return &tagIndividuController{tagIndividuService}
}

func (c *tagIndividuController) GetTagIndividus(cntx *gin.Context) {
	var tagIndividus, err = c.tagIndividuService.FindAll()
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	var tagIndividusResponse []responses.TagIndividuResponse

	for _, tagIndividu := range tagIndividus {
		var tagIndividuResponse = helper.ConvertToTagIndividuResponse(tagIndividu)
		tagIndividusResponse = append(tagIndividusResponse, tagIndividuResponse)
	}

	cntx.JSON(http.StatusOK, tagIndividusResponse)
}

func (c *tagIndividuController) GetTagIndividu(cntx *gin.Context) {
	var idString = cntx.Param("id")
	var id, _ = strconv.Atoi(idString)

	var tagIndividu, err = c.tagIndividuService.FindById(id)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": "Data tidak ditemukan",
		})
		return
	}

	var tagIndividuResponse = helper.ConvertToTagIndividuResponse(tagIndividu)

	cntx.JSON(http.StatusOK, tagIndividuResponse)
}

func (c *tagIndividuController) GetTagIndividuBySearch(cntx *gin.Context) {
	var whereClauseString = cntx.Request.URL.Query()
	var whereClauseInterface = make(map[string]interface{})

	for k, v := range whereClauseString {
		interfaceKey := k
		interfaceVal := v

		whereClauseInterface[interfaceKey] = interfaceVal
	}

	var tagIndividus, err = c.tagIndividuService.FindBySearch(whereClauseInterface)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
	}

	var tagIndividusResponse []responses.TagIndividuResponse

	for _, tagIndividu := range tagIndividus {
		var tagIndividuResponse = helper.ConvertToTagIndividuResponse(tagIndividu)
		tagIndividusResponse = append(tagIndividusResponse, tagIndividuResponse)
	}

	cntx.JSON(http.StatusOK, tagIndividusResponse)
}

func (c *tagIndividuController) CreateTagIndividu(cntx *gin.Context) {
	var tagIndividuRequest request.CreateTagIndividuRequest

	var err = cntx.ShouldBindJSON(&tagIndividuRequest)
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

	tagIndividu, err := c.tagIndividuService.Create(tagIndividuRequest)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	var tagIndividuResponse = helper.ConvertToTagIndividuResponse(tagIndividu)

	cntx.JSON(http.StatusOK, tagIndividuResponse)
}

func (c *tagIndividuController) UpdateTagIndividu(cntx *gin.Context) {
	var tagIndividuRequest request.UpdateTagIndividuRequest

	var err = cntx.ShouldBindJSON(&tagIndividuRequest)
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

	tagIndividu, err := c.tagIndividuService.Update(id, tagIndividuRequest)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	var tagIndividuResponse = helper.ConvertToTagIndividuResponse(tagIndividu)

	cntx.JSON(http.StatusOK, tagIndividuResponse)
}

func (c *tagIndividuController) DeleteTagIndividu(cntx *gin.Context) {
	var idString = cntx.Param("id")
	var id, _ = strconv.Atoi(idString)

	_, err := c.tagIndividuService.Delete(id)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": "Data gagal dihapus",
		})
		return
	}

	cntx.JSON(http.StatusOK, "Data berhasil dihapus")
}
