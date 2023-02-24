package main

import (
	"kemiskinan/config"
	"kemiskinan/controller"
	"kemiskinan/middleware"
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
	appConfig.DatabaseHost = os.Getenv("APP_DATABASE_HOST")
	appConfig.AppPort = os.Getenv("APP_PORT")

	config.ConnectDB(appConfig)

	var bidangUrusanRepository = repository.NewBidangUrusanRepository(config.DB)
	var bidangUrusanService = service.NewBidangUrusanService(bidangUrusanRepository)
	var bidangUrusanController = controller.NewBidangUrusanController(bidangUrusanService)

	var instansiRepository = repository.NewInstansiRepository(config.DB)
	var instansiService = service.NewInstansiService(instansiRepository)
	var instansiController = controller.NewInstansiController(instansiService)

	var detailInstansiRepository = repository.NewDetailInstansiRepository(config.DB)
	var detailInstansiService = service.NewDetailInstansiService(detailInstansiRepository)
	var detailInstansiController = controller.NewDetailInstansiController(detailInstansiService)

	var programRepository = repository.NewProgramRepository(config.DB)
	var programService = service.NewProgramService(programRepository)
	var programController = controller.NewProgramController(programService)

	var detailProgramRepository = repository.NewDetailProgramRepository(config.DB)
	var detailProgramService = service.NewDetailProgramService(detailProgramRepository)
	var detailProgramController = controller.NewDetailProgramController(detailProgramService)

	var indikatorProgramRepository = repository.NewIndikatorProgramRepository(config.DB)
	var indikatorProgramService = service.NewIndikatorProgramService(indikatorProgramRepository)
	var indikatorProgramController = controller.NewIndikatorProgramController(indikatorProgramService)

	var kegiatanRepository = repository.NewKegiatanRepository(config.DB)
	var kegiatanService = service.NewKegiatanService(kegiatanRepository)
	var kegiatanController = controller.NewKegiatanController(kegiatanService)

	var detailKegiatanRepository = repository.NewDetailKegiatanRepository(config.DB)
	var detailKegiatanService = service.NewDetailKegiatanService(detailKegiatanRepository)
	var detailKegiatanController = controller.NewDetailKegiatanController(detailKegiatanService)

	var indikatorKegiatanRepository = repository.NewIndikatorKegiatanRepository(config.DB)
	var indikatorKegiatanService = service.NewIndikatorKegiatanService(indikatorKegiatanRepository)
	var indikatorKegiatanController = controller.NewIndikatorKegiatanController(indikatorKegiatanService)

	var subKegiatanRepository = repository.NewSubKegiatanRepository(config.DB)
	var subKegiatanService = service.NewSubKegiatanService(subKegiatanRepository)
	var subKegiatanController = controller.NewSubKegiatanController(subKegiatanService)

	var indikatorSubKegiatanRepository = repository.NewIndikatorSubKegiatanRepository(config.DB)
	var indikatorSubKegiatanService = service.NewIndikatorSubKegiatanService(indikatorSubKegiatanRepository)
	var indikatorSubKegiatanController = controller.NewIndikatorSubKegiatanController(indikatorSubKegiatanService)

	var detailSubKegiatanRepository = repository.NewDetailSubKegiatanRepository(config.DB)
	var detailSubKegiatanService = service.NewDetailSubKegiatanService(detailSubKegiatanRepository)
	var detailSubKegiatanController = controller.NewDetailSubKegiatanController(detailSubKegiatanService)

	var fokusBelanjaRepository = repository.NewFokusBelanjaRepository(config.DB)
	var fokusBelanjaService = service.NewFokusBelanjaService(fokusBelanjaRepository)
	var fokusBelanjaController = controller.NewFokusBelanjaController(fokusBelanjaService)

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

	var monevRepository = repository.NewMonevRepository(config.DB)
	var monevService = service.NewMonevService(monevRepository)
	var monevController = controller.NewMonevController(monevService)

	var adminRepository = repository.NewAdminRepository(config.DB)
	var adminService = service.NewAdminService(adminRepository)
	var adminController = controller.NewAdminController(adminService)

	var analisRepository = repository.NewAnalisRepository(config.DB)
	var analisService = service.NewAnalisService(analisRepository)
	var analisController = controller.NewAnalisController(analisService)

	var lokasiDosenRepository = repository.NewLokasiDosenRepository(config.DB)
	var lokasiDosenService = service.NewLokasiDosenService(lokasiDosenRepository)
	var lokasiDosenController = controller.NewLokasiDosenController(lokasiDosenService)

	var keluargaVerifikasiRepository = repository.NewKeluargaVerifikasiRepository(config.DB)
	var keluargaVerifikasiService = service.NewKeluargaVerifikasiService(keluargaVerifikasiRepository)
	var keluargaVerifikasiController = controller.NewKeluargaVerifikasiController(keluargaVerifikasiService)

	var individuRepository = repository.NewIndividuRepository(config.DB)
	var individuService = service.NewIndividuService(individuRepository)
	var individuController = controller.NewIndividuController(individuService)

	var keluargaRepository = repository.NewKeluargaRepository(config.DB)
	var keluargaService = service.NewKeluargaService(keluargaRepository)
	var keluargaController = controller.NewKeluargaController(keluargaService, individuService)

	var individuVerifikasiRepository = repository.NewIndividuVerifikasiRepository(config.DB)
	var individuVerifikasiService = service.NewIndividuVerifikasiService(individuVerifikasiRepository)
	var individuVerifikasiController = controller.NewIndividuVerifikasiController(individuVerifikasiService)

	var mahasiswaRepository = repository.NewMahasiswaRepository(config.DB)
	var mahasiswaService = service.NewMahasiswaService(mahasiswaRepository)
	var mahasiswaController = controller.NewMahasiswaController(mahasiswaService, keluargaService, individuService)

	var dosenRepository = repository.NewDosenRepository(config.DB)
	var dosenService = service.NewDosenService(dosenRepository)
	var dosenController = controller.NewDosenController(dosenService, mahasiswaService)

	var pusbangRepository = repository.NewPusbangRepository(config.DB)
	var pusbangService = service.NewPusbangService(pusbangRepository)
	var pusbangController = controller.NewPusbangController(pusbangService, dosenService, keluargaService, individuService)

	var userRepository = repository.NewUserRepository(config.DB)
	var userService = service.NewUserService(userRepository)
	var userController = controller.NewUserController(userService, adminService, pusbangService, dosenService, analisService, mahasiswaService)

	var statistikController = controller.NewStatistikController(keluargaService, individuService)

	var server = gin.Default()

	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3003", "http://localhost:3003", "https://www.e-sejahtera.info", "https://e-sejahtera.info"},
		AllowCredentials: true,
		AllowMethods:     []string{"POST", "PUT", "DELETE", "GET", "PATCH"},
		AllowHeaders:     []string{"Access-Control-Allow-Headers", "Access-Control-Allow-Origin", "access-control-allow-origin, access-control-allow-headers", "Content-Type", "Accept", "Origin", "Authorization"},
	}))

	server.GET("/bidangurusan", bidangUrusanController.GetBidangUrusans)
	server.GET("/bidangurusan/:id", bidangUrusanController.GetBidangUrusan)
	server.POST("/bidangurusan", bidangUrusanController.CreateBidangUrusan)
	server.PATCH("/bidangurusan/:id", bidangUrusanController.UpdateBidangUrusan)
	server.DELETE("/bidangurusan/:id", bidangUrusanController.DeleteBidangUrusan)

	server.GET("/instansi", instansiController.GetInstansis)
	server.GET("/instansi/:id", instansiController.GetInstansi)
	server.POST("/instansi", instansiController.CreateInstansi)
	server.PATCH("/instansi/:id", instansiController.UpdateInstansi)
	server.DELETE("/instansi/:id", instansiController.DeleteInstansi)

	server.GET("/detailinstansi", detailInstansiController.GetDetailInstansis)
	server.GET("/detailinstansi/:id", detailInstansiController.GetDetailInstansi)
	server.POST("/detailinstansi", detailInstansiController.CreateDetailInstansi)
	server.PATCH("/detailinstansi/:id", detailInstansiController.UpdateDetailInstansi)
	server.DELETE("/detailinstansi/:id", detailInstansiController.DeleteDetailInstansi)

	server.GET("/program", programController.GetPrograms)
	server.GET("/program/:id", programController.GetProgram)
	server.POST("/program", programController.CreateProgram)
	server.PATCH("/program/:id", programController.UpdateProgram)
	server.DELETE("/program/:id", programController.DeleteProgram)

	server.GET("/detailprogram", detailProgramController.GetDetailPrograms)
	server.GET("/detailprogram/:id", detailProgramController.GetDetailProgram)
	server.POST("/detailprogram", detailProgramController.CreateDetailProgram)
	server.PATCH("/detailprogram/:id", detailProgramController.UpdateDetailProgram)
	server.DELETE("/detailprogram/:id", detailProgramController.DeleteDetailProgram)

	server.GET("/indikatorprogram", indikatorProgramController.GetIndikatorPrograms)
	server.GET("/indikatorprogram/:id", indikatorProgramController.GetIndikatorProgram)
	server.POST("/indikatorprogram", indikatorProgramController.CreateIndikatorProgram)
	server.PATCH("/indikatorprogram/:id", indikatorProgramController.UpdateIndikatorProgram)
	server.DELETE("/indikatorprogram/:id", indikatorProgramController.DeleteIndikatorProgram)

	server.GET("/kegiatan", kegiatanController.GetKegiatans)
	server.GET("/kegiatan/:id", kegiatanController.GetKegiatan)
	server.POST("/kegiatan", kegiatanController.CreateKegiatan)
	server.PATCH("/kegiatan/:id", kegiatanController.UpdateKegiatan)
	server.DELETE("/kegiatan/:id", kegiatanController.DeleteKegiatan)

	server.GET("/indikatorsubkegiatan", indikatorSubKegiatanController.GetIndikatorSubKegiatans)
	server.GET("/indikatorsubkegiatan/:id", indikatorSubKegiatanController.GetIndikatorSubKegiatan)
	server.POST("/indikatorsubkegiatan", indikatorSubKegiatanController.CreateIndikatorSubKegiatan)
	server.PATCH("/indikatorsubkegiatan/:id", indikatorSubKegiatanController.UpdateIndikatorSubKegiatan)
	server.DELETE("/indikatorsubkegiatan/:id", indikatorSubKegiatanController.DeleteIndikatorSubKegiatan)

	server.GET("/detailkegiatan", detailKegiatanController.GetDetailKegiatans)
	server.GET("/detailkegiatan/:id", detailKegiatanController.GetDetailKegiatan)
	server.POST("/detailkegiatan", detailKegiatanController.CreateDetailKegiatan)
	server.PATCH("/detailkegiatan/:id", detailKegiatanController.UpdateDetailKegiatan)
	server.DELETE("/detailkegiatan/:id", detailKegiatanController.DeleteDetailKegiatan)

	server.GET("/indikatorkegiatan", indikatorKegiatanController.GetIndikatorKegiatans)
	server.GET("/indikatorkegiatan/:id", indikatorKegiatanController.GetIndikatorKegiatan)
	server.POST("/indikatorkegiatan", indikatorKegiatanController.CreateIndikatorKegiatan)
	server.PATCH("/indikatorkegiatan/:id", indikatorKegiatanController.UpdateIndikatorKegiatan)
	server.DELETE("/indikatorkegiatan/:id", indikatorKegiatanController.DeleteIndikatorKegiatan)

	server.GET("/subkegiatan", subKegiatanController.GetSubKegiatans)
	server.GET("/subkegiatan/:id", subKegiatanController.GetSubKegiatan)
	server.POST("/subkegiatan", subKegiatanController.CreateSubKegiatan)
	server.PATCH("/subkegiatan/:id", subKegiatanController.UpdateSubKegiatan)
	server.DELETE("/subkegiatan/:id", subKegiatanController.DeleteSubKegiatan)

	server.GET("/detailsubkegiatan", detailSubKegiatanController.GetDetailSubKegiatans)
	server.GET("/detailsubkegiatan/:id", detailSubKegiatanController.GetDetailSubKegiatan)
	server.POST("/detailsubkegiatan", detailSubKegiatanController.CreateDetailSubKegiatan)
	server.PATCH("/detailsubkegiatan/:id", detailSubKegiatanController.UpdateDetailSubKegiatan)
	server.DELETE("/detailsubkegiatan/:id", detailSubKegiatanController.DeleteDetailSubKegiatan)

	server.GET("/fokusbelanja", fokusBelanjaController.GetFokusBelanjas)
	server.GET("/fokusbelanja/:id", fokusBelanjaController.GetFokusBelanja)
	server.POST("/fokusbelanja", fokusBelanjaController.CreateFokusBelanja)
	server.PATCH("/fokusbelanja/:id", fokusBelanjaController.UpdateFokusBelanja)
	server.DELETE("/fokusbelanja/:id", fokusBelanjaController.DeleteFokusBelanja)

	server.GET("/detaillokasi", detailLokasiController.GetDetailLokasis)
	server.GET("/detaillokasi/:id", detailLokasiController.GetDetailLokasi)
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

	server.GET("/keluarga/", keluargaController.GetKeluargas)
	server.GET("/keluarga/:id", keluargaController.GetKeluargaById)
	server.GET("/keluarga/idkeluarga/:idkeluarga", keluargaController.GetKeluargaByIdKeluarga)
	server.GET("/keluarga/search", keluargaController.GetKeluargaBySearch)
	server.PATCH("/keluarga/:id", keluargaController.UpdateKeluarga)
	// server.GET("/testdistinct/", keluargaController.TestDistinct)

	server.GET("/monev/:kabupatenkotaid", monevController.GetMonevs)
	server.GET("/monev/detail/:kabupatenkotaid/:id", monevController.GetMonev)

	server.GET("/user", userController.GetUsers)
	server.GET("/user/:id", userController.GetUser)
	server.GET("/auth/session", middleware.CheckAuth, userController.Validate)
	server.GET("/auth/profile", middleware.CheckAuth, userController.GetUserProfile)
	server.POST("/auth/login", userController.LoginUser)
	server.POST("/auth/logout", userController.LogoutUser)
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
	server.GET("/dosen/findmahasiswa", dosenController.GetMahasiswa)
	server.POST("/dosen", dosenController.CreateDosen)
	server.PATCH("/dosen/:id", dosenController.UpdateDosen)
	server.DELETE("/dosen/:id", dosenController.DeleteDosen)

	server.GET("/pusbang", pusbangController.GetPusbangs)
	server.GET("/pusbang/:id", pusbangController.GetPusbang)
	server.GET("/pusbangrelasi", pusbangController.GetPusbangWithRelation)
	server.GET("/pusbang/dosen", pusbangController.GetDistinctDosen)
	server.GET("/pusbang/lokasidosen", pusbangController.GetDistinctLokasiDosen)
	server.POST("/pusbang", pusbangController.CreatePusbang)
	server.PATCH("/pusbang:id", pusbangController.UpdatePusbang)
	server.DELETE("/pusbang/:id", pusbangController.DeletePusbang)

	server.GET("/mahasiswa", mahasiswaController.GetMahasiswas)
	server.GET("/mahasiswa/:id", mahasiswaController.GetMahasiswa)
	server.GET("/mahasiswarelasi", mahasiswaController.GetMahasiswaWithRelation)
	server.GET("/mahasiswa/verifying", mahasiswaController.GetVerifiedByMahasiswa)
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

	server.GET("/keluargaverifikasi/", keluargaVerifikasiController.GetKeluargaVerifikasis)
	server.GET("/keluargaverifikasi/:id", keluargaVerifikasiController.GetKeluargaVerifikasi)
	server.GET("/keluargaverifikasi/idkeluarga/:idkeluarga", keluargaVerifikasiController.GetKeluargaVerifikasiByIdKeluarga)
	server.GET("/keluargaverifikasi/search", keluargaVerifikasiController.GetKeluargaVerifikasiBySearch)
	server.POST("/keluargaverifikasi", keluargaVerifikasiController.CreateKeluargaVerifikasi)
	server.PATCH("/keluargaverifikasi/:id", keluargaVerifikasiController.UpdateKeluargaVerifikasi)
	server.DELETE("/keluargaverifikasi/:id", keluargaVerifikasiController.DeleteKeluargaVerifikasi)

	server.GET("/individu/", individuController.GetIndividus)
	server.GET("/individu/:id", individuController.GetIndividu)
	server.GET("/individu/idkeluarga/:idkeluarga", individuController.GetIndividuByIdKeluarga)
	server.GET("/individu/search", individuController.GetIndividuBySearch)
	server.PATCH("/individu/:id", individuController.UpdateIndividu)

	server.GET("/individuverifikasi", individuVerifikasiController.GetIndividuVerifikasis)
	server.GET("/individuverifikasi/:id", individuVerifikasiController.GetIndividuVerifikasi)
	server.GET("/individuverifikasi/idkeluarga/:idkeluarga", individuVerifikasiController.GetIndividuVerifikasiByIdKeluarga)
	server.GET("/individuverifikasi/search", individuVerifikasiController.GetIndividuVerifikasiBySearch)
	server.POST("/individuverifikasi", individuVerifikasiController.CreateIndividuVerifikasi)
	server.PATCH("/individuverifikasi/:id", individuVerifikasiController.UpdateIndividuVerifikasi)
	server.DELETE("/individuverifikasi/:id", individuVerifikasiController.DeleteIndividuVerifikasi)

	server.GET("/statistik/provinsi", statistikController.StatistikProvinsi)
	server.GET("/statistik/kabupatenkota", statistikController.StatistikKabupatenKota)
	server.GET("/statistik/kecamatan", statistikController.StatistikKecamatan)
	server.GET("/statistik/kelurahan", statistikController.StatistikKelurahan)
	server.GET("/statistik/kelurahan/hitung", statistikController.StatistikSearch)

	server.Run(":" + appConfig.AppPort)
}
