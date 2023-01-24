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

	var keluargaRepository = repository.NewKeluargaRepository(config.DB)
	var keluargaService = service.NewKeluargaService(keluargaRepository)
	var keluargaController = controller.NewKeluargaController(keluargaService)

	var monevRepository = repository.NewMonevRepository(config.DB)
	var monevService = service.NewMonevService(monevRepository)
	var monevController = controller.NewMonevController(monevService)

	var userRepository = repository.NewUserRepository(config.DB)
	var userService = service.NewUserService(userRepository)
	var userController = controller.NewUserController(userService)

	var adminRepository = repository.NewAdminRepository(config.DB)
	var adminService = service.NewAdminService(adminRepository)
	var adminController = controller.NewAdminController(adminService)

	var analisRepository = repository.NewAnalisRepository(config.DB)
	var analisService = service.NewAnalisService(analisRepository)
	var analisController = controller.NewAnalisController(analisService)

	var dosenRepository = repository.NewDosenRepository(config.DB)
	var dosenService = service.NewDosenService(dosenRepository)
	var dosenController = controller.NewDosenController(dosenService)

	var pusbangRepository = repository.NewPusbangRepository(config.DB)
	var pusbangService = service.NewPusbangService(pusbangRepository)
	var pusbangController = controller.NewPusbangController(pusbangService)

	var mahasiswaRepository = repository.NewMahasiswaRepository(config.DB)
	var mahasiswaService = service.NewMahasiswaService(mahasiswaRepository)
	var mahasiswaController = controller.NewMahasiswaController(mahasiswaService)

	var lokasiDosenRepository = repository.NewLokasiDosenRepository(config.DB)
	var lokasiDosenService = service.NewLokasiDosenService(lokasiDosenRepository)
	var lokasiDosenController = controller.NewLokasiDosenController(lokasiDosenService)

	var keluargaVerifikasiRepository = repository.NewKeluargaVerifikasiRepository(config.DB)
	var keluargaVerifikasiService = service.NewKeluargaVerifikasiService(keluargaVerifikasiRepository)
	var keluargaVerifikasiController = controller.NewKeluargaVerifikasiController(keluargaVerifikasiService)

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

	server.GET("/keluarga/:kabupatenkotaid", keluargaController.GetKeluargas)
	server.GET("/keluarga/detail/:kabupatenkotaid/:id", keluargaController.GetKeluargaById)
	server.GET("/keluarga/idkeluarga/:kabupatenkotaid/:idkeluarga", keluargaController.GetIdKeluargaByKabupatenKota)
	server.GET("/keluarga/jumlah/penerima/kabupatenkota/:kabupatenkotaid/:penerimaparameter/:nilai", keluargaController.CountPenerimaByKabupatenKota)
	server.GET("/keluarga/jumlah/penerima/kecamatan/:kecamatanid/:penerimaparameter/:nilai", keluargaController.CountPenerimaByKecamatan)
	server.GET("/keluarga/jumlah/penerima/kelurahan/:kelurahanid/:penerimaparameter/:nilai", keluargaController.CountPenerimaByKelurahan)
	server.GET("/keluarga/jumlah/desil/kabupatenkota/:kabupatenkotaid/:nilaidesil", keluargaController.CountDesilByKabupatenKota)
	server.GET("/keluarga/jumlah/desil/kecamatan/:kecamatanid/:nilaidesil", keluargaController.CountDesilByKecamatan)
	server.GET("/keluarga/jumlah/desil/kelurahan/:kelurahanid/:nilaidesil", keluargaController.CountDesilByKelurahan)
	server.PATCH("/keluarga/:kabupatenkotaid/:id", keluargaController.UpdateKeluarga)

	server.GET("/monev/:kabupatenkotaid", monevController.GetMonevs)
	server.GET("/monev/detail/:kabupatenkotaid/:id", monevController.GetMonev)

	server.GET("/user", userController.GetUsers)
	server.GET("/user/:id", userController.GetUser)
	server.POST("/user", userController.CreateUser)
	server.PATCH("/user/:id", userController.UpdateUser)
	server.DELETE("/user/:id", userController.DeleteUser)

	server.GET("/admin", adminController.GetAdmins)
	server.GET("/admin/:id", adminController.GetAdmin)
	server.GET("/adminrelasi", adminController.GetAdminWithRelation)
	server.POST("/admin", adminController.CreateAdmin)
	server.PATCH("/admin/:id", adminController.UpdateAdmin)
	server.DELETE("/admin/:id", adminController.DeleteAdmin)

	server.GET("/analis", analisController.GetAnaliss)
	server.GET("/analis/:id", analisController.GetAnalis)
	server.GET("/analisrelasi", analisController.GetAnalisWithRelation)
	server.POST("/analis", analisController.CreateAnalis)
	server.PATCH("/analis/:id", analisController.UpdateAnalis)
	server.DELETE("/analis/:id", analisController.DeleteAnalis)

	server.GET("/dosen", dosenController.GetDosens)
	server.GET("/dosen/:id", dosenController.GetDosen)
	server.GET("/dosenrelasi", dosenController.GetDosenWithRelation)
	server.POST("/dosen", dosenController.CreateDosen)
	server.PATCH("/dosen/:id", dosenController.UpdateDosen)
	server.DELETE("/dosen/:id", dosenController.DeleteDosen)

	server.GET("/pusbang", pusbangController.GetPusbangs)
	server.GET("/pusbang/:id", pusbangController.GetPusbang)
	server.GET("/pusbangrelasi", pusbangController.GetPusbangWithRelation)
	server.POST("/pusbang", pusbangController.CreatePusbang)
	server.PATCH("/pusbang:id", pusbangController.UpdatePusbang)
	server.DELETE("/pusbang/:id", pusbangController.DeletePusbang)

	server.GET("/mahasiswa", mahasiswaController.GetMahasiswas)
	server.GET("/mahasiswa/:id", mahasiswaController.GetMahasiswa)
	server.GET("/mahasiswarelasi", mahasiswaController.GetMahasiswaWithRelation)
	server.POST("/mahasiswa", mahasiswaController.CreateMahasiswa)
	server.PATCH("/mahasiswa/:id", mahasiswaController.UpdateMahasiswa)
	server.DELETE("/mahasiswa/:id", mahasiswaController.DeleteMahasiswa)

	server.GET("/lokasidosen", lokasiDosenController.GetLokasiDosens)
	server.GET("/lokasidosen/:id", lokasiDosenController.GetLokasiDosen)
	server.GET("/lokasidosenrelasi", lokasiDosenController.GetLokasiDosenWithRelation)
	server.GET("/lokasidosen/dosen/:dosenid", lokasiDosenController.GetLokasiDosenRelationByDosenId)
	server.POST("/lokasidosen", lokasiDosenController.CreateLokasiDosen)
	server.PATCH("/lokasidosen/:id", lokasiDosenController.UpdateLokasiDosen)
	server.DELETE("/lokasidosen/:id", lokasiDosenController.DeleteLokasiDosen)

	server.GET("/keluargaverifikasi/:kabupatenkotaid", keluargaVerifikasiController.GetKeluargaVerifikasis)
	server.GET("/keluargaverifikasi/detail/:kabupatenkotaid/:id", keluargaVerifikasiController.GetKeluargaVerifikasi)
	server.POST("/keluargaverifikasi", keluargaVerifikasiController.CreateKeluargaVerifikasi)
	server.PATCH("/keluargaverifikasi/:kabupatenkotaid/:id", keluargaVerifikasiController.UpdateKeluargaVerifikasi)
	server.DELETE("/keluargaverifikasi/:kabupatenkotaid/:id", keluargaVerifikasiController.DeleteKeluargaVerifikasi)

	server.Run(":" + appConfig.AppPort)
}
