package service

import (
	"kemiskinan/model"
	"kemiskinan/repository"
	"kemiskinan/request"
)

type KebijakanService interface {
	FindAll() ([]model.Kebijakan, error)
	FindById(id int) (model.Kebijakan, error)
	Create(kebijakanRequest request.CreateKebijakanRequest) (model.Kebijakan, error)
	Update(id int, kebijakanRequest request.UpdateKebijakanRequest) (model.Kebijakan, error)
	Delete(id int) (model.Kebijakan, error)
}

type kebijakanService struct {
	kebijakanRepository repository.KebijakanRepository
}

func NewKebijakanService(kebijakanRepository repository.KebijakanRepository) *kebijakanService {
	return &kebijakanService{kebijakanRepository}
}

func (s *kebijakanService) FindAll() ([]model.Kebijakan, error) {
	var kebijakans, err = s.kebijakanRepository.FindAll()

	return kebijakans, err
}

func (s *kebijakanService) FindById(id int) (model.Kebijakan, error) {
	var kebijakan, err = s.kebijakanRepository.FindById(id)

	return kebijakan, err
}

func (s *kebijakanService) Create(kebijakanRequest request.CreateKebijakanRequest) (model.Kebijakan, error) {
	var kebijakan = model.Kebijakan{
		ProgramId:     kebijakanRequest.ProgramId,
		NamaKebijakan: kebijakanRequest.NamaKebijakan,
	}

	newKebijakan, err := s.kebijakanRepository.Create(kebijakan)

	return newKebijakan, err
}

func (s *kebijakanService) Update(id int, kebijakanRequest request.UpdateKebijakanRequest) (model.Kebijakan, error) {
	var kebijakan, err = s.kebijakanRepository.FindById(id)

	kebijakan.ProgramId = kebijakanRequest.ProgramId
	kebijakan.NamaKebijakan = kebijakanRequest.NamaKebijakan

	updatedKebijakan, err := s.kebijakanRepository.Update(kebijakan)

	return updatedKebijakan, err
}

func (s *kebijakanService) Delete(id int) (model.Kebijakan, error) {
	var kebijakan, err = s.kebijakanRepository.FindById(id)

	deletedKebijakan, err := s.kebijakanRepository.Delete(kebijakan)

	return deletedKebijakan, err
}
