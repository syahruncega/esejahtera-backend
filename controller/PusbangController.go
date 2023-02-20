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

type pusbangController struct {
	pusbangService  service.PusbangService
	dosenService    service.DosenService
	keluargaService service.KeluargaService
	individuService service.IndividuService
}

func NewPusbangController(pusbangService service.PusbangService, dosenService service.DosenService, keluargaService service.KeluargaService, individuService service.IndividuService) *pusbangController {
	return &pusbangController{
		pusbangService,
		dosenService,
		keluargaService,
		individuService,
	}
}

func (c *pusbangController) GetPusbangs(cntx *gin.Context) {
	var pusbangs, err = c.pusbangService.FindAll()
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	var pusbangsResponse []responses.PusbangResponse

	for _, pusbang := range pusbangs {
		var pusbangResponse = helper.ConvertToPusbangResponse(pusbang)
		pusbangsResponse = append(pusbangsResponse, pusbangResponse)
	}

	cntx.JSON(http.StatusOK, pusbangsResponse)
}

func (c *pusbangController) GetPusbang(cntx *gin.Context) {
	var idString = cntx.Param("id")
	var id, _ = strconv.Atoi(idString)

	var pusbang, err = c.pusbangService.FindById(id)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": "Data tidak ditemukan",
		})
		return
	}

	var pusbangResponse = helper.ConvertToPusbangResponse(pusbang)

	cntx.JSON(http.StatusOK, pusbangResponse)
}

func (c *pusbangController) GetPusbangWithRelation(cntx *gin.Context) {
	var pusbangs, err = c.pusbangService.FindAllRelation()
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	var pusbangsResponse []responses.PusbangWithRelationResponse

	for _, pusbang := range pusbangs {
		var pusbangResponse = helper.ConvertToPusbangWithRelationResponse(pusbang)
		pusbangsResponse = append(pusbangsResponse, pusbangResponse)
	}

	cntx.JSON(http.StatusOK, pusbangsResponse)
}

func (c *pusbangController) GetDistinctDosen(cntx *gin.Context) {
	var kelurahanId = cntx.Query("kelurahanid")
	var places = "kelurahanId"
	var statusVerified = 1

	var distinctDosen, err = c.dosenService.DistinctDosen(kelurahanId)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	jumlahKeluarga, err := c.keluargaService.CountJumlahKeluarga(places, kelurahanId)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	jumlahKeluargaVerified, err := c.keluargaService.CountVerified(places, kelurahanId, statusVerified)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	jumlahIndividu, err := c.individuService.CountJumlahIndividu(places, kelurahanId)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	jumlahIndividuVerified, err := c.individuService.CountVerified(places, kelurahanId, statusVerified)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	if len(distinctDosen) > 1 {
		var arrDosenId = []int{}

		for _, dosen := range distinctDosen {
			arrDosenId = append(arrDosenId, dosen.DosenId)
		}

		var dosens []model.Dosen

		for _, v := range arrDosenId {
			var dosen, _ = c.dosenService.FindById(v)
			dosens = append(dosens, dosen)
		}

		var dosensResponse []responses.DosenResponse

		for _, dosen := range dosens {
			var dosenResponse = helper.ConvertToDosenResponse(dosen)
			dosensResponse = append(dosensResponse, dosenResponse)
		}

		cntx.JSON(http.StatusOK, gin.H{
			"data": dosensResponse,
			"keluarga": gin.H{
				"jumlahKeluarga":         jumlahKeluarga,
				"jumlahKeluargaVerified": jumlahKeluargaVerified,
			},
			"individu": gin.H{
				"jumlahIndividu":         jumlahIndividu,
				"jumlahIndividuVerified": jumlahIndividuVerified,
			},
		})

	} else {
		var dosen, _ = c.dosenService.FindById(distinctDosen[0].DosenId)

		var dosenResponse = helper.ConvertToDosenResponse(dosen)

		cntx.JSON(http.StatusOK, gin.H{
			"data": dosenResponse,
			"keluarga": gin.H{
				"jumlahKeluarga":         jumlahKeluarga,
				"jumlahKeluargaVerified": jumlahKeluargaVerified,
			},
			"individu": gin.H{
				"jumlahIndividu":         jumlahIndividu,
				"jumlahIndividuVerified": jumlahIndividuVerified,
			},
		})
	}

}

func (c *pusbangController) GetDistinctLokasiDosen(cntx *gin.Context) {
	var dosenId = 0

	var distinctLokasiDosen, err = c.dosenService.DistinctLokasiDosen(dosenId)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	var arrLokasi = []string{}

	for _, lokasi := range distinctLokasiDosen {
		arrLokasi = append(arrLokasi, lokasi.KelurahanId)
	}

	cntx.JSON(http.StatusOK, arrLokasi)
}

func (c *pusbangController) CreatePusbang(cntx *gin.Context) {
	var pusbangRequest request.CreatePusbangRequest

	var err = cntx.ShouldBindJSON(&pusbangRequest)
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

	pusbang, err := c.pusbangService.Create(pusbangRequest)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	var pusbangResponse = helper.ConvertToPusbangResponse(pusbang)

	cntx.JSON(http.StatusCreated, pusbangResponse)
}

func (c *pusbangController) UpdatePusbang(cntx *gin.Context) {
	var pusbangRequest request.UpdatePusbangRequest

	var err = cntx.ShouldBindJSON(&pusbangRequest)
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

	pusbang, err := c.pusbangService.Update(id, pusbangRequest)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"errors": cntx.Error(err),
		})
		return
	}

	var pusbangResponse = helper.ConvertToPusbangResponse(pusbang)

	cntx.JSON(http.StatusOK, pusbangResponse)
}

func (c *pusbangController) DeletePusbang(cntx *gin.Context) {
	var idString = cntx.Param("id")
	var id, _ = strconv.Atoi(idString)

	_, err := c.pusbangService.Delete(id)
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
