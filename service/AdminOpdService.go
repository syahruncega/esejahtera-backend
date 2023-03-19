package service

import (
	"kemiskinan/model"
	"kemiskinan/repository"
	"kemiskinan/request"
)

type AdminOpdService interface {
	FindAll() ([]model.AdminOpd, error)
	FindById(id int) (model.AdminOpd, error)
	FindByUserId(userId int) (model.AdminOpd, error)
	Create(adminOpdRequest request.CreateAdminOpdRequest) (model.AdminOpd, error)
	Update(id int, adminOpdRequest request.UpdateAdminOpdRequest) (model.AdminOpd, error)
	Delete(id int) (model.AdminOpd, error)
}

type adminOpdService struct {
	adminOpdRepository repository.AdminOpdRepository
}

func NewAdminOpdService(adminOpdRepository repository.AdminOpdRepository) *adminOpdService {
	return &adminOpdService{adminOpdRepository}
}

func (s *adminOpdService) FindAll() ([]model.AdminOpd, error) {
	var adminOpds, err = s.adminOpdRepository.FindAll()

	return adminOpds, err
}

func (s *adminOpdService) FindById(id int) (model.AdminOpd, error) {
	var adminOpd, err = s.adminOpdRepository.FindById(id)

	return adminOpd, err
}

func (s *adminOpdService) FindByUserId(userId int) (model.AdminOpd, error) {
	var adminOpd, err = s.adminOpdRepository.FindByUserId(userId)

	return adminOpd, err
}

func (s *adminOpdService) Create(adminOpdRequest request.CreateAdminOpdRequest) (model.AdminOpd, error) {
	var adminOpd = model.AdminOpd{
		UserId:        adminOpdRequest.UserId,
		NamaLengkap:   adminOpdRequest.NamaLengkap,
		InstansiId:    adminOpdRequest.InstansiId,
		UrlFotoProfil: adminOpdRequest.UrlFotoProfil,
	}

	newAdminOpd, err := s.adminOpdRepository.Create(adminOpd)

	return newAdminOpd, err
}

func (s *adminOpdService) Update(id int, adminOpdRequest request.UpdateAdminOpdRequest) (model.AdminOpd, error) {
	var adminOpd, err = s.adminOpdRepository.FindById(id)

	adminOpd.UserId = adminOpdRequest.UserId
	adminOpd.NamaLengkap = adminOpdRequest.NamaLengkap
	adminOpd.InstansiId = adminOpdRequest.InstansiId
	adminOpd.UrlFotoProfil = adminOpdRequest.UrlFotoProfil

	updatedAdminOpd, err := s.adminOpdRepository.Update(adminOpd)

	return updatedAdminOpd, err
}

func (s *adminOpdService) Delete(id int) (model.AdminOpd, error) {
	var adminOpd, err = s.adminOpdRepository.FindById(id)

	deletedAdminOpd, err := s.adminOpdRepository.Delete(adminOpd)

	return deletedAdminOpd, err
}
