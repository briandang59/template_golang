package service

import (
	"fmt"
	"io"

	"github.com/briandang59/be_scada/internal/dto"
	"github.com/briandang59/be_scada/internal/model"
	"github.com/briandang59/be_scada/internal/repository"
	"github.com/gocarina/gocsv"
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

// ImportFromCSV imports equipment data from CSV file
func (s *EquipmentService) ImportFromCSV(reader io.Reader) (*ImportResult, error) {
	var csvData []*dto.EquipmentCSV
	
	// Parse CSV
	if err := gocsv.Unmarshal(reader, &csvData); err != nil {
		return nil, fmt.Errorf("failed to parse CSV: %w", err)
	}

	if len(csvData) == 0 {
		return nil, fmt.Errorf("CSV file is empty")
	}

	result := &ImportResult{
		Total:     len(csvData),
		Success:   0,
		Failed:    0,
		Errors:    make([]string, 0),
		Created:   make([]*model.Equipment, 0),
	}

	// Process each row
	for i, csvRow := range csvData {
		equipment, err := csvRow.ToEquipmentModel()
		if err != nil {
			result.Failed++
			result.Errors = append(result.Errors, fmt.Sprintf("Row %d: Failed to parse data - %s", i+1, err.Error()))
			continue
		}

		// Validate required fields
		if equipment.NameEn == "" {
			result.Failed++
			result.Errors = append(result.Errors, fmt.Sprintf("Row %d: NameEn is required", i+1))
			continue
		}

		if equipment.Code == "" {
			result.Failed++
			result.Errors = append(result.Errors, fmt.Sprintf("Row %d: Code is required", i+1))
			continue
		}

		// Create equipment
		if err := s.repo.Create(equipment); err != nil {
			result.Failed++
			result.Errors = append(result.Errors, fmt.Sprintf("Row %d: Failed to create equipment - %s", i+1, err.Error()))
			continue
		}

		result.Success++
		result.Created = append(result.Created, equipment)
	}

	return result, nil
}

// ImportResult represents the result of CSV import
type ImportResult struct {
	Total   int                `json:"total"`
	Success int                `json:"success"`
	Failed  int                `json:"failed"`
	Errors  []string           `json:"errors"`
	Created []*model.Equipment `json:"created"`
}
