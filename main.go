package main

import (
	"kemiskinan/config"
	"kemiskinan/controller"
	"kemiskinan/repository"
	"kemiskinan/service"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	var appConfig = config.AppConfig{}

	var err = godotenv.Load()
	if err != nil {
		log.Fatal("Terdapat kesalahan memuat file .env")
	}

	appConfig.DatabaseUsername = os.Getenv("APP_DATABASE_USERNAME")
	appConfig.DatabasePassword = os.Getenv("APP_DATABASE_PASSWORD")
	appConfig.DatabasePort = os.Getenv("APP_DATABASE_PORT")
	appConfig.DatabaseName = os.Getenv("APP_DATABASE_NAME")
	appConfig.AppPort = os.Getenv("APP_PORT")

	config.ConnectDB(appConfig)

	var bidangUrusanRepository = repository.NewBidangUrusanRepository(config.DB)
	var bidangUrusanService = service.NewBidangUrusanService(bidangUrusanRepository)
	var bidangUrusanController = controller.NewBidangUrusanController(bidangUrusanService)

	var instansiRepository = repository.NewInstansiRepository(config.DB)
	var instansiService = service.NewInstansiService(instansiRepository)
	var instansiController = controller.NewInstansiController(instansiService)

	var programRepository = repository.NewProgramRepository(config.DB)
	var programService = service.NewProgramService(programRepository)
	var programController = controller.NewProgramController(programService)

	var kegiatanRepository = repository.NewKegiatanRepository(config.DB)
	var kegiatanService = service.NewKegiatanService(kegiatanRepository)
	var kegiatanController = controller.NewKegiatanController(kegiatanService)

	var subKegiatanRepository = repository.NewSubKegiatanRepository(config.DB)
	var subKegiatanService = service.NewSubKegiatanService(subKegiatanRepository)
	var subKegiatanController = controller.NewSubKegiatanController(subKegiatanService)

	var detailSubKegiatanRepository = repository.NewDetailSubKegiatanRepository(config.DB)
	var detailSubKegiatanService = service.NewDetailSubKegiatanService(detailSubKegiatanRepository)
	var detailSubKegiatanController = controller.NewDetailSubKegiatanController(detailSubKegiatanService)

	var detailLokasiRepository = repository.NewDetailLokasiRepository(config.DB)
	var detailLokasiService = service.NewDetailLokasiService(detailLokasiRepository)
	var detailLokasiController = controller.NewDetailLokasiController(detailLokasiService)

	var provinsiRepository = repository.NewProvinsiRepository(config.DB)
	var provinsiService = service.NewProvinsiService(provinsiRepository)
	var provinsiController = controller.NewProvinsiController(provinsiService)

	var kabupatenKotaRepository = repository.NewKabupatenKotaRepository(config.DB)
	var kabupatenKotaService = service.NewKabupatenKotaService(kabupatenKotaRepository)
	var kabupatenKotaController = controller.NewKabupatenKotaController(kabupatenKotaService)

	var kecamatanRepository = repository.NewKecamatanRepository(config.DB)
	var kecamatanService = service.NewKecamatanService(kecamatanRepository)
	var kecamatanController = controller.NewKecamatanController(kecamatanService)

	var kelurahanRepository = repository.NewKelurahanRepository(config.DB)
	var kelurahanService = service.NewKelurahanService(kelurahanRepository)
	var kelurahanController = controller.NewKelurahanController(kelurahanService)

	var server = gin.Default()

	server.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"POST", "PUT", "DELETE", "GET", "PATCH"},
		AllowHeaders: []string{"Content-Type, access-control-allow-origin, access-control-allow-headers"},
	}))

	server.GET("/bidangurusan", bidangUrusanController.GetBidangUrusans)
	server.GET("/bidangurusan/:id", bidangUrusanController.GetBidangUrusan)
	server.POST("/bidangurusan", bidangUrusanController.CreateBidangUrusan)
	server.PATCH("/bidangurusan/:id", bidangUrusanController.UpdateBidangUrusan)
	server.DELETE("/bidangurusan/:id", bidangUrusanController.DeleteBidangUrusan)

	server.GET("/instansi", instansiController.GetInstansis)
	server.GET("/instansi/:id", instansiController.GetInstansi)
	server.GET("/instansirelasi", instansiController.GetInstansiWithRelation)
	server.POST("/instansi", instansiController.CreateInstansi)
	server.PATCH("/instansi/:id", instansiController.UpdateInstansi)
	server.DELETE("/instansi/:id", instansiController.DeleteInstansi)

	server.GET("/program", programController.GetPrograms)
	server.GET("/program/:id", programController.GetProgram)
	server.GET("/programrelasi", programController.GetProgramWithRelation)
	server.POST("/program", programController.CreateProgram)
	server.PATCH("/program/:id", programController.UpdateProgram)
	server.DELETE("/program/:id", programController.DeleteProgram)

	server.GET("/kegiatan", kegiatanController.GetKegiatans)
	server.GET("/kegiatan/:id", kegiatanController.GetKegiatan)
	server.GET("/kegiatanrelasi", kegiatanController.GetKegiatanWithRelation)
	server.POST("/kegiatan", kegiatanController.CreateKegiatan)
	server.PATCH("/kegiatan/:id", kegiatanController.UpdateKegiatan)
	server.DELETE("/kegiatan/:id", kegiatanController.DeleteKegiatan)

	server.GET("/subkegiatan", subKegiatanController.GetSubKegiatans)
	server.GET("/subkegiatan/:id", subKegiatanController.GetSubKegiatan)
	server.GET("/subkegiatanrelasi", subKegiatanController.GetSubKegiatanWithRelation)
	server.POST("/subkegiatan", subKegiatanController.CreateSubKegiatan)
	server.PATCH("/subkegiatan/:id", subKegiatanController.UpdateSubKegiatan)
	server.DELETE("/subkegiatan/:id", subKegiatanController.DeleteSubKegiatan)

	server.GET("/detailsubkegiatan", detailSubKegiatanController.GetDetailSubKegiatans)
	server.GET("/detailsubkegiatan/:id", detailSubKegiatanController.GetDetailSubKegiatan)
	server.GET("/detailsubkegiatanrelasi", detailSubKegiatanController.GetDetailSubKegiatanWithRelation)
	server.POST("/detailsubkegiatan", detailSubKegiatanController.CreateDetailSubKegiatan)
	server.PATCH("/detailsubkegiatan/:id", detailSubKegiatanController.UpdateDetailSubKegiatan)
	server.DELETE("/detailsubkegiatan/:id", detailSubKegiatanController.DeleteDetailSubKegitan)

	server.GET("/detaillokasi", detailLokasiController.GetDetailLokasis)
	server.GET("/detaillokasi/:id", detailLokasiController.GetDetailLokasi)
	server.GET("/detaillokasirelasi", detailLokasiController.GetDetailLokasiWithRelation)
	server.POST("/detaillokasi", detailLokasiController.CreateDetailLokasi)
	server.PATCH("/detaillokasi/:id", detailLokasiController.UpdateDetailLokasi)
	server.DELETE("/detaillokasi/:id", detailLokasiController.DeleteDetailLokasi)

	server.GET("/provinsi", provinsiController.GetProvinsis)
	server.GET("/provinsi/:id", provinsiController.GetProvinsi)

	server.GET("/kabupatenkota", kabupatenKotaController.GetKabupatenKotas)
	server.GET("/kabupatenkota/:id", kabupatenKotaController.GetKabupatenKota)

	server.GET("/kecamatan", kecamatanController.GetKecamatans)
	server.GET("/kecamatan/:id", kecamatanController.GetKecamatan)

	server.GET("/kelurahan", kelurahanController.GetKelurahans)
	server.GET("/kelurahan/:id", kelurahanController.GetKelurahan)

	server.Run(":" + appConfig.AppPort)
}
