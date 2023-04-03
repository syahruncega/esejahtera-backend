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

type mahasiswaController struct {
	mahasiswaService service.MahasiswaService
	keluargaService  service.KeluargaService
	individuService  service.IndividuService
}

func NewMahasiswaController(mahasiswaService service.MahasiswaService, keluargaService service.KeluargaService, individuService service.IndividuService) *mahasiswaController {
	return &mahasiswaController{
		mahasiswaService,
		keluargaService,
		individuService,
	}
}

func (c *mahasiswaController) GetMahasiswas(cntx *gin.Context) {
	var mahasiswas, err = c.mahasiswaService.FindAll()
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	var mahasiswasResponse []responses.MahasiswaResponse

	for _, mahasiswa := range mahasiswas {
		var mahasiswaResponse = helper.ConvertToMahasiswaResponse(mahasiswa)
		mahasiswasResponse = append(mahasiswasResponse, mahasiswaResponse)
	}

	cntx.JSON(http.StatusOK, mahasiswasResponse)
}

func (c *mahasiswaController) GetMahasiswa(cntx *gin.Context) {
	var idString = cntx.Param("id")
	var id, _ = strconv.Atoi(idString)

	var mahasiswa, err = c.mahasiswaService.FindById(id)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": "Data tidak ditemukan",
		})
		return
	}

	var mahasiswaResponse = helper.ConvertToMahasiswaResponse(mahasiswa)

	cntx.JSON(http.StatusOK, mahasiswaResponse)
}

func (c *mahasiswaController) GetMahasiswaWithRelation(cntx *gin.Context) {
	var mahasiswas, err = c.mahasiswaService.FindAllRelation()
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	var mahasiswasResponse []responses.MahasiswaWithRelationResponse

	for _, mahasiswa := range mahasiswas {
		var mahasiswaResponse = helper.ConvertToMahasiswaWithRelationResponse(mahasiswa)
		mahasiswasResponse = append(mahasiswasResponse, mahasiswaResponse)
	}

	cntx.JSON(http.StatusOK, mahasiswasResponse)
}

func (c *mahasiswaController) GetVerifiedByMahasiswa(cntx *gin.Context) {
	var mahasiswaIdString = cntx.Query("mahasiswaid")
	var mahasiswaId, _ = strconv.Atoi(mahasiswaIdString)

	var mahasiswa, err = c.mahasiswaService.FindById(mahasiswaId)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": "Data tidak ditemukan",
		})
		return
	}

	var mahasiswaResponse = helper.ConvertToMahasiswaResponse(mahasiswa)

	jumlahKeluargaVerified, err := c.keluargaService.CountVerifiedByMahasiswa(mahasiswaId)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	jumlahIndividuVerified, err := c.individuService.CountVerifiedByMahasiswa(mahasiswaId)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	cntx.JSON(http.StatusOK, gin.H{
		"data":                     mahasiswaResponse,
		"jumlahVerifikasiKeluarga": jumlahKeluargaVerified,
		"jumlahVerifikasiIndividu": jumlahIndividuVerified,
	})
}

func (c *mahasiswaController) DistinctKelurahan(cntx *gin.Context) {
	var distinctsKelurahan, err = c.mahasiswaService.DistinctKelurahan()
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	var distinctsKelurahanResponse []responses.DistinctKelurahan

	for _, distinctKelurahan := range distinctsKelurahan {
		var distinctKelurahanResponse = helper.ConvertToDistinctKelurahanResponse(distinctKelurahan)
		distinctsKelurahanResponse = append(distinctsKelurahanResponse, distinctKelurahanResponse)
	}

	cntx.JSON(http.StatusOK, distinctsKelurahanResponse)
}

func (c *mahasiswaController) CreateMahasiswa(cntx *gin.Context) {
	var mahasiswaRequest request.CreateMahasiswaRequest

	var err = cntx.ShouldBindJSON(&mahasiswaRequest)
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

	mahasiswa, err := c.mahasiswaService.Create(mahasiswaRequest)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	var mahasiswaResponse = helper.ConvertToMahasiswaResponse(mahasiswa)

	cntx.JSON(http.StatusCreated, mahasiswaResponse)
}

func (c *mahasiswaController) CreateBatchMahasiswa(cntx *gin.Context) {
	var mahasiswasRequest []request.CreateMahasiswaRequest

	var err = cntx.ShouldBindJSON(&mahasiswasRequest)
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

	mahasiswas, err := c.mahasiswaService.CreateBatch(mahasiswasRequest)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	var mahasiswasResponse []responses.MahasiswaResponse

	for _, mahasiswa := range mahasiswas {
		var mahasiswaResponse = helper.ConvertToMahasiswaResponse(mahasiswa)
		mahasiswasResponse = append(mahasiswasResponse, mahasiswaResponse)
	}

	cntx.JSON(http.StatusOK, mahasiswasResponse)
}

func (c *mahasiswaController) UpdateMahasiswa(cntx *gin.Context) {
	var mahasiswaRequest request.UpdateMahasiswaRequest

	var err = cntx.ShouldBindJSON(&mahasiswaRequest)
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

	mahasiswa, err := c.mahasiswaService.Update(id, mahasiswaRequest)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"errors": cntx.Error(err),
		})
		return
	}

	var mahasiswaResponse = helper.ConvertToMahasiswaResponse(mahasiswa)

	cntx.JSON(http.StatusOK, mahasiswaResponse)
}

func (c *mahasiswaController) DeleteMahasiswa(cntx *gin.Context) {
	var idString = cntx.Param("id")
	var id, _ = strconv.Atoi(idString)

	_, err := c.mahasiswaService.Delete(id)
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
