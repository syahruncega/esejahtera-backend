package request

type CreateDetailProgramRequest struct {
	InstansiId  int   `json:"instansiId" binding:"required"`
	ProgramId   int   `json:"programId" binding:"required"`
	PaguProgram int64 `json:"paguProgram" binding:"required"`
}

type UpdateDetailProgramRequest struct {
	InstansiId  int   `json:"instansiId"`
	ProgramId   int   `json:"programId"`
	PaguProgram int64 `json:"paguProgram"`
}
