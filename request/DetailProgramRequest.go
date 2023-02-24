package request

type CreateDetailProgramRequest struct {
	InstansiId  string `json:"instansiId" binding:"required"`
	ProgramId   string `json:"programId" binding:"required"`
	PaguProgram int64  `json:"paguProgram" binding:"required"`
}

type UpdateDetailProgramRequest struct {
	InstansiId  string `json:"instansiId"`
	ProgramId   string `json:"programId"`
	PaguProgram int64  `json:"paguProgram"`
}
