package request

type CreateRencanaSubKegiatanRequest struct {
	RencanaKegiatanId int    `json:"rencanaKegiatanId" binding:"required"`
	SubKegiatanId     int    `json:"subKegiatanId" binding:"required"`
	PaguSubKegiatan   int64  `json:"paguSubKegiatan" binding:"required"`
	Tipe              string `json:"tipe" binding:"required"`
	Tahun             string `json:"tahun" binding:"required"`
}

type UpdateRencanaSubKegiatanRequest struct {
	RencanaKegiatanId int    `json:"rencanaKegiatanId"`
	SubKegiatanId     int    `json:"subKegiatanId"`
	PaguSubKegiatan   int64  `json:"paguSubKegiatan"`
	Tipe              string `json:"tipe"`
	Tahun             string `json:"tahun"`
}
