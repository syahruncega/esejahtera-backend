package model

import "time"

type BidangUrusan struct {
	Id               int       `gorm:"column:id;primaryKey;not null;autoIncrement"`
	NamaBidangUrusan string    `gorm:"column:namaBidangUrusan;not null"`
	CreatedAt        time.Time `gorm:"column:createdAt;not null"`
	UpdatedAt        time.Time `gorm:"column:updatedAt;not null"`
}

type Instansi struct {
	Id             int       `gorm:"column:id;primaryKey;not null;autoIncrement"`
	NamaInstansi   string    `gorm:"column:namaInstansi;not null"`
	BidangUrusanId int       `gorm:"column:bidangUrusanId;not null"`
	CreatedAt      time.Time `gorm:"column:createdAt;not null"`
	UpdatedAt      time.Time `gorm:"column:updatedAt;not null"`
	BidangUrusan   BidangUrusan
}

type Program struct {
	Id                      int       `gorm:"column:id;primaryKey;not null;autoIncrement"`
	Sasaran                 string    `gorm:"column:sasaran;not null"`
	IndikatorSasaran        string    `gorm:"column:indikatorSasaran;not null"`
	Kebijakan               string    `gorm:"column:kebijakan;not null"`
	NamaProgram             string    `gorm:"column:namaProgram;not null"`
	IndikatorKinerjaProgram string    `gorm:"column:indikatorKinerjaProgram;not null"`
	PaguProgram             int       `gorm:"column:paguProgram"`
	InstansiId              int       `gorm:"column:instansiId;not null"`
	CreatedAt               time.Time `gorm:"column:createdAt;not null"`
	UpdatedAt               time.Time `gorm:"column:updatedAt;not null"`
	Instansi                Instansi
}

type Kegiatan struct {
	Id                       int       `gorm:"column:id;primaryKey;not null;autoIncrement"`
	NamaKegiatan             string    `gorm:"column:namaKegiatan;not null"`
	IndikatorKinerjaKegiatan string    `gorm:"column:indikatorKinerjaKegiatan;not null"`
	PaguKegiatan             int       `gorm:"column:paguKegiatan"`
	ProgramId                int       `gorm:"column:programId;not null"`
	CreatedAt                time.Time `gorm:"column:createdAt;not null"`
	UpdatedAt                time.Time `gorm:"column:updatedAt;not null"`
	Program                  Program
}

type Sub_Kegiatan struct {
	Id                          int       `gorm:"column:id;primaryKey;not null;autoIncrement"`
	NamaSubKegiatan             string    `gorm:"column:namaSubKegiatan;not null"`
	IndikatorKinerjaSubKegiatan string    `gorm:"column:indikatorKinerjaSubKegiatan;not null"`
	PaguSubKegiatan             int       `gorm:"column:paguSubKegiatan"`
	KegiatanId                  int       `gorm:"column:kegiatanId;not null"`
	CreatedAt                   time.Time `gorm:"column:createdAt;not null"`
	UpdatedAt                   time.Time `gorm:"column:updatedAt;not null"`
	Kegiatan                    Kegiatan
}

type Detail_Sub_Kegiatan struct {
	Id               int       `gorm:"column:id;primaryKey;not null;autoIncrement"`
	FokusBelanja     string    `gorm:"column:fokusBelanja;not null"`
	Indikator        string    `gorm:"column:indikator;not null"`
	Target           float32   `gorm:"column:target"`
	Satuan           string    `gorm:"column:satuan;not null"`
	PaguFokusBelanja int       `gorm:"column:paguFokusBelanja"`
	Keterangan       string    `gorm:"column:keterangan;not null"`
	SubKegiatanId    int       `gorm:"column:subKegiatanId;not null"`
	CreatedAt        time.Time `gorm:"column:createdAt;not null"`
	UpdatedAt        time.Time `gorm:"column:updatedAt;not null"`
	SubKegiatan      Sub_Kegiatan
}

type Detail_Lokasi struct {
	Id                  int       `gorm:"column:id;primaryKey;not null;autoIncrement"`
	KabupatenKotaId     string    `gorm:"column:kabupatenKotaId;not null"`
	KecamatanId         string    `gorm:"column:kecamatanId;not null"`
	KelurahanId         string    `gorm:"column:kelurahanId;not null"`
	DetailSubKegiatanId int       `gorm:"column:detailSubKegiatanId;not null"`
	CreatedAt           time.Time `gorm:"column:createdAt;not null"`
	UpdatedAt           time.Time `gorm:"column:updatedAt;not null"`
	DetailSubKegiatan   Detail_Sub_Kegiatan
	KabupatenKota       Kabupaten_Kota
	Kecamatan           Kecamatan
	Kelurahan           Kelurahan
}

type Provinsi struct {
	Id   string `gorm:"column:id;primaryKey;not null"`
	Nama string `gorm:"column:nama;not null"`
}

type Kabupaten_Kota struct {
	Id         string `gorm:"column:id;primaryKey;not null"`
	ProvinsiId string `gorm:"column:provinsiId;not null"`
	Nama       string `gorm:"column:nama;not null"`
}

type Kecamatan struct {
	Id              string `gorm:"column:id;primaryKey;not null"`
	KabupatenKotaId string `gorm:"column:kabupatenKotaId;not null"`
	Nama            string `gorm:"column:nama;not null"`
}

type Kelurahan struct {
	Id          string `gorm:"column:id;primaryKey;not null"`
	KecamatanId string `gorm:"column:kecamatanId;not null"`
	Nama        string `gorm:"column:nama;not null"`
}

type Keluarga struct {
	Id                     int    `gorm:"column:id;primaryKey;not null"`
	IdKeluarga             string `gorm:"column:idKeluarga"`
	ProvinsiId             string `gorm:"column:provinsiId"`
	Provinsi               Provinsi
	KabupatenKotaId        string `gorm:"column:kabupatenKotaId"`
	KabupatenKota          Kabupaten_Kota
	KecamatanId            string `gorm:"column:kecamatanId"`
	Kecamatan              Kecamatan
	KelurahanId            string `gorm:"column:kelurahanId"`
	Kelurahan              Kelurahan
	DesilKesejahteraan     string `gorm:"column:desilKesejahteraan"`
	Alamat                 string `gorm:"column:alamat"`
	KepalaKeluarga         string `gorm:"column:kepalaKeluarga"`
	Nik                    string `gorm:"column:nik"`
	PadanDukcapil          string `gorm:"column:padanDukcapil"`
	JenisKelamin           string `gorm:"column:jenisKelamin"`
	TanggalLahir           string `gorm:"column:tanggalLahir"`
	Pekerjaan              string `gorm:"column:pekerjaan"`
	Pendidikan             string `gorm:"column:pendidikan"`
	KepemilikanRumah       string `gorm:"column:kepemilikanRumah"`
	Simpanan               string `gorm:"column:simpanan"`
	JenisAtap              string `gorm:"column:jenisAtap"`
	JenisDinding           string `gorm:"column:jenisDinding"`
	JenisLantai            string `gorm:"column:jenisLantai"`
	SumberPenerangan       string `gorm:"column:sumberPenerangan"`
	BahanBakarMemasak      string `gorm:"column:bahanBakarMemasak"`
	SumberAirMinum         string `gorm:"column:sumberAirMinum"`
	FasilitasBuangAirBesar string `gorm:"column:fasilitasBuangAirBesar"`
	PenerimaBPNT           string `gorm:"column:penerimaBPNT"`
	PenerimaBPUM           string `gorm:"column:penerimaBPUM"`
	PenerimaBST            string `gorm:"column:penerimaBST"`
	PenerimaPKH            string `gorm:"column:penerimaPKH"`
	PenerimaSembako        string `gorm:"column:penerimaSembako"`
	UserId                 int    `gorm:"column:userId"`
	User                   User
	StatusVerifikasi       int `gorm:"column:statusVerifikasi"`
}

type Monev struct {
	Id              int    `gorm:"column:id;primaryKey;not null"`
	NamaPenerima    string `gorm:"column:namaPenerima"`
	KabupatenKotaId string `gorm:"column:kabupatenKotaId"`
	KabupatenKota   Kabupaten_Kota
	KecamatanId     string `gorm:"column:kecamatanId"`
	Kecamatan       Kecamatan
	KelurahanId     string `gorm:"column:kelurahanId"`
	Kelurahan       Kelurahan
	JenisBantuan    string  `gorm:"column:jenisBantuan"`
	VolumeBantuan   float32 `gorm:"column:volumeBantuan"`
	SatuanVolume    string  `gorm:"column:satuanVolume"`
}

type User struct {
	Id        int       `gorm:"column:id;primaryKey;not null;autoIncrement"`
	Username  string    `gorm:"column:username;not null"`
	Password  string    `gorm:"column:password;not null"`
	Email     string    `gorm:"column:email; not null"`
	NoHp      string    `gorm:"column:noHp; not null"`
	Role      string    `gorm:"column:role;not null"`
	CreatedAt time.Time `gorm:"column:createdAt;not null"`
	UpdatedAt time.Time `gorm:"column:updatedAt;not null"`
}

type Admin struct {
	Id          int `gorm:"column:id;primaryKey;not null;autoIncrement"`
	UserId      int `gorm:"column:userId;not null"`
	User        User
	NamaLengkap string    `gorm:"column:namaLengkap"`
	CreatedAt   time.Time `gorm:"column:createdAt;not null"`
	UpdatedAt   time.Time `gorm:"column:updatedAt;not null"`
}

type Analis struct {
	Id          int `gorm:"column:id;primaryKey;not null;autoIncrement"`
	UserId      int `gorm:"column:userId;not null"`
	User        User
	NamaLengkap string    `gorm:"column:namaLengkap"`
	Universitas string    `gorm:"column:universitas"`
	LokasiKkn   string    `gorm:"column:lokasiKkn"`
	CreatedAt   time.Time `gorm:"column:createdAt;not null"`
	UpdatedAt   time.Time `gorm:"column:updatedAt;not null"`
}

type Pusbang struct {
	Id          int `gorm:"column:id;primaryKey;not null;autoIncrement"`
	UserId      int `gorm:"column:userId;not null"`
	User        User
	NamaLengkap string    `gorm:"column:namaLengkap"`
	Universitas string    `gorm:"column:universitas"`
	CreatedAt   time.Time `gorm:"column:createdAt;not null"`
	UpdatedAt   time.Time `gorm:"column:updatedAt;not null"`
}

type Dosen struct {
	Id          int `gorm:"column:id;primaryKey;not null;autoIncrement"`
	UserId      int `gorm:"column:userId;not null"`
	User        User
	NamaLengkap string    `gorm:"column:namaLengkap"`
	Universitas string    `gorm:"column:universitas"`
	CreatedAt   time.Time `gorm:"column:createdAt;not null"`
	UpdatedAt   time.Time `gorm:"column:updatedAt;not null"`
}

type Mahasiswa struct {
	Id           int `gorm:"column:id;primaryKey;not null;autoIncrement"`
	UserId       int `gorm:"column:userId;not null"`
	User         User
	NamaLengkap  string `gorm:"column:namaLengkap"`
	Universitas  string `gorm:"column:universitas"`
	JenisKelamin string `gorm:"column:jenisKelamin"`
	TanggalLahir string `gorm:"column:tanggalLahir"`
	KelurahanId  string `gorm:"column:kelurahanId"`
	Kelurahan    Kelurahan
	CreatedAt    time.Time `gorm:"column:createdAt;not null"`
	UpdatedAt    time.Time `gorm:"column:updatedAt;not null"`
}

type Lokasi_Dosen struct {
	Id          int `gorm:"column:id;primaryKey;not null; autoIncrement"`
	DosenId     int `gorm:"column:dosenId;not null"`
	Dosen       Dosen
	KelurahanId string `gorm:"column:kelurahanId;not null"`
	Kelurahan   Kelurahan
	CreatedAt   time.Time `gorm:"column:createdAt;not null"`
	UpdatedAt   time.Time `gorm:"column:updatedAt;not null"`
}

type KeluargaVerifikasi struct {
	Id                     int    `gorm:"column:id;primaryKey;not null"`
	IdKeluarga             string `gorm:"column:idKeluarga;not null"`
	ProvinsiId             string `gorm:"column:provinsiId;not null"`
	Provinsi               Provinsi
	KabupatenKotaId        string `gorm:"column:kabupatenKotaId;not null"`
	KabupatenKota          Kabupaten_Kota
	KecamatanId            string `gorm:"column:kecamatanId;not null"`
	Kecamatan              Kecamatan
	KelurahanId            string `gorm:"column:kelurahanId; not null"`
	DesilKesejahteraan     string `gorm:"column:desilKesejahteraan; not null"`
	Alamat                 string `gorm:"column:alamat;not null"`
	KepalaKeluarga         string `gorm:"column:kepalaKeluarga;not null"`
	Nik                    string `gorm:"column:nik;not null"`
	PadanDukcapil          string `gorm:"column:padanDukcapil;not null"`
	JenisKelamin           string `gorm:"column:jenisKelamin;not null"`
	TanggalLahir           string `gorm:"column:tanggalLahir;not null"`
	Pekerjaan              string `gorm:"column:pekerjaan;not null"`
	Pendidikan             string `gorm:"column:pendidikan;not null"`
	KepemilikanRumah       string `gorm:"column:kepemilikanRumah;not null"`
	Simpanan               string `gorm:"column:simpanan;not null"`
	JenisAtap              string `gorm:"column:jenisAtap;not null"`
	JenisDinding           string `gorm:"column:jenisDinding;not null"`
	JenisLantai            string `gorm:"column:jenisLantai;not null"`
	SumberPenerangan       string `gorm:"column:sumberPenerangan;not null"`
	BahanBakarMemasak      string `gorm:"column:bahanBakarMemasak;not null"`
	SumberAirMinum         string `gorm:"column:sumberAirMinum;not null"`
	FasilitasBuangAirBesar string `gorm:"column:fasilitasBuangAirBesar;not null"`
	PenerimaBPNT           string `gorm:"column:penerimaBPNT;not null"`
	PenerimaBPUM           string `gorm:"column:penerimaBPUM;not null"`
	PenerimaBST            string `gorm:"column:penerimaBST;not null"`
	PenerimaPKH            string `gorm:"column:penerimaPKH;not null"`
	PenerimaSembako        string `gorm:"column:penerimaSembako;not null"`
	PenerimaLainnya        string `gorm:"column:penerimaLainnya"`
	StatusResponden        string `gorm:"column:statusResponden;not null"`
	UserId                 int    `gorm:"column:userId;not null"`
	User                   User
	MahasiswaId            int `gorm:"column:mahasiswaId;not null"`
	Mahasiswa              Mahasiswa
	Kelurahan              Kelurahan
	CreatedAt              time.Time `gorm:"column:createdAt;not null"`
	UpdatedAt              time.Time `gorm:"column:updatedAt;not null"`
}
