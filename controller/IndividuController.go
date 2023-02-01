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

type individuController struct {
	individuService service.IndividuService
}

func NewIndividuController(individuService service.IndividuService) *individuController {
	return &individuController{individuService}
}

func (c *individuController) GetIndividus(cntx *gin.Context) {
	var individus, err = c.individuService.FindAll()
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
	}

	var individusResponse []responses.IndividuResponse

	for _, individu := range individus {
		var individuResponse = helper.ConvertToIndividuResponse(individu)
		individusResponse = append(individusResponse, individuResponse)
	}

	cntx.JSON(http.StatusOK, individusResponse)
}

func (c *individuController) GetIndividu(cntx *gin.Context) {
	var idString = cntx.Param("id")
	var id, _ = strconv.Atoi(idString)

	var individu, err = c.individuService.FindById(id)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": "Data tidak ditemukan",
		})
		return
	}

	var individuResponse = helper.ConvertToIndividuResponse(individu)

	cntx.JSON(http.StatusOK, individuResponse)
}

func (c *individuController) GetIndividuByIdKeluarga(cntx *gin.Context) {
	var idKeluarga = cntx.Param("idkeluarga")

	var individus, err = c.individuService.FindByIdKeluarga(idKeluarga)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
	}

	var individusResponse []responses.IndividuResponse

	for _, individu := range individus {
		var individuResponse = helper.ConvertToIndividuResponse(individu)
		individusResponse = append(individusResponse, individuResponse)
	}

	cntx.JSON(http.StatusOK, individusResponse)
}

func (c *individuController) GetIndividuBySearch(cntx *gin.Context) {
	var whereClauseString = cntx.Request.URL.Query()
	var whereClauseInterface = make(map[string]interface{})

	for k, v := range whereClauseString {
		interfaceKey := k
		interfaceValue := v

		whereClauseInterface[interfaceKey] = interfaceValue
	}

	var individus, err = c.individuService.FindBySearch(whereClauseInterface)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
	}

	var individusResponse []responses.IndividuResponse

	for _, individu := range individus {
		var individuResponse = helper.ConvertToIndividuResponse(individu)
		individusResponse = append(individusResponse, individuResponse)
	}

	cntx.JSON(http.StatusOK, individusResponse)
}

func (c *individuController) UpdateIndividu(cntx *gin.Context) {
	var individuRequest request.UpdateIndividuRequest

	var err = cntx.ShouldBindJSON(&individuRequest)
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

	individu, err := c.individuService.Update(id, individuRequest)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"errors": cntx.Error(err),
		})
		return
	}

	var individuResponse = helper.ConvertToIndividuResponse(individu)

	cntx.JSON(http.StatusOK, individuResponse)
}
