package request

type CreateKegiatanRequest struct {
	NamaKegiatan             string `json:"namaKegiatan" binding:"required"`
	IndikatorKinerjaKegiatan string `json:"indikatorKinerjaKegiatan" binding:"required"`
	PaguKegiatan             int    `json:"paguKegiatan" binding:"required"`
	ProgramId                int    `json:"programId" binding:"required"`
}

type UpdateKegiatanRequest struct {
	NamaKegiatan             string `json:"namaKegiatan"`
	IndikatorKinerjaKegiatan string `json:"indikatorKinerjaKegiatan"`
	PaguKegiatan             int    `json:"paguKegiatan"`
	ProgramId                int    `json:"programId"`
}
