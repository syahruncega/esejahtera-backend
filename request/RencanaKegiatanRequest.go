package request

type CreateRencanaKegiatanRequest struct {
	RencanaProgramId int    `json:"rencanaProgramId" binding:"required"`
	KegiatanId       int    `json:"kegiatanId" binding:"required"`
	PaguKegiatan     int64  `json:"paguKegiatan" binding:"required"`
	Tipe             string `json:"tipe" binding:"required"`
	Tahun            string `json:"tahun" binding:"required"`
}

type UpdateRencanaKegiatanRequest struct {
	RencanaProgramId int    `json:"rencanaProgramId"`
	KegiatanId       int    `json:"kegiatanId"`
	PaguKegiatan     int64  `json:"paguKegiatan"`
	Tipe             string `json:"tipe"`
	Tahun            string `json:"tahun" binding:"required"`
}
