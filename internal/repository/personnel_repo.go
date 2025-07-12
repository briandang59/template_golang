package repository

import (
	"github.com/briandang59/be_scada/config"
	"github.com/briandang59/be_scada/internal/model"
)

type PersonnelRepository interface {
	Create(f *model.Personnel) error
	FindAll(page, pageSize int, preloadFields []string) ([]model.Personnel, int64, error)
	UpdatePartial(id uint, data map[string]interface{}) error
	FindByID(id uint) (*model.Personnel, error)
}

type personnelRepo struct{}

func NewPersonnelRepository() PersonnelRepository {
	return &personnelRepo{}
}

func (r *personnelRepo) FindAll(page, pageSize int, preloadFields []string) ([]model.Personnel, int64, error) {
	var list []model.Personnel
	var total int64
	offset := (page - 1) * pageSize

	query := config.DB.Model(&model.Personnel{})

	// preload theo danh sách
	for _, field := range preloadFields {
		query = query.Preload(field)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Limit(pageSize).Offset(offset).Find(&list).Error; err != nil {
		return nil, 0, err
	}

	return list, total, nil
}

func (r *personnelRepo) Create(f *model.Personnel) error {
	return config.DB.Create(f).Error
}

func (r *personnelRepo) UpdatePartial(id uint, data map[string]interface{}) error {
	return config.DB.Model(&model.Personnel{}).Where("id = ?", id).Updates(data).Error
}
func (r *personnelRepo) FindByID(id uint) (*model.Personnel, error) {
	var personnel model.Personnel
	if err := config.DB.First(&personnel, id).Error; err != nil {
		return nil, err
	}
	return &personnel, nil
}
