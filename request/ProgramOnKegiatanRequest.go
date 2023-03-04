package request

type CreateProgramOnKegiatanRequest struct {
	ProgramId  int `json:"programId" binding:"required"`
	KegiatanId int `json:"kegiatanId" binding:"required"`
}

type UpdateProgramOnKegiatanRequest struct {
	ProgramId  int `json:"programId"`
	KegiatanId int `json:"kegiatanId"`
}
