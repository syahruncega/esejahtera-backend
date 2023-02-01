package request

type CreateIndividuVerifikasiRequest struct {
	Id                 int    `json:"id" binding:"required"`
	IdKeluarga         string `json:"idKeluarga" binding:"required"`
	ProvinsiId         string `json:"provinsiId" binding:"required"`
	KabupatenKotaId    string `json:"kabupatenKotaId" binding:"required"`
	KecamatanId        string `json:"kecamatanId" binding:"required"`
	KelurahanId        string `json:"kelurahanId" binding:"required"`
	DesilKesejahteraan string `json:"desilKesejahteraan" binding:"required"`
	Alamat             string `json:"alamat" binding:"required"`
	IdIndividu         string `json:"idIndividu" binding:"required"`
	Nama               string `json:"nama" binding:"required"`
	Nik                string `json:"nik" binding:"required"`
	PadanDukcapil      string `json:"padanDukcapil" binding:"required"`
	JenisKelamin       string `json:"jenisKelamin" binding:"required"`
	Hubungan           string `json:"hubungan" binding:"required"`
	TanggalLahir       string `json:"tanggalLahir" binding:"required"`
	StatusKawin        string `json:"statusKawin" binding:"required"`
	Pekerjaan          string `json:"pekerjaan" binding:"required"`
	Pendidikan         string `json:"pendidikan" binding:"required"`
	KategoriUsia       string `json:"kategoriUsia" binding:"required"`
	PenerimaBPNT       string `json:"penerimaBPNT" binding:"required"`
	PenerimaBPUM       string `json:"penerimaBPUM" binding:"required"`
	PenerimaBST        string `json:"penerimaBST" binding:"required"`
	PenerimaPKH        string `json:"penerimaPKH" binding:"required"`
	PenerimaSembako    string `json:"penerimaSembako" binding:"required"`
	PenerimaLainnya    string `json:"penerimaLainnya" binding:"required"`
	StatusResponden    string `json:"statusResponden" binding:"required"`
	UserId             int    `json:"userId" binding:"required"`
	MahasiswaId        int    `json:"mahasiswaId" binding:"required"`
}

type UpdateIndividuVerifikasiRequest struct {
	Id                 int    `json:"id"`
	IdKeluarga         string `json:"idKeluarga"`
	ProvinsiId         string `json:"provinsiId"`
	KabupatenKotaId    string `json:"kabupatenKotaId"`
	KecamatanId        string `json:"kecamatanId"`
	KelurahanId        string `json:"kelurahanId"`
	DesilKesejahteraan string `json:"desilKesejahteraan"`
	Alamat             string `json:"alamat"`
	IdIndividu         string `json:"idIndividu"`
	Nama               string `json:"nama"`
	Nik                string `json:"nik"`
	PadanDukcapil      string `json:"padanDukcapil"`
	JenisKelamin       string `json:"jenisKelamin"`
	Hubungan           string `json:"hubungan"`
	TanggalLahir       string `json:"tanggalLahir"`
	StatusKawin        string `json:"statusKawin"`
	Pekerjaan          string `json:"pekerjaan"`
	Pendidikan         string `json:"pendidikan"`
	KategoriUsia       string `json:"kategoriUsia"`
	PenerimaBPNT       string `json:"penerimaBPNT"`
	PenerimaBPUM       string `json:"penerimaBPUM"`
	PenerimaBST        string `json:"penerimaBST"`
	PenerimaPKH        string `json:"penerimaPKH"`
	PenerimaSembako    string `json:"penerimaSembako"`
	PenerimaLainnya    string `json:"penerimaLainnya"`
	StatusResponden    string `json:"statusResponden"`
	UserId             int    `json:"userId"`
	MahasiswaId        int    `json:"mahasiswaId"`
}
