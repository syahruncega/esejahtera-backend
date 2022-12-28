package model

import "time"

type Instansi struct {
	Id           int       `gorm:"column:id;primaryKey;not null;autoIncrement"`
	NamaInstansi string    `gorm:"column:namaInstansi;not null"`
	CreatedAt    time.Time `gorm:"column:createdAt;not null"`
	UpdatedAt    time.Time `gorm:"column:updatedAt;not null"`
}

type Program struct {
	Id                      int       `gorm:"column:id;primaryKey;not null;autoIncrement"`
	NamaProgram             string    `gorm:"column:namaProgram;not null"`
	IndikatorKinerjaProgram string    `gorm:"column:indikatorKinerjaProgram;not null"`
	PaguProgram             int       `gorm:"column:paguProgram;not null"`
	InstansiId              int       `gorm:"column:instansiId;not null"`
	CreatedAt               time.Time `gorm:"column:createdAt;not null"`
	UpdatedAt               time.Time `gorm:"column:updatedAt;not null"`
	Instansi                Instansi
}

type Kegiatan struct {
	Id                       int       `gorm:"column:id;primaryKey;not null;autoIncrement"`
	NamaKegiatan             string    `gorm:"column:namaKegiatan;not null"`
	IndikatorKinerjaKegiatan string    `gorm:"column:indikatorKinerjaKegiatan;not null"`
	PaguKegiatan             int       `gorm:"column:paguKegiatan;not null"`
	ProgramId                int       `gorm:"column:programId;not null"`
	CreatedAt                time.Time `gorm:"column:createdAt;not null"`
	UpdatedAt                time.Time `gorm:"column:updatedAt;not null"`
	Program                  Program
}

type Sub_Kegiatan struct {
	Id                          int       `gorm:"column:id;primaryKey;not null;autoIncrement"`
	NamaSubKegiatan             string    `gorm:"column:namaSubKegiatan;not null"`
	IndikatorKinerjaSubKegiatan string    `gorm:"column:indikatorKinerjaSubKegiatan;not null"`
	PaguSubKegiatan             int       `gorm:"column:paguSubKegiatan;not null"`
	KegiatanId                  int       `gorm:"column:kegiatanId;not null"`
	CreatedAt                   time.Time `gorm:"column:createdAt;not null"`
	UpdatedAt                   time.Time `gorm:"column:updatedAt;not null"`
	Kegiatan                    Kegiatan
}

type Detail_Sub_Kegiatan struct {
	Id               int       `gorm:"column:id;primaryKey;not null;autoIncrement"`
	FokusBelanja     string    `gorm:"column:fokusBelanja;not null"`
	Indikator        string    `gorm:"column:indikator;not null"`
	Target           float32   `gorm:"column:target;not null"`
	Satuan           string    `gorm:"column:satuan;not null"`
	PaguFokusBelanja int       `gorm:"column:paguFokusBelanja;not null"`
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
