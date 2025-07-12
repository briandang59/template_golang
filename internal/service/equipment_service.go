package service

import (
	"github.com/briandang59/be_scada/internal/model"
	"github.com/briandang59/be_scada/internal/repository"
)

type EquipmentService struct {
	repo repository.EquipmentRepository
}

func NewEquipmentService(r repository.EquipmentRepository) *EquipmentService {
	return &EquipmentService{repo: r}
}

func (s *EquipmentService) GetAll(page, pageSize int, preloadFields []string) ([]model.Equipment, int64, error) {
	return s.repo.FindAll(page, pageSize, preloadFields)
}

func (s *EquipmentService) UpdatePartial(id uint, data map[string]interface{}) error {
	return s.repo.UpdatePartial(id, data)
}

func (s *EquipmentService) GetByID(id uint) (*model.Equipment, error) {
	return s.repo.FindByID(id)
}

func (s *EquipmentService) Delete(id uint) (*model.Equipment, error) {
	return s.repo.Delete(id)
}

func (s *EquipmentService) Create(f *model.Equipment) error {
	return s.repo.Create(f)
}
