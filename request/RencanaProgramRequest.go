package request

type CreateRencanaProgramRequest struct {
	InstansiId  int    `json:"instansiId" binding:"required"`
	ProgramId   int    `json:"programId" binding:"required"`
	PaguProgram int64  `json:"paguProgram" binding:"required"`
	Tipe        string `json:"tipe" binding:"required"`
	Tahun       string `json:"tahun" binding:"required"`
}

type UpdateRencanaProgramRequest struct {
	InstansiId  int    `json:"instansiId"`
	ProgramId   int    `json:"programId"`
	PaguProgram int64  `json:"paguProgram"`
	Tipe        string `json:"tipe"`
	Tahun       string `json:"tahun"`
}
