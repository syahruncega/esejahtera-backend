package request

type CreateIndikatorSasaranRequest struct {
	RencanaProgramId     int    `json:"rencanaProgramId" binding:"required"`
	NamaIndikatorSasaran string `json:"namaIndikatorSasaran" binding:"required"`
}

type UpdateIndikatorSasaranRequest struct {
	RencanaProgramId     int    `json:"rencanaProgramId"`
	NamaIndikatorSasaran string `json:"namaIndikatorSasaran"`
}
