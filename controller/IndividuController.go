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

func (c *individuController) CountDataIndividuProvinsi(cntx *gin.Context) {
	var places = "provinsiId"
	var placesId = "72"

	jumlahIndividu, err := c.individuService.CountJumlahIndividu(places, placesId)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	jumlahDesil1, err := c.individuService.CountJumlahDesil(places, placesId, "1")
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	jumlahDesil2, err := c.individuService.CountJumlahDesil(places, placesId, "2")
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	jumlahDesil3, err := c.individuService.CountJumlahDesil(places, placesId, "3")
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}
	jumlahDesil4, err := c.individuService.CountJumlahDesil(places, placesId, "4")
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}
	jumlahDesil5, err := c.individuService.CountJumlahDesil(places, placesId, "5")
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	jumlahPenerimaBPNTYa, err := c.individuService.CountJumlahPenerima(places, placesId, "penerimaBPNT", "Ya")
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	jumlahPenerimaBPNTTidak, err := c.individuService.CountJumlahPenerima(places, placesId, "penerimaBPNT", "Tidak")
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	jumlahPenerimaBPUMYa, err := c.individuService.CountJumlahPenerima(places, placesId, "penerimaBPUM", "Ya")
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	jumlahPenerimaBPUMTidak, err := c.individuService.CountJumlahPenerima(places, placesId, "penerimaBPUM", "Tidak")
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	jumlahPenerimaBSTYa, err := c.individuService.CountJumlahPenerima(places, placesId, "penerimaBST", "Ya")
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	jumlahPenerimaBSTTidak, err := c.individuService.CountJumlahPenerima(places, placesId, "penerimaBST", "Tidak")
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	jumlahPenerimaPKHYa, err := c.individuService.CountJumlahPenerima(places, placesId, "penerimaPKH", "Ya")
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	jumlahPenerimaPKHTidak, err := c.individuService.CountJumlahPenerima(places, placesId, "penerimaPKH", "Tidak")
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	jumlahPenerimaSembakoYa, err := c.individuService.CountJumlahPenerima(places, placesId, "penerimaSembako", "Ya")
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	jumlahPenerimaSembakoTidak, err := c.individuService.CountJumlahPenerima(places, placesId, "penerimaSembako", "Tidak")
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	kabupatenKotaCountIndividu, err := c.individuService.DistinctCountKabupatenKota()
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
	}

	cntx.JSON(http.StatusOK, gin.H{
		"dataHitungIndividuProvinsi": gin.H{
			"jumlahKeseluruhan": gin.H{
				"jumlahIndividu": jumlahIndividu,
			},
			"desilKesejahteraan": gin.H{
				"desil1": jumlahDesil1,
				"desil2": jumlahDesil2,
				"desil3": jumlahDesil3,
				"desil4": jumlahDesil4,
				"desil5": jumlahDesil5,
			},
			"penerimaBantuan": gin.H{
				"jumlahPenerimaBPNT": gin.H{
					"ya":    jumlahPenerimaBPNTYa,
					"tidak": jumlahPenerimaBPNTTidak,
				},
				"jumlahPenerimaBPUM": gin.H{
					"ya":    jumlahPenerimaBPUMYa,
					"tidak": jumlahPenerimaBPUMTidak,
				},
				"jumlahPenerimaBST": gin.H{
					"ya":    jumlahPenerimaBSTYa,
					"tidak": jumlahPenerimaBSTTidak,
				},
				"jumlahPenerimaPKH": gin.H{
					"ya":    jumlahPenerimaPKHYa,
					"tidak": jumlahPenerimaPKHTidak,
				},
				"jumlahPenerimaSembako": gin.H{
					"ya":    jumlahPenerimaSembakoYa,
					"tidak": jumlahPenerimaSembakoTidak,
				},
			},
			"jumlahIndividuKabupatenKota": kabupatenKotaCountIndividu,
		},
	})
}

func (c *individuController) CountDataIndividuKabupatenKota(cntx *gin.Context) {
	var err error
	kabupatenKotaDistinct, err := c.individuService.DistinctKabupatenKota()
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
	}

	var places = "kabupatenKotaId"

	var jumlahIndividu, jumlahDesil1, jumlahDesil2, jumlahDesil3, jumlahDesil4, jumlahDesil5, jumlahPenerimaBPNTYa, jumlahPenerimaBPNTTidak, jumlahPenerimaBPUMYa, jumlahPenerimaBPUMTidak, jumlahPenerimaBSTYa, jumlahPenerimaBSTTidak, jumlahPenerimaPKHYa, jumlahPenerimaPKHTidak, jumlahPenerimaSembakoYa, jumlahPenerimaSembakoTidak int64

	var nestedJSON = map[string]map[string]interface{}{}

	for _, kabupatenKota := range kabupatenKotaDistinct {

		jumlahIndividu, err = c.individuService.CountJumlahIndividu(places, kabupatenKota.KabupatenKotaId)
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahDesil1, err = c.individuService.CountJumlahDesil(places, kabupatenKota.KabupatenKotaId, "1")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahDesil2, err = c.individuService.CountJumlahDesil(places, kabupatenKota.KabupatenKotaId, "2")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahDesil3, err = c.individuService.CountJumlahDesil(places, kabupatenKota.KabupatenKotaId, "3")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahDesil4, err = c.individuService.CountJumlahDesil(places, kabupatenKota.KabupatenKotaId, "4")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahDesil5, err = c.individuService.CountJumlahDesil(places, kabupatenKota.KabupatenKotaId, "5")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahPenerimaBPNTYa, err = c.individuService.CountJumlahPenerima(places, kabupatenKota.KabupatenKotaId, "penerimaBPNT", "Ya")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahPenerimaBPNTTidak, err = c.individuService.CountJumlahPenerima(places, kabupatenKota.KabupatenKotaId, "penerimaBPNT", "Tidak")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahPenerimaBPUMYa, err = c.individuService.CountJumlahPenerima(places, kabupatenKota.KabupatenKotaId, "penerimaBPUM", "Ya")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahPenerimaBPUMTidak, err = c.individuService.CountJumlahPenerima(places, kabupatenKota.KabupatenKotaId, "penerimaBPUM", "Tidak")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahPenerimaBSTYa, err = c.individuService.CountJumlahPenerima(places, kabupatenKota.KabupatenKotaId, "penerimaBST", "Ya")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahPenerimaBSTTidak, err = c.individuService.CountJumlahPenerima(places, kabupatenKota.KabupatenKotaId, "penerimaBST", "Tidak")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahPenerimaPKHYa, err = c.individuService.CountJumlahPenerima(places, kabupatenKota.KabupatenKotaId, "penerimaPKH", "Ya")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahPenerimaPKHTidak, err = c.individuService.CountJumlahPenerima(places, kabupatenKota.KabupatenKotaId, "penerimaPKH", "Tidak")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahPenerimaSembakoYa, err = c.individuService.CountJumlahPenerima(places, kabupatenKota.KabupatenKotaId, "penerimaSembako", "Ya")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahPenerimaSembakoTidak, err = c.individuService.CountJumlahPenerima(places, kabupatenKota.KabupatenKotaId, "penerimaSembako", "Tidak")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		kecamatanCountIndividu, err := c.individuService.DistinctCountKecamatan(kabupatenKota.KabupatenKotaId)
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
		}

		nestedJSON[kabupatenKota.Nama] = map[string]interface{}{
			"jumlahKeseluruhan": map[string]int64{
				"jumlahIndividu": jumlahIndividu,
			},
			"desilKesejahteraan": map[string]int64{
				"desil1": jumlahDesil1,
				"desil2": jumlahDesil2,
				"desil3": jumlahDesil3,
				"desil4": jumlahDesil4,
				"desil5": jumlahDesil5,
			},
			"penerimaBantuan": map[string]interface{}{
				"jumlahPenerimaBPNT": map[string]int64{
					"ya":    jumlahPenerimaBPNTYa,
					"tidak": jumlahPenerimaBPNTTidak,
				},
				"jumlahPenerimaBPUM": map[string]int64{
					"ya":    jumlahPenerimaBPUMYa,
					"tidak": jumlahPenerimaBPUMTidak,
				},
				"jumlahPenerimaBST": map[string]int64{
					"ya":    jumlahPenerimaBSTYa,
					"tidak": jumlahPenerimaBSTTidak,
				},
				"jumlahPenerimaPKH": map[string]int64{
					"ya":    jumlahPenerimaPKHYa,
					"tidak": jumlahPenerimaPKHTidak,
				},
				"jumlahPenerimaSembako": map[string]int64{
					"ya":    jumlahPenerimaSembakoYa,
					"tidak": jumlahPenerimaSembakoTidak,
				},
			},
			"jumlahIndividuKecamatan": kecamatanCountIndividu,
		}
	}

	cntx.JSON(http.StatusOK, gin.H{
		"dataHitungIndividuKabupatenKota": nestedJSON,
	})
}

func (c *individuController) CountDataIndividuKecamatan(cntx *gin.Context) {
	var kabupatenKotaId = cntx.Query("kabupatenkotaid")
	var err error
	kecamatanDistinct, err := c.individuService.DistinctKecamatanByKabupatenKota(kabupatenKotaId)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
	}

	var places = "kecamatanId"

	var jumlahIndividu, jumlahDesil1, jumlahDesil2, jumlahDesil3, jumlahDesil4, jumlahDesil5, jumlahPenerimaBPNTYa, jumlahPenerimaBPNTTidak, jumlahPenerimaBPUMYa, jumlahPenerimaBPUMTidak, jumlahPenerimaBSTYa, jumlahPenerimaBSTTidak, jumlahPenerimaPKHYa, jumlahPenerimaPKHTidak, jumlahPenerimaSembakoYa, jumlahPenerimaSembakoTidak int64

	var nestedJSON = map[string]map[string]interface{}{}

	for _, kecamatan := range kecamatanDistinct {

		jumlahIndividu, err = c.individuService.CountJumlahIndividu(places, kecamatan.KecamatanId)
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahDesil1, err = c.individuService.CountJumlahDesil(places, kecamatan.KecamatanId, "1")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahDesil2, err = c.individuService.CountJumlahDesil(places, kecamatan.KecamatanId, "2")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahDesil3, err = c.individuService.CountJumlahDesil(places, kecamatan.KecamatanId, "3")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahDesil4, err = c.individuService.CountJumlahDesil(places, kecamatan.KecamatanId, "4")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahDesil5, err = c.individuService.CountJumlahDesil(places, kecamatan.KecamatanId, "5")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahPenerimaBPNTYa, err = c.individuService.CountJumlahPenerima(places, kecamatan.KecamatanId, "penerimaBPNT", "Ya")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahPenerimaBPNTTidak, err = c.individuService.CountJumlahPenerima(places, kecamatan.KecamatanId, "penerimaBPNT", "Tidak")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahPenerimaBPUMYa, err = c.individuService.CountJumlahPenerima(places, kecamatan.KecamatanId, "penerimaBPUM", "Ya")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahPenerimaBPUMTidak, err = c.individuService.CountJumlahPenerima(places, kecamatan.KecamatanId, "penerimaBPUM", "Tidak")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahPenerimaBSTYa, err = c.individuService.CountJumlahPenerima(places, kecamatan.KecamatanId, "penerimaBST", "Ya")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahPenerimaBSTTidak, err = c.individuService.CountJumlahPenerima(places, kecamatan.KecamatanId, "penerimaBST", "Tidak")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahPenerimaPKHYa, err = c.individuService.CountJumlahPenerima(places, kecamatan.KecamatanId, "penerimaPKH", "Ya")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahPenerimaPKHTidak, err = c.individuService.CountJumlahPenerima(places, kecamatan.KecamatanId, "penerimaPKH", "Tidak")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahPenerimaSembakoYa, err = c.individuService.CountJumlahPenerima(places, kecamatan.KecamatanId, "penerimaSembako", "Ya")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahPenerimaSembakoTidak, err = c.individuService.CountJumlahPenerima(places, kecamatan.KecamatanId, "penerimaSembako", "Tidak")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		kelurahanCountIndividu, err := c.individuService.DistinctCountKelurahan(kecamatan.KecamatanId)
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
		}

		nestedJSON[kecamatan.Nama] = map[string]interface{}{
			"jumlahKeseluruhan": map[string]int64{
				"jumlahIndividu": jumlahIndividu,
			},
			"desilKesejahteraan": map[string]int64{
				"desil1": jumlahDesil1,
				"desil2": jumlahDesil2,
				"desil3": jumlahDesil3,
				"desil4": jumlahDesil4,
				"desil5": jumlahDesil5,
			},
			"penerimaBantuan": map[string]interface{}{
				"jumlahPenerimaBPNT": map[string]int64{
					"ya":    jumlahPenerimaBPNTYa,
					"tidak": jumlahPenerimaBPNTTidak,
				},
				"jumlahPenerimaBPUM": map[string]int64{
					"ya":    jumlahPenerimaBPUMYa,
					"tidak": jumlahPenerimaBPUMTidak,
				},
				"jumlahPenerimaBST": map[string]int64{
					"ya":    jumlahPenerimaBSTYa,
					"tidak": jumlahPenerimaBSTTidak,
				},
				"jumlahPenerimaPKH": map[string]int64{
					"ya":    jumlahPenerimaPKHYa,
					"tidak": jumlahPenerimaPKHTidak,
				},
				"jumlahPenerimaSembako": map[string]int64{
					"ya":    jumlahPenerimaSembakoYa,
					"tidak": jumlahPenerimaSembakoTidak,
				},
			},
			"jumlahIndividuKelurahan": kelurahanCountIndividu,
		}
	}

	cntx.JSON(http.StatusOK, gin.H{
		"dataHitungIndividuKecamatan": nestedJSON,
	})
}
