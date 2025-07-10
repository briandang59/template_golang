package service

import (
	"github.com/briandang59/be_scada/internal/model"
	"github.com/briandang59/be_scada/internal/repository"
)

type FactoryService struct {
	repo repository.FactoryRepository
}

func NewFactoryService(r repository.FactoryRepository) *FactoryService {
	return &FactoryService{repo: r}
}

func (s *FactoryService) GetAll(page, pageSize int) ([]model.Factory, int64, error) {
	return s.repo.FindAll(page, pageSize)
}
func (s *FactoryService) Create(f *model.Factory) error {
	return s.repo.Create(f)
}

func (s *FactoryService) UpdatePartial(id uint, data map[string]interface{}) error {
	return s.repo.UpdatePartial(id, data)
}

func (s *FactoryService) GetByID(id uint) (*model.Factory, error) {
	return s.repo.FindByID(id)
}
