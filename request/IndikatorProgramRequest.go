package request

type CreateIndikatorProgramRequest struct {
	InstansiId              string `json:"instansiId" binding:"required"`
	ProgramId               string `json:"programId" binding:"required"`
	Sasaran                 string `json:"sasaran" binding:"required"`
	IndikatorSasaran        string `json:"indikatorSasaran" binding:"required"`
	Kebijakan               string `json:"kebijakan" binding:"required"`
	IndikatorKinerjaProgram string `json:"indikatorKinerjaProgram" binding:"required"`
	PaguProgram             int    `json:"paguProgram"`
}

type UpdateIndikatorProgramRequest struct {
	Sasaran                 string `json:"sasaran"`
	IndikatorSasaran        string `json:"indikatorSasaran"`
	Kebijakan               string `json:"kebijakan"`
	IndikatorKinerjaProgram string `json:"indikatorKinerjaProgram"`
	PaguProgram             int    `json:"paguProgram"`
	InstansiId              string `json:"instansiId"`
	ProgramId               string `json:"programId"`
}
