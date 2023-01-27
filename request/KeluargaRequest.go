package request

type UpdateKeluargaRequest struct {
	UserId           int `json:"userId"`
	MahasiswaId      int `json:"mahasiswaId"`
	StatusVerifikasi int `json:"statusVerifikasi"`
}
