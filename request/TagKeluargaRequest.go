package request

type CreateTagKeluargaRequest struct {
	FokusBelanjaId int `json:"fokusBelanjaId" binding:"required"`
	KeluargaId     int `json:"keluargaId" binding:"required"`
}

type UpdateTagKeluargaRequest struct {
	FokusBelanjaId int `json:"fokusBelanjaId"`
	KeluargaId     int `json:"keluargaId"`
}
