package request

type CreateKebijakanRequest struct {
	RencanaProgramId int    `json:"rencanaProgramId" binding:"required"`
	NamaKebijakan    string `json:"namaKebijakan" binding:"required"`
}

type UpdateKebijakanRequest struct {
	RencanaProgramId int    `json:"rencanaProgramId"`
	NamaKebijakan    string `json:"namaKebijakan"`
}
