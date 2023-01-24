package request

type UpdateKeluargaRequest struct {
	UserId           int `json:"userId"`
	StatusVerifikasi int `json:"statusVerifikasi"`
}
