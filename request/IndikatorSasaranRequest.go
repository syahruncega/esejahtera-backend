package request

type CreateIndikatorSasaranRequest struct {
	ProgramId            int    `json:"programId" binding:"required"`
	NamaIndikatorSasaran string `json:"namaIndikatorSasaran" binding:"required"`
}

type UpdateIndikatorSasaranRequest struct {
	ProgramId            int    `json:"programId"`
	NamaIndikatorSasaran string `json:"namaIndikatorSasaran"`
}
