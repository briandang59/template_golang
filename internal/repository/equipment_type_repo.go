package repository

import (
	"github.com/briandang59/be_scada/config"
	"github.com/briandang59/be_scada/internal/model"
)

type EquipmentTypeRepository interface {
	FindAll(page, pageSize int, preloadFields []string) ([]model.EquipmentType, int64, error)
	Create(f *model.EquipmentType) error
	UpdatePartial(id uint, data map[string]interface{}) error
	FindByID(id uint) (*model.EquipmentType, error)
	Delete(id uint) (*model.EquipmentType, error)
}

type equipmentTypeRepo struct {
}

func NewEquipmentTypeRepo() EquipmentTypeRepository {
	return &equipmentTypeRepo{}
}

func (r *equipmentTypeRepo) FindAll(page, pageSize int, preloadFields []string) ([]model.EquipmentType, int64, error) {
	var list []model.EquipmentType
	var total int64
	offset := (page - 1) * pageSize

	query := config.DB.Model(&model.EquipmentType{}).Where("active = ?", true)

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

func (r *equipmentTypeRepo) Create(f *model.EquipmentType) error {
	return config.DB.Create(f).Error

}

func (r *equipmentTypeRepo) UpdatePartial(id uint, data map[string]interface{}) error {
	return config.DB.Model(&model.EquipmentType{}).Where("id = ?", id).Updates(data).Error
}
func (r *equipmentTypeRepo) FindByID(id uint) (*model.EquipmentType, error) {
	var equipmentType model.EquipmentType
	if err := config.DB.First(&equipmentType, id).Error; err != nil {
		return nil, err
	}
	return &equipmentType, nil
}

func (r *equipmentTypeRepo) Delete(id uint) (*model.EquipmentType, error) {
	var equipmentType model.EquipmentType

	if err := config.DB.First(&equipmentType, id).Error; err != nil {
		return nil, err
	}

	if err := config.DB.Delete(&equipmentType).Error; err != nil {
		return nil, err
	}

	return &equipmentType, nil
}
