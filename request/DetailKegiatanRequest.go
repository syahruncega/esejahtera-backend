package request

type CreateDetailKegiatanRequest struct {
	ProgramId    int   `json:"programId" binding:"required"`
	KegiatanId   int   `json:"kegiatanId" binding:"required"`
	PaguKegiatan int64 `json:"paguKegiatan" binding:"required"`
}

type UpdateDetailKegiatanRequest struct {
	ProgramId    int   `json:"programId"`
	KegiatanId   int   `json:"kegiatanId"`
	PaguKegiatan int64 `json:"paguKegiatan"`
}
