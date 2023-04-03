package main

// import (
// 	"kemiskinan/config"
// 	"kemiskinan/controller"
// 	"kemiskinan/middleware"
// 	"kemiskinan/repository"
// 	"kemiskinan/service"
// 	"log"
// 	"os"

// 	"github.com/gin-contrib/cors"
// 	"github.com/gin-gonic/gin"
// 	"github.com/joho/godotenv"
// )

// func main() {

// 	var appConfig = config.AppConfig{}

// 	var err = godotenv.Load()
// 	if err != nil {
// 		log.Fatal("Terdapat kesalahan memuat file .env")
// 	}

// 	appConfig.DatabaseUsername = os.Getenv("APP_DATABASE_USERNAME")
// 	appConfig.DatabasePassword = os.Getenv("APP_DATABASE_PASSWORD")
// 	appConfig.DatabasePort = os.Getenv("APP_DATABASE_PORT")
// 	appConfig.DatabaseName = os.Getenv("APP_DATABASE_NAME")
// 	appConfig.DatabaseHost = os.Getenv("APP_DATABASE_HOST")
// 	appConfig.AppPort = os.Getenv("APP_PORT")

// 	config.ConnectDB(appConfig)

// 	var bidangUrusanRepository = repository.NewBidangUrusanRepository(config.DB)
// 	var bidangUrusanService = service.NewBidangUrusanService(bidangUrusanRepository)
// 	var bidangUrusanController = controller.NewBidangUrusanController(bidangUrusanService)

// 	var instansiRepository = repository.NewInstansiRepository(config.DB)
// 	var instansiService = service.NewInstansiService(instansiRepository)
// 	var instansiController = controller.NewInstansiController(instansiService)

// 	var bidangUrusanOnInstansiRepository = repository.NewBidangUrusanOnInstansiRepository(config.DB)
// 	var bidangUrusanOnInstansiService = service.NewBidangUrusanOnInstansiService(bidangUrusanOnInstansiRepository)
// 	var bidangUrusanOnInstansiController = controller.NewBidangUrusanOnInstansiController(bidangUrusanOnInstansiService)

// 	var programRepository = repository.NewProgramRepository(config.DB)
// 	var programService = service.NewProgramService(programRepository)
// 	var programController = controller.NewProgramController(programService)

// 	var rencanaProgramRepository = repository.NewRencanaProgramRepository(config.DB)
// 	var rencanaProgramService = service.NewRencanaProgramService(rencanaProgramRepository)
// 	var rencanaProgramController = controller.NewRencanaProgramController(rencanaProgramService)

// 	var instansiOnProgramRepository = repository.NewInstansiOnProgramRepository(config.DB)
// 	var instansiOnProgramService = service.NewInstansiOnProgramService(instansiOnProgramRepository)
// 	var instansiOnProgramController = controller.NewInstansiOnProgramController(instansiOnProgramService)

// 	var sasaranRepository = repository.NewSasaranRepository(config.DB)
// 	var sasaranService = service.NewSasaranService(sasaranRepository)
// 	var sasaranController = controller.NewSasaranController(sasaranService)

// 	var indikatorSasaranRepository = repository.NewIndikatorSasaranRepository(config.DB)
// 	var indikatorSasaranService = service.NewIndikatorSasaranService(indikatorSasaranRepository)
// 	var indikatorSasaranController = controller.NewIndikatorSasaranController(indikatorSasaranService)

// 	var kebijakanRepository = repository.NewKebijakanRepository(config.DB)
// 	var kebijakanService = service.NewKebijakanService(kebijakanRepository)
// 	var kebijakanController = controller.NewKebijakanController(kebijakanService)

// 	var kegiatanRepository = repository.NewKegiatanRepository(config.DB)
// 	var kegiatanService = service.NewKegiatanService(kegiatanRepository)
// 	var kegiatanController = controller.NewKegiatanController(kegiatanService)

// 	var rencanaKegiatanRepository = repository.NewRencanaKegiatanRepository(config.DB)
// 	var rencanaKegiatanService = service.NewRencanaKegiatanService(rencanaKegiatanRepository, rencanaProgramRepository)
// 	var rencanaKegiatanController = controller.NewRencanaKegiatanController(rencanaKegiatanService)

// 	var programOnKegiatanRepository = repository.NewProgramOnKegiatanRepository(config.DB)
// 	var programOnKegiatanService = service.NewProgramOnKegiatanService(programOnKegiatanRepository)
// 	var programOnKegiatanController = controller.NewProgramOnKegiatanController(programOnKegiatanService)

// 	var subKegiatanRepository = repository.NewSubKegiatanRepository(config.DB)
// 	var subKegiatanService = service.NewSubKegiatanService(subKegiatanRepository)
// 	var subKegiatanController = controller.NewSubKegiatanController(subKegiatanService)

// 	var rencanaSubKegiatanRepository = repository.NewRencanaSubKegiatanRepository(config.DB)
// 	var rencanaSubKegiatanService = service.NewRencanaSubKegiatanService(rencanaSubKegiatanRepository, rencanaKegiatanRepository)
// 	var rencanaSubKegiatanController = controller.NewRencanaSubKegiatanController(rencanaSubKegiatanService)

// 	var kegiatanOnSubKegiatanRepository = repository.NewKegiatanOnSubKegiatanRepository(config.DB)
// 	var kegiatanOnSubKegiatanService = service.NewKegiatanOnSubKegiatanService(kegiatanOnSubKegiatanRepository)
// 	var kegiatanOnSubKegiatanController = controller.NewKegiatanOnSubKegiatanController(kegiatanOnSubKegiatanService)

// 	var fokusBelanjaRepository = repository.NewFokusBelanjaRepository(config.DB)
// 	var fokusBelanjaService = service.NewFokusBelanjaService(fokusBelanjaRepository, rencanaSubKegiatanRepository)
// 	var fokusBelanjaController = controller.NewFokusBelanjaController(fokusBelanjaService)

// 	var detailLokasiRepository = repository.NewDetailLokasiRepository(config.DB)
// 	var detailLokasiService = service.NewDetailLokasiService(detailLokasiRepository)
// 	var detailLokasiController = controller.NewDetailLokasiController(detailLokasiService)

// 	var provinsiRepository = repository.NewProvinsiRepository(config.DB)
// 	var provinsiService = service.NewProvinsiService(provinsiRepository)
// 	var provinsiController = controller.NewProvinsiController(provinsiService)

// 	var kabupatenKotaRepository = repository.NewKabupatenKotaRepository(config.DB)
// 	var kabupatenKotaService = service.NewKabupatenKotaService(kabupatenKotaRepository)
// 	var kabupatenKotaController = controller.NewKabupatenKotaController(kabupatenKotaService)

// 	var kecamatanRepository = repository.NewKecamatanRepository(config.DB)
// 	var kecamatanService = service.NewKecamatanService(kecamatanRepository)
// 	var kecamatanController = controller.NewKecamatanController(kecamatanService)

// 	var kelurahanRepository = repository.NewKelurahanRepository(config.DB)
// 	var kelurahanService = service.NewKelurahanService(kelurahanRepository)
// 	var kelurahanController = controller.NewKelurahanController(kelurahanService)

// 	var monevRepository = repository.NewMonevRepository(config.DB)
// 	var monevService = service.NewMonevService(monevRepository)
// 	var monevController = controller.NewMonevController(monevService)

// 	var adminRepository = repository.NewAdminRepository(config.DB)
// 	var adminService = service.NewAdminService(adminRepository)
// 	var adminController = controller.NewAdminController(adminService)

// 	var adminOpdRepository = repository.NewAdminOpdRepository(config.DB)
// 	var adminOpdService = service.NewAdminOpdService(adminOpdRepository)
// 	var adminOpdController = controller.NewAdminOpdController(adminOpdService)

// 	var analisRepository = repository.NewAnalisRepository(config.DB)
// 	var analisService = service.NewAnalisService(analisRepository)
// 	var analisController = controller.NewAnalisController(analisService)

// 	var lokasiDosenRepository = repository.NewLokasiDosenRepository(config.DB)
// 	var lokasiDosenService = service.NewLokasiDosenService(lokasiDosenRepository)
// 	var lokasiDosenController = controller.NewLokasiDosenController(lokasiDosenService)

// 	var keluargaVerifikasiRepository = repository.NewKeluargaVerifikasiRepository(config.DB)
// 	var keluargaVerifikasiService = service.NewKeluargaVerifikasiService(keluargaVerifikasiRepository)
// 	var keluargaVerifikasiController = controller.NewKeluargaVerifikasiController(keluargaVerifikasiService)

// 	var individuRepository = repository.NewIndividuRepository(config.DB)
// 	var individuService = service.NewIndividuService(individuRepository)
// 	var individuController = controller.NewIndividuController(individuService)

// 	var keluargaRepository = repository.NewKeluargaRepository(config.DB)
// 	var keluargaService = service.NewKeluargaService(keluargaRepository)
// 	var keluargaController = controller.NewKeluargaController(keluargaService, individuService)

// 	var individuVerifikasiRepository = repository.NewIndividuVerifikasiRepository(config.DB)
// 	var individuVerifikasiService = service.NewIndividuVerifikasiService(individuVerifikasiRepository)
// 	var individuVerifikasiController = controller.NewIndividuVerifikasiController(individuVerifikasiService)

// 	var mahasiswaRepository = repository.NewMahasiswaRepository(config.DB)
// 	var mahasiswaService = service.NewMahasiswaService(mahasiswaRepository)
// 	var mahasiswaController = controller.NewMahasiswaController(mahasiswaService, keluargaService, individuService)

// 	var dosenRepository = repository.NewDosenRepository(config.DB)
// 	var dosenService = service.NewDosenService(dosenRepository)
// 	var dosenController = controller.NewDosenController(dosenService, mahasiswaService)

// 	var pusbangRepository = repository.NewPusbangRepository(config.DB)
// 	var pusbangService = service.NewPusbangService(pusbangRepository)
// 	var pusbangController = controller.NewPusbangController(pusbangService, dosenService, keluargaService, individuService)

// 	var userRepository = repository.NewUserRepository(config.DB)
// 	var userService = service.NewUserService(userRepository)
// 	var userController = controller.NewUserController(userService, adminService, pusbangService, dosenService, analisService, mahasiswaService)

// 	var statistikController = controller.NewStatistikController(keluargaService, individuService)

// 	var server = gin.Default()

// 	server.Use(cors.New(cors.Config{
// 		AllowOrigins:     []string{"http://localhost:4003", "http://localhost:3003", "https://www.e-sejahtera.info", "https://e-sejahtera.info"},
// 		AllowCredentials: true,
// 		AllowMethods:     []string{"POST", "PUT", "DELETE", "GET", "PATCH"},
// 		AllowHeaders:     []string{"Access-Control-Allow-Headers", "Access-Control-Allow-Origin", "access-control-allow-origin, access-control-allow-headers", "Content-Type", "Accept", "Origin", "Authorization"},
// 	}))

// 	server.GET("/bidangurusan", middleware.CheckAuth, bidangUrusanController.GetBidangUrusans)
// 	server.GET("/bidangurusan/:id", middleware.CheckAuth, bidangUrusanController.GetBidangUrusan)
// 	server.POST("/bidangurusan", middleware.CheckAuth, bidangUrusanController.CreateBidangUrusan)
// 	server.PATCH("/bidangurusan/:id", middleware.CheckAuth, bidangUrusanController.UpdateBidangUrusan)
// 	server.DELETE("/bidangurusan/:id", middleware.CheckAuth, bidangUrusanController.DeleteBidangUrusan)

// 	server.GET("/instansi", middleware.CheckAuth, instansiController.GetInstansis)
// 	server.GET("/instansi/:id", middleware.CheckAuth, instansiController.GetInstansi)
// 	server.POST("/instansi", middleware.CheckAuth, instansiController.CreateInstansi)
// 	server.PATCH("/instansi/:id", middleware.CheckAuth, instansiController.UpdateInstansi)
// 	server.DELETE("/instansi/:id", middleware.CheckAuth, instansiController.DeleteInstansi)

// 	server.GET("/bidangurusanoninstansi", middleware.CheckAuth, bidangUrusanOnInstansiController.GetBidangUrusanOnInstansis)
// 	server.GET("/bidangurusanoninstansi/:id", middleware.CheckAuth, bidangUrusanOnInstansiController.GetBidangUrusanOnInstansi)
// 	server.POST("/bidangurusanoninstansi", middleware.CheckAuth, bidangUrusanOnInstansiController.CreateBidangUrusanOnInstansi)
// 	server.PATCH("/bidangurusanoninstansi/:id", middleware.CheckAuth, bidangUrusanOnInstansiController.UpdateBidangUrusanOnInstansi)
// 	server.DELETE("/bidangurusanoninstansi/:id", middleware.CheckAuth, bidangUrusanOnInstansiController.DeleteBidangUrusanOnInstansi)

// 	server.GET("/program", middleware.CheckAuth, programController.GetPrograms)
// 	server.GET("/program/:id", middleware.CheckAuth, programController.GetProgram)
// 	server.POST("/program", middleware.CheckAuth, programController.CreateProgram)
// 	server.PATCH("/program/:id", middleware.CheckAuth, programController.UpdateProgram)
// 	server.DELETE("/program/:id", middleware.CheckAuth, programController.DeleteProgram)

// 	server.GET("/rencanaprogram", middleware.CheckAuth, rencanaProgramController.GetRencanaPrograms)
// 	server.GET("/rencanaprogram/:id", middleware.CheckAuth, rencanaProgramController.GetRencanaProgram)
// 	server.GET("/rencanaprogram/search", middleware.CheckAuth, rencanaProgramController.GetRencanaProgramsBySearch)
// 	server.POST("/rencanaprogram", middleware.CheckAuth, rencanaProgramController.CreateRencanaProgram)
// 	server.PATCH("/rencanaprogram/:id", middleware.CheckAuth, rencanaProgramController.UpdateRencanaProgram)
// 	server.DELETE("/rencanaprogram/:id", middleware.CheckAuth, rencanaProgramController.DeleteRencanaProgram)

// 	server.GET("/instansionprogram", middleware.CheckAuth, instansiOnProgramController.GetInstansiOnPrograms)
// 	server.GET("/instansionprogram/:id", middleware.CheckAuth, instansiOnProgramController.GetInstansiOnProgram)
// 	server.POST("/instansionprogram", middleware.CheckAuth, instansiOnProgramController.CreateInstansiOnProgram)
// 	server.PATCH("/instansionprogram/:id", middleware.CheckAuth, instansiOnProgramController.UpdateInstansiOnProgram)
// 	server.DELETE("/instansionprogram/:id", middleware.CheckAuth, instansiOnProgramController.DeleteInstansiOnProgram)

// 	server.GET("/sasaran", middleware.CheckAuth, sasaranController.GetSasarans)
// 	server.GET("/sasaran/:id", middleware.CheckAuth, sasaranController.GetSasaran)
// 	server.POST("/sasaran", middleware.CheckAuth, sasaranController.CreateSasaran)
// 	server.PATCH("/sasaran/:id", middleware.CheckAuth, sasaranController.UpdateSasaran)
// 	server.DELETE("/sasaran/:id", middleware.CheckAuth, sasaranController.DeleteSasaran)

// 	server.GET("/indikatorsasaran", middleware.CheckAuth, indikatorSasaranController.GetIndikatorSasarans)
// 	server.GET("/indikatorsasaran/:id", middleware.CheckAuth, indikatorSasaranController.GetIndikatorSasaran)
// 	server.POST("/indikatorsasaran", middleware.CheckAuth, indikatorSasaranController.CreateIndikatorSasaran)
// 	server.PATCH("/indikatorsasaran/:id", middleware.CheckAuth, indikatorSasaranController.UpdateIndikatorSasaran)
// 	server.DELETE("/indikatorsasaran/:id", middleware.CheckAuth, indikatorSasaranController.DeleteIndikatorSasaran)

// 	server.GET("/kebijakan", middleware.CheckAuth, kebijakanController.GetKebijakans)
// 	server.GET("/kebijakan/:id", middleware.CheckAuth, kebijakanController.GetKebijakan)
// 	server.POST("/kebijakan", middleware.CheckAuth, kebijakanController.CreateKebijakan)
// 	server.PATCH("/kebijakan/:id", middleware.CheckAuth, kebijakanController.UpdateKebijakan)
// 	server.DELETE("/kebijakan/:id", middleware.CheckAuth, kebijakanController.DeleteKebijakan)

// 	server.GET("/kegiatan", middleware.CheckAuth, kegiatanController.GetKegiatans)
// 	server.GET("/kegiatan/:id", middleware.CheckAuth, kegiatanController.GetKegiatan)
// 	server.POST("/kegiatan", middleware.CheckAuth, kegiatanController.CreateKegiatan)
// 	server.PATCH("/kegiatan/:id", middleware.CheckAuth, kegiatanController.UpdateKegiatan)
// 	server.DELETE("/kegiatan/:id", middleware.CheckAuth, kegiatanController.DeleteKegiatan)

// 	server.GET("/rencanakegiatan", middleware.CheckAuth, rencanaKegiatanController.GetRencanaKegiatans)
// 	server.GET("/rencanakegiatan/:id", middleware.CheckAuth, rencanaKegiatanController.GetRencanaKegiatan)
// 	server.GET("/rencanakegiatan/search", middleware.CheckAuth, rencanaKegiatanController.GetRencanaKegiatansBySearch)
// 	server.POST("/rencanakegiatan", middleware.CheckAuth, rencanaKegiatanController.CreateRencanaKegiatan)
// 	server.PATCH("/rencanakegiatan/:id", middleware.CheckAuth, rencanaKegiatanController.UpdateRencanaKegiatan)
// 	server.DELETE("/rencanakegiatan/:id", middleware.CheckAuth, rencanaKegiatanController.DeleteRencanaKegiatan)

// 	server.GET("/programonkegiatan", middleware.CheckAuth, programOnKegiatanController.GetProgramOnKegiatans)
// 	server.GET("/programonkegiatan/:id", middleware.CheckAuth, programOnKegiatanController.GetProgramOnKegiatan)
// 	server.POST("/programonkegiatan", middleware.CheckAuth, programOnKegiatanController.CreateProgramOnKegiatan)
// 	server.PATCH("/programonkegiatan/:id", middleware.CheckAuth, programOnKegiatanController.UpdateProgramOnKegiatan)
// 	server.DELETE("/programonkegiatan/:id", middleware.CheckAuth, programOnKegiatanController.DeleteProgramOnKegiatan)

// 	server.GET("/subkegiatan", middleware.CheckAuth, subKegiatanController.GetSubKegiatans)
// 	server.GET("/subkegiatan/:id", middleware.CheckAuth, subKegiatanController.GetSubKegiatan)
// 	server.POST("/subkegiatan", middleware.CheckAuth, subKegiatanController.CreateSubKegiatan)
// 	server.PATCH("/subkegiatan/:id", middleware.CheckAuth, subKegiatanController.UpdateSubKegiatan)
// 	server.DELETE("/subkegiatan/:id", middleware.CheckAuth, subKegiatanController.DeleteSubKegiatan)

// 	server.GET("/rencanasubkegiatan", middleware.CheckAuth, rencanaSubKegiatanController.GetRencanaSubKegiatans)
// 	server.GET("/rencanasubkegiatan/:id", middleware.CheckAuth, rencanaSubKegiatanController.GetRencanaSubKegiatan)
// 	server.GET("/rencanasubkegiatan/search", middleware.CheckAuth, rencanaSubKegiatanController.GetRencanaSubKegiatansBySearch)
// 	server.POST("/rencanasubkegiatan", middleware.CheckAuth, rencanaSubKegiatanController.CreateRencanaSubKegiatan)
// 	server.PATCH("/rencanasubkegiatan/:id", middleware.CheckAuth, rencanaSubKegiatanController.UpdateRencanaSubKegiatan)
// 	server.DELETE("/rencanasubkegiatan/:id", middleware.CheckAuth, rencanaSubKegiatanController.DeleteRencanaSubkKegiatan)

// 	server.GET("/kegiatanonsubkegiatan", middleware.CheckAuth, kegiatanOnSubKegiatanController.GetKegiatanOnSubKegiatans)
// 	server.GET("/kegiatanonsubkegiatan/:id", middleware.CheckAuth, kegiatanOnSubKegiatanController.GetKegiatanOnSubKegiatan)
// 	server.POST("/kegiatanonsubkegiatan", middleware.CheckAuth, kegiatanOnSubKegiatanController.CreateKegiatanOnSubKegiatan)
// 	server.PATCH("/kegiatanonsubkegiatan/:id", middleware.CheckAuth, kegiatanOnSubKegiatanController.UpdateKegiatanOnSubKegiatan)
// 	server.DELETE("/kegiatanonsubkegiatan/:id", middleware.CheckAuth, kegiatanOnSubKegiatanController.DeleteKegiatanOnSubKegiatan)

// 	server.GET("/fokusbelanja", middleware.CheckAuth, fokusBelanjaController.GetFokusBelanjas)
// 	server.GET("/fokusbelanja/:id", middleware.CheckAuth, fokusBelanjaController.GetFokusBelanja)
// 	server.POST("/fokusbelanja", middleware.CheckAuth, fokusBelanjaController.CreateFokusBelanja)
// 	server.PATCH("/fokusbelanja/:id", middleware.CheckAuth, fokusBelanjaController.UpdateFokusBelanja)
// 	server.DELETE("/fokusbelanja/:id", middleware.CheckAuth, fokusBelanjaController.DeleteFokusBelanja)

// 	server.GET("/detaillokasi", middleware.CheckAuth, detailLokasiController.GetDetailLokasis)
// 	server.GET("/detaillokasi/:id", middleware.CheckAuth, detailLokasiController.GetDetailLokasi)
// 	server.POST("/detaillokasi", middleware.CheckAuth, detailLokasiController.CreateDetailLokasi)
// 	server.PATCH("/detaillokasi/:id", middleware.CheckAuth, detailLokasiController.UpdateDetailLokasi)
// 	server.DELETE("/detaillokasi/:id", middleware.CheckAuth, detailLokasiController.DeleteDetailLokasi)

// 	server.GET("/provinsi", middleware.CheckAuth, provinsiController.GetProvinsis)
// 	server.GET("/provinsi/:id", middleware.CheckAuth, provinsiController.GetProvinsi)

// 	server.GET("/kabupatenkota", middleware.CheckAuth, kabupatenKotaController.GetKabupatenKotas)
// 	server.GET("/kabupatenkota/:id", middleware.CheckAuth, kabupatenKotaController.GetKabupatenKota)

// 	server.GET("/kecamatan", middleware.CheckAuth, kecamatanController.GetKecamatans)
// 	server.GET("/kecamatan/:id", middleware.CheckAuth, kecamatanController.GetKecamatan)

// 	server.GET("/kelurahan", middleware.CheckAuth, kelurahanController.GetKelurahans)
// 	server.GET("/kelurahan/:id", middleware.CheckAuth, kelurahanController.GetKelurahan)

// 	server.GET("/keluarga/", middleware.CheckAuth, keluargaController.GetKeluargas)
// 	server.GET("/keluarga/:id", middleware.CheckAuth, keluargaController.GetKeluargaById)
// 	server.GET("/keluarga/idkeluarga/:idkeluarga", middleware.CheckAuth, keluargaController.GetKeluargaByIdKeluarga)
// 	server.GET("/keluarga/search", middleware.CheckAuth, keluargaController.GetKeluargaBySearch)
// 	server.PATCH("/keluarga/:id", middleware.CheckAuth, keluargaController.UpdateKeluarga)
// 	// server.GET("/testdistinct/", keluargaController.TestDistinct)

// 	server.GET("/monev/:kabupatenkotaid", middleware.CheckAuth, monevController.GetMonevs)
// 	server.GET("/monev/detail/:kabupatenkotaid/:id", middleware.CheckAuth, monevController.GetMonev)

// 	server.GET("/user", middleware.CheckAuth, userController.GetUsers)
// 	server.GET("/user/:id", middleware.CheckAuth, userController.GetUser)
// 	server.GET("/auth/session", middleware.CheckAuth, userController.Validate)
// 	server.GET("/auth/profile", middleware.CheckAuth, userController.GetUserProfile)
// 	server.POST("/auth/login", userController.LoginUser)
// 	server.POST("/auth/logout", userController.LogoutUser)
// 	server.POST("/user", middleware.CheckAuth, userController.CreateUser)
// 	server.POST("/user/batch", middleware.CheckAuth, userController.CreateBatchUser)
// 	server.PATCH("/user/:id", middleware.CheckAuth, userController.UpdateUser)
// 	server.DELETE("/user/:id", middleware.CheckAuth, userController.DeleteUser)

// 	server.GET("/admin", middleware.CheckAuth, adminController.GetAdmins)
// 	server.GET("/admin/:id", middleware.CheckAuth, adminController.GetAdmin)
// 	server.GET("/adminrelasi", middleware.CheckAuth, adminController.GetAdminWithRelation)
// 	server.POST("/admin", middleware.CheckAuth, adminController.CreateAdmin)
// 	server.PATCH("/admin/:id", middleware.CheckAuth, adminController.UpdateAdmin)
// 	server.DELETE("/admin/:id", middleware.CheckAuth, adminController.DeleteAdmin)

// 	server.GET("/adminopd", middleware.CheckAuth, adminOpdController.GetAdminOpd)
// 	server.GET("/adminopd/:id", middleware.CheckAuth, adminOpdController.GetAdminOpd)
// 	server.GET("/adminopd/user/:userid", middleware.CheckAuth, adminOpdController.GetAdminOpdByUserId)
// 	server.POST("/adminopd", middleware.CheckAuth, adminOpdController.CreateAdminOpd)
// 	server.PATCH("/adminopd/:id", middleware.CheckAuth, adminOpdController.UpdateAdminOpd)
// 	server.DELETE("/adminopd/:id", middleware.CheckAuth, adminOpdController.DeleteAdminOpd)

// 	server.GET("/analis", middleware.CheckAuth, analisController.GetAnaliss)
// 	server.GET("/analis/:id", middleware.CheckAuth, analisController.GetAnalis)
// 	server.GET("/analisrelasi", middleware.CheckAuth, analisController.GetAnalisWithRelation)
// 	server.POST("/analis", middleware.CheckAuth, analisController.CreateAnalis)
// 	server.PATCH("/analis/:id", middleware.CheckAuth, analisController.UpdateAnalis)
// 	server.DELETE("/analis/:id", middleware.CheckAuth, analisController.DeleteAnalis)

// 	server.GET("/dosen", middleware.CheckAuth, dosenController.GetDosens)
// 	server.GET("/dosen/:id", middleware.CheckAuth, dosenController.GetDosen)
// 	server.GET("/dosenrelasi", middleware.CheckAuth, dosenController.GetDosenWithRelation)
// 	server.GET("/dosen/findmahasiswa", middleware.CheckAuth, dosenController.GetMahasiswa)
// 	server.POST("/dosen", middleware.CheckAuth, dosenController.CreateDosen)
// 	server.PATCH("/dosen/:id", middleware.CheckAuth, dosenController.UpdateDosen)
// 	server.DELETE("/dosen/:id", middleware.CheckAuth, dosenController.DeleteDosen)

// 	server.GET("/pusbang", middleware.CheckAuth, pusbangController.GetPusbangs)
// 	server.GET("/pusbang/:id", middleware.CheckAuth, pusbangController.GetPusbang)
// 	server.GET("/pusbangrelasi", middleware.CheckAuth, pusbangController.GetPusbangWithRelation)
// 	server.GET("/pusbang/dosen", middleware.CheckAuth, pusbangController.GetDistinctDosen)
// 	server.GET("/pusbang/lokasidosen", middleware.CheckAuth, pusbangController.GetDistinctLokasiDosen)
// 	server.POST("/pusbang", middleware.CheckAuth, pusbangController.CreatePusbang)
// 	server.PATCH("/pusbang:id", middleware.CheckAuth, pusbangController.UpdatePusbang)
// 	server.DELETE("/pusbang/:id", middleware.CheckAuth, pusbangController.DeletePusbang)

// 	server.GET("/mahasiswa", middleware.CheckAuth, mahasiswaController.GetMahasiswas)
// 	server.GET("/mahasiswa/:id", middleware.CheckAuth, mahasiswaController.GetMahasiswa)
// 	server.GET("/mahasiswarelasi", middleware.CheckAuth, mahasiswaController.GetMahasiswaWithRelation)
// 	server.GET("/mahasiswa/verifying", middleware.CheckAuth, mahasiswaController.GetVerifiedByMahasiswa)
//  server.GET("/lokasikkn", middleware.CheckAuth, mahasiswaController.DistinctKelurahan)
// 	server.POST("/mahasiswa", middleware.CheckAuth, mahasiswaController.CreateMahasiswa)
// 	server.POST("/mahasiswa/batch", middleware.CheckAuth, mahasiswaController.CreateBatchMahasiswa)
// 	server.PATCH("/mahasiswa/:id", middleware.CheckAuth, mahasiswaController.UpdateMahasiswa)
// 	server.DELETE("/mahasiswa/:id", middleware.CheckAuth, mahasiswaController.DeleteMahasiswa)

// 	server.GET("/lokasidosen", middleware.CheckAuth, lokasiDosenController.GetLokasiDosens)
// 	server.GET("/lokasidosen/:id", middleware.CheckAuth, lokasiDosenController.GetLokasiDosen)
// 	server.GET("/lokasidosenrelasi", middleware.CheckAuth, lokasiDosenController.GetLokasiDosenWithRelation)
// 	server.GET("/lokasidosen/dosen/:dosenid", middleware.CheckAuth, lokasiDosenController.GetLokasiDosenRelationByDosenId)
// 	server.POST("/lokasidosen", middleware.CheckAuth, lokasiDosenController.CreateLokasiDosen)
// 	server.PATCH("/lokasidosen/:id", middleware.CheckAuth, lokasiDosenController.UpdateLokasiDosen)
// 	server.DELETE("/lokasidosen/:id", middleware.CheckAuth, lokasiDosenController.DeleteLokasiDosen)

// 	server.GET("/keluargaverifikasi/", middleware.CheckAuth, keluargaVerifikasiController.GetKeluargaVerifikasis)
// 	server.GET("/keluargaverifikasi/:id", middleware.CheckAuth, keluargaVerifikasiController.GetKeluargaVerifikasi)
// 	server.GET("/keluargaverifikasi/idkeluarga/:idkeluarga", middleware.CheckAuth, keluargaVerifikasiController.GetKeluargaVerifikasiByIdKeluarga)
// 	server.GET("/keluargaverifikasi/search", middleware.CheckAuth, keluargaVerifikasiController.GetKeluargaVerifikasiBySearch)
// 	server.POST("/keluargaverifikasi", middleware.CheckAuth, keluargaVerifikasiController.CreateKeluargaVerifikasi)
// 	server.PATCH("/keluargaverifikasi/:id", middleware.CheckAuth, keluargaVerifikasiController.UpdateKeluargaVerifikasi)
// 	server.DELETE("/keluargaverifikasi/:id", middleware.CheckAuth, keluargaVerifikasiController.DeleteKeluargaVerifikasi)

// 	server.GET("/individu/", middleware.CheckAuth, individuController.GetIndividus)
// 	server.GET("/individu/:id", middleware.CheckAuth, individuController.GetIndividu)
// 	server.GET("/individu/idkeluarga/:idkeluarga", middleware.CheckAuth, individuController.GetIndividuByIdKeluarga)
// 	server.GET("/individu/search", middleware.CheckAuth, individuController.GetIndividuBySearch)
// 	server.PATCH("/individu/:id", middleware.CheckAuth, individuController.UpdateIndividu)

// 	server.GET("/individuverifikasi", middleware.CheckAuth, individuVerifikasiController.GetIndividuVerifikasis)
// 	server.GET("/individuverifikasi/:id", middleware.CheckAuth, individuVerifikasiController.GetIndividuVerifikasi)
// 	server.GET("/individuverifikasi/idkeluarga/:idkeluarga", middleware.CheckAuth, individuVerifikasiController.GetIndividuVerifikasiByIdKeluarga)
// 	server.GET("/individuverifikasi/search", middleware.CheckAuth, individuVerifikasiController.GetIndividuVerifikasiBySearch)
// 	server.POST("/individuverifikasi", middleware.CheckAuth, individuVerifikasiController.CreateIndividuVerifikasi)
// 	server.PATCH("/individuverifikasi/:id", middleware.CheckAuth, individuVerifikasiController.UpdateIndividuVerifikasi)
// 	server.DELETE("/individuverifikasi/:id", middleware.CheckAuth, individuVerifikasiController.DeleteIndividuVerifikasi)

// 	server.GET("/statistik/provinsi", middleware.CheckAuth, statistikController.StatistikProvinsi)
// 	server.GET("/statistik/kabupatenkota", middleware.CheckAuth, statistikController.StatistikKabupatenKota)
// 	server.GET("/statistik/kecamatan", middleware.CheckAuth, statistikController.StatistikKecamatan)
// 	server.GET("/statistik/kelurahan", middleware.CheckAuth, statistikController.StatistikKelurahan)
// 	server.GET("/statistik/kelurahan/hitung", middleware.CheckAuth, statistikController.StatistikSearch)

// 	server.Run(":" + appConfig.AppPort)
// }
