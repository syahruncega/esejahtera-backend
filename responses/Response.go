package responses

import "time"

type BidangUrusanResponse struct {
	Id               int       `json:"id"`
	KodeBidangUrusan string    `json:"kodeBidangUrusan"`
	NamaBidangUrusan string    `json:"namaBidangUrusan"`
	CreatedAt        time.Time `json:"createdAt"`
	UpdatedAt        time.Time `json:"updatedAt"`
}

type InstansiResponse struct {
	Id           int       `json:"id"`
	KodeInstansi string    `json:"kodeInstansi"`
	NamaInstansi string    `json:"namaInstansi"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

type BidangUrusanOnInstansiResponse struct {
	Id             int                  `json:"id"`
	InstansiId     int                  `json:"instansiId"`
	Instansi       InstansiResponse     `json:"instansi"`
	BidangUrusanId int                  `json:"bidangUrusanId"`
	BidangUrusan   BidangUrusanResponse `json:"bidangUrusan"`
	CreatedAt      time.Time            `json:"createdAt"`
	UpdatedAt      time.Time            `json:"updatedAt"`
}

type ProgramResponse struct {
	Id          int       `json:"id"`
	Tahun       string    `json:"tahun"`
	KodeProgram string    `json:"kodeProgram"`
	NamaProgram string    `json:"namaProgram"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type RencanaProgramResponse struct {
	Id          int              `json:"id"`
	InstansiId  int              `json:"instansiId"`
	Instansi    InstansiResponse `json:"instansi"`
	ProgramId   int              `json:"programId"`
	Program     ProgramResponse  `json:"program"`
	PaguProgram int64            `json:"paguProgram"`
	Tipe        string           `json:"tipe"`
	Tahun       string           `json:"tahun"`
	CreatedAt   time.Time        `json:"createdAt"`
	UpdatedAt   time.Time        `json:"updatedAt"`
}

type InstansiOnProgramResponse struct {
	Id         int              `json:"id"`
	InstansiId int              `json:"instansiId"`
	Instansi   InstansiResponse `json:"instansi"`
	ProgramId  int              `json:"programId"`
	Program    ProgramResponse  `json:"program"`
	CreatedAt  time.Time        `json:"createdAt"`
	UpdatedAt  time.Time        `json:"updatedAt"`
}

type IndikatorSasaranResponse struct {
	Id                   int                    `json:"id"`
	RencanaProgramId     int                    `json:"rencanaProgramId"`
	RencanaProgram       RencanaProgramResponse `json:"rencanaProgram"`
	NamaIndikatorSasaran string                 `json:"namaIndikatorSasaran"`
	CreatedAt            time.Time              `json:"createdAt"`
	UpdatedAt            time.Time              `json:"updatedAt"`
}

type KebijakanResponse struct {
	Id               int                    `json:"id"`
	RencanaProgramId int                    `json:"rencanaProgramId"`
	RencanaProgram   RencanaProgramResponse `json:"rencanaProgram"`
	NamaKebijakan    string                 `json:"namaKebijakan"`
	CreatedAt        time.Time              `json:"createdAt"`
	UpdatedAt        time.Time              `json:"updatedAt"`
}

type KegiatanResponse struct {
	Id           int       `json:"id"`
	Tahun        string    `json:"tahun"`
	KodeKegiatan string    `json:"kodeKegiatan"`
	NamaKegiatan string    `json:"namaKegiatan"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

type RencanaKegiatanResponse struct {
	Id               int                    `json:"id"`
	RencanaProgramId int                    `json:"rencanaProgramId"`
	RencanaProgram   RencanaProgramResponse `json:"rencanaProgram"`
	KegiatanId       int                    `json:"kegiatanId"`
	Kegiatan         KegiatanResponse       `json:"kegiatan"`
	Tipe             string                 `json:"tipe"`
	PaguKegiatan     int64                  `json:"paguKegiatan"`
	Tahun            string                 `json:"tahun"`
	CreatedAt        time.Time              `json:"createdAt"`
	UpdatedAt        time.Time              `json:"updatedAt"`
}

type ProgramOnKegiatanResponse struct {
	Id         int              `json:"id"`
	ProgramId  int              `json:"programId"`
	Program    ProgramResponse  `json:"program"`
	KegiatanId int              `json:"kegiatanId"`
	Kegiatan   KegiatanResponse `json:"kegiatan"`
	CreatedAt  time.Time        `json:"createdAt"`
	UpdatedAt  time.Time        `json:"updatedAt"`
}

type SasaranResponse struct {
	Id               int                    `json:"id"`
	RencanaProgramId int                    `json:"rencanaProgramId"`
	RencanaProgram   RencanaProgramResponse `json:"rencanaProgram"`
	NamaSasaran      string                 `json:"namaSasaran"`
	CreatedAt        time.Time              `json:"createdAt"`
	UpdatedAt        time.Time              `json:"updatedAt"`
}

type SubKegiatanResponse struct {
	Id              int       `json:"id"`
	Tahun           string    `json:"tahun"`
	KodeSubKegiatan string    `json:"kodeSubKegiatan"`
	NamaSubKegiatan string    `json:"namaSubKegiatan"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
}

type RencanaSubKegiatanResponse struct {
	Id                int                     `json:"id"`
	RencanaKegiatanId int                     `json:"rencanaKegiatanId"`
	RencanaKegiatan   RencanaKegiatanResponse `json:"rencanaKegiatan"`
	SubKegiatanId     int                     `json:"subKegiatanId"`
	SubKegiatan       SubKegiatanResponse     `json:"subKegiatan"`
	Tipe              string                  `json:"tipe"`
	PaguSubKegiatan   int64                   `json:"paguSubKegiatan"`
	Tahun             string                  `json:"tahun"`
	CreatedAt         time.Time               `json:"createdAt"`
	UpdatedAt         time.Time               `json:"updatedAt"`
}

type KegiatanOnSubKegiatanResponse struct {
	Id            int                 `json:"id"`
	KegiatanId    int                 `json:"kegiatanId"`
	Kegiatan      KegiatanResponse    `json:"kegiatan"`
	SubKegiatanId int                 `json:"subKegiatanId"`
	SubKegiatan   SubKegiatanResponse `json:"subKegiatan"`
	CreatedAt     time.Time           `json:"createdAt"`
	UpdatedAt     time.Time           `json:"updatedAt"`
}

type FokusBelanjaResponse struct {
	Id                   int                        `json:"id"`
	RencanaSubKegiatanId int                        `json:"rencanaSubKegiatanId"`
	RencanaSubKegiatan   RencanaSubKegiatanResponse `json:"rencanaSubKegiatan"`
	NamaFokusBelanja     string                     `json:"namaFokusBelanja"`
	Indikator            string                     `json:"indikator"`
	Target               float32                    `json:"target"`
	Satuan               string                     `json:"satuan"`
	PaguFokusBelanja     int64                      `json:"paguFokusBelanja"`
	Keterangan           string                     `json:"keterangan"`
	Tahun                string                     `json:"tahun"`
	CreatedAt            time.Time                  `json:"createdAt"`
	UpdatedAt            time.Time                  `json:"updatedAt"`
}

type Detail_LokasiResponse struct {
	Id              int                    `json:"id"`
	FokusBelanjaid  int                    `json:"fokusBelanjaId"`
	FokusBelanja    FokusBelanjaResponse   `json:"fokusBelanja"`
	KabupatenKotaId string                 `json:"kabupatenKotaId"`
	KabupatenKota   Kabupaten_KotaResponse `json:"kabupatenKota"`
	KecamatanId     string                 `json:"kecamatanId"`
	Kecamatan       KecamatanResponse      `json:"kecamatan"`
	KelurahanId     string                 `json:"kelurahanId"`
	Kelurahan       KelurahanResponse      `json:"kelurahan"`
	CreatedAt       time.Time              `json:"createdAt"`
	UpdatedAt       time.Time              `json:"updatedAt"`
}

type ProvinsiResponse struct {
	Id   string `json:"id"`
	Nama string `json:"nama"`
}

type Kabupaten_KotaResponse struct {
	Id         string `json:"id"`
	ProvinsiId string `json:"provinsiId"`
	Nama       string `json:"nama"`
}

type KecamatanResponse struct {
	Id              string `json:"id"`
	KabupatenKotaId string `json:"kabupatenKotaId"`
	Nama            string `json:"nama"`
}

type KelurahanResponse struct {
	Id          string `json:"id"`
	KecamatanId string `json:"kecamatanId"`
	Nama        string `json:"nama"`
}

type DistinctKelurahan struct {
	KelurahanId string `json:"kelurahanId"`
	Nama        string `json:"nama"`
}

type KeluargaResponse struct {
	Id                     int                    `json:"id"`
	IdKeluarga             string                 `json:"idKeluarga"`
	ProvinsiId             string                 `json:"provinsiId"`
	Provinsi               ProvinsiResponse       `json:"provinsi"`
	KabupatenKotaId        string                 `json:"kabupatenKotaId"`
	KabupatenKota          Kabupaten_KotaResponse `json:"kabupatenKota"`
	KecamatanId            string                 `json:"kecamatanId"`
	Kecamatan              KecamatanResponse      `json:"kecamatan"`
	KelurahanId            string                 `json:"kelurahanId"`
	Kelurahan              KelurahanResponse      `json:"kelurahan"`
	DesilKesejahteraan     string                 `json:"desilKesejahteraan"`
	Alamat                 string                 `json:"alamat"`
	KepalaKeluarga         string                 `json:"kepalaKeluarga"`
	Nik                    string                 `json:"nik"`
	PadanDukcapil          string                 `json:"padanDukcapil"`
	JenisKelamin           string                 `json:"jenisKelamin"`
	TanggalLahir           string                 `json:"tanggalLahir"`
	Pekerjaan              string                 `json:"pekerjaan"`
	Pendidikan             string                 `json:"pendidikan"`
	KepemilikanRumah       string                 `json:"kepemilikanRumah"`
	Simpanan               string                 `json:"simpanan"`
	JenisAtap              string                 `json:"jenisAtap"`
	JenisDinding           string                 `json:"jenisDinding"`
	JenisLantai            string                 `json:"jenisLantai"`
	SumberPenerangan       string                 `json:"sumberPenerangan"`
	BahanBakarMemasak      string                 `json:"bahanBakarMemasak"`
	SumberAirMinum         string                 `json:"sumberAirMinum"`
	FasilitasBuangAirBesar string                 `json:"fasilitasBuangAirBesar"`
	PenerimaBPNT           string                 `json:"penerimaBPNT"`
	PenerimaBPUM           string                 `json:"penerimaBPUM"`
	PenerimaBST            string                 `json:"penerimaBST"`
	PenerimaPKH            string                 `json:"penerimaPKH"`
	PenerimaSembako        string                 `json:"penerimaSembako"`
	UserId                 int                    `json:"userId"`
	User                   UserResponse           `json:"user"`
	MahasiswaId            int                    `json:"mahasiswaId"`
	Mahasiswa              MahasiswaResponse      `json:"mahasiswa"`
	StatusVerifikasi       int                    `json:"statusVerifikasi"`
	JumlahIndividu         int                    `json:"jumlahIndividu"`
	JumlahIndividuVerified int                    `json:"jumlahIndividuVerified"`
}

type MonevResponse struct {
	Id              int                    `json:"id"`
	NamaPenerima    string                 `json:"namaPenerima"`
	KabupatenKotaId string                 `json:"kabupatenKotaId"`
	KabupatenKota   Kabupaten_KotaResponse `json:"kabupatenKota"`
	KecamatanId     string                 `json:"kecamatanId"`
	Kecamatan       KecamatanResponse      `json:"kecamatan"`
	KelurahanId     string                 `json:"kelurahanId"`
	Kelurahan       KelurahanResponse      `json:"kelurahan"`
	JenisBantuan    string                 `json:"jenisBantuan"`
	VolumeBantuan   float32                `json:"volumeBantuan"`
	SatuanVolume    string                 `json:"satuanVolume"`
}

type JumlahResponse struct {
	Jumlah int64 `json:"jumlah"`
}

type UserResponse struct {
	Id        int       `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	NoHp      string    `json:"noHp"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type AdminResponse struct {
	Id            int       `json:"id"`
	UserId        int       `json:"userId"`
	NamaLengkap   string    `json:"namaLengkap"`
	UrlFotoProfil string    `json:"urlFotoProfil"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}

type AdminOpdResponse struct {
	Id            int              `json:"id"`
	UserId        int              `json:"userId"`
	User          UserResponse     `json:"user"`
	NamaLengkap   string           `json:"namaLengkap"`
	InstansiId    int              `json:"instansiId"`
	Instansi      InstansiResponse `json:"instansi"`
	UrlFotoProfil string           `json:"urlFotoProfil"`
	CreatedAt     time.Time        `json:"createdAt"`
	UpdatedAt     time.Time        `json:"updatedAt"`
}

type AdminWithRelationResponse struct {
	Id            int          `json:"id"`
	UserId        int          `json:"userId"`
	User          UserResponse `json:"user"`
	NamaLengkap   string       `json:"namaLengkap"`
	UrlFotoProfil string       `json:"urlFotoProfil"`
	CreatedAt     time.Time    `json:"createdAt"`
	UpdatedAt     time.Time    `json:"updatedAt"`
}

type AnalisResponse struct {
	Id            int       `json:"id"`
	UserId        int       `json:"userId"`
	NamaLengkap   string    `json:"namaLengkap"`
	Universitas   string    `json:"universitas"`
	UrlFotoProfil string    `json:"urlFotoProfil"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}

type AnalisWithRelationResponse struct {
	Id            int          `json:"id"`
	UserId        int          `json:"userId"`
	User          UserResponse `json:"user"`
	NamaLengkap   string       `json:"namaLengkap"`
	Universitas   string       `json:"universitas"`
	UrlFotoProfil string       `json:"urlFotoProfil"`
	CreatedAt     time.Time    `json:"createdAt"`
	UpdatedAt     time.Time    `json:"updatedAt"`
}

type PusbangResponse struct {
	Id            int       `json:"id"`
	UserId        int       `json:"userId"`
	NamaLengkap   string    `json:"namaLengkap"`
	Universitas   string    `json:"universitas"`
	UrlFotoProfil string    `json:"urlFotoProfil"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}

type PusbangWithRelationResponse struct {
	Id            int          `json:"id"`
	UserId        int          `json:"userId"`
	User          UserResponse `json:"user"`
	NamaLengkap   string       `json:"namaLengkap"`
	Universitas   string       `json:"universitas"`
	UrlFotoProfil string       `json:"urlFotoProfil"`
	CreatedAt     time.Time    `json:"createdAt"`
	UpdatedAt     time.Time    `json:"updatedAt"`
}

type DosenResponse struct {
	Id            int       `json:"id"`
	UserId        int       `json:"userId"`
	NamaLengkap   string    `json:"namaLengkap"`
	Universitas   string    `json:"universitas"`
	UrlFotoProfil string    `json:"urlFotoProfil"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}

type DosenWithRelationResponse struct {
	Id            int          `json:"id"`
	UserId        int          `json:"userId"`
	User          UserResponse `json:"user,omitempty"`
	NamaLengkap   string       `json:"namaLengkap"`
	Universitas   string       `json:"universitas"`
	UrlFotoProfil string       `json:"urlFotoProfil"`
	CreatedAt     time.Time    `json:"createdAt"`
	UpdatedAt     time.Time    `json:"updatedAt"`
}

type MahasiswaResponse struct {
	Id              int       `json:"id"`
	UserId          int       `json:"userId"`
	NamaLengkap     string    `json:"namaLengkap"`
	Universitas     string    `json:"universitas"`
	JenisKelamin    string    `json:"jenisKelamin"`
	TanggalLahir    string    `json:"tanggalLahir"`
	KabupatenKotaId string    `json:"kabupatenKotaId"`
	KecamatanId     string    `json:"kecamatanId"`
	KelurahanId     string    `json:"kelurahanId"`
	UrlFotoProfil   string    `json:"urlFotoProfil"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
}

type MahasiswaWithRelationResponse struct {
	Id              int                    `json:"id"`
	UserId          int                    `json:"userId"`
	User            UserResponse           `json:"user"`
	NamaLengkap     string                 `json:"namaLengkap"`
	Universitas     string                 `json:"universitas"`
	JenisKelamin    string                 `json:"jenisKelamin"`
	TanggalLahir    string                 `json:"tanggalLahir"`
	KabupatenKotaId string                 `json:"kabupatenKotaId"`
	KabupatenKota   Kabupaten_KotaResponse `json:"kabupatenKota"`
	KecamatanId     string                 `json:"kecamatanId"`
	Kecamatan       KecamatanResponse      `json:"kecamatan"`
	KelurahanId     string                 `json:"kelurahanId"`
	Kelurahan       KelurahanResponse      `json:"kelurahan"`
	UrlFotoProfil   string                 `json:"urlFotoProfil"`
	CreatedAt       time.Time              `json:"createdAt"`
	UpdatedAt       time.Time              `json:"updatedAt"`
}

type MahasiswaWithVerifiedCountResponse struct {
	Id                     int                    `json:"id"`
	UserId                 int                    `json:"userId"`
	User                   UserResponse           `json:"user"`
	NamaLengkap            string                 `json:"namaLengkap"`
	Universitas            string                 `json:"universitas"`
	JenisKelamin           string                 `json:"jenisKelamin"`
	TanggalLahir           string                 `json:"tanggalLahir"`
	KabupatenKotaId        string                 `json:"kabupatenKotaId"`
	KabupatenKota          Kabupaten_KotaResponse `json:"kabupatenKota"`
	KecamatanId            string                 `json:"kecamatanId"`
	Kecamatan              KecamatanResponse      `json:"kecamatan"`
	KelurahanId            string                 `json:"kelurahanId"`
	Kelurahan              KelurahanResponse      `json:"kelurahan"`
	JumlahIndividuVerified int64                  `json:"jumlahIndividuVerified"`
	JumlahKeluargaVerified int64                  `json:"jumlahKeluargaVerified"`
	UrlFotoProfil          string                 `json:"urlFotoProfil"`
	CreatedAt              time.Time              `json:"createdAt"`
	UpdatedAt              time.Time              `json:"updatedAt"`
}

type LokasiDosenResponse struct {
	Id              int       `json:"id"`
	DosenId         int       `json:"dosenId"`
	KabupatenKotaId string    `json:"kabupatenKotaId"`
	KecamatanId     string    `json:"kecamatanId"`
	KelurahanId     string    `json:"kelurahanId"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
}

type LokasiDosenWithRelationResponse struct {
	Id              int                    `json:"id"`
	DosenId         int                    `json:"dosenId"`
	Dosen           DosenResponse          `json:"dosen"`
	KabupatenKotaId string                 `json:"kabupatenKotaId"`
	KabupatenKota   Kabupaten_KotaResponse `json:"kabupatenKota"`
	KecamatanId     string                 `json:"kecamatanId"`
	Kecamatan       KecamatanResponse      `json:"kecamatan"`
	KelurahanId     string                 `json:"kelurahanId"`
	Kelurahan       KelurahanResponse      `json:"kelurahan"`
	CreatedAt       time.Time              `json:"createdAt"`
	UpdatedAt       time.Time              `json:"updatedAt"`
}

type KeluargaVerifikasiResponse struct {
	Id                     int                    `json:"id"`
	IdKeluarga             string                 `json:"idKeluarga"`
	ProvinsiId             string                 `json:"provinsiId"`
	Provinsi               ProvinsiResponse       `json:"provinsi"`
	KabupatenKotaId        string                 `json:"kabupatenKotaId"`
	KabupatenKota          Kabupaten_KotaResponse `json:"kabupatenKota"`
	KecamatanId            string                 `json:"kecamatanId"`
	Kecamatan              KecamatanResponse      `json:"kecamatan"`
	KelurahanId            string                 `json:"kelurahanId"`
	DesilKesejahteraan     string                 `json:"desilKesejahteraan"`
	Alamat                 string                 `json:"alamat"`
	KepalaKeluarga         string                 `json:"kepalaKeluarga"`
	Nik                    string                 `json:"nik"`
	PadanDukcapil          string                 `json:"padanDukcapil"`
	JenisKelamin           string                 `json:"jenisKelamin"`
	TanggalLahir           string                 `json:"tanggalLahir"`
	Pekerjaan              string                 `json:"pekerjaan"`
	Pendidikan             string                 `json:"pendidikan"`
	KepemilikanRumah       string                 `json:"kepemilikanRumah"`
	Simpanan               string                 `json:"simpanan"`
	JenisAtap              string                 `json:"jenisAtap"`
	JenisDinding           string                 `json:"jenisDinding"`
	JenisLantai            string                 `json:"jenisLantai"`
	SumberPenerangan       string                 `json:"sumberPenerangan"`
	BahanBakarMemasak      string                 `json:"bahanBakarMemasak"`
	SumberAirMinum         string                 `json:"sumberAirMinum"`
	FasilitasBuangAirBesar string                 `json:"fasilitasBuangAirBesar"`
	PenerimaBPNT           string                 `json:"penerimaBPNT"`
	PenerimaBPUM           string                 `json:"penerimaBPUM"`
	PenerimaBST            string                 `json:"penerimaBST"`
	PenerimaPKH            string                 `json:"penerimaPKH"`
	PenerimaSembako        string                 `json:"penerimaSembako"`
	PenerimaLainnya        string                 `json:"penerimaLainnya"`
	StatusResponden        string                 `json:"statusResponden"`
	UserId                 int                    `json:"userId"`
	User                   UserResponse           `json:"user"`
	MahasiswaId            int                    `json:"mahasiswaId"`
	Mahasiswa              MahasiswaResponse      `json:"mahasiswa"`
	Kelurahan              KelurahanResponse      `json:"kelurahan"`
	UrlBukti               string                 `json:"urlBukti"`
	CreatedAt              time.Time              `json:"createdAt"`
	UpdatedAt              time.Time              `json:"updatedAt"`
}

type IndividuResponse struct {
	Id                 int                    `json:"id"`
	IdKeluarga         string                 `json:"idKeluarga"`
	ProvinsiId         string                 `json:"provinsiId"`
	Provinsi           ProvinsiResponse       `json:"provinsi"`
	KabupatenKotaId    string                 `json:"kabupatenKotaId"`
	Kabupaten_Kota     Kabupaten_KotaResponse `json:"kabupatenKota"`
	KecamatanId        string                 `json:"kecamatanId"`
	Kecamatan          KecamatanResponse      `json:"kecamatan"`
	KelurahanId        string                 `json:"kelurahanId"`
	Kelurahan          KelurahanResponse      `json:"kelurahan"`
	DesilKesejahteraan string                 `json:"desilKesejahteraan"`
	Alamat             string                 `json:"alamat"`
	IdIndividu         string                 `json:"idIndividu"`
	Nama               string                 `json:"nama"`
	Nik                string                 `json:"nik"`
	PadanDukcapil      string                 `json:"padanDukcapil"`
	JenisKelamin       string                 `json:"jenisKelamin"`
	Hubungan           string                 `json:"hubungan"`
	TanggalLahir       string                 `json:"tanggalLahir"`
	StatusKawin        string                 `json:"statusKawin"`
	Pekerjaan          string                 `json:"pekerjaan"`
	Pendidikan         string                 `json:"pendidikan"`
	KategoriUsia       string                 `json:"kategoriUsia"`
	PenerimaBPNT       string                 `json:"penerimaBPNT"`
	PenerimaBPUM       string                 `json:"penerimaBPUM"`
	PenerimaBST        string                 `json:"penerimaBST"`
	PenerimaPKH        string                 `json:"penerimaPKH"`
	PenerimaSembako    string                 `json:"penerimaSembako"`
	UserId             int                    `json:"userId"`
	User               UserResponse           `json:"user"`
	MahasiswaId        int                    `json:"mahasiswaId"`
	Mahasiswa          MahasiswaResponse      `json:"mahasiswa"`
	StatusVerifikasi   int                    `json:"statusVerifikasi"`
}

type IndividuVerifikasiResponse struct {
	Id                 int                    `json:"id"`
	IdKeluarga         string                 `json:"idKeluarga"`
	ProvinsiId         string                 `json:"provinsiId"`
	Provinsi           ProvinsiResponse       `json:"provinsi"`
	KabupatenKotaId    string                 `json:"kabupatenKotaId"`
	KabupatenKota      Kabupaten_KotaResponse `json:"kabupatenKota"`
	KecamatanId        string                 `json:"kecamatanId"`
	Kecamatan          KecamatanResponse      `json:"kecamatan"`
	KelurahanId        string                 `json:"kelurahanId"`
	Kelurahan          KelurahanResponse      `json:"kelurahan"`
	DesilKesejahteraan string                 `json:"desilKesejahteraan"`
	Alamat             string                 `json:"alamat"`
	IdIndividu         string                 `json:"idIndividu"`
	Nama               string                 `json:"nama"`
	Nik                string                 `json:"nik"`
	PadanDukcapil      string                 `json:"padanDukcapil"`
	JenisKelamin       string                 `json:"jenisKelamin"`
	Hubungan           string                 `json:"hubungan"`
	TanggalLahir       string                 `json:"tanggalLahir"`
	StatusKawin        string                 `json:"statusKawin"`
	Pekerjaan          string                 `json:"pekerjaan"`
	Pendidikan         string                 `json:"pendidikan"`
	KategoriUsia       string                 `json:"kategoriUsia"`
	PenerimaBPNT       string                 `json:"penerimaBPNT"`
	PenerimaBPUM       string                 `json:"penerimaBPUM"`
	PenerimaBST        string                 `json:"penerimaBST"`
	PenerimaPKH        string                 `json:"penerimaPKH"`
	PenerimaSembako    string                 `json:"penerimaSembako"`
	PenerimaLainnya    string                 `json:"penerimaLainnya"`
	StatusResponden    string                 `json:"statusResponden"`
	UserId             int                    `json:"userId"`
	User               UserResponse           `json:"user"`
	MahasiswaId        int                    `json:"mahasiswaId"`
	Mahasiswa          MahasiswaResponse      `json:"mahasiswa"`
	UrlBukti           string                 `json:"urlBukti"`
	CreatedAt          time.Time              `json:"createdAt"`
	UpdatedAt          time.Time              `json:"updatedAt"`
}
