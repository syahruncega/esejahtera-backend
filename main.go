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

	var bidangUrusanOnInstansiRepository = repository.NewBidangUrusanOnInstansiRepository(config.DB)
	var bidangUrusanOnInstansiService = service.NewBidangUrusanOnInstansiService(bidangUrusanOnInstansiRepository)
	var bidangUrusanOnInstansiController = controller.NewBidangUrusanOnInstansiController(bidangUrusanOnInstansiService)

	var programRepository = repository.NewProgramRepository(config.DB)
	var programService = service.NewProgramService(programRepository)
	var programController = controller.NewProgramController(programService)

	var rencanaProgramRepository = repository.NewRencanaProgramRepository(config.DB)
	var rencanaProgramService = service.NewRencanaProgramService(rencanaProgramRepository)
	var rencanaProgramController = controller.NewRencanaProgramController(rencanaProgramService)

	var instansiOnProgramRepository = repository.NewInstansiOnProgramRepository(config.DB)
	var instansiOnProgramService = service.NewInstansiOnProgramService(instansiOnProgramRepository)
	var instansiOnProgramController = controller.NewInstansiOnProgramController(instansiOnProgramService)

	var sasaranRepository = repository.NewSasaranRepository(config.DB)
	var sasaranService = service.NewSasaranService(sasaranRepository)
	var sasaranController = controller.NewSasaranController(sasaranService)

	var indikatorSasaranRepository = repository.NewIndikatorSasaranRepository(config.DB)
	var indikatorSasaranService = service.NewIndikatorSasaranService(indikatorSasaranRepository)
	var indikatorSasaranController = controller.NewIndikatorSasaranController(indikatorSasaranService)

	var kebijakanRepository = repository.NewKebijakanRepository(config.DB)
	var kebijakanService = service.NewKebijakanService(kebijakanRepository)
	var kebijakanController = controller.NewKebijakanController(kebijakanService)

	var kegiatanRepository = repository.NewKegiatanRepository(config.DB)
	var kegiatanService = service.NewKegiatanService(kegiatanRepository)
	var kegiatanController = controller.NewKegiatanController(kegiatanService)

	var rencanaKegiatanRepository = repository.NewRencanaKegiatanRepository(config.DB)
	var rencanaKegiatanService = service.NewRencanaKegiatanService(rencanaKegiatanRepository)
	var rencanaKegiatanController = controller.NewRencanaKegiatanController(rencanaKegiatanService)

	var programOnKegiatanRepository = repository.NewProgramOnKegiatanRepository(config.DB)
	var programOnKegiatanService = service.NewProgramOnKegiatanService(programOnKegiatanRepository)
	var programOnKegiatanController = controller.NewProgramOnKegiatanController(programOnKegiatanService)

	var subKegiatanRepository = repository.NewSubKegiatanRepository(config.DB)
	var subKegiatanService = service.NewSubKegiatanService(subKegiatanRepository)
	var subKegiatanController = controller.NewSubKegiatanController(subKegiatanService)

	var rencanaSubKegiatanRepository = repository.NewRencanaSubKegiatanRepository(config.DB)
	var rencanaSubKegiatanService = service.NewRencanaSubKegiatanService(rencanaSubKegiatanRepository)
	var rencanaSubKegiatanController = controller.NewRencanaSubKegiatanController(rencanaSubKegiatanService)

	var kegiatanOnSubKegiatanRepository = repository.NewKegiatanOnSubKegiatanRepository(config.DB)
	var kegiatanOnSubKegiatanService = service.NewKegiatanOnSubKegiatanService(kegiatanOnSubKegiatanRepository)
	var kegiatanOnSubKegiatanController = controller.NewKegiatanOnSubKegiatanController(kegiatanOnSubKegiatanService)

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

	var adminOpdRepository = repository.NewAdminOpdRepository(config.DB)
	var adminOpdService = service.NewAdminOpdService(adminOpdRepository)
	var adminOpdController = controller.NewAdminOpdController(adminOpdService)

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
		AllowOrigins:     []string{"http://localhost:4003", "http://localhost:3003", "https://www.e-sejahtera.info", "https://e-sejahtera.info"},
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

	server.GET("/bidangurusanoninstansi", bidangUrusanOnInstansiController.GetBidangUrusanOnInstansis)
	server.GET("/bidangurusanoninstansi/:id", bidangUrusanOnInstansiController.GetBidangUrusanOnInstansi)
	server.POST("/bidangurusanoninstansi", bidangUrusanOnInstansiController.CreateBidangUrusanOnInstansi)
	server.PATCH("/bidangurusanoninstansi/:id", bidangUrusanOnInstansiController.UpdateBidangUrusanOnInstansi)
	server.DELETE("/bidangurusanoninstansi/:id", bidangUrusanOnInstansiController.DeleteBidangUrusanOnInstansi)

	server.GET("/program", programController.GetPrograms)
	server.GET("/program/:id", programController.GetProgram)
	server.POST("/program", programController.CreateProgram)
	server.PATCH("/program/:id", programController.UpdateProgram)
	server.DELETE("/program/:id", programController.DeleteProgram)

	server.GET("/rencanaprogram", rencanaProgramController.GetRencanaPrograms)
	server.GET("/rencanaprogram/:id", rencanaProgramController.GetRencanaProgram)
	server.GET("/rencanaprogram/search", rencanaProgramController.GetRencanaProgramsBySearch)
	server.POST("/rencanaprogram", rencanaProgramController.CreateRencanaProgram)
	server.PATCH("/rencanaprogram/:id", rencanaProgramController.UpdateRencanaProgram)
	server.DELETE("/rencanaprogram/:id", rencanaProgramController.DeleteRencanaProgram)

	server.GET("/instansionprogram", instansiOnProgramController.GetInstansiOnPrograms)
	server.GET("/instansionprogram/:id", instansiOnProgramController.GetInstansiOnProgram)
	server.POST("/instansionprogram", instansiOnProgramController.CreateInstansiOnProgram)
	server.PATCH("/instansionprogram/:id", instansiOnProgramController.UpdateInstansiOnProgram)
	server.DELETE("/instansionprogram/:id", instansiOnProgramController.DeleteInstansiOnProgram)

	server.GET("/sasaran", sasaranController.GetSasarans)
	server.GET("/sasaran/:id", sasaranController.GetSasaran)
	server.POST("/sasaran", sasaranController.CreateSasaran)
	server.PATCH("/sasaran/:id", sasaranController.UpdateSasaran)
	server.DELETE("/sasaran/:id", sasaranController.DeleteSasaran)

	server.GET("/indikatorsasaran", indikatorSasaranController.GetIndikatorSasarans)
	server.GET("/indikatorsasaran/:id", indikatorSasaranController.GetIndikatorSasaran)
	server.POST("/indikatorsasaran", indikatorSasaranController.CreateIndikatorSasaran)
	server.PATCH("/indikatorsasaran/:id", indikatorSasaranController.UpdateIndikatorSasaran)
	server.DELETE("/indikatorsasaran/:id", indikatorSasaranController.DeleteIndikatorSasaran)

	server.GET("/kebijakan", kebijakanController.GetKebijakans)
	server.GET("/kebijakan/:id", kebijakanController.GetKebijakan)
	server.POST("/kebijakan", kebijakanController.CreateKebijakan)
	server.PATCH("/kebijakan/:id", kebijakanController.UpdateKebijakan)
	server.DELETE("/kebijakan/:id", kebijakanController.DeleteKebijakan)

	server.GET("/kegiatan", kegiatanController.GetKegiatans)
	server.GET("/kegiatan/:id", kegiatanController.GetKegiatan)
	server.POST("/kegiatan", kegiatanController.CreateKegiatan)
	server.PATCH("/kegiatan/:id", kegiatanController.UpdateKegiatan)
	server.DELETE("/kegiatan/:id", kegiatanController.DeleteKegiatan)

	server.GET("/rencanakegiatan", rencanaKegiatanController.GetRencanaKegiatans)
	server.GET("/rencanakegiatan/:id", rencanaKegiatanController.GetRencanaKegiatan)
	server.GET("/rencanakegiatan/search", rencanaKegiatanController.GetRencanaKegiatansBySearch)
	server.POST("/rencanakegiatan", rencanaKegiatanController.CreateRencanaKegiatan)
	server.PATCH("/rencanakegiatan/:id", rencanaKegiatanController.UpdateRencanaKegiatan)
	server.DELETE("/rencanakegiatan/:id", rencanaKegiatanController.DeleteRencanaKegiatan)

	server.GET("/programonkegiatan", programOnKegiatanController.GetProgramOnKegiatans)
	server.GET("/programonkegiatan/:id", programOnKegiatanController.GetProgramOnKegiatan)
	server.POST("/programonkegiatan", programOnKegiatanController.CreateProgramOnKegiatan)
	server.PATCH("/programonkegiatan/:id", programOnKegiatanController.UpdateProgramOnKegiatan)
	server.DELETE("/programonkegiatan/:id", programOnKegiatanController.DeleteProgramOnKegiatan)

	server.GET("/subkegiatan", subKegiatanController.GetSubKegiatans)
	server.GET("/subkegiatan/:id", subKegiatanController.GetSubKegiatan)
	server.POST("/subkegiatan", subKegiatanController.CreateSubKegiatan)
	server.PATCH("/subkegiatan/:id", subKegiatanController.UpdateSubKegiatan)
	server.DELETE("/subkegiatan/:id", subKegiatanController.DeleteSubKegiatan)

	server.GET("/rencanasubkegiatan", rencanaSubKegiatanController.GetRencanaSubKegiatans)
	server.GET("/rencanasubkegiatan/:id", rencanaSubKegiatanController.GetRencanaSubKegiatan)
	server.GET("/rencanasubkegiatan/search", rencanaSubKegiatanController.GetRencanaSubKegiatansBySearch)
	server.POST("/rencanasubkegiatan", rencanaSubKegiatanController.CreateRencanaSubKegiatan)
	server.PATCH("/rencanasubkegiatan/:id", rencanaSubKegiatanController.UpdateRencanaSubKegiatan)
	server.DELETE("/rencanasubkegiatan/:id", rencanaSubKegiatanController.DeleteRencanaSubkKegiatan)

	server.GET("/kegiatanonsubkegiatan", kegiatanOnSubKegiatanController.GetKegiatanOnSubKegiatans)
	server.GET("/kegiatanonsubkegiatan/:id", kegiatanOnSubKegiatanController.GetKegiatanOnSubKegiatan)
	server.POST("/kegiatanonsubkegiatan", kegiatanOnSubKegiatanController.CreateKegiatanOnSubKegiatan)
	server.PATCH("/kegiatanonsubkegiatan/:id", kegiatanOnSubKegiatanController.UpdateKegiatanOnSubKegiatan)
	server.DELETE("/kegiatanonsubkegiatan/:id", kegiatanOnSubKegiatanController.DeleteKegiatanOnSubKegiatan)

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
	server.POST("/user/batch", userController.CreateBatchUser)
	server.PATCH("/user/:id", userController.UpdateUser)
	server.DELETE("/user/:id", userController.DeleteUser)

	server.GET("/admin", adminController.GetAdmins)
	server.GET("/admin/:id", adminController.GetAdmin)
	server.GET("/adminrelasi", adminController.GetAdminWithRelation)
	server.POST("/admin", adminController.CreateAdmin)
	server.PATCH("/admin/:id", adminController.UpdateAdmin)
	server.DELETE("/admin/:id", adminController.DeleteAdmin)

	server.GET("/adminopd", adminOpdController.GetAdminOpd)
	server.GET("/adminopd/:id", adminOpdController.GetAdminOpd)
	server.GET("/adminopd/user/:userid", adminOpdController.GetAdminOpdByUserId)
	server.POST("/adminopd", adminOpdController.CreateAdminOpd)
	server.PATCH("/adminopd/:id", adminOpdController.UpdateAdminOpd)
	server.DELETE("/adminopd/:id", adminOpdController.DeleteAdminOpd)

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
	server.POST("/mahasiswa/batch", mahasiswaController.CreateBatchMahasiswa)
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
