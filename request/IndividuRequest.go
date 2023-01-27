package request

type UpdateIndividuRequest struct {
	UserId           int `json:"userId"`
	MahasiswaId      int `json:"mahasiswaId"`
	StatusVerifikasi int `json:"statusVerifikasi"`
}
