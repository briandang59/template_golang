package repository

import (
	"github.com/briandang59/be_scada/config"
	"github.com/briandang59/be_scada/internal/model"
)

type EquipmentRepository interface {
	FindAll(page, pageSize int, preloadFields []string) ([]model.Equipment, int64, error)
	Create(f *model.Equipment) error
	UpdatePartial(id uint, data map[string]interface{}) error
	FindByID(id uint) (*model.Equipment, error)
	Delete(id uint) (*model.Equipment, error)
}

type equipmentRepo struct {
}

func NewEquipmentRepo() EquipmentRepository {
	return &equipmentRepo{}
}

func (r *equipmentRepo) FindAll(page, pageSize int, preloadFields []string) ([]model.Equipment, int64, error) {
	var list []model.Equipment
	var total int64
	offset := (page - 1) * pageSize

	query := config.DB.Model(&model.Equipment{}).Where("active = ?", true)

	for _, field := range preloadFields {
		query = query.Preload((field))
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Limit(pageSize).Offset(offset).Find(&list).Error; err != nil {
		return nil, 0, err
	}

	return list, total, nil
}

func (r *equipmentRepo) Create(f *model.Equipment) error {
	return config.DB.Create(f).Error

}

func (r *equipmentRepo) UpdatePartial(id uint, data map[string]interface{}) error {
	return config.DB.Model(&model.Equipment{}).Where("id = ?", id).Updates(data).Error
}
func (r *equipmentRepo) FindByID(id uint) (*model.Equipment, error) {
	var equipment model.Equipment
	if err := config.DB.First(&equipment, id).Error; err != nil {
		return nil, err
	}
	return &equipment, nil
}

func (r *equipmentRepo) Delete(id uint) (*model.Equipment, error) {
	var equipment model.Equipment

	if err := config.DB.First(&equipment, id).Error; err != nil {
		return nil, err
	}

	if err := config.DB.Delete(&equipment).Error; err != nil {
		return nil, err
	}

	return &equipment, nil
}
