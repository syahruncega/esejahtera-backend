package request

type CreateSasaranRequest struct {
	ProgramId   int    `json:"programId" binding:"required"`
	NamaSasaran string `json:"namaSasaran" binding:"required"`
}

type UpdateSasaranRequest struct {
	ProgramId   int    `json:"programId"`
	NamaSasaran string `json:"namaSasaran"`
}
