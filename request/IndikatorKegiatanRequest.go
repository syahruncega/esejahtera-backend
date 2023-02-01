package request

type CreateIndikatorKegiatanRequest struct {
	ProgramId                string `json:"programId" binding:"required"`
	KegiatanId               string `json:"kegiatanId" binding:"required"`
	IndikatorKinerjaKegiatan string `json:"indikatorKinerjaKegiatan" binding:"required"`
	PaguKegiatan             int    `json:"paguKegiatan"`
}

type UpdateIndikatorKegiatanRequest struct {
	ProgramId                string `json:"programId"`
	KegiatanId               string `json:"kegiatanId"`
	IndikatorKinerjaKegiatan string `json:"indikatorKinerjaKegiatan"`
	PaguKegiatan             int    `json:"paguKegiatan"`
}
