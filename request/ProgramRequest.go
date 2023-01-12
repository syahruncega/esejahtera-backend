package request

type CreateProgramRequest struct {
	Sasaran                 string `json:"sasaran" binding:"required"`
	IndikatorSasaran        string `json:"indikatorSasaran" binding:"required"`
	Kebijakan               string `json:"kebijakan" binding:"required"`
	NamaProgram             string `json:"namaProgram" binding:"required"`
	IndikatorKinerjaProgram string `json:"indikatorKinerjaProgram" binding:"required"`
	PaguProgram             int    `json:"paguProgram"`
	BidangUrusanId          int    `json:"bidangUrusanId" binding:"required"`
	InstansiId              int    `json:"instansiId" binding:"required"`
}

type UpdateProgramRequest struct {
	Sasaran                 string `json:"sasaran"`
	IndikatorSasaran        string `json:"indikatorSasaran"`
	Kebijakan               string `json:"kebijakan"`
	NamaProgram             string `json:"namaProgram"`
	IndikatorKinerjaProgram string `json:"indikatorKinerjaProgram"`
	PaguProgram             int    `json:"paguProgram"`
	BidangUrusanId          int    `json:"bidangUrusanId"`
	InstansiId              int    `json:"instansiId"`
}
