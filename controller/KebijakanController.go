package controller

import (
	"fmt"
	"kemiskinan/helper"
	"kemiskinan/model"
	"kemiskinan/request"
	"kemiskinan/responses"
	"kemiskinan/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type kebijakanController struct {
	kebijakanService service.KebijakanService
}

func NewKebijakanController(kebijakanService service.KebijakanService) *kebijakanController {
	return &kebijakanController{kebijakanService}
}

func (c *kebijakanController) GetKebijakans(cntx *gin.Context) {
	var programIdString = cntx.Query("programid")
	var programId, _ = strconv.Atoi(programIdString)

	var kebijakans []model.Kebijakan
	var err error

	if programIdString != "" {
		kebijakans, err = c.kebijakanService.FindByProgramId(programId)
	} else {
		kebijakans, err = c.kebijakanService.FindAll()
	}

	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
	}

	var kebijakansResponse []responses.KebijakanResponse

	for _, kebijakan := range kebijakans {
		var kebijakanResponse = helper.ConvertToKebijakanResponse(kebijakan)
		kebijakansResponse = append(kebijakansResponse, kebijakanResponse)
	}

	cntx.JSON(http.StatusOK, kebijakansResponse)
}

func (c *kebijakanController) GetKebijakan(cntx *gin.Context) {
	var idString = cntx.Param("id")
	var id, _ = strconv.Atoi(idString)

	var kebijakan, err = c.kebijakanService.FindById(id)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": "Data tidak ditemukan",
		})
		return
	}

	var kebijakanResponse = helper.ConvertToKebijakanResponse(kebijakan)

	cntx.JSON(http.StatusOK, kebijakanResponse)
}

func (c *kebijakanController) CreateKebijakan(cntx *gin.Context) {
	var kebijakanRequest request.CreateKebijakanRequest

	var err = cntx.ShouldBindJSON(&kebijakanRequest)
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

	kebijakan, err := c.kebijakanService.Create(kebijakanRequest)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	var kebijakanResponse = helper.ConvertToKebijakanResponse(kebijakan)

	cntx.JSON(http.StatusOK, kebijakanResponse)
}

func (c *kebijakanController) UpdateKebijakan(cntx *gin.Context) {
	var kebijakanRequest request.UpdateKebijakanRequest

	var err = cntx.ShouldBindJSON(&kebijakanRequest)
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

	kebijakan, err := c.kebijakanService.Update(id, kebijakanRequest)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	var kebijakanResponse = helper.ConvertToKebijakanResponse(kebijakan)

	cntx.JSON(http.StatusOK, kebijakanResponse)
}

func (c *kebijakanController) DeleteKebijakan(cntx *gin.Context) {
	var idString = cntx.Param("id")
	var id, _ = strconv.Atoi(idString)

	_, err := c.kebijakanService.Delete(id)
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
