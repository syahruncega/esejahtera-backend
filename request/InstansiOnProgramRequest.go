package request

type CreateInstansiOnProgramRequest struct {
	InstansiId int `json:"instansiId" binding:"required"`
	ProgramId  int `json:"programId" binding:"required"`
}

type UpdateInstansiOnProgramRequest struct {
	InstansiId int `json:"instansiId" binding:"required"`
	ProgramId  int `json:"programId" binding:"required"`
}
