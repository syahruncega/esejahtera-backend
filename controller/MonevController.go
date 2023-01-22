package controller

import (
	"kemiskinan/helper"
	"kemiskinan/responses"
	"kemiskinan/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type monevController struct {
	monevService service.MonevService
}

func NewMonevController(monevService service.MonevService) *monevController {
	return &monevController{monevService}
}

func (c *monevController) GetMonevs(cntx *gin.Context) {
	var kabupatenKotaId = cntx.Param("kabupatenkotaid")

	var monevs, err = c.monevService.FindAll(kabupatenKotaId)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
	}

	var monevsResponse []responses.MonevResponse

	for _, monev := range monevs {
		var monevResponse = helper.ConvertToMonevResponse(monev)
		monevsResponse = append(monevsResponse, monevResponse)
	}

	cntx.JSON(http.StatusOK, monevsResponse)
}

func (c *monevController) GetMonev(cntx *gin.Context) {
	var kabupatenKotaId = cntx.Param("kabupatenkotaid")
	var idString = cntx.Param("id")
	var id, _ = strconv.Atoi(idString)

	var monev, err = c.monevService.FindById(kabupatenKotaId, id)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": "Data tidak ditemukan",
		})
	}

	var monevResponse = helper.ConvertToMonevResponse(monev)

	cntx.JSON(http.StatusOK, monevResponse)
}
