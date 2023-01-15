package helper

import (
	"kemiskinan/model"
	"kemiskinan/responses"
)

func ConvertToBidangUrusanResponse(b model.BidangUrusan) responses.BidangUrusanResponse {
	return responses.BidangUrusanResponse{
		Id:               b.Id,
		NamaBidangUrusan: b.NamaBidangUrusan,
		CreatedAt:        b.CreatedAt,
		UpdatedAt:        b.UpdatedAt,
	}
}

func ConvertToInstansiResponse(i model.Instansi) responses.InstansiResponse {
	return responses.InstansiResponse{
		Id:             i.Id,
		BidangUrusanId: i.BidangUrusanId,
		NamaInstansi:   i.NamaInstansi,
		CreatedAt:      i.CreatedAt,
		UpdatedAt:      i.UpdatedAt,
	}
}

func ConvertToInstansiWithBidangUrusanResponse(i model.Instansi) responses.InstansiWithBidangUrusanResponse {
	return responses.InstansiWithBidangUrusanResponse{
		Id:             i.Id,
		NamaInstansi:   i.NamaInstansi,
		CreatedAt:      i.CreatedAt,
		UpdatedAt:      i.UpdatedAt,
		BidangUrusanId: i.BidangUrusanId,
		BidangUrusan: responses.BidangUrusanResponse{
			Id:               i.BidangUrusan.Id,
			NamaBidangUrusan: i.BidangUrusan.NamaBidangUrusan,
			CreatedAt:        i.BidangUrusan.CreatedAt,
			UpdatedAt:        i.BidangUrusan.UpdatedAt,
		},
	}
}

func ConvertToProgramResponse(p model.Program) responses.ProgramResponse {
	return responses.ProgramResponse{
		Id:                      p.Id,
		Sasaran:                 p.Sasaran,
		IndikatorSasaran:        p.IndikatorSasaran,
		Kebijakan:               p.Kebijakan,
		NamaProgram:             p.NamaProgram,
		IndikatorKinerjaProgram: p.IndikatorKinerjaProgram,
		PaguProgram:             p.PaguProgram,
		InstansiId:              p.InstansiId,
		CreatedAt:               p.CreatedAt,
		UpdatedAt:               p.UpdatedAt,
	}
}

func ConvertToProgramWithInstansiResponse(p model.Program) responses.ProgramWithInstansiResponse {
	return responses.ProgramWithInstansiResponse{
		Id:                      p.Id,
		Sasaran:                 p.Sasaran,
		IndikatorSasaran:        p.IndikatorSasaran,
		Kebijakan:               p.Kebijakan,
		NamaProgram:             p.NamaProgram,
		IndikatorKinerjaProgram: p.IndikatorKinerjaProgram,
		PaguProgram:             p.PaguProgram,
		InstansiId:              p.InstansiId,
		CreatedAt:               p.CreatedAt,
		UpdatedAt:               p.UpdatedAt,
		Instansi: responses.InstansiResponse{
			Id:           p.Instansi.Id,
			NamaInstansi: p.Instansi.NamaInstansi,
			CreatedAt:    p.Instansi.CreatedAt,
			UpdatedAt:    p.Instansi.UpdatedAt,
		},
	}
}

func ConvertToKegiatanResponse(k model.Kegiatan) responses.KegiatanResponse {
	return responses.KegiatanResponse{
		Id:                       k.Id,
		NamaKegiatan:             k.NamaKegiatan,
		IndikatorKinerjaKegiatan: k.IndikatorKinerjaKegiatan,
		PaguKegiatan:             k.PaguKegiatan,
		ProgramId:                k.ProgramId,
		CreatedAt:                k.CreatedAt,
		UpdatedAt:                k.UpdatedAt,
	}
}

func ConvertToKegiatanWithProgramResponse(k model.Kegiatan) responses.KegiatanWithProgramResponse {
	return responses.KegiatanWithProgramResponse{
		Id:                       k.Id,
		NamaKegiatan:             k.NamaKegiatan,
		IndikatorKinerjaKegiatan: k.IndikatorKinerjaKegiatan,
		PaguKegiatan:             k.PaguKegiatan,
		ProgramId:                k.ProgramId,
		CreatedAt:                k.CreatedAt,
		UpdatedAt:                k.UpdatedAt,
		Program: responses.ProgramResponse{
			Id:                      k.Program.Id,
			NamaProgram:             k.Program.NamaProgram,
			IndikatorKinerjaProgram: k.Program.IndikatorKinerjaProgram,
			PaguProgram:             k.Program.PaguProgram,
			InstansiId:              k.Program.InstansiId,
			CreatedAt:               k.Program.CreatedAt,
			UpdatedAt:               k.Program.UpdatedAt,
		},
	}
}

func ConvertToSubKegiatanResponse(s model.Sub_Kegiatan) responses.Sub_KegiatanResponse {
	return responses.Sub_KegiatanResponse{
		Id:                          s.Id,
		NamaSubKegiatan:             s.NamaSubKegiatan,
		IndikatorKinerjaSubKegiatan: s.IndikatorKinerjaSubKegiatan,
		PaguSubKegiatan:             s.PaguSubKegiatan,
		KegiatanId:                  s.KegiatanId,
		CreatedAt:                   s.CreatedAt,
		UpdatedAt:                   s.UpdatedAt,
	}
}

func ConvertToSubKegiatanWithKegiatanResponse(s model.Sub_Kegiatan) responses.Sub_KegiatanWithKegiatanResponse {
	return responses.Sub_KegiatanWithKegiatanResponse{
		Id:                          s.Id,
		NamaSubKegiatan:             s.NamaSubKegiatan,
		IndikatorKinerjaSubKegiatan: s.IndikatorKinerjaSubKegiatan,
		PaguSubKegiatan:             s.PaguSubKegiatan,
		CreatedAt:                   s.CreatedAt,
		UpdatedAt:                   s.UpdatedAt,
		KegiatanId:                  s.KegiatanId,
		Kegiatan: responses.KegiatanResponse{
			Id:                       s.Kegiatan.Id,
			NamaKegiatan:             s.Kegiatan.NamaKegiatan,
			IndikatorKinerjaKegiatan: s.Kegiatan.IndikatorKinerjaKegiatan,
			PaguKegiatan:             s.Kegiatan.PaguKegiatan,
			CreatedAt:                s.Kegiatan.CreatedAt,
			UpdatedAt:                s.Kegiatan.UpdatedAt,
			ProgramId:                s.Kegiatan.ProgramId,
		},
	}
}

func ConvertToDetailSubKegiatanResponse(d model.Detail_Sub_Kegiatan) responses.Detail_Sub_KegiatanResponse {
	return responses.Detail_Sub_KegiatanResponse{
		Id:               d.Id,
		FokusBelanja:     d.FokusBelanja,
		Indikator:        d.Indikator,
		Target:           d.Target,
		Satuan:           d.Satuan,
		PaguFokusBelanja: d.PaguFokusBelanja,
		Keterangan:       d.Keterangan,
		SubKegiatanId:    d.SubKegiatanId,
		CreatedAt:        d.CreatedAt,
		UpdatedAt:        d.UpdatedAt,
	}
}

func ConvertToDetailSubKegiatanWithSubKegiatanResponse(d model.Detail_Sub_Kegiatan) responses.Detail_Sub_KegiatanWithSub_KegiatanResponse {
	return responses.Detail_Sub_KegiatanWithSub_KegiatanResponse{
		Id:               d.Id,
		FokusBelanja:     d.FokusBelanja,
		Indikator:        d.Indikator,
		Target:           d.Target,
		Satuan:           d.Satuan,
		PaguFokusBelanja: d.PaguFokusBelanja,
		Keterangan:       d.Keterangan,
		CreatedAt:        d.CreatedAt,
		UpdatedAt:        d.UpdatedAt,
		SubKegiatanId:    d.SubKegiatanId,
		SubKegiatan: responses.Sub_KegiatanResponse{
			Id:                          d.SubKegiatan.Id,
			NamaSubKegiatan:             d.SubKegiatan.NamaSubKegiatan,
			IndikatorKinerjaSubKegiatan: d.SubKegiatan.IndikatorKinerjaSubKegiatan,
			PaguSubKegiatan:             d.SubKegiatan.PaguSubKegiatan,
			CreatedAt:                   d.SubKegiatan.CreatedAt,
			UpdatedAt:                   d.SubKegiatan.UpdatedAt,
			KegiatanId:                  d.SubKegiatan.KegiatanId,
		},
	}
}

func ConvertToDetailLokasiResponse(d model.Detail_Lokasi) responses.Detail_LokasiResponse {
	return responses.Detail_LokasiResponse{
		Id:                  d.Id,
		KabupatenKotaId:     d.KabupatenKotaId,
		KecamatanId:         d.KecamatanId,
		KelurahanId:         d.KelurahanId,
		DetailSubKegiatanId: d.DetailSubKegiatanId,
		CreatedAt:           d.CreatedAt,
		UpdatedAt:           d.UpdatedAt,
	}
}

func ConvertToDetailLokasiWithRelationResponse(d model.Detail_Lokasi) responses.Detail_LokasiWithRelationResponse {
	return responses.Detail_LokasiWithRelationResponse{
		Id:                  d.Id,
		KabupatenKotaId:     d.KabupatenKotaId,
		KecamatanId:         d.KecamatanId,
		KelurahanId:         d.KelurahanId,
		CreatedAt:           d.CreatedAt,
		UpdatedAt:           d.UpdatedAt,
		DetailSubKegiatanId: d.DetailSubKegiatanId,
		DetailSubKegiatan: responses.Detail_Sub_KegiatanResponse{
			Id:               d.DetailSubKegiatan.Id,
			FokusBelanja:     d.DetailSubKegiatan.FokusBelanja,
			Indikator:        d.DetailSubKegiatan.Indikator,
			Target:           d.DetailSubKegiatan.Target,
			Satuan:           d.DetailSubKegiatan.Satuan,
			PaguFokusBelanja: d.DetailSubKegiatan.PaguFokusBelanja,
			Keterangan:       d.DetailSubKegiatan.Keterangan,
			CreatedAt:        d.DetailSubKegiatan.CreatedAt,
			UpdatedAt:        d.DetailSubKegiatan.UpdatedAt,
			SubKegiatanId:    d.DetailSubKegiatan.SubKegiatanId,
		},
		KabupatenKota: responses.Kabupaten_KotaResponse{
			Id:         d.KabupatenKota.Id,
			ProvinsiId: d.KabupatenKota.ProvinsiId,
			Nama:       d.KabupatenKota.Nama,
		},
		Kecamatan: responses.KecamatanResponse{
			Id:              d.Kecamatan.Id,
			KabupatenKotaId: d.Kecamatan.KabupatenKotaId,
			Nama:            d.Kecamatan.Nama,
		},
		Kelurahan: responses.KelurahanResponse{
			Id:          d.Kelurahan.Id,
			KecamatanId: d.Kelurahan.KecamatanId,
			Nama:        d.Kelurahan.Nama,
		},
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

func ConvertToKeluargaSigiResponse(k model.KeluargaSigi) responses.KeluargaResponse {
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
	}
}

func ConvertToKeluargaDonggalaResponse(k model.KeluargaDonggala) responses.KeluargaResponse {
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
	}
}
