package request

type CreateProgramRequest struct {
	ProgramId   string `json:"programId" binding:"required"`
	NamaProgram string `json:"namaProgram" binding:"required"`
}

type UpdateProgramRequest struct {
	ProgramId   string `json:"programId"`
	NamaProgram string `json:"namaProgram" binding:"required"`
}
