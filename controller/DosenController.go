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

type dosenController struct {
	dosenService     service.DosenService
	mahasiswaService service.MahasiswaService
}

func NewDosenController(dosenService service.DosenService, mahasiswaService service.MahasiswaService) *dosenController {
	return &dosenController{
		dosenService,
		mahasiswaService,
	}
}

func (c *dosenController) GetDosens(cntx *gin.Context) {
	var dosens, err = c.dosenService.FindAll()
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	var dosensResponse []responses.DosenResponse

	for _, dosen := range dosens {
		var dosenResponse = helper.ConvertToDosenResponse(dosen)
		dosensResponse = append(dosensResponse, dosenResponse)
	}

	cntx.JSON(http.StatusOK, dosensResponse)
}

func (c *dosenController) GetDosen(cntx *gin.Context) {
	var idString = cntx.Param("id")
	var id, _ = strconv.Atoi(idString)

	var dosen, err = c.dosenService.FindById(id)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": "Data tidak ditemukan",
		})
		return
	}

	var dosenResponse = helper.ConvertToDosenResponse(dosen)

	cntx.JSON(http.StatusOK, dosenResponse)
}

func (c *dosenController) GetDosenWithRelation(cntx *gin.Context) {
	var dosens, err = c.dosenService.FindAllRelation()
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	var dosensResponse []responses.DosenWithRelationResponse

	for _, dosen := range dosens {
		var dosenResponse = helper.ConvertToDosenWithRelationResponse(dosen)
		dosensResponse = append(dosensResponse, dosenResponse)
	}

	cntx.JSON(http.StatusOK, dosensResponse)
}

func (c *dosenController) CreateDosen(cntx *gin.Context) {
	var dosenRequest request.CreateDosenRequest

	var err = cntx.ShouldBindJSON(&dosenRequest)
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

	dosen, err := c.dosenService.Create(dosenRequest)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	var dosenResponse = helper.ConvertToDosenResponse(dosen)

	cntx.JSON(http.StatusCreated, gin.H{
		"data": dosenResponse,
	})
}

func (c *dosenController) UpdateDosen(cntx *gin.Context) {
	var dosenRequest request.UpdateDosenRequest

	var err = cntx.ShouldBindJSON(&dosenRequest)
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

	dosen, err := c.dosenService.Update(id, dosenRequest)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"errors": cntx.Error(err),
		})
		return
	}

	var dosenResponse = helper.ConvertToDosenResponse(dosen)

	cntx.JSON(http.StatusOK, dosenResponse)
}

func (c *dosenController) DeleteDosen(cntx *gin.Context) {
	var idString = cntx.Param("id")
	var id, _ = strconv.Atoi(idString)

	_, err := c.dosenService.Delete(id)
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

func (c *dosenController) GetMahasiswa(cntx *gin.Context) {
	var kelurahanId = cntx.Query("kelurahanid")

	var mahasiswas, err = c.dosenService.FindMahasiswa(kelurahanId)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": "Data tidak ditemukan",
		})
		return
	}

	var mahasiswasResponse []responses.MahasiswaWithVerifiedCountResponse

	for _, mahasiswa := range mahasiswas {
		var jumlahVerifiedIndividu, _ = c.mahasiswaService.CountVerifiedIndividu(mahasiswa.Id, kelurahanId)
		var jumlahVerifiedKeluarga, _ = c.mahasiswaService.CountVerifiedKeluarga(mahasiswa.Id, kelurahanId)

		var mahasiswaResponse = helper.ConvertToMahasiswaWithVerifiedCountResponse(mahasiswa, jumlahVerifiedIndividu, jumlahVerifiedKeluarga)
		mahasiswasResponse = append(mahasiswasResponse, mahasiswaResponse)
	}

	cntx.JSON(http.StatusOK, mahasiswasResponse)
}
