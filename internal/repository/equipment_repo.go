package repository

import (
	"github.com/briandang59/be_scada/config"
	"github.com/briandang59/be_scada/internal/model"
)

// Định nghĩa interface
type EquipmentRepository interface {
	Create(e *model.Equipment) error
	FindByID(id uint) (*model.Equipment, error)
	FindAll() ([]model.Equipment, error)
	FindAllPaginate(page, pageSize int) ([]model.Equipment, int64, error)
	Update(e *model.Equipment) error
	Delete(id uint) error
}

// Cấu trúc triển khai interface
type equipmentRepo struct{}

func NewEquipmentRepo() EquipmentRepository {
	return &equipmentRepo{}
}
func (r *equipmentRepo) FindAllPaginate(page, pageSize int) ([]model.Equipment, int64, error) {
	var list []model.Equipment
	var total int64
	offset := (page - 1) * pageSize

	if err := config.DB.Model(&model.Equipment{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if err := config.DB.Limit(pageSize).Offset(offset).Find(&list).Error; err != nil {
		return nil, 0, err
	}
	return list, total, nil
}
func (r *equipmentRepo) Create(e *model.Equipment) error {
	return config.DB.Create(e).Error
}

func (r *equipmentRepo) FindByID(id uint) (*model.Equipment, error) {
	var eq model.Equipment
	err := config.DB.First(&eq, id).Error
	return &eq, err
}

func (r *equipmentRepo) FindAll() ([]model.Equipment, error) {
	var list []model.Equipment
	err := config.DB.Find(&list).Error
	return list, err
}

func (r *equipmentRepo) Update(e *model.Equipment) error {
	return config.DB.Save(e).Error
}

func (r *equipmentRepo) Delete(id uint) error {
	return config.DB.Delete(&model.Equipment{}, id).Error
}
