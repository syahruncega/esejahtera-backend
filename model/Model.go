package model

import "time"

type BidangUrusan struct {
	Id               int       `gorm:"column:id;primaryKey;not null;autoIncrement"`
	KodeBidangUrusan string    `gorm:"column:kodeBidangUrusan;not null;unique"`
	NamaBidangUrusan string    `gorm:"column:namaBidangUrusan;not null"`
	CreatedAt        time.Time `gorm:"column:createdAt;not null"`
	UpdatedAt        time.Time `gorm:"column:updatedAt;not null"`
}

type Instansi struct {
	Id             int `gorm:"column:id;primaryKey;not null;autoIncrement"`
	BidangUrusanId int `gorm:"column:bidangUrusanId;not null"`
	BidangUrusan   BidangUrusan
	KodeInstansi   string    `gorm:"column:kodeInstansi;not null"`
	NamaInstansi   string    `gorm:"column:namaInstansi;not null"`
	CreatedAt      time.Time `gorm:"column:createdAt;not null"`
	UpdatedAt      time.Time `gorm:"column:updatedAt;not null"`
}

type Program struct {
	Id          int `gorm:"column:id;primaryKey;not null;autoIncrement"`
	InstansiId  int `gorm:"column:instansiId;not null"`
	Instansi    Instansi
	Tahun       string    `gorm:"column:tahun;not null"`
	KodeProgram string    `gorm:"column:kodeProgram;not null"`
	NamaProgram string    `gorm:"column:namaProgram;not null"`
	CreatedAt   time.Time `gorm:"column:createdAt;not null"`
	UpdatedAt   time.Time `gorm:"column:updatedAt;not null"`
}

type DetailProgram struct {
	Id          int `gorm:"column:id;primaryKey;not null;autoIncrement"`
	ProgramId   int `gorm:"column:programId;not null"`
	Program     Program
	PaguProgram int64     `gorm:"column:paguProgram;not null"`
	Tipe        string    `gorm:"column:tipe;not null"`
	CreatedAt   time.Time `gorm:"column:createdAt;not null"`
	UpdatedAt   time.Time `gorm:"column:updatedAt;not null"`
}

type Sasaran struct {
	Id          int `gorm:"column:id;primaryKey;not null;autoIncrement"`
	ProgramId   int `gorm:"column:programId;not null"`
	Program     Program
	NamaSasaran string    `gorm:"column:namaSasaran;not null"`
	CreatedAt   time.Time `gorm:"column:createdAt;not null"`
	UpdatedAt   time.Time `gorm:"column:updatedAt;not null"`
}

type IndikatorSasaran struct {
	Id                   int `gorm:"column:id;primaryKey;not null;autoIncrement"`
	ProgramId            int `gorm:"column:programId;not null"`
	Program              Program
	NamaIndikatorSasaran string    `gorm:"column:namaIndikatorSasaran;not null"`
	CreatedAt            time.Time `gorm:"column:createdAt;not null"`
	UpdatedAt            time.Time `gorm:"column:updatedAt;not null"`
}

type Kegiatan struct {
	Id           int `gorm:"column:id;primaryKey;not null;autoIncrement"`
	ProgramId    int `gorm:"column:programId;not null"`
	Program      Program
	Tahun        string    `gorm:"column:tahun;not null"`
	KodeKegiatan string    `gorm:"column:kodeKegiatan;not null"`
	NamaKegiatan string    `gorm:"column:namaKegiatan;not null"`
	CreatedAt    time.Time `gorm:"column:createdAt;not null"`
	UpdatedAt    time.Time `gorm:"column:updatedAt;not null"`
}

type DetailKegiatan struct {
	Id           int `gorm:"column:id;primaryKey;not null;autoIncrement"`
	KegiatanId   int `gorm:"column:kegiatanId;not null"`
	Kegiatan     Kegiatan
	PaguKegiatan int64     `gorm:"column:paguKegiatan;not null"`
	Tipe         string    `gorm:"column:tipe;not null"`
	CreatedAt    time.Time `gorm:"column:createdAt;not null"`
	UpdatedAt    time.Time `gorm:"column:updatedAt;not null"`
}

type SubKegiatan struct {
	Id              int `gorm:"column:id;primaryKey;not null;autoIncrement"`
	KegiatanId      int `gorm:"column:kegiatanId;not null"`
	Kegiatan        Kegiatan
	Tahun           string    `gorm:"column:tahun;not null"`
	KodeSubKegiatan string    `gorm:"column:kodeSubKegiatan;not null"`
	NamaSubKegiatan string    `gorm:"column:namaSubKegiatan;not null"`
	CreatedAt       time.Time `gorm:"column:createdAt;not null"`
	UpdatedAt       time.Time `gorm:"column:updatedAt;not null"`
}

type Kebijakan struct {
	Id            int `gorm:"column:id;primaryKey;not null;autoIncrement"`
	ProgramId     int `gorm:"column:programId;not null"`
	Program       Program
	NamaKebijakan string    `gorm:"column:namaKebijakan;not null"`
	CreatedAt     time.Time `gorm:"column:createdAt;not null"`
	UpdatedAt     time.Time `gorm:"column:updatedAt;not null"`
}

type DetailSubKegiatan struct {
	Id              int `gorm:"column:id;primaryKey;not null;autoIncrement"`
	SubKegiatanId   int `gorm:"column:subKegiatanId;not null"`
	SubKegiatan     SubKegiatan
	PaguSubKegiatan int64     `gorm:"column:paguSubKegiatan;not null"`
	Tipe            string    `gorm:"column:tipe;not null"`
	CreatedAt       time.Time `gorm:"column:createdAt;not null"`
	UpdatedAt       time.Time `gorm:"column:updatedAt;not null"`
}

type FokusBelanja struct {
	Id               int `gorm:"column:id;primaryKey;not null;autoIncrement"`
	SubKegiatanId    int `gorm:"column:subKegiatanId;not null"`
	SubKegiatan      SubKegiatan
	NamaFokusBelanja string    `gorm:"column:namaFokusBelanja;not null"`
	Indikator        string    `gorm:"column:indikator;not null"`
	Target           float32   `gorm:"column:target;not null"`
	Satuan           string    `gorm:"column:satuan;not null"`
	PaguFokusBelanja int       `gorm:"column:paguFokusBelanja;not null"`
	Keterangan       string    `gorm:"column:keterangan;not null"`
	CreatedAt        time.Time `gorm:"column:createdAt;not null"`
	UpdatedAt        time.Time `gorm:"column:updatedAt;not null"`
}

type Detail_Lokasi struct {
	Id              int `gorm:"column:id;primaryKey;not null;autoIncrement"`
	FokusBelanjaId  int `gorm:"column:fokusBelanjaId;not null"`
	FokusBelanja    FokusBelanja
	KabupatenKotaId string `gorm:"column:kabupatenKotaId;not null"`
	KabupatenKota   Kabupaten_Kota
	KecamatanId     string `gorm:"column:kecamatanId;not null"`
	Kecamatan       Kecamatan
	KelurahanId     string `gorm:"column:kelurahanId;not null"`
	Kelurahan       Kelurahan
	CreatedAt       time.Time `gorm:"column:createdAt;not null"`
	UpdatedAt       time.Time `gorm:"column:updatedAt;not null"`
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
	MahasiswaId            int `gorm:"column:mahasiswaId"`
	Mahasiswa              Mahasiswa
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
	Id            int `gorm:"column:id;primaryKey;not null;autoIncrement"`
	UserId        int `gorm:"column:userId;not null"`
	User          User
	NamaLengkap   string    `gorm:"column:namaLengkap"`
	UrlFotoProfil string    `gorm:"column:urlFotoProfil"`
	CreatedAt     time.Time `gorm:"column:createdAt;not null"`
	UpdatedAt     time.Time `gorm:"column:updatedAt;not null"`
}

type Analis struct {
	Id            int `gorm:"column:id;primaryKey;not null;autoIncrement"`
	UserId        int `gorm:"column:userId;not null"`
	User          User
	NamaLengkap   string    `gorm:"column:namaLengkap"`
	Universitas   string    `gorm:"column:universitas"`
	UrlFotoProfil string    `gorm:"column:urlFotoProfil"`
	CreatedAt     time.Time `gorm:"column:createdAt;not null"`
	UpdatedAt     time.Time `gorm:"column:updatedAt;not null"`
}

type Pusbang struct {
	Id            int `gorm:"column:id;primaryKey;not null;autoIncrement"`
	UserId        int `gorm:"column:userId;not null"`
	User          User
	NamaLengkap   string    `gorm:"column:namaLengkap"`
	Universitas   string    `gorm:"column:universitas"`
	UrlFotoProfil string    `gorm:"column:urlFotoProfil"`
	CreatedAt     time.Time `gorm:"column:createdAt;not null"`
	UpdatedAt     time.Time `gorm:"column:updatedAt;not null"`
}

type Dosen struct {
	Id            int `gorm:"column:id;primaryKey;not null;autoIncrement"`
	UserId        int `gorm:"column:userId;not null"`
	User          User
	NamaLengkap   string    `gorm:"column:namaLengkap"`
	Universitas   string    `gorm:"column:universitas"`
	UrlFotoProfil string    `gorm:"column:urlFotoProfil"`
	CreatedAt     time.Time `gorm:"column:createdAt;not null"`
	UpdatedAt     time.Time `gorm:"column:updatedAt;not null"`
}

type Mahasiswa struct {
	Id              int `gorm:"column:id;primaryKey;not null;autoIncrement"`
	UserId          int `gorm:"column:userId;not null"`
	User            User
	NamaLengkap     string `gorm:"column:namaLengkap"`
	Universitas     string `gorm:"column:universitas"`
	JenisKelamin    string `gorm:"column:jenisKelamin"`
	TanggalLahir    string `gorm:"column:tanggalLahir"`
	KabupatenKotaId string `gorm:"column:kabupatenKotaId;not null"`
	KabupatenKota   Kabupaten_Kota
	KecamatanId     string `gorm:"column:kecamatanId;not null"`
	Kecamatan       Kecamatan
	KelurahanId     string `gorm:"column:kelurahanId;not null"`
	Kelurahan       Kelurahan
	UrlFotoProfil   string    `gorm:"column:urlFotoProfil"`
	CreatedAt       time.Time `gorm:"column:createdAt;not null"`
	UpdatedAt       time.Time `gorm:"column:updatedAt;not null"`
}

type LokasiDosen struct {
	Id              int `gorm:"column:id;primaryKey;not null;autoIncrement"`
	DosenId         int `gorm:"column:dosenId;not null"`
	Dosen           Dosen
	KabupatenKotaId string `gorm:"column:kabupatenKotaId;not null"`
	KabupatenKota   Kabupaten_Kota
	KecamatanId     string `gorm:"column:kecamatanId;not null"`
	Kecamatan       Kecamatan
	KelurahanId     string `gorm:"column:kelurahanId;not null"`
	Kelurahan       Kelurahan
	CreatedAt       time.Time `gorm:"column:createdAt;not null"`
	UpdatedAt       time.Time `gorm:"column:updatedAt;not null"`
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
	UrlBukti               string    `gorm:"column:urlBukti;not null"`
	CreatedAt              time.Time `gorm:"column:createdAt;not null"`
	UpdatedAt              time.Time `gorm:"column:updatedAt;not null"`
}

type Individu struct {
	Id                 int    `gorm:"column:id;primaryKey;not null"`
	IdKeluarga         string `gorm:"column:idKeluarga"`
	ProvinsiId         string `gorm:"column:provinsiId"`
	Provinsi           Provinsi
	KabupatenKotaId    string `gorm:"column:kabupatenKotaId"`
	KabupatenKota      Kabupaten_Kota
	KecamatanId        string `gorm:"column:kecamatanId"`
	Kecamatan          Kecamatan
	KelurahanId        string `gorm:"column:kelurahanId"`
	Kelurahan          Kelurahan
	DesilKesejahteraan string `gorm:"column:desilKesejahteraan"`
	Alamat             string `gorm:"column:alamat"`
	IdIndividu         string `gorm:"column:idIndividu"`
	Nama               string `gorm:"column:nama"`
	Nik                string `gorm:"column:nik"`
	PadanDukcapil      string `gorm:"column:padanDukcapil"`
	JenisKelamin       string `gorm:"column:jenisKelamin"`
	Hubungan           string `gorm:"column:hubungan"`
	TanggalLahir       string `gorm:"column:tanggalLahir"`
	StatusKawin        string `gorm:"column:statusKawin"`
	Pekerjaan          string `gorm:"column:pekerjaan"`
	Pendidikan         string `gorm:"column:pendidikan"`
	KategoriUsia       string `gorm:"column:kategoriUsia"`
	PenerimaBPNT       string `gorm:"column:penerimaBPNT"`
	PenerimaBPUM       string `gorm:"column:penerimaBPUM"`
	PenerimaBST        string `gorm:"column:penerimaBST"`
	PenerimaPKH        string `gorm:"column:penerimaPKH"`
	PenerimaSembako    string `gorm:"column:penerimaSembako"`
	UserId             int    `gorm:"column:userId"`
	User               User
	MahasiswaId        int `gorm:"column:mahasiswaId"`
	Mahasiswa          Mahasiswa
	StatusVerifikasi   int `gorm:"column:statusVerifikasi"`
}

type IndividuVerifikasi struct {
	Id                 int    `gorm:"column:id;primaryKey;not null"`
	IdKeluarga         string `gorm:"column:idKeluarga;not null"`
	ProvinsiId         string `gorm:"column:provinsiId;not null"`
	Provinsi           Provinsi
	KabupatenKotaId    string `gorm:"column:kabupatenKotaId;not null"`
	KabupatenKota      Kabupaten_Kota
	KecamatanId        string `gorm:"column:kecamatanId;not null"`
	Kecamatan          Kecamatan
	KelurahanId        string `gorm:"column:kelurahanId;not null"`
	Kelurahan          Kelurahan
	DesilKesejahteraan string `gorm:"column:desilKesejahteraan;not null"`
	Alamat             string `gorm:"column:alamat;not null"`
	IdIndividu         string `gorm:"column:idIndividu;not null"`
	Nama               string `gorm:"column:nama;not null"`
	Nik                string `gorm:"column:nik;not null"`
	PadanDukcapil      string `gorm:"column:padanDukcapil;not null"`
	JenisKelamin       string `gorm:"column:jenisKelamin;not null"`
	Hubungan           string `gorm:"column:hubungan;not null"`
	TanggalLahir       string `gorm:"column:tanggalLahir;not null"`
	StatusKawin        string `gorm:"column:statusKawin;not null"`
	Pekerjaan          string `gorm:"column:pekerjaan;not null"`
	Pendidikan         string `gorm:"column:pendidikan;not null"`
	KategoriUsia       string `gorm:"column:kategoriUsia;not null"`
	PenerimaBPNT       string `gorm:"column:penerimaBPNT;not null"`
	PenerimaBPUM       string `gorm:"column:penerimaBPUM;not null"`
	PenerimaBST        string `gorm:"column:penerimaBST;not null"`
	PenerimaPKH        string `gorm:"column:penerimaPKH;not null"`
	PenerimaSembako    string `gorm:"column:penerimaSembako;not null"`
	PenerimaLainnya    string `gorm:"column:penerimaLainnya;not null"`
	StatusResponden    string `gorm:"column:statusResponden;not null"`
	UserId             int    `gorm:"column:userId;not null"`
	User               User
	MahasiswaId        int `gorm:"column:mahasiswaId;not null"`
	Mahasiswa          Mahasiswa
	UrlBukti           string    `gorm:"column:urlBukti;not null"`
	CreatedAt          time.Time `gorm:"column:createdAt;not null"`
	UpdatedAt          time.Time `gorm:"column:updatedAt;not null"`
}

type DistinctKabupatenKota struct {
	KabupatenKotaId string `gorm:"column:kabupatenKotaId"`
	Nama            string `gorm:"column:nama"`
}

type DistinctKecamatan struct {
	KecamatanId string `gorm:"column:kecamatanId"`
	Nama        string `gorm:"column:nama"`
}

type DistinctKelurahan struct {
	KelurahanId string `gorm:"column:kelurahanId"`
	Nama        string `gorm:"column:nama"`
}
