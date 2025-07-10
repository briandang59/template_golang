package repository

import (
	"github.com/briandang59/be_scada/config"
	"github.com/briandang59/be_scada/internal/model"
)

type FactoryRepository interface {
	Create(f *model.Factory) error
	FindAll(page, pageSize int, preloadFields []string) ([]model.Factory, int64, error)
	UpdatePartial(id uint, data map[string]interface{}) error
	FindByID(id uint) (*model.Factory, error)
	Delete(id uint) (*model.Factory, error)
}

type factoryRepo struct{}

func NewFactoryRepo() FactoryRepository {
	return &factoryRepo{}
}

func (r *factoryRepo) FindAll(page, pageSize int, preloadFields []string) ([]model.Factory, int64, error) {
	var list []model.Factory
	var total int64
	offset := (page - 1) * pageSize

	query := config.DB.Model(&model.Factory{}).Where("active = ?", true)

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

func (r *factoryRepo) Create(f *model.Factory) error {
	return config.DB.Create(f).Error
}

func (r *factoryRepo) UpdatePartial(id uint, data map[string]interface{}) error {
	return config.DB.Model(&model.Factory{}).Where("id = ?", id).Updates(data).Error
}
func (r *factoryRepo) FindByID(id uint) (*model.Factory, error) {
	var factory model.Factory
	if err := config.DB.First(&factory, id).Error; err != nil {
		return nil, err
	}
	return &factory, nil
}

func (r *factoryRepo) Delete(id uint) (*model.Factory, error) {
	var factory model.Factory

	if err := config.DB.First(&factory, id).Error; err != nil {
		return nil, err
	}

	if err := config.DB.Delete(&factory).Error; err != nil {
		return nil, err
	}

	return &factory, nil
}
