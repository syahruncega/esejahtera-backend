package controller

import (
	"kemiskinan/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type statistikController struct {
	keluargaService service.KeluargaService
	individuService service.IndividuService
}

func NewStatistikController(keluargaService service.KeluargaService, individuService service.IndividuService) *statistikController {
	return &statistikController{
		keluargaService,
		individuService,
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
