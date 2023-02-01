package request

type CreateDetailKegiatanRequest struct {
	ProgramId  string `json:"programId" binding:"required"`
	KegiatanId string `json:"kegiatanId" binding:"required"`
}

type UpdateDetailKegiatanRequest struct {
	ProgramId  string `json:"programId"`
	KegiatanId string `json:"kegiatanId"`
}
