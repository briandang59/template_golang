package service

import (
	"github.com/briandang59/be_scada/internal/model"
	"github.com/briandang59/be_scada/internal/repository"
)

type PersonnelService struct {
	repo repository.PersonnelRepository
}

func NewPersonnelService(r repository.PersonnelRepository) *PersonnelService {
	return &PersonnelService{repo: r}
}

func (s *PersonnelService) GetAll(page, pageSize int, preloadFields []string) ([]model.Personnel, int64, error) {
	return s.repo.FindAll(page, pageSize, preloadFields)
}

func (s *PersonnelService) Create(f *model.Personnel) error {
	return s.repo.Create(f)
}

func (s *PersonnelService) UpdatePartial(id uint, data map[string]interface{}) error {
	return s.repo.UpdatePartial(id, data)
}
func (s *PersonnelService) GetByID(id uint) (*model.Personnel, error) {
	return s.repo.FindByID(id)
}
