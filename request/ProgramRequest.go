package request

type CreateProgramRequest struct {
	Id          string `json:"id" binding:"required"`
	NamaProgram string `json:"namaProgram" binding:"required"`
}

type UpdateProgramRequest struct {
	Id          string `json:"id"`
	NamaProgram string `json:"namaProgram" binding:"required"`
}
