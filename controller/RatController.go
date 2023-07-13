package controller

import (
	"kemiskinan/responses"
	"kemiskinan/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ratController struct {
	rencanaProgramService     service.RencanaProgramService
	rencanaKegiatanService    service.RencanaKegiatanService
	rencanaSubKegiatanService service.RencanaSubKegiatanService
	fokusBelanjaService       service.FokusBelanjaService
	detailLokasiService       service.DetailLokasiService
}

func NewRatController(rencanaProgramService service.RencanaProgramService, rencanaKegiatanService service.RencanaKegiatanService, rencanaSubKegiatanService service.RencanaSubKegiatanService, fokusBelanjaService service.FokusBelanjaService, detailLokasiService service.DetailLokasiService) *ratController {
	return &ratController{
		rencanaProgramService,
		rencanaKegiatanService,
		rencanaSubKegiatanService,
		fokusBelanjaService,
		detailLokasiService,
	}
}

func (c *ratController) TestingRatController(cntx *gin.Context) {
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

									detailLokasis, err := c.detailLokasiService.FindByFokusBelanjaId(fokusBelanjas[l].Id)
									if err != nil {
										cntx.JSON(http.StatusBadRequest, gin.H{
											"error": cntx.Error(err),
										})
										return
									}

									if len(detailLokasis) != 0 {
										ratResponse.Program[i].Kegiatan[j].SubKegiatan[k].FokusBelanja[l].DetailLokasi = make([]responses.DetailLokasi, len(detailLokasis))

										for m := 0; m < len(detailLokasis); m++ {
											ratResponse.Program[i].Kegiatan[j].SubKegiatan[k].FokusBelanja[l].DetailLokasi[m].KabupatenKota = detailLokasis[m].KabupatenKota.Nama
											ratResponse.Program[i].Kegiatan[j].SubKegiatan[k].FokusBelanja[l].DetailLokasi[m].Kecamatan = detailLokasis[m].Kecamatan.Nama
											ratResponse.Program[i].Kegiatan[j].SubKegiatan[k].FokusBelanja[l].DetailLokasi[m].KelurahanDesa = detailLokasis[m].Kelurahan.Nama
											ratResponse.Program[i].Kegiatan[j].SubKegiatan[k].FokusBelanja[l].DetailLokasi[m].TipePenerima = detailLokasis[m].TipePenerima
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
