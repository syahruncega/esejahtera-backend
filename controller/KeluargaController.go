package controller

import (
	"fmt"
	"kemiskinan/helper"
	"kemiskinan/request"
	"kemiskinan/responses"
	"kemiskinan/service"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type keluargaController struct {
	keluargaService service.KeluargaService
}

func NewKeluargaController(keluargaService service.KeluargaService) *keluargaController {
	return &keluargaController{keluargaService}
}

func (c *keluargaController) GetKeluargas(cntx *gin.Context) {
	var keluargas, err = c.keluargaService.FindAll()
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
	}

	var keluargasResponse []responses.KeluargaResponse

	for _, keluarga := range keluargas {
		var keluargaResponse = helper.ConvertToKeluargaResponse(keluarga)
		keluargasResponse = append(keluargasResponse, keluargaResponse)
	}

	cntx.JSON(http.StatusOK, keluargasResponse)
}

func (c *keluargaController) GetKeluargaById(cntx *gin.Context) {
	var idString = cntx.Param("id")
	var id, _ = strconv.Atoi(idString)

	var keluarga, err = c.keluargaService.FindById(id)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": "Data tidak ditemukan",
		})
		return
	}

	var keluargaResponse = helper.ConvertToKeluargaResponse(keluarga)

	cntx.JSON(http.StatusOK, keluargaResponse)
}

func (c *keluargaController) GetKeluargaByIdKeluarga(cntx *gin.Context) {
	var idKeluarga = cntx.Param("idkeluarga")

	var keluarga, err = c.keluargaService.FindByIdKeluarga(idKeluarga)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": "Data tidak ditemukan",
		})
		return
	}

	var keluargaResponse = helper.ConvertToKeluargaResponse(keluarga)

	cntx.JSON(http.StatusOK, keluargaResponse)
}

func (c *keluargaController) GetKeluargaBySearch(cntx *gin.Context) {
	var pageRowString = cntx.Query("pagerow")
	var limit, _ = strconv.Atoi(pageRowString)
	var pageRowFloat, _ = strconv.ParseFloat(pageRowString, 64)

	var halamanString = cntx.Query("halaman")
	var halaman, _ = strconv.Atoi(halamanString)

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

	var jumlahData, err = c.keluargaService.CountData(whereClauseInterface)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
	}

	var offset = limit * (halaman - 1)

	var hasilBagi = float64(jumlahData) / float64(pageRowFloat)
	var jumlahHalaman = int(math.Ceil(hasilBagi))

	fmt.Println(whereClauseInterface)
	fmt.Println(jumlahData)
	fmt.Println(hasilBagi)

	keluargas, err := c.keluargaService.FindBySearch(whereClauseInterface, limit, offset)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
	}

	var keluargasResponse []responses.KeluargaResponse

	for _, keluarga := range keluargas {
		var keluargaResponse = helper.ConvertToKeluargaResponse(keluarga)
		keluargasResponse = append(keluargasResponse, keluargaResponse)
	}

	cntx.JSON(http.StatusOK, gin.H{
		"data":          keluargasResponse,
		"jumlahHalaman": jumlahHalaman,
	})
}

func (c *keluargaController) UpdateKeluarga(cntx *gin.Context) {
	var keluargaRequest request.UpdateKeluargaRequest

	var err = cntx.ShouldBindJSON(&keluargaRequest)
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

	keluarga, err := c.keluargaService.Update(id, keluargaRequest)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"errors": cntx.Error(err),
		})
		return
	}

	var keluargaResponse = helper.ConvertToKeluargaResponse(keluarga)

	cntx.JSON(http.StatusOK, keluargaResponse)
}

func (c *keluargaController) CountDataKeluargaProvinsi(cntx *gin.Context) {
	var places = "provinsiId"
	var placesId = "72"

	jumlahKeluarga, err := c.keluargaService.CountJumlahKeluarga(places, placesId)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	jumlahDesil1, err := c.keluargaService.CountJumlahDesil(places, placesId, "1")
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	jumlahDesil2, err := c.keluargaService.CountJumlahDesil(places, placesId, "2")
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	jumlahDesil3, err := c.keluargaService.CountJumlahDesil(places, placesId, "3")
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}
	jumlahDesil4, err := c.keluargaService.CountJumlahDesil(places, placesId, "4")
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}
	jumlahDesil5, err := c.keluargaService.CountJumlahDesil(places, placesId, "5")
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	jumlahPenerimaBPNTYa, err := c.keluargaService.CountJumlahPenerima(places, placesId, "penerimaBPNT", "Ya")
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	jumlahPenerimaBPNTTidak, err := c.keluargaService.CountJumlahPenerima(places, placesId, "penerimaBPNT", "Tidak")
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	jumlahPenerimaBPUMYa, err := c.keluargaService.CountJumlahPenerima(places, placesId, "penerimaBPUM", "Ya")
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	jumlahPenerimaBPUMTidak, err := c.keluargaService.CountJumlahPenerima(places, placesId, "penerimaBPUM", "Tidak")
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	jumlahPenerimaBSTYa, err := c.keluargaService.CountJumlahPenerima(places, placesId, "penerimaBST", "Ya")
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	jumlahPenerimaBSTTidak, err := c.keluargaService.CountJumlahPenerima(places, placesId, "penerimaBST", "Tidak")
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	jumlahPenerimaPKHYa, err := c.keluargaService.CountJumlahPenerima(places, placesId, "penerimaPKH", "Ya")
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	jumlahPenerimaPKHTidak, err := c.keluargaService.CountJumlahPenerima(places, placesId, "penerimaPKH", "Tidak")
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	jumlahPenerimaSembakoYa, err := c.keluargaService.CountJumlahPenerima(places, placesId, "penerimaSembako", "Ya")
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	jumlahPenerimaSembakoTidak, err := c.keluargaService.CountJumlahPenerima(places, placesId, "penerimaSembako", "Tidak")
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	kabupatenKotaCountKeluarga, err := c.keluargaService.DistinctCountKabupatenKota()
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
	}

	cntx.JSON(http.StatusOK, gin.H{
		"dataHitungKeluargaProvinsi": gin.H{
			"jumlahKeseluruhan": gin.H{
				"jumlahKeluarga": jumlahKeluarga,
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
			"jumlahKeluargaKabupatenKota": kabupatenKotaCountKeluarga,
		},
	})
}

func (c *keluargaController) CountDataKeluargaKabupatenKota(cntx *gin.Context) {
	var err error
	kabupatenKotaDistinct, err := c.keluargaService.DistinctKabupatenKota()
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
	}

	var places = "kabupatenKotaId"

	var jumlahKeluarga, jumlahDesil1, jumlahDesil2, jumlahDesil3, jumlahDesil4, jumlahDesil5, jumlahPenerimaBPNTYa, jumlahPenerimaBPNTTidak, jumlahPenerimaBPUMYa, jumlahPenerimaBPUMTidak, jumlahPenerimaBSTYa, jumlahPenerimaBSTTidak, jumlahPenerimaPKHYa, jumlahPenerimaPKHTidak, jumlahPenerimaSembakoYa, jumlahPenerimaSembakoTidak int64

	var nestedJSON = map[string]map[string]interface{}{}

	for _, kabupatenKota := range kabupatenKotaDistinct {

		jumlahKeluarga, err = c.keluargaService.CountJumlahKeluarga(places, kabupatenKota.KabupatenKotaId)
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahDesil1, err = c.keluargaService.CountJumlahDesil(places, kabupatenKota.KabupatenKotaId, "1")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahDesil2, err = c.keluargaService.CountJumlahDesil(places, kabupatenKota.KabupatenKotaId, "2")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahDesil3, err = c.keluargaService.CountJumlahDesil(places, kabupatenKota.KabupatenKotaId, "3")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahDesil4, err = c.keluargaService.CountJumlahDesil(places, kabupatenKota.KabupatenKotaId, "4")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahDesil5, err = c.keluargaService.CountJumlahDesil(places, kabupatenKota.KabupatenKotaId, "5")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahPenerimaBPNTYa, err = c.keluargaService.CountJumlahPenerima(places, kabupatenKota.KabupatenKotaId, "penerimaBPNT", "Ya")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahPenerimaBPNTTidak, err = c.keluargaService.CountJumlahPenerima(places, kabupatenKota.KabupatenKotaId, "penerimaBPNT", "Tidak")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahPenerimaBPUMYa, err = c.keluargaService.CountJumlahPenerima(places, kabupatenKota.KabupatenKotaId, "penerimaBPUM", "Ya")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahPenerimaBPUMTidak, err = c.keluargaService.CountJumlahPenerima(places, kabupatenKota.KabupatenKotaId, "penerimaBPUM", "Tidak")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahPenerimaBSTYa, err = c.keluargaService.CountJumlahPenerima(places, kabupatenKota.KabupatenKotaId, "penerimaBST", "Ya")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahPenerimaBSTTidak, err = c.keluargaService.CountJumlahPenerima(places, kabupatenKota.KabupatenKotaId, "penerimaBST", "Tidak")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahPenerimaPKHYa, err = c.keluargaService.CountJumlahPenerima(places, kabupatenKota.KabupatenKotaId, "penerimaPKH", "Ya")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahPenerimaPKHTidak, err = c.keluargaService.CountJumlahPenerima(places, kabupatenKota.KabupatenKotaId, "penerimaPKH", "Tidak")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahPenerimaSembakoYa, err = c.keluargaService.CountJumlahPenerima(places, kabupatenKota.KabupatenKotaId, "penerimaSembako", "Ya")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahPenerimaSembakoTidak, err = c.keluargaService.CountJumlahPenerima(places, kabupatenKota.KabupatenKotaId, "penerimaSembako", "Tidak")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		kecamatanCountKeluarga, err := c.keluargaService.DistinctCountKecamatan(kabupatenKota.KabupatenKotaId)
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
		}

		nestedJSON[kabupatenKota.Nama] = map[string]interface{}{
			"jumlahKeseluruhan": map[string]int64{
				"jumlahKeluarga": jumlahKeluarga,
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
			"jumlahKeluargaKecamatan": kecamatanCountKeluarga,
		}
	}

	cntx.JSON(http.StatusOK, gin.H{
		"dataHitungKeluargaKabupatenKota": nestedJSON,
	})
}

func (c *keluargaController) CountDataKeluargaKecamatan(cntx *gin.Context) {
	var kabupatenKotaId = cntx.Query("kabupatenkotaid")
	var err error
	kecamatanDistinct, err := c.keluargaService.DistinctKecamatanByKabupatenKota(kabupatenKotaId)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
	}

	var places = "kecamatanId"

	var jumlahKeluarga, jumlahDesil1, jumlahDesil2, jumlahDesil3, jumlahDesil4, jumlahDesil5, jumlahPenerimaBPNTYa, jumlahPenerimaBPNTTidak, jumlahPenerimaBPUMYa, jumlahPenerimaBPUMTidak, jumlahPenerimaBSTYa, jumlahPenerimaBSTTidak, jumlahPenerimaPKHYa, jumlahPenerimaPKHTidak, jumlahPenerimaSembakoYa, jumlahPenerimaSembakoTidak int64

	var nestedJSON = map[string]map[string]interface{}{}

	for _, kecamatan := range kecamatanDistinct {

		jumlahKeluarga, err = c.keluargaService.CountJumlahKeluarga(places, kecamatan.KecamatanId)
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahDesil1, err = c.keluargaService.CountJumlahDesil(places, kecamatan.KecamatanId, "1")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahDesil2, err = c.keluargaService.CountJumlahDesil(places, kecamatan.KecamatanId, "2")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahDesil3, err = c.keluargaService.CountJumlahDesil(places, kecamatan.KecamatanId, "3")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahDesil4, err = c.keluargaService.CountJumlahDesil(places, kecamatan.KecamatanId, "4")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahDesil5, err = c.keluargaService.CountJumlahDesil(places, kecamatan.KecamatanId, "5")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahPenerimaBPNTYa, err = c.keluargaService.CountJumlahPenerima(places, kecamatan.KecamatanId, "penerimaBPNT", "Ya")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahPenerimaBPNTTidak, err = c.keluargaService.CountJumlahPenerima(places, kecamatan.KecamatanId, "penerimaBPNT", "Tidak")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahPenerimaBPUMYa, err = c.keluargaService.CountJumlahPenerima(places, kecamatan.KecamatanId, "penerimaBPUM", "Ya")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahPenerimaBPUMTidak, err = c.keluargaService.CountJumlahPenerima(places, kecamatan.KecamatanId, "penerimaBPUM", "Tidak")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahPenerimaBSTYa, err = c.keluargaService.CountJumlahPenerima(places, kecamatan.KecamatanId, "penerimaBST", "Ya")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahPenerimaBSTTidak, err = c.keluargaService.CountJumlahPenerima(places, kecamatan.KecamatanId, "penerimaBST", "Tidak")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahPenerimaPKHYa, err = c.keluargaService.CountJumlahPenerima(places, kecamatan.KecamatanId, "penerimaPKH", "Ya")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahPenerimaPKHTidak, err = c.keluargaService.CountJumlahPenerima(places, kecamatan.KecamatanId, "penerimaPKH", "Tidak")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahPenerimaSembakoYa, err = c.keluargaService.CountJumlahPenerima(places, kecamatan.KecamatanId, "penerimaSembako", "Ya")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		jumlahPenerimaSembakoTidak, err = c.keluargaService.CountJumlahPenerima(places, kecamatan.KecamatanId, "penerimaSembako", "Tidak")
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		kelurahanCountKeluarga, err := c.keluargaService.DistinctCountKelurahan(kecamatan.KecamatanId)
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
		}

		nestedJSON[kecamatan.Nama] = map[string]interface{}{
			"jumlahKeseluruhan": map[string]int64{
				"jumlahKeluarga": jumlahKeluarga,
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
			"jumlahKeluargaKelurahan": kelurahanCountKeluarga,
		}
	}

	cntx.JSON(http.StatusOK, gin.H{
		"dataHitungKeluargaKecamatan": nestedJSON,
	})
}

// func (c *keluargaController) TestDistinct(cntx *gin.Context) {
// 	var distinctsKelurahan, err = c.keluargaService.DistinctKelurahan("7205140")
// 	if err != nil {
// 		cntx.JSON(http.StatusBadRequest, gin.H{
// 			"error": cntx.Error(err),
// 		})
// 		return
// 	}

// 	var distinctsKelurahanResponse []responses.DistinctKelurahan

// 	for _, distinctKelurahan := range distinctsKelurahan {
// 		var distinctKelurahanResponse = helper.ConvertToDistinctKelurahanResponse(distinctKelurahan)
// 		distinctsKelurahanResponse = append(distinctsKelurahanResponse, distinctKelurahanResponse)
// 	}

// 	cntx.JSON(http.StatusOK, distinctsKelurahanResponse)

// }
