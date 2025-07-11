package service

import (
	"github.com/briandang59/be_scada/internal/model"
	"github.com/briandang59/be_scada/internal/repository"
)

type EquipmentTypeService struct {
	repo repository.EquipmentTypeRepository
}

func NewEquipmentTypeService(r repository.EquipmentTypeRepository) *EquipmentTypeService {
	return &EquipmentTypeService{repo: r}
}

func (s *EquipmentTypeService) GetAll(page, pageSize int, preloadFields []string) ([]model.EquipmentType, int64, error) {
	return s.repo.FindAll(page, pageSize, preloadFields)
}

func (s *EquipmentTypeService) Create(f *model.EquipmentType) error {
	return s.repo.Create(f)
}

func (s *EquipmentTypeService) UpdatePartial(id uint, data map[string]interface{}) error {
	return s.repo.UpdatePartial(id, data)
}

func (s *EquipmentTypeService) GetByID(id uint) (*model.EquipmentType, error) {
	return s.repo.FindByID(id)
}

func (s *EquipmentTypeService) Delete(id uint) (*model.EquipmentType, error) {
	return s.repo.Delete(id)
}
