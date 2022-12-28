package request

type CreateProgramRequest struct {
	NamaProgram             string `json:"namaProgram" binding:"required"`
	IndikatorKinerjaProgram string `json:"indikatorKinerjaProgram" binding:"required"`
	PaguProgram             int    `json:"paguProgram"`
	InstansiId              int    `json:"instansiId" binding:"required"`
}

type UpdateProgramRequest struct {
	NamaProgram             string `json:"namaProgram"`
	IndikatorKinerjaProgram string `json:"indikatorKinerjaProgram"`
	PaguProgram             int    `json:"paguProgram"`
	InstansiId              int    `json:"instansiId"`
}
