package responses

import "time"

type BidangUrusanResponse struct {
	Id               string    `json:"id"`
	NamaBidangUrusan string    `json:"namaBidangUrusan"`
	CreatedAt        time.Time `json:"createdAt"`
	UpdatedAt        time.Time `json:"updatedAt"`
}

type InstansiResponse struct {
	Id           string    `json:"id"`
	NamaInstansi string    `json:"namaInstansi"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

type InstansiWithBidangUrusanResponse struct {
	Id             string               `json:"id"`
	NamaInstansi   string               `json:"namaInstansi"`
	BidangUrusanId string               `json:"bidangUrusanId"`
	CreatedAt      time.Time            `json:"createdAt"`
	UpdatedAt      time.Time            `json:"updatedAt"`
	BidangUrusan   BidangUrusanResponse `json:"bidangUrusan"`
}

type ProgramResponse struct {
	Id                      int       `json:"id"`
	Sasaran                 string    `json:"sasaran"`
	IndikatorSasaran        string    `json:"indikatorSasaran"`
	Kebijakan               string    `json:"kebijakan"`
	NamaProgram             string    `json:"namaProgram"`
	IndikatorKinerjaProgram string    `json:"indikatorKinerjaProgram"`
	PaguProgram             int       `json:"paguProgram"`
	InstansiId              string    `json:"instansiId"`
	CreatedAt               time.Time `json:"createdAt"`
	UpdatedAt               time.Time `json:"updatedAt"`
}

type ProgramWithInstansiResponse struct {
	Id                      int              `json:"id"`
	Sasaran                 string           `json:"sasaran"`
	IndikatorSasaran        string           `json:"indikatorSasaran"`
	Kebijakan               string           `json:"kebijakan"`
	NamaProgram             string           `json:"namaProgram"`
	IndikatorKinerjaProgram string           `json:"indikatorKinerjaProgram"`
	PaguProgram             int              `json:"paguProgram"`
	InstansiId              string           `json:"instansiId"`
	CreatedAt               time.Time        `json:"createdAt"`
	UpdatedAt               time.Time        `json:"updatedAt"`
	Instansi                InstansiResponse `json:"instansi"`
}

type KegiatanResponse struct {
	Id                       int       `json:"id"`
	NamaKegiatan             string    `json:"namaKegiatan"`
	IndikatorKinerjaKegiatan string    `json:"indikatorKinerjaKegiatan"`
	PaguKegiatan             int       `json:"paguKegiatan"`
	ProgramId                int       `json:"programId"`
	CreatedAt                time.Time `json:"createdAt"`
	UpdatedAt                time.Time `json:"updatedAt"`
}

type KegiatanWithProgramResponse struct {
	Id                       int             `json:"id"`
	NamaKegiatan             string          `json:"namaKegiatan"`
	IndikatorKinerjaKegiatan string          `json:"indikatorKinerjaKegiatan"`
	PaguKegiatan             int             `json:"paguKegiatan"`
	ProgramId                int             `json:"programId"`
	CreatedAt                time.Time       `json:"createdAt"`
	UpdatedAt                time.Time       `json:"updatedAt"`
	Program                  ProgramResponse `json:"program"`
}

type Sub_KegiatanResponse struct {
	Id                          int       `json:"id"`
	NamaSubKegiatan             string    `json:"namaSubKegiatan"`
	IndikatorKinerjaSubKegiatan string    `json:"indikatorKinerjaSubKegiatan"`
	PaguSubKegiatan             int       `json:"paguSubKegiatan"`
	KegiatanId                  int       `json:"kegiatanId"`
	CreatedAt                   time.Time `json:"createdAt"`
	UpdatedAt                   time.Time `json:"updatedAt"`
}

type Sub_KegiatanWithKegiatanResponse struct {
	Id                          int              `json:"id"`
	NamaSubKegiatan             string           `json:"namaSubKegiatan"`
	IndikatorKinerjaSubKegiatan string           `json:"indikatorKinerjaSubKegiatan"`
	PaguSubKegiatan             int              `json:"paguSubKegiatan"`
	KegiatanId                  int              `json:"kegiatanId"`
	CreatedAt                   time.Time        `json:"createdAt"`
	UpdatedAt                   time.Time        `json:"updatedAt"`
	Kegiatan                    KegiatanResponse `json:"kegiatan"`
}

type Detail_Sub_KegiatanResponse struct {
	Id               int       `json:"id"`
	FokusBelanja     string    `json:"fokusBelanja"`
	Indikator        string    `json:"indikator"`
	Target           float32   `json:"target"`
	Satuan           string    `json:"satuan"`
	PaguFokusBelanja int       `json:"paguFokusBelanja"`
	Keterangan       string    `json:"keterangan"`
	SubKegiatanId    int       `json:"subKegiatanId"`
	CreatedAt        time.Time `json:"createdAt"`
	UpdatedAt        time.Time `json:"updatedAt"`
}

type Detail_Sub_KegiatanWithSub_KegiatanResponse struct {
	Id               int                  `json:"id"`
	FokusBelanja     string               `json:"fokusBelanja"`
	Indikator        string               `json:"indikator"`
	Target           float32              `json:"target"`
	Satuan           string               `json:"satuan"`
	PaguFokusBelanja int                  `json:"paguFokusBelanja"`
	Keterangan       string               `json:"keterangan"`
	SubKegiatanId    int                  `json:"subKegiatanId"`
	CreatedAt        time.Time            `json:"createdAt"`
	UpdatedAt        time.Time            `json:"updatedAt"`
	SubKegiatan      Sub_KegiatanResponse `json:"subKegiatan"`
}

type Detail_LokasiResponse struct {
	Id                  int       `json:"id"`
	KabupatenKotaId     string    `json:"kabupatenKotaId"`
	KecamatanId         string    `json:"kecamatanId"`
	KelurahanId         string    `json:"kelurahanId"`
	DetailSubKegiatanId int       `json:"detailSubKegiatanId"`
	CreatedAt           time.Time `json:"createdAt"`
	UpdatedAt           time.Time `json:"updatedAt"`
}

type Detail_LokasiWithRelationResponse struct {
	Id                  int                         `json:"id"`
	KabupatenKotaId     string                      `json:"kabupatenKotaId"`
	KecamatanId         string                      `json:"kecamatanId"`
	KelurahanId         string                      `json:"kelurahanId"`
	DetailSubKegiatanId int                         `json:"detailSubKegiatanId"`
	CreatedAt           time.Time                   `json:"createdAt"`
	UpdatedAt           time.Time                   `json:"updatedAt"`
	DetailSubKegiatan   Detail_Sub_KegiatanResponse `json:"detailSubKegiatan"`
	KabupatenKota       Kabupaten_KotaResponse      `json:"kabupatenKota"`
	Kecamatan           KecamatanResponse           `json:"kecamatan"`
	Kelurahan           KelurahanResponse           `json:"kelurahan"`
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
	Id          int       `json:"id"`
	UserId      int       `json:"userId"`
	NamaLengkap string    `json:"namaLengkap"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type AdminWithRelationResponse struct {
	Id          int          `json:"id"`
	UserId      int          `json:"userId"`
	User        UserResponse `json:"user"`
	NamaLengkap string       `json:"namaLengkap"`
	CreatedAt   time.Time    `json:"createdAt"`
	UpdatedAt   time.Time    `json:"updatedAt"`
}

type AnalisResponse struct {
	Id          int       `json:"id"`
	UserId      int       `json:"userId"`
	NamaLengkap string    `json:"namaLengkap"`
	Universitas string    `json:"universitas"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type AnalisWithRelationResponse struct {
	Id          int          `json:"id"`
	UserId      int          `json:"userId"`
	User        UserResponse `json:"user"`
	NamaLengkap string       `json:"namaLengkap"`
	Universitas string       `json:"universitas"`
	CreatedAt   time.Time    `json:"createdAt"`
	UpdatedAt   time.Time    `json:"updatedAt"`
}

type PusbangResponse struct {
	Id          int       `json:"id"`
	UserId      int       `json:"userId"`
	NamaLengkap string    `json:"namaLengkap"`
	Universitas string    `json:"universitas"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type PusbangWithRelationResponse struct {
	Id          int          `json:"id"`
	UserId      int          `json:"userId"`
	User        UserResponse `json:"user"`
	NamaLengkap string       `json:"namaLengkap"`
	Universitas string       `json:"universitas"`
	CreatedAt   time.Time    `json:"createdAt"`
	UpdatedAt   time.Time    `json:"updatedAt"`
}

type DosenResponse struct {
	Id          int       `json:"id"`
	UserId      int       `json:"userId"`
	NamaLengkap string    `json:"namaLengkap"`
	Universitas string    `json:"universitas"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type DosenWithRelationResponse struct {
	Id          int          `json:"id"`
	UserId      int          `json:"userId"`
	User        UserResponse `json:"user"`
	NamaLengkap string       `json:"namaLengkap"`
	Universitas string       `json:"universitas"`
	CreatedAt   time.Time    `json:"createdAt"`
	UpdatedAt   time.Time    `json:"updatedAt"`
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
	KelurahanId     string    `json:"KelurahanId"`
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
	CreatedAt       time.Time              `json:"createdAt"`
	UpdatedAt       time.Time              `json:"updatedAt"`
}

type LokasiDosenResponse struct {
	Id          int       `json:"id"`
	DosenId     int       `json:"dosenId"`
	KelurahanId string    `json:"kelurahanId"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type LokasiDosenWithRelationResponse struct {
	Id          int               `json:"id"`
	DosenId     int               `json:"dosenId"`
	Dosen       DosenResponse     `json:"dosen"`
	KelurahanId string            `json:"kelurahanId"`
	Kelurahan   KelurahanResponse `json:"kelurahan"`
	CreatedAt   time.Time         `json:"createdAt"`
	UpdatedAt   time.Time         `json:"updatedAt"`
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
