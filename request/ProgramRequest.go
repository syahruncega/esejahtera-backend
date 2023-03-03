package request

type CreateProgramRequest struct {
	InstansiId  int    `json:"instansiId" binding:"required"`
	Tahun       string `json:"tahun" binding:"required"`
	KodeProgram string `json:"kodeProgram" binding:"required"`
	NamaProgram string `json:"namaProgram" binding:"required"`
}

type UpdateProgramRequest struct {
	InstansiId  int    `json:"instansiId"`
	Tahun       string `json:"tahun"`
	KodeProgram string `json:"kodeProgram"`
	NamaProgram string `json:"namaProgram"`
}
