package request

type CreateSasaranRequest struct {
	RencanaProgramId int    `json:"rencanaProgramId" binding:"required"`
	NamaSasaran      string `json:"namaSasaran" binding:"required"`
}

type UpdateSasaranRequest struct {
	RencanaProgramId int    `json:"rencanaProgramId"`
	NamaSasaran      string `json:"namaSasaran"`
}
