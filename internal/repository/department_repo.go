package repository

import (
	"github.com/briandang59/be_scada/config"
	"github.com/briandang59/be_scada/internal/model"
)

type DepartmentRepository interface {
	Create(f *model.Department) error
	FindAll(page, pageSize int, preloadFields []string) ([]model.Department, int64, error)
	UpdatePartial(id uint, data map[string]interface{}) error
	FindByID(id uint) (*model.Department, error)
	Delete(id uint) (*model.Department, error)
}

type departmentRepo struct{}

func NewDepartmentRepo() DepartmentRepository {
	return &departmentRepo{}
}

func (r *departmentRepo) FindAll(page, pageSize int, preloadFields []string) ([]model.Department, int64, error) {
	var list []model.Department
	var total int64
	offset := (page - 1) * pageSize

	query := config.DB.Model(&model.Department{}).Where("active = ?", true)

	for _, field := range preloadFields {
		query = query.Preload(field)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.
		Limit(pageSize).
		Offset(offset).
		Find(&list).Error; err != nil {
		return nil, 0, err
	}

	return list, total, nil
}

func (r *departmentRepo) Create(f *model.Department) error {
	return config.DB.Create(f).Error
}

func (r *departmentRepo) UpdatePartial(id uint, data map[string]interface{}) error {
	return config.DB.Model(&model.Department{}).Where("id = ?", id).Updates(data).Error
}
func (r *departmentRepo) FindByID(id uint) (*model.Department, error) {
	var department model.Department
	if err := config.DB.First(&department, id).Error; err != nil {
		return nil, err
	}
	return &department, nil
}

func (r *departmentRepo) Delete(id uint) (*model.Department, error) {
	var department model.Department

	if err := config.DB.First(&department, id).Error; err != nil {
		return nil, err
	}

	if err := config.DB.Delete(&department).Error; err != nil {
		return nil, err
	}

	return &department, nil
}
