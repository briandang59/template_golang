package service

import (
	"github.com/briandang59/be_scada/internal/model"
	"github.com/briandang59/be_scada/internal/repository"
)

type DepartmentService struct {
	repo repository.DepartmentRepository
}

func NewDepartmentService(r repository.DepartmentRepository) *DepartmentService {
	return &DepartmentService{repo: r}
}

func (s *DepartmentService) GetAll(page, pageSize int, preloadFields []string) ([]model.Department, int64, error) {
	return s.repo.FindAll(page, pageSize, preloadFields)
}

func (s *DepartmentService) Create(f *model.Department) error {
	return s.repo.Create(f)
}

func (s *DepartmentService) UpdatePartial(id uint, data map[string]interface{}) error {
	return s.repo.UpdatePartial(id, data)
}

func (s *DepartmentService) GetByID(id uint) (*model.Department, error) {
	return s.repo.FindByID(id)
}

func (s *DepartmentService) Delete(id uint) (*model.Department, error) {
	return s.repo.Delete(id)
}
