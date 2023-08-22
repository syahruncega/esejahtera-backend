package request

type CreateRealisasiRequest struct {
	FokusBelanjaId  int    `json:"fokusBelanjaId" binding:"required"`
	Tanggal         string `json:"tanggal" binding:"required"`
	NomorSp2d       string `json:"nomorSp2d" binding:"required"`
	RealisasiPagu   int64  `json:"realisasiPagu" binding:"required"`
	RealisasiTarget string `json:"realisasiTarget" binding:"required"`
	Bulan           string `json:"bulan"`
	Keterangan      string `json:"keterangan" binding:"required"`
}

type UpdateRealisasiRequest struct {
	FokusBelanjaId  int    `json:"fokusBelanjaId"`
	Tanggal         string `json:"tanggal"`
	NomorSp2d       string `json:"nomorSp2d"`
	RealisasiPagu   int64  `json:"realisasiPagu"`
	RealisasiTarget string `json:"realisasiTarget"`
	Bulan           string `json:"bulan"`
	Keterangan      string `json:"keterangan"`
}
