package request

type CreateDetailProgramRequest struct {
	ProgramId   int    `json:"programId" binding:"required"`
	PaguProgram int64  `json:"paguProgram" binding:"required"`
	Tipe        string `json:"tipe" binding:"required"`
}

type UpdateDetailProgramRequest struct {
	ProgramId   int    `json:"programId"`
	PaguProgram int64  `json:"paguProgram"`
	Tipe        string `json:"tipe"`
}
