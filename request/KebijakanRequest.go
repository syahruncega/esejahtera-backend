package request

type CreateKebijakanRequest struct {
	ProgramId     int    `json:"programId" binding:"required"`
	NamaKebijakan string `json:"namaKebijakan" binding:"required"`
}

type UpdateKebijakanRequest struct {
	ProgramId     int    `json:"programId"`
	NamaKebijakan string `json:"namaKebijakan"`
}
