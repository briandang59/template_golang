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

func (s *EquipmentService) GetAllPaginate(page, pageSize int) ([]model.Equipment, int64, error) {
	return s.repo.FindAllPaginate(page, pageSize)
}

func (s *EquipmentService) Get(id uint) (*model.Equipment, error) {
	return s.repo.FindByID(id)
}

func (s *EquipmentService) Create(e *model.Equipment) error {
	return s.repo.Create(e)
}

func (s *EquipmentService) Update(e *model.Equipment) error {
	return s.repo.Update(e)
}

func (s *EquipmentService) Delete(id uint) error {
	return s.repo.Delete(id)
}
