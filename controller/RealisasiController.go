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

type realisasiController struct {
	rencanaProgramService     service.RencanaProgramService
	rencanaKegiatanService    service.RencanaKegiatanService
	rencanaSubKegiatanService service.RencanaSubKegiatanService
	fokusBelanjaService       service.FokusBelanjaService
	detailLokasiService       service.DetailLokasiService
	realisasiService          service.RealisasiService
}

func NewRealisasiController(rencanaProgramService service.RencanaProgramService, rencanaKegiatanService service.RencanaKegiatanService, rencanaSubKegiatanService service.RencanaSubKegiatanService, fokusBelanjaService service.FokusBelanjaService, detailLokasiService service.DetailLokasiService, realisasiService service.RealisasiService) *realisasiController {
	return &realisasiController{
		rencanaProgramService,
		rencanaKegiatanService,
		rencanaSubKegiatanService,
		fokusBelanjaService,
		detailLokasiService,
		realisasiService,
	}
}

func (c *realisasiController) GetRealisasis(cntx *gin.Context) {
	var realisasis []model.Realisasi
	var err error

	var fokusBelanjaIdString = cntx.Query("fokusBelanjaId")
	var fokusBelanjaId, _ = strconv.Atoi(fokusBelanjaIdString)

	if fokusBelanjaIdString == "" {
		realisasis, err = c.realisasiService.FindAll()
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}
	} else {
		realisasis, err = c.realisasiService.FindByFokusBelanjaId(fokusBelanjaId)
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}
	}

	var realisasisResponse []responses.RealisasiResponse

	for _, realisasi := range realisasis {
		var realisasiResponse = helper.ConvertToRealisasiResponse(realisasi)
		realisasisResponse = append(realisasisResponse, realisasiResponse)
	}

	cntx.JSON(http.StatusOK, realisasisResponse)
}

func (c *realisasiController) GetRealisasi(cntx *gin.Context) {
	var idString = cntx.Param("id")
	var id, _ = strconv.Atoi(idString)

	var realisasi, err = c.realisasiService.FindById(id)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": "Data tidak ditemukan",
		})
		return
	}

	var realisasiResponse = helper.ConvertToRealisasiResponse(realisasi)

	cntx.JSON(http.StatusOK, realisasiResponse)
}

func (c *realisasiController) GetRealisasiBySearch(cntx *gin.Context) {
	var whereClauseString = cntx.Request.URL.Query()
	var whereClauseInterface = map[string]interface{}{}

	for k, v := range whereClauseString {
		interfaceKey := k
		interfaceVal := v

		whereClauseInterface[interfaceKey] = interfaceVal
	}

	var realisasis, err = c.realisasiService.FindBySearch(whereClauseInterface)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	var realisasisResponse []responses.RealisasiResponse

	for _, realisasi := range realisasis {
		var realisasiResponse = helper.ConvertToRealisasiResponse(realisasi)
		realisasisResponse = append(realisasisResponse, realisasiResponse)
	}

	cntx.JSON(http.StatusOK, realisasisResponse)
}

func (c *realisasiController) CreateRealiasasi(cntx *gin.Context) {
	var realisasiRequest request.CreateRealisasiRequest

	var err = cntx.ShouldBindJSON(&realisasiRequest)
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

	realisasi, err := c.realisasiService.Create(realisasiRequest)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	var realisasiResponse = helper.ConvertToRealisasiResponse(realisasi)

	cntx.JSON(http.StatusOK, realisasiResponse)
}

func (c *realisasiController) UpdateRealisasi(cntx *gin.Context) {
	var realisasiRequest request.UpdateRealisasiRequest

	var err = cntx.ShouldBindJSON(&realisasiRequest)
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

	realisasi, err := c.realisasiService.Update(id, realisasiRequest)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"errors": cntx.Error(err),
		})
		return
	}

	var realisasiResponse = helper.ConvertToRealisasiResponse(realisasi)

	cntx.JSON(http.StatusOK, realisasiResponse)
}

func (c *realisasiController) DeleteRealisasi(cntx *gin.Context) {
	var idString = cntx.Param("id")
	var id, _ = strconv.Atoi(idString)

	_, err := c.realisasiService.Delete(id)
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

func (c *realisasiController) GetRealisasiCascading(cntx *gin.Context) {
	var tahun = cntx.Query("tahun")
	var tipe = cntx.Query("tipe")
	var instansiIdString = cntx.Query("instansiId")
	var instansiId, _ = strconv.Atoi(instansiIdString)

	var err error

	if tahun == "" || tipe == "" || instansiIdString == "" {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": "missing parameter tahun, tipe, atau instansiId",
		})
		return
	}

	var ratResponse responses.RatResponse

	rencanaPrograms, err := c.rencanaProgramService.FindRatRencanaProgram(tahun, tipe, instansiId)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	if len(rencanaPrograms) != 0 {
		ratResponse.NamaInstansi = rencanaPrograms[0].Instansi.NamaInstansi

		ratResponse.Program = make([]responses.Program, len(rencanaPrograms))

		for i := 0; i < len(rencanaPrograms); i++ {
			ratResponse.Program[i].NamaProgram = rencanaPrograms[i].Program.NamaProgram
			ratResponse.Program[i].PaguProgram = rencanaPrograms[i].PaguProgram

			var whereClauseKegiatan = make(map[string]interface{})
			whereClauseKegiatan["rencanaProgramId"] = rencanaPrograms[i].Id
			rencanaKegiatans, err := c.rencanaKegiatanService.FindBySearch(whereClauseKegiatan)
			if err != nil {
				cntx.JSON(http.StatusBadRequest, gin.H{
					"error": cntx.Error(err),
				})
				return
			}

			if len(rencanaKegiatans) != 0 {
				ratResponse.Program[i].Kegiatan = make([]responses.Kegiatan, len(rencanaKegiatans))

				for j := 0; j < len(rencanaKegiatans); j++ {

					ratResponse.Program[i].Kegiatan[j].NamaKegiatan = rencanaKegiatans[j].Kegiatan.NamaKegiatan
					ratResponse.Program[i].Kegiatan[j].PaguKegiatan = rencanaKegiatans[j].PaguKegiatan

					var whereClauseSubKegiatan = make(map[string]interface{})
					whereClauseSubKegiatan["rencanaKegiatanId"] = rencanaKegiatans[j].Id
					rencanaSubKegiatans, err := c.rencanaSubKegiatanService.FindBySearch(whereClauseSubKegiatan)
					if err != nil {
						cntx.JSON(http.StatusBadRequest, gin.H{
							"error": cntx.Error(err),
						})
						return
					}

					if len(rencanaSubKegiatans) != 0 {
						ratResponse.Program[i].Kegiatan[j].SubKegiatan = make([]responses.SubKegiatan, len(rencanaSubKegiatans))

						for k := 0; k < len(rencanaSubKegiatans); k++ {

							ratResponse.Program[i].Kegiatan[j].SubKegiatan[k].NamaSubKegiatan = rencanaSubKegiatans[k].SubKegiatan.NamaSubKegiatan
							ratResponse.Program[i].Kegiatan[j].SubKegiatan[k].PaguSubKegiatan = rencanaSubKegiatans[k].PaguSubKegiatan

							var whereClauseFokusBelanja = make(map[string]interface{})
							whereClauseFokusBelanja["rencanaSubKegiatanId"] = rencanaSubKegiatans[k].Id
							fokusBelanjas, err := c.fokusBelanjaService.FindBySearch(whereClauseFokusBelanja)
							if err != nil {
								cntx.JSON(http.StatusBadRequest, gin.H{
									"error": cntx.Error(err),
								})
								return
							}

							if len(fokusBelanjas) != 0 {
								ratResponse.Program[i].Kegiatan[j].SubKegiatan[k].FokusBelanja = make([]responses.FokusBelanja, len(fokusBelanjas))

								for l := 0; l < len(fokusBelanjas); l++ {
									ratResponse.Program[i].Kegiatan[j].SubKegiatan[k].FokusBelanja[l].NamaFokusBelanja = fokusBelanjas[l].NamaFokusBelanja
									ratResponse.Program[i].Kegiatan[j].SubKegiatan[k].FokusBelanja[l].PaguFokusBelanja = fokusBelanjas[l].PaguFokusBelanja
									ratResponse.Program[i].Kegiatan[j].SubKegiatan[k].FokusBelanja[l].Indikator = fokusBelanjas[l].Indikator
									ratResponse.Program[i].Kegiatan[j].SubKegiatan[k].FokusBelanja[l].Target = fokusBelanjas[l].Target
									ratResponse.Program[i].Kegiatan[j].SubKegiatan[k].FokusBelanja[l].Satuan = fokusBelanjas[l].Satuan
									ratResponse.Program[i].Kegiatan[j].SubKegiatan[k].FokusBelanja[l].Keterangan = fokusBelanjas[l].Keterangan

									realisasis, err := c.realisasiService.FindByFokusBelanjaId(fokusBelanjas[l].Id)
									if err != nil {
										cntx.JSON(http.StatusBadRequest, gin.H{
											"error": cntx.Error(err),
										})
										return
									}

									if len(realisasis) != 0 {
										ratResponse.Program[i].Kegiatan[j].SubKegiatan[k].FokusBelanja[l].Realisasi = make([]responses.Realisasi, len(realisasis))

										for m := 0; m < len(realisasis); m++ {
											ratResponse.Program[i].Kegiatan[j].SubKegiatan[k].FokusBelanja[l].Realisasi[m].Tanggal = realisasis[m].Tanggal
											ratResponse.Program[i].Kegiatan[j].SubKegiatan[k].FokusBelanja[l].Realisasi[m].NomorSp2d = realisasis[m].NomorSp2d
											ratResponse.Program[i].Kegiatan[j].SubKegiatan[k].FokusBelanja[l].Realisasi[m].RealisasiPagu = realisasis[m].RealisasiPagu
											ratResponse.Program[i].Kegiatan[j].SubKegiatan[k].FokusBelanja[l].Realisasi[m].RealisasiTarget = realisasis[m].RealisasiTarget
											ratResponse.Program[i].Kegiatan[j].SubKegiatan[k].FokusBelanja[l].Realisasi[m].Bulan = realisasis[m].Bulan
											ratResponse.Program[i].Kegiatan[j].SubKegiatan[k].FokusBelanja[l].Realisasi[m].Keterangan = realisasis[m].Keterangan
										}
									}

									detailLokasis, err := c.detailLokasiService.FindByFokusBelanjaId(fokusBelanjas[l].Id)
									if err != nil {
										cntx.JSON(http.StatusBadRequest, gin.H{
											"error": cntx.Error(err),
										})
										return
									}

									if len(detailLokasis) != 0 {
										ratResponse.Program[i].Kegiatan[j].SubKegiatan[k].FokusBelanja[l].DetailLokasi = make([]responses.DetailLokasi, len(detailLokasis))

										for n := 0; n < len(detailLokasis); n++ {
											ratResponse.Program[i].Kegiatan[j].SubKegiatan[k].FokusBelanja[l].DetailLokasi[n].KabupatenKota = detailLokasis[n].KabupatenKota.Nama
											ratResponse.Program[i].Kegiatan[j].SubKegiatan[k].FokusBelanja[l].DetailLokasi[n].Kecamatan = detailLokasis[n].Kecamatan.Nama
											ratResponse.Program[i].Kegiatan[j].SubKegiatan[k].FokusBelanja[l].DetailLokasi[n].KelurahanDesa = detailLokasis[n].Kelurahan.Nama
											ratResponse.Program[i].Kegiatan[j].SubKegiatan[k].FokusBelanja[l].DetailLokasi[n].TipePenerima = detailLokasis[n].TipePenerima
										}
									}
								}
							}
						}
					}
				}
			}
		}
		cntx.JSON(http.StatusOK, ratResponse)
	} else {
		cntx.JSON(http.StatusOK, nil)
	}
}
