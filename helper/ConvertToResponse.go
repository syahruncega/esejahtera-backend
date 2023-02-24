package helper

import (
	"kemiskinan/model"
	"kemiskinan/responses"
)

func ConvertToBidangUrusanResponse(b model.BidangUrusan) responses.BidangUrusanResponse {
	return responses.BidangUrusanResponse{
		Id:               b.Id,
		BidangUrusanId:   b.BidangUrusanId,
		NamaBidangUrusan: b.NamaBidangUrusan,
		CreatedAt:        b.CreatedAt,
		UpdatedAt:        b.UpdatedAt,
	}
}

func ConvertToInstansiResponse(i model.Instansi) responses.InstansiResponse {
	return responses.InstansiResponse{
		Id:           i.Id,
		InstansiId:   i.InstansiId,
		NamaInstansi: i.NamaInstansi,
		CreatedAt:    i.CreatedAt,
		UpdatedAt:    i.UpdatedAt,
	}
}

func ConvertToDetailInstansiResponse(d model.DetailInstansi) responses.DetailInstansiResponse {
	return responses.DetailInstansiResponse{
		Id:         d.Id,
		InstansiId: d.InstansiId,
		Instansi: responses.InstansiResponse{
			Id:           d.Instansi.Id,
			InstansiId:   d.Instansi.InstansiId,
			NamaInstansi: d.Instansi.NamaInstansi,
			CreatedAt:    d.Instansi.CreatedAt,
			UpdatedAt:    d.Instansi.UpdatedAt,
		},
		BidangUrusanId: d.BidangUrusanId,
		BidangUrusan: responses.BidangUrusanResponse{
			Id:               d.BidangUrusan.Id,
			BidangUrusanId:   d.BidangUrusan.BidangUrusanId,
			NamaBidangUrusan: d.BidangUrusan.NamaBidangUrusan,
			CreatedAt:        d.BidangUrusan.CreatedAt,
			UpdatedAt:        d.BidangUrusan.UpdatedAt,
		},
		CreatedAt: d.CreatedAt,
		UpdatedAt: d.UpdatedAt,
	}
}

func ConvertToProgramResponse(p model.Program) responses.ProgramResponse {
	return responses.ProgramResponse{
		Id:          p.Id,
		ProgramId:   p.ProgramId,
		NamaProgram: p.NamaProgram,
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   p.UpdatedAt,
	}
}

func ConvertToDetailProgramResponse(d model.DetailProgram) responses.DetailProgramResponse {
	return responses.DetailProgramResponse{
		Id:         d.Id,
		InstansiId: d.InstansiId,
		Instansi: responses.InstansiResponse{
			Id:           d.Instansi.Id,
			InstansiId:   d.Instansi.InstansiId,
			NamaInstansi: d.Instansi.NamaInstansi,
			CreatedAt:    d.Instansi.CreatedAt,
			UpdatedAt:    d.Instansi.UpdatedAt,
		},
		ProgramId: d.ProgramId,
		Program: responses.ProgramResponse{
			Id:          d.Program.Id,
			ProgramId:   d.Program.ProgramId,
			NamaProgram: d.Program.NamaProgram,
			CreatedAt:   d.Program.CreatedAt,
			UpdatedAt:   d.UpdatedAt,
		},
		PaguProgram: d.PaguProgram,
		CreatedAt:   d.CreatedAt,
		UpdatedAt:   d.UpdatedAt,
	}
}

func ConvertToIndikatorProgramResponse(p model.IndikatorProgram) responses.IndikatorProgramResponse {
	return responses.IndikatorProgramResponse{
		Id:         p.Id,
		InstansiId: p.InstansiId,
		Instansi: responses.InstansiResponse{
			Id:           p.Instansi.Id,
			InstansiId:   p.Instansi.InstansiId,
			NamaInstansi: p.Instansi.NamaInstansi,
			CreatedAt:    p.Instansi.CreatedAt,
			UpdatedAt:    p.Instansi.UpdatedAt,
		},
		ProgramId: p.ProgramId,
		Program: responses.ProgramResponse{
			Id:          p.Program.Id,
			ProgramId:   p.Program.ProgramId,
			NamaProgram: p.Program.NamaProgram,
			CreatedAt:   p.Program.CreatedAt,
			UpdatedAt:   p.Program.UpdatedAt,
		},
		Sasaran:                 p.Sasaran,
		IndikatorSasaran:        p.IndikatorSasaran,
		Kebijakan:               p.Kebijakan,
		IndikatorKinerjaProgram: p.IndikatorKinerjaProgram,
		PaguProgram:             p.PaguProgram,
		CreatedAt:               p.CreatedAt,
		UpdatedAt:               p.UpdatedAt,
	}
}

func ConvertToKegiatanResponse(k model.Kegiatan) responses.KegiatanResponse {
	return responses.KegiatanResponse{
		Id:           k.Id,
		KegiatanId:   k.KegiatanId,
		NamaKegiatan: k.NamaKegiatan,
		CreatedAt:    k.CreatedAt,
		UpdatedAt:    k.UpdatedAt,
	}
}

func ConvertToDetailKegiatanResponse(d model.DetailKegiatan) responses.DetailKegiatanResponse {
	return responses.DetailKegiatanResponse{
		Id:        d.Id,
		ProgramId: d.ProgramId,
		Program: responses.ProgramResponse{
			Id:          d.Program.Id,
			ProgramId:   d.Program.ProgramId,
			NamaProgram: d.Program.NamaProgram,
			CreatedAt:   d.Program.CreatedAt,
			UpdatedAt:   d.Program.UpdatedAt,
		},
		KegiatanId: d.KegiatanId,
		Kegiatan: responses.KegiatanResponse{
			Id:           d.Kegiatan.Id,
			KegiatanId:   d.Kegiatan.KegiatanId,
			NamaKegiatan: d.Kegiatan.NamaKegiatan,
			CreatedAt:    d.Kegiatan.CreatedAt,
			UpdatedAt:    d.Kegiatan.UpdatedAt,
		},
		PaguKegiatan: d.PaguKegiatan,
		CreatedAt:    d.CreatedAt,
		UpdatedAt:    d.UpdatedAt,
	}
}

func ConvertToIndikatorKegiatanResponse(i model.IndikatorKegiatan) responses.IndikatorKegiatanResponse {
	return responses.IndikatorKegiatanResponse{
		Id:        i.Id,
		ProgramId: i.ProgramId,
		Program: responses.ProgramResponse{
			Id:          i.Program.Id,
			ProgramId:   i.Program.ProgramId,
			NamaProgram: i.Program.NamaProgram,
			CreatedAt:   i.Program.CreatedAt,
			UpdatedAt:   i.Program.UpdatedAt,
		},
		KegiatanId: i.KegiatanId,
		Kegiatan: responses.KegiatanResponse{
			Id:           i.Kegiatan.Id,
			KegiatanId:   i.Kegiatan.KegiatanId,
			NamaKegiatan: i.Kegiatan.NamaKegiatan,
			CreatedAt:    i.Kegiatan.CreatedAt,
			UpdatedAt:    i.Kegiatan.UpdatedAt,
		},
		IndikatorKinerjaKegiatan: i.IndikatorKinerjaKegiatan,
		PaguKegiatan:             i.PaguKegiatan,
		CreatedAt:                i.CreatedAt,
		UpdatedAt:                i.UpdatedAt,
	}
}

func ConvertToSubKegiatanResponse(s model.SubKegiatan) responses.SubKegiatanResponse {
	return responses.SubKegiatanResponse{
		Id:              s.Id,
		SubKegiatanId:   s.SubKegiatanId,
		NamaSubKegiatan: s.NamaSubKegiatan,
		CreatedAt:       s.CreatedAt,
		UpdatedAt:       s.UpdatedAt,
	}
}

func ConvertToDetailSubKegiatanResponse(d model.DetailSubKegiatan) responses.DetailSubKegiatanResponse {
	return responses.DetailSubKegiatanResponse{
		Id:         d.Id,
		KegiatanId: d.KegiatanId,
		Kegiatan: responses.KegiatanResponse{
			Id:           d.Kegiatan.Id,
			KegiatanId:   d.Kegiatan.KegiatanId,
			NamaKegiatan: d.Kegiatan.NamaKegiatan,
			CreatedAt:    d.Kegiatan.CreatedAt,
			UpdatedAt:    d.Kegiatan.UpdatedAt,
		},
		SubKegiatanId: d.SubKegiatanId,
		SubKegiatan: responses.SubKegiatanResponse{
			Id:              d.SubKegiatan.Id,
			SubKegiatanId:   d.SubKegiatanId,
			NamaSubKegiatan: d.SubKegiatan.NamaSubKegiatan,
			CreatedAt:       d.SubKegiatan.CreatedAt,
			UpdatedAt:       d.SubKegiatan.UpdatedAt,
		},
		PaguSubKegiatan: d.PaguSubKegiatan,
		CreatedAt:       d.CreatedAt,
		UpdatedAt:       d.UpdatedAt,
	}
}

func ConvertToIndikatorSubKegiatanResponse(i model.IndikatorSubKegiatan) responses.IndikatorSubKegiatanResponse {
	return responses.IndikatorSubKegiatanResponse{
		Id:         i.Id,
		KegiatanId: i.KegiatanId,
		Kegiatan: responses.KegiatanResponse{
			Id:           i.Kegiatan.Id,
			KegiatanId:   i.Kegiatan.KegiatanId,
			NamaKegiatan: i.Kegiatan.NamaKegiatan,
			CreatedAt:    i.CreatedAt,
			UpdatedAt:    i.UpdatedAt,
		},
		SubKegiatanId: i.SubKegiatanId,
		SubKegiatan: responses.SubKegiatanResponse{
			Id:              i.SubKegiatan.Id,
			SubKegiatanId:   i.SubKegiatanId,
			NamaSubKegiatan: i.SubKegiatan.NamaSubKegiatan,
			CreatedAt:       i.SubKegiatan.CreatedAt,
			UpdatedAt:       i.SubKegiatan.UpdatedAt,
		},
		IndikatorKinerjaSubKegiatan: i.IndikatorKinerjaSubKegiatan,
		PaguSubKegiatan:             i.PaguSubKegiatan,
		CreatedAt:                   i.CreatedAt,
		UpdatedAt:                   i.UpdatedAt,
	}
}

func ConvertToFokusBelanjaResponse(f model.FokusBelanja) responses.FokusBelanjaResponse {
	return responses.FokusBelanjaResponse{
		Id:            f.Id,
		SubKegiatanId: f.SubKegiatanId,
		SubKegiatan: responses.SubKegiatanResponse{
			Id:              f.SubKegiatan.Id,
			SubKegiatanId:   f.SubKegiatanId,
			NamaSubKegiatan: f.SubKegiatan.NamaSubKegiatan,
			CreatedAt:       f.SubKegiatan.CreatedAt,
			UpdatedAt:       f.SubKegiatan.UpdatedAt,
		},
		FokusBelanja:     f.FokusBelanja,
		Indikator:        f.Indikator,
		Target:           f.Target,
		Satuan:           f.Satuan,
		PaguFokusBelanja: f.PaguFokusBelanja,
		Keterangan:       f.Keterangan,
		CreatedAt:        f.CreatedAt,
		UpdatedAt:        f.UpdatedAt,
	}
}

func ConvertToDetailLokasiResponse(d model.Detail_Lokasi) responses.Detail_LokasiResponse {
	return responses.Detail_LokasiResponse{
		Id:             d.Id,
		FokusBelanjaid: d.FokusBelanjaId,
		FokusBelanja: responses.FokusBelanjaResponse{
			Id:               d.FokusBelanja.Id,
			SubKegiatanId:    d.FokusBelanja.SubKegiatanId,
			FokusBelanja:     d.FokusBelanja.FokusBelanja,
			Indikator:        d.FokusBelanja.Indikator,
			Target:           d.FokusBelanja.Target,
			Satuan:           d.FokusBelanja.Satuan,
			PaguFokusBelanja: d.FokusBelanja.PaguFokusBelanja,
			Keterangan:       d.FokusBelanja.Keterangan,
			CreatedAt:        d.FokusBelanja.CreatedAt,
			UpdatedAt:        d.FokusBelanja.UpdatedAt,
		},
		KabupatenKotaId: d.KabupatenKotaId,
		KabupatenKota: responses.Kabupaten_KotaResponse{
			Id:         d.KabupatenKota.Id,
			ProvinsiId: d.KabupatenKota.ProvinsiId,
			Nama:       d.KabupatenKota.Nama,
		},
		KecamatanId: d.KecamatanId,
		Kecamatan: responses.KecamatanResponse{
			Id:              d.Kecamatan.Id,
			KabupatenKotaId: d.Kecamatan.KabupatenKotaId,
			Nama:            d.Kecamatan.Nama,
		},
		KelurahanId: d.KelurahanId,
		Kelurahan: responses.KelurahanResponse{
			Id:          d.Kelurahan.Id,
			KecamatanId: d.Kelurahan.KecamatanId,
			Nama:        d.Kelurahan.Nama,
		},
		CreatedAt: d.CreatedAt,
		UpdatedAt: d.UpdatedAt,
	}
}

func ConvertToProvinsiResponse(p model.Provinsi) responses.ProvinsiResponse {
	return responses.ProvinsiResponse{
		Id:   p.Id,
		Nama: p.Nama,
	}
}

func ConvertToKabupatenKotaResponse(k model.Kabupaten_Kota) responses.Kabupaten_KotaResponse {
	return responses.Kabupaten_KotaResponse{
		Id:         k.Id,
		ProvinsiId: k.ProvinsiId,
		Nama:       k.Nama,
	}
}

func ConvertToKecamatanResponse(k model.Kecamatan) responses.KecamatanResponse {
	return responses.KecamatanResponse{
		Id:              k.Id,
		KabupatenKotaId: k.KabupatenKotaId,
		Nama:            k.Nama,
	}
}

func ConvertToKelurahanResponse(k model.Kelurahan) responses.KelurahanResponse {
	return responses.KelurahanResponse{
		Id:          k.Id,
		KecamatanId: k.KecamatanId,
		Nama:        k.Nama,
	}
}

func ConvertToKeluargaResponse(k model.Keluarga) responses.KeluargaResponse {
	return responses.KeluargaResponse{
		Id:         k.Id,
		IdKeluarga: k.IdKeluarga,
		ProvinsiId: k.ProvinsiId,
		Provinsi: responses.ProvinsiResponse{
			Id:   k.Provinsi.Id,
			Nama: k.Provinsi.Nama,
		},

		KabupatenKotaId: k.KabupatenKotaId,
		KabupatenKota: responses.Kabupaten_KotaResponse{
			Id:         k.KabupatenKota.Id,
			ProvinsiId: k.KabupatenKota.ProvinsiId,
			Nama:       k.KabupatenKota.Nama,
		},

		KecamatanId: k.KecamatanId,
		Kecamatan: responses.KecamatanResponse{
			Id:              k.Kecamatan.Id,
			KabupatenKotaId: k.Kecamatan.KabupatenKotaId,
			Nama:            k.Kecamatan.Nama,
		},

		KelurahanId: k.KelurahanId,
		Kelurahan: responses.KelurahanResponse{
			Id:          k.Kelurahan.Id,
			KecamatanId: k.Kelurahan.KecamatanId,
			Nama:        k.Kelurahan.Nama,
		},

		DesilKesejahteraan:     k.DesilKesejahteraan,
		Alamat:                 k.Alamat,
		KepalaKeluarga:         k.KepalaKeluarga,
		Nik:                    k.Nik,
		PadanDukcapil:          k.PadanDukcapil,
		JenisKelamin:           k.JenisKelamin,
		TanggalLahir:           k.TanggalLahir,
		Pekerjaan:              k.Pekerjaan,
		Pendidikan:             k.Pendidikan,
		KepemilikanRumah:       k.KepemilikanRumah,
		Simpanan:               k.Simpanan,
		JenisAtap:              k.JenisAtap,
		JenisDinding:           k.JenisDinding,
		JenisLantai:            k.JenisLantai,
		SumberPenerangan:       k.SumberPenerangan,
		BahanBakarMemasak:      k.BahanBakarMemasak,
		SumberAirMinum:         k.SumberAirMinum,
		FasilitasBuangAirBesar: k.FasilitasBuangAirBesar,
		PenerimaBPNT:           k.PenerimaBPNT,
		PenerimaBPUM:           k.PenerimaBPUM,
		PenerimaBST:            k.PenerimaBST,
		PenerimaPKH:            k.PenerimaPKH,
		PenerimaSembako:        k.PenerimaSembako,

		UserId: k.UserId,
		User: responses.UserResponse{
			Id:        k.User.Id,
			Username:  k.User.Username,
			Password:  k.User.Password,
			Email:     k.User.Email,
			NoHp:      k.User.NoHp,
			Role:      k.User.Role,
			CreatedAt: k.User.CreatedAt,
			UpdatedAt: k.User.UpdatedAt,
		},

		MahasiswaId: k.MahasiswaId,
		Mahasiswa: responses.MahasiswaResponse{
			Id:           k.Mahasiswa.Id,
			UserId:       k.Mahasiswa.UserId,
			NamaLengkap:  k.Mahasiswa.NamaLengkap,
			Universitas:  k.Mahasiswa.Universitas,
			JenisKelamin: k.Mahasiswa.JenisKelamin,
			TanggalLahir: k.Mahasiswa.TanggalLahir,
			KelurahanId:  k.Mahasiswa.KelurahanId,
			CreatedAt:    k.Mahasiswa.CreatedAt,
			UpdatedAt:    k.Mahasiswa.UpdatedAt,
		},

		StatusVerifikasi: k.StatusVerifikasi,
	}
}

func ConvertToMonevResponse(m model.Monev) responses.MonevResponse {
	return responses.MonevResponse{
		Id:              m.Id,
		NamaPenerima:    m.NamaPenerima,
		KabupatenKotaId: m.KabupatenKotaId,
		KabupatenKota: responses.Kabupaten_KotaResponse{
			Id:         m.KabupatenKota.Id,
			ProvinsiId: m.KabupatenKota.ProvinsiId,
			Nama:       m.KabupatenKota.Nama,
		},
		KecamatanId: m.KecamatanId,
		Kecamatan: responses.KecamatanResponse{
			Id:              m.Kecamatan.Id,
			KabupatenKotaId: m.Kecamatan.KabupatenKotaId,
			Nama:            m.Kecamatan.Nama,
		},
		KelurahanId: m.KelurahanId,
		Kelurahan: responses.KelurahanResponse{
			Id:          m.Kelurahan.Id,
			KecamatanId: m.Kelurahan.KecamatanId,
			Nama:        m.Kelurahan.Nama,
		},
		JenisBantuan:  m.JenisBantuan,
		VolumeBantuan: m.VolumeBantuan,
		SatuanVolume:  m.SatuanVolume,
	}
}

func ConvertToJumlahResponse(jumlah int64) responses.JumlahResponse {
	return responses.JumlahResponse{
		Jumlah: jumlah,
	}
}

func ConvertToUserResponse(u model.User) responses.UserResponse {
	return responses.UserResponse{
		Id:        u.Id,
		Username:  u.Username,
		Password:  u.Password,
		Email:     u.Email,
		NoHp:      u.NoHp,
		Role:      u.Role,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func ConvertToAdminResponse(a model.Admin) responses.AdminResponse {
	return responses.AdminResponse{
		Id:            a.Id,
		UserId:        a.UserId,
		NamaLengkap:   a.NamaLengkap,
		UrlFotoProfil: a.UrlFotoProfil,
		CreatedAt:     a.CreatedAt,
		UpdatedAt:     a.UpdatedAt,
	}
}

func ConvertToAdminWithRelationResponse(a model.Admin) responses.AdminWithRelationResponse {
	return responses.AdminWithRelationResponse{
		Id:     a.Id,
		UserId: a.UserId,
		User: responses.UserResponse{
			Id:        a.User.Id,
			Username:  a.User.Username,
			Password:  a.User.Password,
			Email:     a.User.Email,
			NoHp:      a.User.NoHp,
			Role:      a.User.Role,
			CreatedAt: a.User.CreatedAt,
			UpdatedAt: a.User.UpdatedAt,
		},
		NamaLengkap:   a.NamaLengkap,
		UrlFotoProfil: a.UrlFotoProfil,
		CreatedAt:     a.CreatedAt,
		UpdatedAt:     a.UpdatedAt,
	}
}

func ConvertToAnalisResponse(a model.Analis) responses.AnalisResponse {
	return responses.AnalisResponse{
		Id:            a.Id,
		UserId:        a.UserId,
		NamaLengkap:   a.NamaLengkap,
		Universitas:   a.Universitas,
		UrlFotoProfil: a.UrlFotoProfil,
		CreatedAt:     a.CreatedAt,
		UpdatedAt:     a.UpdatedAt,
	}
}

func ConvertToAnalisWithRelationResponse(a model.Analis) responses.AnalisWithRelationResponse {
	return responses.AnalisWithRelationResponse{
		Id:     a.Id,
		UserId: a.UserId,
		User: responses.UserResponse{
			Id:        a.User.Id,
			Username:  a.User.Username,
			Password:  a.User.Password,
			Email:     a.User.Email,
			NoHp:      a.User.NoHp,
			Role:      a.User.Role,
			CreatedAt: a.User.CreatedAt,
			UpdatedAt: a.User.UpdatedAt,
		},
		NamaLengkap:   a.NamaLengkap,
		Universitas:   a.Universitas,
		UrlFotoProfil: a.UrlFotoProfil,
		CreatedAt:     a.CreatedAt,
		UpdatedAt:     a.UpdatedAt,
	}
}

func ConvertToPusbangResponse(p model.Pusbang) responses.PusbangResponse {
	return responses.PusbangResponse{
		Id:            p.Id,
		UserId:        p.UserId,
		NamaLengkap:   p.NamaLengkap,
		Universitas:   p.Universitas,
		UrlFotoProfil: p.UrlFotoProfil,
		CreatedAt:     p.CreatedAt,
		UpdatedAt:     p.UpdatedAt,
	}
}

func ConvertToPusbangWithRelationResponse(p model.Pusbang) responses.PusbangWithRelationResponse {
	return responses.PusbangWithRelationResponse{
		Id:     p.Id,
		UserId: p.UserId,
		User: responses.UserResponse{
			Id:        p.User.Id,
			Username:  p.User.Username,
			Password:  p.User.Password,
			Email:     p.User.Email,
			NoHp:      p.User.NoHp,
			Role:      p.User.Role,
			CreatedAt: p.User.CreatedAt,
			UpdatedAt: p.User.UpdatedAt,
		},
		NamaLengkap:   p.NamaLengkap,
		Universitas:   p.Universitas,
		UrlFotoProfil: p.UrlFotoProfil,
		CreatedAt:     p.CreatedAt,
		UpdatedAt:     p.UpdatedAt,
	}
}

func ConvertToDosenResponse(d model.Dosen) responses.DosenResponse {
	return responses.DosenResponse{
		Id:            d.Id,
		UserId:        d.UserId,
		NamaLengkap:   d.NamaLengkap,
		Universitas:   d.Universitas,
		UrlFotoProfil: d.UrlFotoProfil,
		CreatedAt:     d.CreatedAt,
		UpdatedAt:     d.UpdatedAt,
	}
}

func ConvertToDosenWithRelationResponse(d model.Dosen) responses.DosenWithRelationResponse {
	return responses.DosenWithRelationResponse{
		Id:     d.Id,
		UserId: d.UserId,
		User: responses.UserResponse{
			Id:        d.User.Id,
			Username:  d.User.Username,
			Password:  d.User.Password,
			Email:     d.User.Email,
			NoHp:      d.User.NoHp,
			Role:      d.User.Role,
			CreatedAt: d.User.CreatedAt,
			UpdatedAt: d.User.UpdatedAt,
		},
		NamaLengkap:   d.NamaLengkap,
		Universitas:   d.Universitas,
		UrlFotoProfil: d.UrlFotoProfil,
		CreatedAt:     d.CreatedAt,
		UpdatedAt:     d.UpdatedAt,
	}
}

func ConvertToMahasiswaResponse(m model.Mahasiswa) responses.MahasiswaResponse {
	return responses.MahasiswaResponse{
		Id:              m.Id,
		UserId:          m.UserId,
		NamaLengkap:     m.NamaLengkap,
		Universitas:     m.Universitas,
		JenisKelamin:    m.JenisKelamin,
		TanggalLahir:    m.TanggalLahir,
		KabupatenKotaId: m.KabupatenKotaId,
		KecamatanId:     m.KecamatanId,
		KelurahanId:     m.KelurahanId,
		UrlFotoProfil:   m.UrlFotoProfil,
		CreatedAt:       m.CreatedAt,
		UpdatedAt:       m.UpdatedAt,
	}
}

func ConvertToMahasiswaWithRelationResponse(m model.Mahasiswa) responses.MahasiswaWithRelationResponse {
	return responses.MahasiswaWithRelationResponse{
		Id:     m.Id,
		UserId: m.UserId,
		User: responses.UserResponse{
			Id:        m.User.Id,
			Username:  m.User.Username,
			Password:  m.User.Password,
			Email:     m.User.Email,
			NoHp:      m.User.NoHp,
			Role:      m.User.Role,
			CreatedAt: m.User.CreatedAt,
			UpdatedAt: m.User.UpdatedAt,
		},
		NamaLengkap:     m.NamaLengkap,
		Universitas:     m.Universitas,
		JenisKelamin:    m.JenisKelamin,
		TanggalLahir:    m.TanggalLahir,
		KabupatenKotaId: m.KabupatenKotaId,
		KabupatenKota: responses.Kabupaten_KotaResponse{
			Id:         m.KabupatenKota.Id,
			ProvinsiId: m.KabupatenKota.ProvinsiId,
			Nama:       m.KabupatenKota.Nama,
		},
		KecamatanId: m.KecamatanId,
		Kecamatan: responses.KecamatanResponse{
			Id:              m.Kecamatan.Id,
			KabupatenKotaId: m.Kecamatan.KabupatenKotaId,
			Nama:            m.KabupatenKota.Nama,
		},
		KelurahanId: m.KelurahanId,
		Kelurahan: responses.KelurahanResponse{
			Id:          m.Kelurahan.Id,
			KecamatanId: m.Kelurahan.KecamatanId,
			Nama:        m.Kelurahan.Nama,
		},
		UrlFotoProfil: m.UrlFotoProfil,
		CreatedAt:     m.CreatedAt,
		UpdatedAt:     m.UpdatedAt,
	}
}

func ConvertToMahasiswaWithVerifiedCountResponse(m model.Mahasiswa, jumlahVerifiedIndividu int64, jumlahVerifiedKeluarga int64) responses.MahasiswaWithVerifiedCountResponse {
	return responses.MahasiswaWithVerifiedCountResponse{
		Id:     m.Id,
		UserId: m.UserId,
		User: responses.UserResponse{
			Id:        m.User.Id,
			Username:  m.User.Username,
			Password:  m.User.Password,
			Email:     m.User.Email,
			NoHp:      m.User.NoHp,
			Role:      m.User.Role,
			CreatedAt: m.User.CreatedAt,
			UpdatedAt: m.User.UpdatedAt,
		},
		NamaLengkap:     m.NamaLengkap,
		Universitas:     m.Universitas,
		JenisKelamin:    m.JenisKelamin,
		TanggalLahir:    m.TanggalLahir,
		KabupatenKotaId: m.KabupatenKotaId,
		KabupatenKota: responses.Kabupaten_KotaResponse{
			Id:         m.KabupatenKota.Id,
			ProvinsiId: m.KabupatenKota.ProvinsiId,
			Nama:       m.KabupatenKota.Nama,
		},
		KecamatanId: m.KecamatanId,
		Kecamatan: responses.KecamatanResponse{
			Id:              m.Kecamatan.Id,
			KabupatenKotaId: m.Kecamatan.KabupatenKotaId,
			Nama:            m.KabupatenKota.Nama,
		},
		KelurahanId: m.KelurahanId,
		Kelurahan: responses.KelurahanResponse{
			Id:          m.Kelurahan.Id,
			KecamatanId: m.Kelurahan.KecamatanId,
			Nama:        m.Kelurahan.Nama,
		},
		JumlahIndividuVerified: jumlahVerifiedIndividu,
		JumlahKeluargaVerified: jumlahVerifiedKeluarga,
		UrlFotoProfil:          m.UrlFotoProfil,
		CreatedAt:              m.CreatedAt,
		UpdatedAt:              m.UpdatedAt,
	}
}

func ConvertToLokasiDosenResponse(l model.LokasiDosen) responses.LokasiDosenResponse {
	return responses.LokasiDosenResponse{
		Id:              l.Id,
		DosenId:         l.DosenId,
		KabupatenKotaId: l.KabupatenKotaId,
		KecamatanId:     l.KecamatanId,
		KelurahanId:     l.KelurahanId,
		CreatedAt:       l.CreatedAt,
		UpdatedAt:       l.UpdatedAt,
	}
}

func ConvertToLokasiDosenWithRelationResponse(l model.LokasiDosen) responses.LokasiDosenWithRelationResponse {
	return responses.LokasiDosenWithRelationResponse{
		Id:      l.Id,
		DosenId: l.DosenId,
		Dosen: responses.DosenResponse{
			Id:          l.Dosen.Id,
			UserId:      l.Dosen.UserId,
			NamaLengkap: l.Dosen.NamaLengkap,
			Universitas: l.Dosen.Universitas,
			CreatedAt:   l.Dosen.CreatedAt,
			UpdatedAt:   l.Dosen.UpdatedAt,
		},
		KabupatenKotaId: l.KabupatenKotaId,
		KabupatenKota: responses.Kabupaten_KotaResponse{
			Id:         l.KabupatenKota.Id,
			ProvinsiId: l.KabupatenKota.ProvinsiId,
			Nama:       l.KabupatenKota.Nama,
		},
		KecamatanId: l.KecamatanId,
		Kecamatan: responses.KecamatanResponse{
			Id:              l.Kecamatan.Id,
			KabupatenKotaId: l.Kecamatan.KabupatenKotaId,
			Nama:            l.Kecamatan.Nama,
		},
		KelurahanId: l.KelurahanId,
		Kelurahan: responses.KelurahanResponse{
			Id:          l.Kelurahan.Id,
			KecamatanId: l.Kelurahan.KecamatanId,
			Nama:        l.Kelurahan.Nama,
		},
		CreatedAt: l.CreatedAt,
		UpdatedAt: l.UpdatedAt,
	}
}

func ConvertToKeluargaVerifikasiResponse(k model.KeluargaVerifikasi) responses.KeluargaVerifikasiResponse {
	return responses.KeluargaVerifikasiResponse{
		Id:         k.Id,
		IdKeluarga: k.IdKeluarga,
		ProvinsiId: k.ProvinsiId,
		Provinsi: responses.ProvinsiResponse{
			Id:   k.Provinsi.Id,
			Nama: k.Provinsi.Nama,
		},

		KabupatenKotaId: k.KabupatenKotaId,
		KabupatenKota: responses.Kabupaten_KotaResponse{
			Id:         k.KabupatenKota.Id,
			ProvinsiId: k.KabupatenKota.ProvinsiId,
			Nama:       k.KabupatenKota.Nama,
		},

		KecamatanId: k.KecamatanId,
		Kecamatan: responses.KecamatanResponse{
			Id:              k.Kecamatan.Id,
			KabupatenKotaId: k.Kecamatan.KabupatenKotaId,
			Nama:            k.Kecamatan.Nama,
		},

		KelurahanId:            k.KelurahanId,
		DesilKesejahteraan:     k.DesilKesejahteraan,
		Alamat:                 k.Alamat,
		KepalaKeluarga:         k.KepalaKeluarga,
		Nik:                    k.Nik,
		PadanDukcapil:          k.PadanDukcapil,
		JenisKelamin:           k.JenisKelamin,
		TanggalLahir:           k.TanggalLahir,
		Pekerjaan:              k.Pekerjaan,
		Pendidikan:             k.Pendidikan,
		KepemilikanRumah:       k.KepemilikanRumah,
		Simpanan:               k.Simpanan,
		JenisAtap:              k.JenisAtap,
		JenisDinding:           k.JenisDinding,
		JenisLantai:            k.JenisLantai,
		SumberPenerangan:       k.SumberPenerangan,
		BahanBakarMemasak:      k.BahanBakarMemasak,
		SumberAirMinum:         k.SumberAirMinum,
		FasilitasBuangAirBesar: k.FasilitasBuangAirBesar,
		PenerimaBPNT:           k.PenerimaBPNT,
		PenerimaBPUM:           k.PenerimaBPUM,
		PenerimaBST:            k.PenerimaBST,
		PenerimaPKH:            k.PenerimaPKH,
		PenerimaSembako:        k.PenerimaSembako,
		PenerimaLainnya:        k.PenerimaLainnya,
		StatusResponden:        k.StatusResponden,

		UserId: k.UserId,
		User: responses.UserResponse{
			Id:        k.User.Id,
			Username:  k.User.Username,
			Password:  k.User.Password,
			Email:     k.User.Email,
			NoHp:      k.User.NoHp,
			Role:      k.User.Role,
			CreatedAt: k.User.CreatedAt,
			UpdatedAt: k.User.UpdatedAt,
		},

		MahasiswaId: k.MahasiswaId,
		Mahasiswa: responses.MahasiswaResponse{
			Id:           k.Mahasiswa.Id,
			UserId:       k.UserId,
			NamaLengkap:  k.Mahasiswa.NamaLengkap,
			Universitas:  k.Mahasiswa.Universitas,
			JenisKelamin: k.JenisKelamin,
			TanggalLahir: k.TanggalLahir,
			KelurahanId:  k.KelurahanId,
		},

		Kelurahan: responses.KelurahanResponse{
			Id:          k.Kelurahan.Id,
			KecamatanId: k.Kelurahan.KecamatanId,
			Nama:        k.Kelurahan.Nama,
		},

		UrlBukti: k.UrlBukti,

		CreatedAt: k.CreatedAt,
		UpdatedAt: k.UpdatedAt,
	}
}

func ConvertToIndividuResponse(i model.Individu) responses.IndividuResponse {
	return responses.IndividuResponse{
		Id:         i.Id,
		IdKeluarga: i.IdKeluarga,

		ProvinsiId: i.ProvinsiId,
		Provinsi: responses.ProvinsiResponse{
			Id:   i.Provinsi.Id,
			Nama: i.Provinsi.Nama,
		},

		KabupatenKotaId: i.KabupatenKotaId,
		Kabupaten_Kota: responses.Kabupaten_KotaResponse{
			Id:         i.KabupatenKota.Id,
			ProvinsiId: i.KabupatenKota.ProvinsiId,
			Nama:       i.KabupatenKota.Nama,
		},

		KecamatanId: i.KecamatanId,
		Kecamatan: responses.KecamatanResponse{
			Id:              i.Kecamatan.Id,
			KabupatenKotaId: i.Kecamatan.KabupatenKotaId,
			Nama:            i.Kecamatan.Nama,
		},

		KelurahanId: i.KelurahanId,
		Kelurahan: responses.KelurahanResponse{
			Id:          i.Kelurahan.Id,
			KecamatanId: i.Kelurahan.KecamatanId,
			Nama:        i.Kelurahan.Nama,
		},

		DesilKesejahteraan: i.DesilKesejahteraan,
		Alamat:             i.Alamat,
		IdIndividu:         i.IdIndividu,
		Nama:               i.Nama,
		Nik:                i.Nik,
		PadanDukcapil:      i.PadanDukcapil,
		JenisKelamin:       i.JenisKelamin,
		Hubungan:           i.Hubungan,
		TanggalLahir:       i.TanggalLahir,
		StatusKawin:        i.StatusKawin,
		Pekerjaan:          i.Pekerjaan,
		Pendidikan:         i.Pendidikan,
		KategoriUsia:       i.KategoriUsia,
		PenerimaBPNT:       i.PenerimaBPNT,
		PenerimaBPUM:       i.PenerimaBPUM,
		PenerimaBST:        i.PenerimaBST,
		PenerimaPKH:        i.PenerimaPKH,
		PenerimaSembako:    i.PenerimaSembako,

		UserId: i.UserId,
		User: responses.UserResponse{
			Id:        i.User.Id,
			Username:  i.User.Username,
			Password:  i.User.Password,
			Email:     i.User.Email,
			NoHp:      i.User.NoHp,
			Role:      i.User.Role,
			CreatedAt: i.User.CreatedAt,
			UpdatedAt: i.Mahasiswa.UpdatedAt,
		},

		MahasiswaId: i.MahasiswaId,
		Mahasiswa: responses.MahasiswaResponse{
			Id:              i.Mahasiswa.Id,
			UserId:          i.Mahasiswa.UserId,
			NamaLengkap:     i.Mahasiswa.NamaLengkap,
			Universitas:     i.Mahasiswa.Universitas,
			JenisKelamin:    i.Mahasiswa.JenisKelamin,
			TanggalLahir:    i.Mahasiswa.TanggalLahir,
			KabupatenKotaId: i.Mahasiswa.KabupatenKotaId,
			KecamatanId:     i.Mahasiswa.KecamatanId,
			KelurahanId:     i.Mahasiswa.KelurahanId,
			CreatedAt:       i.Mahasiswa.CreatedAt,
			UpdatedAt:       i.Mahasiswa.UpdatedAt,
		},

		StatusVerifikasi: i.StatusVerifikasi,
	}
}

func ConvertToIndividuVerifikasiResponse(i model.IndividuVerifikasi) responses.IndividuVerifikasiResponse {
	return responses.IndividuVerifikasiResponse{
		Id:         i.Id,
		IdKeluarga: i.IdKeluarga,

		ProvinsiId: i.ProvinsiId,
		Provinsi: responses.ProvinsiResponse{
			Id:   i.Provinsi.Id,
			Nama: i.Provinsi.Nama,
		},

		KabupatenKotaId: i.KabupatenKotaId,
		KabupatenKota: responses.Kabupaten_KotaResponse{
			Id:         i.KabupatenKota.Id,
			ProvinsiId: i.KabupatenKota.ProvinsiId,
			Nama:       i.KabupatenKota.Nama,
		},

		KecamatanId: i.KecamatanId,
		Kecamatan: responses.KecamatanResponse{
			Id:              i.Kecamatan.Id,
			KabupatenKotaId: i.Kecamatan.KabupatenKotaId,
			Nama:            i.Kecamatan.Nama,
		},

		KelurahanId: i.KelurahanId,
		Kelurahan: responses.KelurahanResponse{
			Id:          i.Kelurahan.Id,
			KecamatanId: i.Kelurahan.KecamatanId,
			Nama:        i.Kelurahan.Nama,
		},

		DesilKesejahteraan: i.DesilKesejahteraan,
		Alamat:             i.Alamat,
		IdIndividu:         i.IdIndividu,
		Nama:               i.Nama,
		Nik:                i.Nik,
		PadanDukcapil:      i.PadanDukcapil,
		JenisKelamin:       i.JenisKelamin,
		Hubungan:           i.Hubungan,
		TanggalLahir:       i.TanggalLahir,
		StatusKawin:        i.StatusKawin,
		Pekerjaan:          i.Pekerjaan,
		Pendidikan:         i.Pendidikan,
		KategoriUsia:       i.KategoriUsia,
		PenerimaBPNT:       i.PenerimaBPNT,
		PenerimaBPUM:       i.PenerimaBPUM,
		PenerimaBST:        i.PenerimaBST,
		PenerimaPKH:        i.PenerimaPKH,
		PenerimaSembako:    i.PenerimaSembako,
		PenerimaLainnya:    i.PenerimaLainnya,
		StatusResponden:    i.StatusResponden,

		UserId: i.UserId,
		User: responses.UserResponse{
			Id:        i.User.Id,
			Username:  i.User.Username,
			Password:  i.User.Password,
			Email:     i.User.Email,
			NoHp:      i.User.NoHp,
			Role:      i.User.Role,
			CreatedAt: i.User.CreatedAt,
			UpdatedAt: i.User.UpdatedAt,
		},

		MahasiswaId: i.MahasiswaId,
		Mahasiswa: responses.MahasiswaResponse{
			Id:              i.Mahasiswa.Id,
			UserId:          i.Mahasiswa.UserId,
			NamaLengkap:     i.Mahasiswa.NamaLengkap,
			Universitas:     i.Mahasiswa.Universitas,
			JenisKelamin:    i.Mahasiswa.JenisKelamin,
			TanggalLahir:    i.Mahasiswa.TanggalLahir,
			KabupatenKotaId: i.Mahasiswa.KabupatenKotaId,
			KecamatanId:     i.Mahasiswa.KecamatanId,
			KelurahanId:     i.Mahasiswa.KelurahanId,
			CreatedAt:       i.Mahasiswa.CreatedAt,
			UpdatedAt:       i.Mahasiswa.UpdatedAt,
		},
		UrlBukti: i.UrlBukti,

		CreatedAt: i.CreatedAt,
		UpdatedAt: i.UpdatedAt,
	}
}
