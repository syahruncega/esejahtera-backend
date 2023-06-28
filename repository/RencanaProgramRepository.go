package repository

import (
	"kemiskinan/model"

	"gorm.io/gorm"
)

type RencanaProgramRepository interface {
	FindAll() ([]model.RencanaProgram, error)
	FindById(id int) (model.RencanaProgram, error)
	FindBySearch(whereClause map[string]interface{}) ([]model.RencanaProgram, error)
	SumAllPaguRencanaProgram(tahun, tipe string) int64
	SumPaguRencanaProgramByInstansi(tahun, tipe string, instansis []model.Instansi) []int64
	CountJumlahRencanaProgram(tahun string, tipe string) (int64, error)
	CountJumlahRencanaProgramByIntansi(tahun string, tipe string, instansis []model.Instansi) []int64
	Create(rencanaProgram model.RencanaProgram) (model.RencanaProgram, error)
	Update(rencanaProgram model.RencanaProgram) (model.RencanaProgram, error)
	Delete(rencanaProgram model.RencanaProgram) (model.RencanaProgram, error)
}

type rencanaProgramRepository struct {
	db *gorm.DB
}

func NewRencanaProgramRepository(db *gorm.DB) *rencanaProgramRepository {
	return &rencanaProgramRepository{db}
}

func (r *rencanaProgramRepository) FindAll() ([]model.RencanaProgram, error) {
	var rencanaPrograms []model.RencanaProgram

	var err = r.db.Model(&rencanaPrograms).Preload("Instansi").Preload("Program").Find(&rencanaPrograms).Error

	return rencanaPrograms, err
}

func (r *rencanaProgramRepository) FindById(id int) (model.RencanaProgram, error) {
	var rencanaProgram model.RencanaProgram

	var err = r.db.Model(&rencanaProgram).Preload("Instansi").Preload("Program").Take(&rencanaProgram, id).Error

	return rencanaProgram, err
}

func (r *rencanaProgramRepository) FindBySearch(whereClause map[string]interface{}) ([]model.RencanaProgram, error) {
	var rencanaPrograms []model.RencanaProgram

	var err = r.db.Where(whereClause).Model(&rencanaPrograms).Preload("Instansi").Preload("Program").Find(&rencanaPrograms).Error

	return rencanaPrograms, err
}

func (r *rencanaProgramRepository) SumAllPaguRencanaProgram(tahun, tipe string) int64 {
	var totalPaguRencanaProgram int64

	var _ = r.db.Table("rencana_programs").Select("sum(paguProgram) as totalPaguRencanaProgram").Where("tahun = ? and tipe = ?", tahun, tipe).Scan(&totalPaguRencanaProgram)

	return totalPaguRencanaProgram
}

func (r *rencanaProgramRepository) SumPaguRencanaProgramByInstansi(tahun string, tipe string, instansis []model.Instansi) []int64 {
	var sumPaguRencanaProgram int64
	var totalPaguRencanaProgram []int64

	for i := 0; i < len(instansis); i++ {

		var rows, _ = r.db.Table("rencana_programs").Where("tahun = ? and tipe = ? and instansiId = ?", tahun, tipe, instansis[i].Id).Select("sum(paguProgram) as sumPaguRencanaProgram").Rows()

		for rows.Next() {
			rows.Scan(&sumPaguRencanaProgram)
		}

		totalPaguRencanaProgram = append(totalPaguRencanaProgram, sumPaguRencanaProgram)
	}

	return totalPaguRencanaProgram
}

func (r *rencanaProgramRepository) CountJumlahRencanaProgram(tahun string, tipe string) (int64, error) {
	var count int64

	var err = r.db.Where("tahun = ? and tipe = ?", tahun, tipe).Table("rencana_programs").Select("count(*)").Count(&count).Error

	return count, err
}

func (r *rencanaProgramRepository) CountJumlahRencanaProgramByIntansi(tahun string, tipe string, instansis []model.Instansi) []int64 {
	var count int64
	var hasil []int64

	for i := 0; i < len(instansis); i++ {
		var _ = r.db.Select("count(*)").Table("rencana_programs as rp").Joins("inner join programs as p on p.id = rp.programId").Where("p.tahun = ? and rp.tipe = ? and rp.instansiId = ?", tahun, tipe, instansis[i].Id).Scan(&count)

		hasil = append(hasil, count)
	}

	return hasil
}

func (r *rencanaProgramRepository) Create(rencanaProgram model.RencanaProgram) (model.RencanaProgram, error) {
	var err = r.db.Create(&rencanaProgram).Error

	return rencanaProgram, err
}

func (r *rencanaProgramRepository) Update(rencanaProgram model.RencanaProgram) (model.RencanaProgram, error) {
	var err = r.db.Model(&rencanaProgram).Updates(model.RencanaProgram{
		InstansiId:  rencanaProgram.InstansiId,
		ProgramId:   rencanaProgram.ProgramId,
		PaguProgram: rencanaProgram.PaguProgram,
		Tipe:        rencanaProgram.Tipe,
		Tahun:       rencanaProgram.Tahun,
	}).Error

	return rencanaProgram, err
}

func (r *rencanaProgramRepository) Delete(rencanaProgram model.RencanaProgram) (model.RencanaProgram, error) {
	var err = r.db.Delete(&rencanaProgram).Error

	return rencanaProgram, err
}
