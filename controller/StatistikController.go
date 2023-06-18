package controller

import (
	"kemiskinan/helper"
	"kemiskinan/model"
	"kemiskinan/responses"
	"kemiskinan/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type statistikController struct {
	keluargaService              service.KeluargaService
	individuService              service.IndividuService
	instansiService              service.InstansiService
	programService               service.ProgramService
	kegiatanService              service.KegiatanService
	subKegiatanService           service.SubKegiatanService
	fokusBelanjaService          service.FokusBelanjaService
	programOnKegiatanService     service.ProgramOnKegiatanService
	instansiOnProgramService     service.InstansiOnProgramService
	kegiatanOnSubKegiatanService service.KegiatanOnSubKegiatanService
	rencanaProgramService        service.RencanaProgramService
	rencanaKegiatanService       service.RencanaKegiatanService
	rencanaSubKegiatanService    service.RencanaSubKegiatanService
}

func NewStatistikController(keluargaService service.KeluargaService, individuService service.IndividuService, instansiService service.InstansiService, instansiOnProgramService service.InstansiOnProgramService, programOnKegiatanService service.ProgramOnKegiatanService, programService service.ProgramService, kegiatanService service.KegiatanService, subKegiatanService service.SubKegiatanService, kegiatanOnSubKegiatanService service.KegiatanOnSubKegiatanService, fokusBelanjaService service.FokusBelanjaService, rencanaProgramService service.RencanaProgramService, rencanaKegiatanService service.RencanaKegiatanService, rencanaSubKegiatanService service.RencanaSubKegiatanService) *statistikController {
	return &statistikController{
		keluargaService,
		individuService,
		instansiService,
		programService,
		kegiatanService,
		subKegiatanService,
		fokusBelanjaService,
		programOnKegiatanService,
		instansiOnProgramService,
		kegiatanOnSubKegiatanService,
		rencanaProgramService,
		rencanaKegiatanService,
		rencanaSubKegiatanService,
	}
}

func (c *statistikController) StatistikProvinsi(cntx *gin.Context) {
	var places = "provinsiId"
	var placesId = "72"

	jumlahKeluarga, err := c.keluargaService.CountJumlahKeluarga(places, placesId)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	jumlahIndividu, err := c.individuService.CountJumlahIndividu(places, placesId)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	jumlahDesil1Keluarga, err := c.keluargaService.CountJumlahDesil(places, placesId, "1")
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}
	jumlahDesil1Individu, err := c.individuService.CountJumlahDesil(places, placesId, "1")
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	jumlahDesil2Keluarga, err := c.keluargaService.CountJumlahDesil(places, placesId, "2")
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}
	jumlahDesil2Individu, err := c.individuService.CountJumlahDesil(places, placesId, "2")
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	jumlahDesil3Keluarga, err := c.keluargaService.CountJumlahDesil(places, placesId, "3")
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}
	jumlahDesil3Individu, err := c.individuService.CountJumlahDesil(places, placesId, "3")
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	jumlahDesil4Keluarga, err := c.keluargaService.CountJumlahDesil(places, placesId, "4")
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}
	jumlahDesil4Individu, err := c.individuService.CountJumlahDesil(places, placesId, "4")
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	jumlahDesil5Keluarga, err := c.keluargaService.CountJumlahDesil(places, placesId, "5")
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}
	jumlahDesil5Individu, err := c.individuService.CountJumlahDesil(places, placesId, "5")
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	jumlahPenerimaBPNTYaKeluarga, err := c.keluargaService.CountJumlahPenerima(places, placesId, "penerimaBPNT", "Ya")
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}
	jumlahPenerimaBPNTYaIndividu, err := c.individuService.CountJumlahPenerima(places, placesId, "penerimaBPNT", "Ya")
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	jumlahPenerimaBPNTTidakKeluarga, err := c.keluargaService.CountJumlahPenerima(places, placesId, "penerimaBPNT", "Tidak")
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}
	jumlahPenerimaBPNTTidakIndividu, err := c.individuService.CountJumlahPenerima(places, placesId, "penerimaBPNT", "Tidak")
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	jumlahPenerimaBPUMYaKeluarga, err := c.keluargaService.CountJumlahPenerima(places, placesId, "penerimaBPUM", "Ya")
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}
	jumlahPenerimaBPUMYaIndividu, err := c.individuService.CountJumlahPenerima(places, placesId, "penerimaBPUM", "Ya")
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	jumlahPenerimaBPUMTidakKeluarga, err := c.keluargaService.CountJumlahPenerima(places, placesId, "penerimaBPUM", "Tidak")
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}
	jumlahPenerimaBPUMTidakIndividu, err := c.individuService.CountJumlahPenerima(places, placesId, "penerimaBPUM", "Tidak")
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	jumlahPenerimaBSTYaKeluarga, err := c.keluargaService.CountJumlahPenerima(places, placesId, "penerimaBST", "Ya")
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}
	jumlahPenerimaBSTYaIndividu, err := c.individuService.CountJumlahPenerima(places, placesId, "penerimaBST", "Ya")
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	jumlahPenerimaBSTTidakKeluarga, err := c.keluargaService.CountJumlahPenerima(places, placesId, "penerimaBST", "Tidak")
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}
	jumlahPenerimaBSTTidakIndividu, err := c.individuService.CountJumlahPenerima(places, placesId, "penerimaBST", "Tidak")
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	jumlahPenerimaPKHYaKeluarga, err := c.keluargaService.CountJumlahPenerima(places, placesId, "penerimaPKH", "Ya")
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}
	jumlahPenerimaPKHYaIndividu, err := c.individuService.CountJumlahPenerima(places, placesId, "penerimaPKH", "Ya")
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	jumlahPenerimaPKHTidakKeluarga, err := c.keluargaService.CountJumlahPenerima(places, placesId, "penerimaPKH", "Tidak")
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}
	jumlahPenerimaPKHTidakIndividu, err := c.individuService.CountJumlahPenerima(places, placesId, "penerimaPKH", "Tidak")
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	jumlahPenerimaSembakoYaKeluarga, err := c.keluargaService.CountJumlahPenerima(places, placesId, "penerimaSembako", "Ya")
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}
	jumlahPenerimaSembakoYaIndividu, err := c.individuService.CountJumlahPenerima(places, placesId, "penerimaSembako", "Ya")
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	jumlahPenerimaSembakoTidakKeluarga, err := c.keluargaService.CountJumlahPenerima(places, placesId, "penerimaSembako", "Tidak")
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}
	jumlahPenerimaSembakoTidakIndividu, err := c.individuService.CountJumlahPenerima(places, placesId, "penerimaSembako", "Tidak")
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	var kabupatenKotaId = ""

	kabupatenKotaCountKeluarga, err := c.keluargaService.DistinctCountKabupatenKota(kabupatenKotaId)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
	}
	kabupatenKotaCountIndividu, err := c.individuService.DistinctCountKabupatenKota(kabupatenKotaId)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
	}

	cntx.JSON(http.StatusOK, gin.H{
		"keluarga": gin.H{
			"jumlahKeseluruhan": gin.H{
				"jumlahKeluarga": jumlahKeluarga,
			},
			"desilKesejahteraan": gin.H{
				"desil1": jumlahDesil1Keluarga,
				"desil2": jumlahDesil2Keluarga,
				"desil3": jumlahDesil3Keluarga,
				"desil4": jumlahDesil4Keluarga,
				"desil5": jumlahDesil5Keluarga,
			},
			"penerimaBantuan": gin.H{
				"jumlahPenerimaBPNT": gin.H{
					"ya":    jumlahPenerimaBPNTYaKeluarga,
					"tidak": jumlahPenerimaBPNTTidakKeluarga,
				},
				"jumlahPenerimaBPUM": gin.H{
					"ya":    jumlahPenerimaBPUMYaKeluarga,
					"tidak": jumlahPenerimaBPUMTidakKeluarga,
				},
				"jumlahPenerimaBST": gin.H{
					"ya":    jumlahPenerimaBSTYaKeluarga,
					"tidak": jumlahPenerimaBSTTidakKeluarga,
				},
				"jumlahPenerimaPKH": gin.H{
					"ya":    jumlahPenerimaPKHYaKeluarga,
					"tidak": jumlahPenerimaPKHTidakKeluarga,
				},
				"jumlahPenerimaSembako": gin.H{
					"ya":    jumlahPenerimaSembakoYaKeluarga,
					"tidak": jumlahPenerimaSembakoTidakKeluarga,
				},
			},
			"jumlahKeluargaKabupatenKota": kabupatenKotaCountKeluarga,
		},
		"individu": gin.H{
			"jumlahKeseluruhan": gin.H{
				"jumlahIndividu": jumlahIndividu,
			},
			"desilKesejahteraan": gin.H{
				"desil1": jumlahDesil1Individu,
				"desil2": jumlahDesil2Individu,
				"desil3": jumlahDesil3Individu,
				"desil4": jumlahDesil4Individu,
				"desil5": jumlahDesil5Individu,
			},
			"penerimaBantuan": gin.H{
				"jumlahPenerimaBPNT": gin.H{
					"ya":    jumlahPenerimaBPNTYaIndividu,
					"tidak": jumlahPenerimaBPNTTidakIndividu,
				},
				"jumlahPenerimaBPUM": gin.H{
					"ya":    jumlahPenerimaBPUMYaIndividu,
					"tidak": jumlahPenerimaBPUMTidakIndividu,
				},
				"jumlahPenerimaBST": gin.H{
					"ya":    jumlahPenerimaBSTYaIndividu,
					"tidak": jumlahPenerimaBSTTidakIndividu,
				},
				"jumlahPenerimaPKH": gin.H{
					"ya":    jumlahPenerimaPKHYaIndividu,
					"tidak": jumlahPenerimaPKHTidakIndividu,
				},
				"jumlahPenerimaSembako": gin.H{
					"ya":    jumlahPenerimaSembakoYaIndividu,
					"tidak": jumlahPenerimaSembakoTidakIndividu,
				},
			},
			"jumlahIndividuKabupatenKota": kabupatenKotaCountIndividu,
		},
	})
}

func (c *statistikController) StatistikKabupatenKota(cntx *gin.Context) {
	var err error
	var kabupatenKotaId = cntx.DefaultQuery("kabupatenkotaid", "")

	if kabupatenKotaId == "" {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": "Missing parameters",
		})
		return
	}

	kabupatenKotaDistinctKeluarga, err := c.keluargaService.DistinctKabupatenKota(kabupatenKotaId)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
	}

	var places = "kabupatenKotaId"

	var jumlahKeluarga, jumlahDesil1Keluarga, jumlahDesil2Keluarga, jumlahDesil3Keluarga, jumlahDesil4Keluarga, jumlahDesil5Keluarga, jumlahPenerimaBPNTYaKeluarga, jumlahPenerimaBPNTTidakKeluarga, jumlahPenerimaBPUMYaKeluarga, jumlahPenerimaBPUMTidakKeluarga, jumlahPenerimaBSTYaKeluarga, jumlahPenerimaBSTTidakKeluarga, jumlahPenerimaPKHYaKeluarga, jumlahPenerimaPKHTidakKeluarga, jumlahPenerimaSembakoYaKeluarga, jumlahPenerimaSembakoTidakKeluarga int64

	var kecamatanCountKeluarga = make(map[string]int64)

	for _, kabupatenKota := range kabupatenKotaDistinctKeluarga {

		jumlahKeluarga, err = c.keluargaService.CountJumlahKeluarga(places, kabupatenKota.KabupatenKotaId)
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahDesil1Keluarga, err = c.keluargaService.CountJumlahDesil(places, kabupatenKota.KabupatenKotaId, "1")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahDesil2Keluarga, err = c.keluargaService.CountJumlahDesil(places, kabupatenKota.KabupatenKotaId, "2")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahDesil3Keluarga, err = c.keluargaService.CountJumlahDesil(places, kabupatenKota.KabupatenKotaId, "3")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahDesil4Keluarga, err = c.keluargaService.CountJumlahDesil(places, kabupatenKota.KabupatenKotaId, "4")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahDesil5Keluarga, err = c.keluargaService.CountJumlahDesil(places, kabupatenKota.KabupatenKotaId, "5")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahPenerimaBPNTYaKeluarga, err = c.keluargaService.CountJumlahPenerima(places, kabupatenKota.KabupatenKotaId, "penerimaBPNT", "Ya")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahPenerimaBPNTTidakKeluarga, err = c.keluargaService.CountJumlahPenerima(places, kabupatenKota.KabupatenKotaId, "penerimaBPNT", "Tidak")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahPenerimaBPUMYaKeluarga, err = c.keluargaService.CountJumlahPenerima(places, kabupatenKota.KabupatenKotaId, "penerimaBPUM", "Ya")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahPenerimaBPUMTidakKeluarga, err = c.keluargaService.CountJumlahPenerima(places, kabupatenKota.KabupatenKotaId, "penerimaBPUM", "Tidak")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahPenerimaBSTYaKeluarga, err = c.keluargaService.CountJumlahPenerima(places, kabupatenKota.KabupatenKotaId, "penerimaBST", "Ya")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahPenerimaBSTTidakKeluarga, err = c.keluargaService.CountJumlahPenerima(places, kabupatenKota.KabupatenKotaId, "penerimaBST", "Tidak")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahPenerimaPKHYaKeluarga, err = c.keluargaService.CountJumlahPenerima(places, kabupatenKota.KabupatenKotaId, "penerimaPKH", "Ya")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahPenerimaPKHTidakKeluarga, err = c.keluargaService.CountJumlahPenerima(places, kabupatenKota.KabupatenKotaId, "penerimaPKH", "Tidak")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahPenerimaSembakoYaKeluarga, err = c.keluargaService.CountJumlahPenerima(places, kabupatenKota.KabupatenKotaId, "penerimaSembako", "Ya")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahPenerimaSembakoTidakKeluarga, err = c.keluargaService.CountJumlahPenerima(places, kabupatenKota.KabupatenKotaId, "penerimaSembako", "Tidak")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		kecamatanCountKeluarga, err = c.keluargaService.DistinctCountKecamatan(kabupatenKota.KabupatenKotaId)
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
		}
	}

	//=================================================//============================================

	kabupatenKotaDistinctIndividu, err := c.individuService.DistinctKabupatenKota(kabupatenKotaId)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
	}

	var jumlahIndividu, jumlahDesil1Individu, jumlahDesil2Individu, jumlahDesil3Individu, jumlahDesil4Individu, jumlahDesil5Individu, jumlahPenerimaBPNTYaIndividu, jumlahPenerimaBPNTTidakIndividu, jumlahPenerimaBPUMYaIndividu, jumlahPenerimaBPUMTidakIndividu, jumlahPenerimaBSTYaIndividu, jumlahPenerimaBSTTidakIndividu, jumlahPenerimaPKHYaIndividu, jumlahPenerimaPKHTidakIndividu, jumlahPenerimaSembakoYaIndividu, jumlahPenerimaSembakoTidakIndividu int64

	var kecamatanCountIndividu = make(map[string]int64)

	for _, kabupatenKota := range kabupatenKotaDistinctIndividu {

		jumlahIndividu, err = c.individuService.CountJumlahIndividu(places, kabupatenKota.KabupatenKotaId)
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahDesil1Individu, err = c.individuService.CountJumlahDesil(places, kabupatenKota.KabupatenKotaId, "1")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahDesil2Individu, err = c.individuService.CountJumlahDesil(places, kabupatenKota.KabupatenKotaId, "2")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahDesil3Individu, err = c.individuService.CountJumlahDesil(places, kabupatenKota.KabupatenKotaId, "3")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahDesil4Individu, err = c.individuService.CountJumlahDesil(places, kabupatenKota.KabupatenKotaId, "4")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahDesil5Individu, err = c.individuService.CountJumlahDesil(places, kabupatenKota.KabupatenKotaId, "5")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahPenerimaBPNTYaIndividu, err = c.individuService.CountJumlahPenerima(places, kabupatenKota.KabupatenKotaId, "penerimaBPNT", "Ya")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahPenerimaBPNTTidakIndividu, err = c.individuService.CountJumlahPenerima(places, kabupatenKota.KabupatenKotaId, "penerimaBPNT", "Tidak")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahPenerimaBPUMYaIndividu, err = c.individuService.CountJumlahPenerima(places, kabupatenKota.KabupatenKotaId, "penerimaBPUM", "Ya")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahPenerimaBPUMTidakIndividu, err = c.individuService.CountJumlahPenerima(places, kabupatenKota.KabupatenKotaId, "penerimaBPUM", "Tidak")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahPenerimaBSTYaIndividu, err = c.individuService.CountJumlahPenerima(places, kabupatenKota.KabupatenKotaId, "penerimaBST", "Ya")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahPenerimaBSTTidakIndividu, err = c.individuService.CountJumlahPenerima(places, kabupatenKota.KabupatenKotaId, "penerimaBST", "Tidak")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahPenerimaPKHYaIndividu, err = c.individuService.CountJumlahPenerima(places, kabupatenKota.KabupatenKotaId, "penerimaPKH", "Ya")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahPenerimaPKHTidakIndividu, err = c.individuService.CountJumlahPenerima(places, kabupatenKota.KabupatenKotaId, "penerimaPKH", "Tidak")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahPenerimaSembakoYaIndividu, err = c.individuService.CountJumlahPenerima(places, kabupatenKota.KabupatenKotaId, "penerimaSembako", "Ya")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahPenerimaSembakoTidakIndividu, err = c.individuService.CountJumlahPenerima(places, kabupatenKota.KabupatenKotaId, "penerimaSembako", "Tidak")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		kecamatanCountIndividu, err = c.individuService.DistinctCountKecamatan(kabupatenKota.KabupatenKotaId)
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
		}
	}

	cntx.JSON(http.StatusOK, gin.H{
		"keluarga": gin.H{
			"desilKesejahteraan": gin.H{
				"desil1": jumlahDesil1Keluarga,
				"desil2": jumlahDesil2Keluarga,
				"desil3": jumlahDesil3Keluarga,
				"desil4": jumlahDesil4Keluarga,
				"desil5": jumlahDesil5Keluarga,
			},
			"jumlahKeluargaKecamatan": kecamatanCountKeluarga,
			"jumlahKeseluruhan": gin.H{
				"jumlahKeluarga": jumlahKeluarga,
			},
			"penerimaBantuan": gin.H{
				"jumlahPenerimaBPNT": gin.H{
					"ya":    jumlahPenerimaBPNTYaKeluarga,
					"tidak": jumlahPenerimaBPNTTidakKeluarga,
				},
				"jumlahPenerimaBPUM": gin.H{
					"ya":    jumlahPenerimaBPUMYaKeluarga,
					"tidak": jumlahPenerimaBPUMTidakKeluarga,
				},
				"jumlahPenerimaBST": gin.H{
					"ya":    jumlahPenerimaBSTYaKeluarga,
					"tidak": jumlahPenerimaBSTTidakKeluarga,
				},
				"jumlahPenerimaPKH": gin.H{
					"ya":    jumlahPenerimaPKHYaKeluarga,
					"tidak": jumlahPenerimaPKHTidakKeluarga,
				},
				"jumlahPenerimaSembako": gin.H{
					"ya":    jumlahPenerimaSembakoYaKeluarga,
					"tidak": jumlahPenerimaSembakoTidakKeluarga,
				},
			},
		},
		"individu": gin.H{
			"desilKesejahteraan": gin.H{
				"desil1": jumlahDesil1Individu,
				"desil2": jumlahDesil2Individu,
				"desil3": jumlahDesil3Individu,
				"desil4": jumlahDesil4Individu,
				"desil5": jumlahDesil5Individu,
			},
			"jumlahIndividuKecamatan": kecamatanCountIndividu,
			"jumlahKeseluruhan": gin.H{
				"jumlahIndividu": jumlahIndividu,
			},
			"penerimaBantuan": gin.H{
				"jumlahPenerimaBPNT": gin.H{
					"ya":    jumlahPenerimaBPNTYaIndividu,
					"tidak": jumlahPenerimaBPNTTidakIndividu,
				},
				"jumlahPenerimaBPUM": gin.H{
					"ya":    jumlahPenerimaBPUMYaIndividu,
					"tidak": jumlahPenerimaBPUMTidakIndividu,
				},
				"jumlahPenerimaBST": gin.H{
					"ya":    jumlahPenerimaBSTYaIndividu,
					"tidak": jumlahPenerimaBSTTidakIndividu,
				},
				"jumlahPenerimaPKH": gin.H{
					"ya":    jumlahPenerimaPKHYaIndividu,
					"tidak": jumlahPenerimaPKHTidakIndividu,
				},
				"jumlahPenerimaSembako": gin.H{
					"ya":    jumlahPenerimaSembakoYaIndividu,
					"tidak": jumlahPenerimaSembakoTidakIndividu,
				},
			},
		},
	})
}

func (c *statistikController) StatistikKecamatan(cntx *gin.Context) {
	var kecamatanId = cntx.DefaultQuery("kecamatanid", "")
	var err error

	if kecamatanId == "" {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": "Missing parameters",
		})

		return
	}

	kecamatanDistinctKeluarga, err := c.keluargaService.DistinctKecamatan(kecamatanId)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
	}

	var places = "kecamatanId"

	var jumlahKeluarga, jumlahDesil1Keluarga, jumlahDesil2Keluarga, jumlahDesil3Keluarga, jumlahDesil4Keluarga, jumlahDesil5Keluarga, jumlahPenerimaBPNTYaKeluarga, jumlahPenerimaBPNTTidakKeluarga, jumlahPenerimaBPUMYaKeluarga, jumlahPenerimaBPUMTidakKeluarga, jumlahPenerimaBSTYaKeluarga, jumlahPenerimaBSTTidakKeluarga, jumlahPenerimaPKHYaKeluarga, jumlahPenerimaPKHTidakKeluarga, jumlahPenerimaSembakoYaKeluarga, jumlahPenerimaSembakoTidakKeluarga int64

	var kelurahanCountKeluarga = make(map[string]int64)

	for _, kecamatan := range kecamatanDistinctKeluarga {

		jumlahKeluarga, err = c.keluargaService.CountJumlahKeluarga(places, kecamatan.KecamatanId)
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahDesil1Keluarga, err = c.keluargaService.CountJumlahDesil(places, kecamatan.KecamatanId, "1")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahDesil2Keluarga, err = c.keluargaService.CountJumlahDesil(places, kecamatan.KecamatanId, "2")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahDesil3Keluarga, err = c.keluargaService.CountJumlahDesil(places, kecamatan.KecamatanId, "3")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahDesil4Keluarga, err = c.keluargaService.CountJumlahDesil(places, kecamatan.KecamatanId, "4")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahDesil5Keluarga, err = c.keluargaService.CountJumlahDesil(places, kecamatan.KecamatanId, "5")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahPenerimaBPNTYaKeluarga, err = c.keluargaService.CountJumlahPenerima(places, kecamatan.KecamatanId, "penerimaBPNT", "Ya")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahPenerimaBPNTTidakKeluarga, err = c.keluargaService.CountJumlahPenerima(places, kecamatan.KecamatanId, "penerimaBPNT", "Tidak")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahPenerimaBPUMYaKeluarga, err = c.keluargaService.CountJumlahPenerima(places, kecamatan.KecamatanId, "penerimaBPUM", "Ya")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahPenerimaBPUMTidakKeluarga, err = c.keluargaService.CountJumlahPenerima(places, kecamatan.KecamatanId, "penerimaBPUM", "Tidak")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahPenerimaBSTYaKeluarga, err = c.keluargaService.CountJumlahPenerima(places, kecamatan.KecamatanId, "penerimaBST", "Ya")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahPenerimaBSTTidakKeluarga, err = c.keluargaService.CountJumlahPenerima(places, kecamatan.KecamatanId, "penerimaBST", "Tidak")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahPenerimaPKHYaKeluarga, err = c.keluargaService.CountJumlahPenerima(places, kecamatan.KecamatanId, "penerimaPKH", "Ya")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahPenerimaPKHTidakKeluarga, err = c.keluargaService.CountJumlahPenerima(places, kecamatan.KecamatanId, "penerimaPKH", "Tidak")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahPenerimaSembakoYaKeluarga, err = c.keluargaService.CountJumlahPenerima(places, kecamatan.KecamatanId, "penerimaSembako", "Ya")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahPenerimaSembakoTidakKeluarga, err = c.keluargaService.CountJumlahPenerima(places, kecamatan.KecamatanId, "penerimaSembako", "Tidak")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		kelurahanCountKeluarga, err = c.keluargaService.DistinctCountKelurahan(kecamatan.KecamatanId)
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
		}
	}

	//=================================================//==========================================

	kecamatanDistinctIndividu, err := c.keluargaService.DistinctKecamatan(kecamatanId)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
	}

	var jumlahIndividu, jumlahDesil1Individu, jumlahDesil2Individu, jumlahDesil3Individu, jumlahDesil4Individu, jumlahDesil5Individu, jumlahPenerimaBPNTYaIndividu, jumlahPenerimaBPNTTidakIndividu, jumlahPenerimaBPUMYaIndividu, jumlahPenerimaBPUMTidakIndividu, jumlahPenerimaBSTYaIndividu, jumlahPenerimaBSTTidakIndividu, jumlahPenerimaPKHYaIndividu, jumlahPenerimaPKHTidakIndividu, jumlahPenerimaSembakoYaIndividu, jumlahPenerimaSembakoTidakIndividu int64

	var kelurahanCountIndividu = make(map[string]int64)

	for _, kecamatan := range kecamatanDistinctIndividu {

		jumlahIndividu, err = c.individuService.CountJumlahIndividu(places, kecamatan.KecamatanId)
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahDesil1Individu, err = c.individuService.CountJumlahDesil(places, kecamatan.KecamatanId, "1")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahDesil2Individu, err = c.individuService.CountJumlahDesil(places, kecamatan.KecamatanId, "2")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahDesil3Individu, err = c.individuService.CountJumlahDesil(places, kecamatan.KecamatanId, "3")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahDesil4Individu, err = c.individuService.CountJumlahDesil(places, kecamatan.KecamatanId, "4")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahDesil5Individu, err = c.individuService.CountJumlahDesil(places, kecamatan.KecamatanId, "5")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahPenerimaBPNTYaIndividu, err = c.individuService.CountJumlahPenerima(places, kecamatan.KecamatanId, "penerimaBPNT", "Ya")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahPenerimaBPNTTidakIndividu, err = c.individuService.CountJumlahPenerima(places, kecamatan.KecamatanId, "penerimaBPNT", "Tidak")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahPenerimaBPUMYaIndividu, err = c.individuService.CountJumlahPenerima(places, kecamatan.KecamatanId, "penerimaBPUM", "Ya")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahPenerimaBPUMTidakIndividu, err = c.individuService.CountJumlahPenerima(places, kecamatan.KecamatanId, "penerimaBPUM", "Tidak")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahPenerimaBSTYaIndividu, err = c.individuService.CountJumlahPenerima(places, kecamatan.KecamatanId, "penerimaBST", "Ya")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahPenerimaBSTTidakIndividu, err = c.individuService.CountJumlahPenerima(places, kecamatan.KecamatanId, "penerimaBST", "Tidak")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahPenerimaPKHYaIndividu, err = c.individuService.CountJumlahPenerima(places, kecamatan.KecamatanId, "penerimaPKH", "Ya")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahPenerimaPKHTidakIndividu, err = c.individuService.CountJumlahPenerima(places, kecamatan.KecamatanId, "penerimaPKH", "Tidak")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahPenerimaSembakoYaIndividu, err = c.individuService.CountJumlahPenerima(places, kecamatan.KecamatanId, "penerimaSembako", "Ya")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahPenerimaSembakoTidakIndividu, err = c.individuService.CountJumlahPenerima(places, kecamatan.KecamatanId, "penerimaSembako", "Tidak")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		kelurahanCountIndividu, err = c.individuService.DistinctCountKelurahan(kecamatan.KecamatanId)
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
		}
	}

	cntx.JSON(http.StatusOK, gin.H{
		"keluarga": gin.H{
			"desilKesejahteraan": gin.H{
				"desil1": jumlahDesil1Keluarga,
				"desil2": jumlahDesil2Keluarga,
				"desil3": jumlahDesil3Keluarga,
				"desil4": jumlahDesil4Keluarga,
				"desil5": jumlahDesil5Keluarga,
			},
			"jumlahKeluargaKelurahan": kelurahanCountKeluarga,
			"jumlahKeseluruhan": gin.H{
				"jumlahKeluarga": jumlahKeluarga,
			},
			"penerimaBantuan": gin.H{
				"jumlahPenerimaBPNT": gin.H{
					"ya":    jumlahPenerimaBPNTYaKeluarga,
					"tidak": jumlahPenerimaBPNTTidakKeluarga,
				},
				"jumlahPenerimaBPUM": gin.H{
					"ya":    jumlahPenerimaBPUMYaKeluarga,
					"tidak": jumlahPenerimaBPUMTidakKeluarga,
				},
				"jumlahPenerimaBST": gin.H{
					"ya":    jumlahPenerimaBSTYaKeluarga,
					"tidak": jumlahPenerimaBSTTidakKeluarga,
				},
				"jumlahPenerimaPKH": gin.H{
					"ya":    jumlahPenerimaPKHYaKeluarga,
					"tidak": jumlahPenerimaPKHTidakKeluarga,
				},
				"jumlahPenerimaSembako": gin.H{
					"ya":    jumlahPenerimaSembakoYaKeluarga,
					"tidak": jumlahPenerimaSembakoTidakKeluarga,
				},
			},
		},
		"individu": gin.H{
			"desilKesejahteraan": gin.H{
				"desil1": jumlahDesil1Individu,
				"desil2": jumlahDesil2Individu,
				"desil3": jumlahDesil3Individu,
				"desil4": jumlahDesil4Individu,
				"desil5": jumlahDesil5Individu,
			},
			"jumlahIndividuKelurahan": kelurahanCountIndividu,
			"jumlahKeseluruhan": gin.H{
				"jumlahIndividu": jumlahIndividu,
			},
			"penerimaBantuan": gin.H{
				"jumlahPenerimaBPNT": gin.H{
					"ya":    jumlahPenerimaBPNTYaIndividu,
					"tidak": jumlahPenerimaBPNTTidakIndividu,
				},
				"jumlahPenerimaBPUM": gin.H{
					"ya":    jumlahPenerimaBPUMYaIndividu,
					"tidak": jumlahPenerimaBPUMTidakIndividu,
				},
				"jumlahPenerimaBST": gin.H{
					"ya":    jumlahPenerimaBSTYaIndividu,
					"tidak": jumlahPenerimaBSTTidakIndividu,
				},
				"jumlahPenerimaPKH": gin.H{
					"ya":    jumlahPenerimaPKHYaIndividu,
					"tidak": jumlahPenerimaPKHTidakIndividu,
				},
				"jumlahPenerimaSembako": gin.H{
					"ya":    jumlahPenerimaSembakoYaIndividu,
					"tidak": jumlahPenerimaSembakoTidakIndividu,
				},
			},
		},
	})
}

func (c *statistikController) StatistikKelurahan(cntx *gin.Context) {
	var kelurahanId = cntx.DefaultQuery("kelurahanid", "")

	if kelurahanId == "" {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": "Missing parameters",
		})
		return
	}

	kelurahanDistinctKeluarga, err := c.keluargaService.DistinctKelurahan(kelurahanId)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
	}

	var places = "kelurahanId"

	var jumlahKeluarga, jumlahDesil1Keluarga, jumlahDesil2Keluarga, jumlahDesil3Keluarga, jumlahDesil4Keluarga, jumlahDesil5Keluarga, jumlahPenerimaBPNTYaKeluarga, jumlahPenerimaBPNTTidakKeluarga, jumlahPenerimaBPUMYaKeluarga, jumlahPenerimaBPUMTidakKeluarga, jumlahPenerimaBSTYaKeluarga, jumlahPenerimaBSTTidakKeluarga, jumlahPenerimaPKHYaKeluarga, jumlahPenerimaPKHTidakKeluarga, jumlahPenerimaSembakoYaKeluarga, jumlahPenerimaSembakoTidakKeluarga int64

	for _, kelurahan := range kelurahanDistinctKeluarga {

		jumlahKeluarga, err = c.keluargaService.CountJumlahKeluarga(places, kelurahan.KelurahanId)
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahDesil1Keluarga, err = c.keluargaService.CountJumlahDesil(places, kelurahan.KelurahanId, "1")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahDesil2Keluarga, err = c.keluargaService.CountJumlahDesil(places, kelurahan.KelurahanId, "2")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahDesil3Keluarga, err = c.keluargaService.CountJumlahDesil(places, kelurahan.KelurahanId, "3")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahDesil4Keluarga, err = c.keluargaService.CountJumlahDesil(places, kelurahan.KelurahanId, "4")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahDesil5Keluarga, err = c.keluargaService.CountJumlahDesil(places, kelurahan.KelurahanId, "5")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahPenerimaBPNTYaKeluarga, err = c.keluargaService.CountJumlahPenerima(places, kelurahan.KelurahanId, "penerimaBPNT", "Ya")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahPenerimaBPNTTidakKeluarga, err = c.keluargaService.CountJumlahPenerima(places, kelurahan.KelurahanId, "penerimaBPNT", "Tidak")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahPenerimaBPUMYaKeluarga, err = c.keluargaService.CountJumlahPenerima(places, kelurahan.KelurahanId, "penerimaBPUM", "Ya")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahPenerimaBPUMTidakKeluarga, err = c.keluargaService.CountJumlahPenerima(places, kelurahan.KelurahanId, "penerimaBPUM", "Tidak")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahPenerimaBSTYaKeluarga, err = c.keluargaService.CountJumlahPenerima(places, kelurahan.KelurahanId, "penerimaBST", "Ya")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahPenerimaBSTTidakKeluarga, err = c.keluargaService.CountJumlahPenerima(places, kelurahan.KelurahanId, "penerimaBST", "Tidak")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahPenerimaPKHYaKeluarga, err = c.keluargaService.CountJumlahPenerima(places, kelurahan.KelurahanId, "penerimaPKH", "Ya")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahPenerimaPKHTidakKeluarga, err = c.keluargaService.CountJumlahPenerima(places, kelurahan.KelurahanId, "penerimaPKH", "Tidak")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahPenerimaSembakoYaKeluarga, err = c.keluargaService.CountJumlahPenerima(places, kelurahan.KelurahanId, "penerimaSembako", "Ya")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahPenerimaSembakoTidakKeluarga, err = c.keluargaService.CountJumlahPenerima(places, kelurahan.KelurahanId, "penerimaSembako", "Tidak")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}
	}

	//================================================//========================================//

	kelurahanDistinctIndividu, err := c.keluargaService.DistinctKelurahan(kelurahanId)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
	}

	var jumlahIndividu, jumlahDesil1Individu, jumlahDesil2Individu, jumlahDesil3Individu, jumlahDesil4Individu, jumlahDesil5Individu, jumlahPenerimaBPNTYaIndividu, jumlahPenerimaBPNTTidakIndividu, jumlahPenerimaBPUMYaIndividu, jumlahPenerimaBPUMTidakIndividu, jumlahPenerimaBSTYaIndividu, jumlahPenerimaBSTTidakIndividu, jumlahPenerimaPKHYaIndividu, jumlahPenerimaPKHTidakIndividu, jumlahPenerimaSembakoYaIndividu, jumlahPenerimaSembakoTidakIndividu int64

	for _, kelurahan := range kelurahanDistinctIndividu {

		jumlahIndividu, err = c.individuService.CountJumlahIndividu(places, kelurahan.KelurahanId)
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahDesil1Individu, err = c.individuService.CountJumlahDesil(places, kelurahan.KelurahanId, "1")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahDesil2Individu, err = c.individuService.CountJumlahDesil(places, kelurahan.KelurahanId, "2")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahDesil3Individu, err = c.individuService.CountJumlahDesil(places, kelurahan.KelurahanId, "3")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahDesil4Individu, err = c.individuService.CountJumlahDesil(places, kelurahan.KelurahanId, "4")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahDesil5Individu, err = c.individuService.CountJumlahDesil(places, kelurahan.KelurahanId, "5")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahPenerimaBPNTYaIndividu, err = c.individuService.CountJumlahPenerima(places, kelurahan.KelurahanId, "penerimaBPNT", "Ya")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahPenerimaBPNTTidakIndividu, err = c.individuService.CountJumlahPenerima(places, kelurahan.KelurahanId, "penerimaBPNT", "Tidak")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahPenerimaBPUMYaIndividu, err = c.individuService.CountJumlahPenerima(places, kelurahan.KelurahanId, "penerimaBPUM", "Ya")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahPenerimaBPUMTidakIndividu, err = c.individuService.CountJumlahPenerima(places, kelurahan.KelurahanId, "penerimaBPUM", "Tidak")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahPenerimaBSTYaIndividu, err = c.individuService.CountJumlahPenerima(places, kelurahan.KelurahanId, "penerimaBST", "Ya")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahPenerimaBSTTidakIndividu, err = c.individuService.CountJumlahPenerima(places, kelurahan.KelurahanId, "penerimaBST", "Tidak")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahPenerimaPKHYaIndividu, err = c.individuService.CountJumlahPenerima(places, kelurahan.KelurahanId, "penerimaPKH", "Ya")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahPenerimaPKHTidakIndividu, err = c.individuService.CountJumlahPenerima(places, kelurahan.KelurahanId, "penerimaPKH", "Tidak")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahPenerimaSembakoYaIndividu, err = c.individuService.CountJumlahPenerima(places, kelurahan.KelurahanId, "penerimaSembako", "Ya")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahPenerimaSembakoTidakIndividu, err = c.individuService.CountJumlahPenerima(places, kelurahan.KelurahanId, "penerimaSembako", "Tidak")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}
	}

	cntx.JSON(http.StatusOK, gin.H{
		"keluarga": gin.H{
			"desilKesejahteraan": gin.H{
				"desil1": jumlahDesil1Keluarga,
				"desil2": jumlahDesil2Keluarga,
				"desil3": jumlahDesil3Keluarga,
				"desil4": jumlahDesil4Keluarga,
				"desil5": jumlahDesil5Keluarga,
			},
			"jumlahKeseluruhan": gin.H{
				"jumlahKeluarga": jumlahKeluarga,
			},
			"penerimaBantuan": gin.H{
				"jumlahPenerimaBPNT": gin.H{
					"ya":    jumlahPenerimaBPNTYaKeluarga,
					"tidak": jumlahPenerimaBPNTTidakKeluarga,
				},
				"jumlahPenerimaBPUM": gin.H{
					"ya":    jumlahPenerimaBPUMYaKeluarga,
					"tidak": jumlahPenerimaBPUMTidakKeluarga,
				},
				"jumlahPenerimaBST": gin.H{
					"ya":    jumlahPenerimaBSTYaKeluarga,
					"tidak": jumlahPenerimaBSTTidakKeluarga,
				},
				"jumlahPenerimaPKH": gin.H{
					"ya":    jumlahPenerimaPKHYaKeluarga,
					"tidak": jumlahPenerimaPKHTidakKeluarga,
				},
				"jumlahPenerimaSembako": gin.H{
					"ya":    jumlahPenerimaSembakoYaKeluarga,
					"tidak": jumlahPenerimaSembakoTidakKeluarga,
				},
			},
		},
		"individu": gin.H{
			"desilKesejahteraan": gin.H{
				"desil1": jumlahDesil1Individu,
				"desil2": jumlahDesil2Individu,
				"desil3": jumlahDesil3Individu,
				"desil4": jumlahDesil4Individu,
				"desil5": jumlahDesil5Individu,
			},
			"jumlahKeseluruhan": gin.H{
				"jumlahIndividu": jumlahIndividu,
			},
			"penerimaBantuan": gin.H{
				"jumlahPenerimaBPNT": gin.H{
					"ya":    jumlahPenerimaBPNTYaIndividu,
					"tidak": jumlahPenerimaBPNTTidakIndividu,
				},
				"jumlahPenerimaBPUM": gin.H{
					"ya":    jumlahPenerimaBPUMYaIndividu,
					"tidak": jumlahPenerimaBPUMTidakIndividu,
				},
				"jumlahPenerimaBST": gin.H{
					"ya":    jumlahPenerimaBSTYaIndividu,
					"tidak": jumlahPenerimaBSTTidakIndividu,
				},
				"jumlahPenerimaPKH": gin.H{
					"ya":    jumlahPenerimaPKHYaIndividu,
					"tidak": jumlahPenerimaPKHTidakIndividu,
				},
				"jumlahPenerimaSembako": gin.H{
					"ya":    jumlahPenerimaSembakoYaIndividu,
					"tidak": jumlahPenerimaSembakoTidakIndividu,
				},
			},
		},
	})

}

func (c *statistikController) StatistikSearch(cntx *gin.Context) {
	var whereClauseString = cntx.Request.URL.Query()
	var whereClauseInterface = make(map[string]interface{})

	for k, v := range whereClauseString {
		if k == "halaman" || k == "pagerow" {
			continue
		}

		interfaceKey := k
		interfaceVal := v

		whereClauseInterface[interfaceKey] = interfaceVal
	}

	jumlahDataKeluarga, err := c.keluargaService.CountData(whereClauseInterface)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
	}

	jumlahDataIndividu, err := c.individuService.CountData(whereClauseInterface)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
	}

	cntx.JSON(http.StatusOK, gin.H{
		"keluarga": jumlahDataKeluarga,
		"individu": jumlahDataIndividu,
	})
}

func (c *statistikController) StatistikProgram(cntx *gin.Context) {
	var tahun = cntx.Query("tahun")

	var jumlahDataProgram, err = c.programService.CountJumlahProgram(tahun)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
	}

	cntx.JSON(http.StatusOK, gin.H{
		"jumlahDataProgram": jumlahDataProgram,
	})
}

func (c *statistikController) StatistikProgramAllInstansi(cntx *gin.Context) {
	var tahun = cntx.Query("tahun")
	var instansiIdString = cntx.Query("instansiId")
	var instansiId, _ = strconv.Atoi(instansiIdString)

	if instansiIdString != "" {

		var instansi, err = c.instansiService.FindById(instansiId)
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
		}

		var hasil = c.instansiOnProgramService.CountJumlahProgramByInstansiId(tahun, instansi)

		var hasilResponse = helper.ConvertToStatistikProgramInstansiResponse(instansi, hasil)

		cntx.JSON(http.StatusOK, hasilResponse)

	} else {

		var instansis, err = c.instansiService.FindAll()
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
		}

		var hasil = c.instansiOnProgramService.CountJumlahProgramAllInstansi(tahun, instansis)

		var statistikProgramAllInstansiResponse []responses.StatistikProgramInstansiResponse

		for i := 0; i < len(instansis); i++ {
			var tempResponse = helper.ConvertToStatistikProgramInstansiResponse(instansis[i], hasil[i])
			statistikProgramAllInstansiResponse = append(statistikProgramAllInstansiResponse, tempResponse)
		}

		cntx.JSON(http.StatusOK, statistikProgramAllInstansiResponse)
	}
}

func (c *statistikController) StatistikKegiatan(cntx *gin.Context) {
	var tahun = cntx.Query("tahun")

	var jumlahDataKegiatan, err = c.kegiatanService.CountJumlahKegiatan(tahun)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
	}

	cntx.JSON(http.StatusOK, gin.H{
		"jumlahDataKegiatan": jumlahDataKegiatan,
	})
}

func (c *statistikController) StatistikKegiatanAllInstansi(cntx *gin.Context) {
	var tahun = cntx.Query("tahun")
	var instansiIdString = cntx.Query("instansiId")
	var instansiId, _ = strconv.Atoi(instansiIdString)

	if instansiIdString != "" {

		var instansi, err = c.instansiService.FindById(instansiId)
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
		}

		var hasil = c.programOnKegiatanService.CountJumlahKegiatanByInstansiId(tahun, instansi)

		var hasilResponse = helper.ConvertToStatistikKegiatanInstansiResponse(instansi, hasil)

		cntx.JSON(http.StatusOK, hasilResponse)

	} else {

		var instansis, err = c.instansiService.FindAll()
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
		}

		var hasil = c.programOnKegiatanService.CountJumlahKegiatanAllInstansi(tahun, instansis)

		var statistikKegiatanAllInstansiResponse []responses.StatistikKegiatanInstansiResponse

		for i := 0; i < len(instansis); i++ {
			var tempResponse = helper.ConvertToStatistikKegiatanInstansiResponse(instansis[i], hasil[i])
			statistikKegiatanAllInstansiResponse = append(statistikKegiatanAllInstansiResponse, tempResponse)
		}

		cntx.JSON(http.StatusOK, statistikKegiatanAllInstansiResponse)
	}
}

func (c *statistikController) StatistikSubKegiatan(cntx *gin.Context) {
	var tahun = cntx.Query("tahun")

	var jumlahDataSubKegiatan, err = c.subKegiatanService.CountJumlahSubKegiatan(tahun)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
	}

	cntx.JSON(http.StatusOK, gin.H{
		"jumlahDataSubKegiatan": jumlahDataSubKegiatan,
	})
}

func (c *statistikController) StatistikSubKegiatanAllInstansi(cntx *gin.Context) {
	var tahun = cntx.Query("tahun")
	var instansiIdString = cntx.Query("instansiId")
	var instansiId, _ = strconv.Atoi(instansiIdString)

	if instansiIdString != "" {

		var instansi, err = c.instansiService.FindById(instansiId)
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
		}

		var hasil = c.kegiatanOnSubKegiatanService.CountJumlahSubKegiatanByInstansiId(tahun, instansi)

		var jumlahSubKegiatanResponse = helper.ConvertToStatistikSubKegiatanInstansiResponse(instansi, hasil)

		cntx.JSON(http.StatusOK, jumlahSubKegiatanResponse)

	} else {

		var instansis, err = c.instansiService.FindAll()
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
		}

		var hasil = c.kegiatanOnSubKegiatanService.CountJumlahSubKegiatanAllInstansi(tahun, instansis)

		var statistikSubKegiatanAllInstansiResponse []responses.StatistikSubKegiatanInstansiResponse

		for i := 0; i < len(instansis); i++ {
			var tempResponse = helper.ConvertToStatistikSubKegiatanInstansiResponse(instansis[i], hasil[i])
			statistikSubKegiatanAllInstansiResponse = append(statistikSubKegiatanAllInstansiResponse, tempResponse)
		}

		cntx.JSON(http.StatusOK, statistikSubKegiatanAllInstansiResponse)
	}
}

func (c *statistikController) StatistikRencanaProgram(cntx *gin.Context) {
	var tahun = cntx.Query("tahun")
	var tipe = cntx.Query("tipe")

	var jumlahRencanaProgram, err = c.rencanaProgramService.CountJumlahRencanaProgram(tahun, tipe)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
	}

	cntx.JSON(http.StatusOK, gin.H{
		"jumlahRencanaProgram": jumlahRencanaProgram,
	})
}

func (c *statistikController) StatistikRencanaProgramByInstansi(cntx *gin.Context) {
	var tahun = cntx.Query("tahun")
	var tipe = cntx.Query("tipe")
	var instansiIdString = cntx.Query("instansiId")
	var instansiId, _ = strconv.Atoi(instansiIdString)

	if instansiIdString != "" {

		var instansis []model.Instansi

		var instansi, err = c.instansiService.FindById(instansiId)
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
		}

		instansis = append(instansis, instansi)

		var jumlahRencanaProgram = c.rencanaProgramService.CountJumlahRencanaProgramByIntansi(tahun, tipe, instansis)

		var jumlahRencanaProgramResponse = helper.ConvertToStatistikRencanaProgramInstansiResponse(instansis[0], jumlahRencanaProgram[0])

		cntx.JSON(http.StatusOK, jumlahRencanaProgramResponse)

	} else {

		var instansis, err = c.instansiService.FindAll()
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
		}

		var jumlahRencanaProgram = c.rencanaProgramService.CountJumlahRencanaProgramByIntansi(tahun, tipe, instansis)

		var jumlahRencanaProgramResponse []responses.StatistikRencanaProgramInstansiResponse

		for i := 0; i < len(instansis); i++ {
			var tempResponse = helper.ConvertToStatistikRencanaProgramInstansiResponse(instansis[i], jumlahRencanaProgram[i])
			jumlahRencanaProgramResponse = append(jumlahRencanaProgramResponse, tempResponse)
		}

		cntx.JSON(http.StatusOK, jumlahRencanaProgramResponse)
	}
}

func (c *statistikController) StatistikRencanaKegiatan(cntx *gin.Context) {
	var tahun = cntx.Query("tahun")
	var tipe = cntx.Query("tipe")

	var jumlahRencanaKegiatan, err = c.rencanaKegiatanService.CountJumlahRencanaKegiatan(tahun, tipe)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
	}

	cntx.JSON(http.StatusOK, gin.H{
		"jumlahRencanaKegiatan": jumlahRencanaKegiatan,
	})
}

func (c *statistikController) StatistikRencanaKegiatanByInstansi(cntx *gin.Context) {
	var tahun = cntx.Query("tahun")
	var tipe = cntx.Query("tipe")
	var instansiIdString = cntx.Query("instansiId")
	var instansiId, _ = strconv.Atoi(instansiIdString)

	if instansiIdString != "" {

		var instansis []model.Instansi

		var instansi, err = c.instansiService.FindById(instansiId)
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
		}

		instansis = append(instansis, instansi)

		var jumlahRencanaKegiatan = c.rencanaKegiatanService.CountJumlahRencanaKegiatanAllInstansi(tahun, tipe, instansis)

		var jumlahDataResponse = helper.ConvertToStatistikRencanaKegiatanInstansiResponse(instansis[0], jumlahRencanaKegiatan[0])

		cntx.JSON(http.StatusOK, jumlahDataResponse)
	} else {

		var instansis, err = c.instansiService.FindAll()
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
		}

		var jumlahRencanaKegiatan = c.rencanaKegiatanService.CountJumlahRencanaKegiatanAllInstansi(tahun, tipe, instansis)

		var jumlahRencanaKegiatanResponse []responses.StatistikRencanaKegiatanInstansiResponse

		for i := 0; i < len(instansis); i++ {
			var tempResponse = helper.ConvertToStatistikRencanaKegiatanInstansiResponse(instansis[i], jumlahRencanaKegiatan[i])
			jumlahRencanaKegiatanResponse = append(jumlahRencanaKegiatanResponse, tempResponse)
		}

		cntx.JSON(http.StatusOK, jumlahRencanaKegiatanResponse)
	}
}

func (c *statistikController) StatistikRencanaSubKegiatan(cntx *gin.Context) {
	var tahun = cntx.Query("tahun")
	var tipe = cntx.Query("tipe")

	var jumlahRencanaSubKegiatan, err = c.rencanaSubKegiatanService.CountJumlahRencanaSubKegiatan(tahun, tipe)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
	}

	cntx.JSON(http.StatusOK, gin.H{
		"jumlahRencanaSubKegiatan": jumlahRencanaSubKegiatan,
	})
}

func (c *statistikController) StatistikRencanaSubKegiatanByInstansi(cntx *gin.Context) {
	var tahun = cntx.Query("tahun")
	var tipe = cntx.Query("tipe")
	var instansiIdString = cntx.Query("instansiId")
	var instansiId, _ = strconv.Atoi(instansiIdString)

	if instansiIdString != "" {

		var instansis []model.Instansi

		var instansi, err = c.instansiService.FindById(instansiId)
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
		}

		instansis = append(instansis, instansi)

		var jumlahRencanaSubKegiatan = c.rencanaSubKegiatanService.CountJumlahRencanaSubKegiatanByInstansi(tahun, tipe, instansis)

		var jumlahRencanaSubKegiatanResponse = helper.ConvertToStatistikRencanaSubKegiatanInstansiResponse(instansis[0], jumlahRencanaSubKegiatan[0])

		cntx.JSON(http.StatusOK, jumlahRencanaSubKegiatanResponse)
	} else {

		var instansis, err = c.instansiService.FindAll()
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
		}

		var jumlahRencanaSubKegiatan = c.rencanaSubKegiatanService.CountJumlahRencanaSubKegiatanByInstansi(tahun, tipe, instansis)

		var jumlahRencanaSubKegiatanResponse []responses.StatistikRencanaSubKegiatanInstansiResponse

		for i := 0; i < len(instansis); i++ {
			var tempResponse = helper.ConvertToStatistikRencanaSubKegiatanInstansiResponse(instansis[i], jumlahRencanaSubKegiatan[i])
			jumlahRencanaSubKegiatanResponse = append(jumlahRencanaSubKegiatanResponse, tempResponse)
		}

		cntx.JSON(http.StatusOK, jumlahRencanaSubKegiatanResponse)
	}
}

func (c *statistikController) StatistikFokusBelanja(cntx *gin.Context) {
	var tahun = cntx.Query("tahun")

	var jumlahFokusBelanja, err = c.fokusBelanjaService.CountJumlahFokusBelanja(tahun)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
	}

	cntx.JSON(http.StatusOK, gin.H{
		"jumlahFokusBelanja": jumlahFokusBelanja,
	})
}

func (c *statistikController) StatistikFokusBelanjaByInstansi(cntx *gin.Context) {
	var tahun = cntx.Query("tahun")
	var tipe = cntx.Query("tipe")
	var instansiIdString = cntx.Query("instansiId")
	var instansiId, _ = strconv.Atoi(instansiIdString)

	if instansiIdString != "" {

		var instansis []model.Instansi

		var instansi, err = c.instansiService.FindById(instansiId)
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
		}

		instansis = append(instansis, instansi)

		var jumlahFokusBelanja = c.fokusBelanjaService.CountJumlahFokusBelanjaByInstansi(tahun, tipe, instansis)

		var jumlahFokusBelanjaResponse = helper.ConvertToStatistikFokusBelanjaInstansiResponse(instansis[0], jumlahFokusBelanja[0])

		cntx.JSON(http.StatusOK, jumlahFokusBelanjaResponse)
	} else {

		var instansis, err = c.instansiService.FindAll()
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
		}

		var jumlahFokusBelanja = c.fokusBelanjaService.CountJumlahFokusBelanjaByInstansi(tahun, tipe, instansis)

		var jumlahFokusBelanjaResponse []responses.StatistikFokusBelanjaInstansiResponse

		for i := 0; i < len(instansis); i++ {
			var tempResponse = helper.ConvertToStatistikFokusBelanjaInstansiResponse(instansis[i], jumlahFokusBelanja[i])
			jumlahFokusBelanjaResponse = append(jumlahFokusBelanjaResponse, tempResponse)
		}

		cntx.JSON(http.StatusOK, jumlahFokusBelanjaResponse)
	}
}
