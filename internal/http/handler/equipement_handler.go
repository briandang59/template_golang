package handler

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/briandang59/be_scada/internal/dto"
	"github.com/briandang59/be_scada/internal/http/response"
	"github.com/briandang59/be_scada/internal/model"
	"github.com/briandang59/be_scada/internal/service"
	"github.com/briandang59/be_scada/utils"
	"github.com/gin-gonic/gin"
)

type EquipementHandler struct {
	svc *service.EquipmentService
}

func NewEquipmentHandler(s *service.EquipmentService) *EquipementHandler {
	return &EquipementHandler{svc: s}
}

// GetAll godoc
// @Summary Get all equipment-types
// @Description Get list of equipment-types with pagination
// @Tags Equipment Types
// @Accept json
// @Produce json
// @Param page query int false "Page number"
// @Param page_size query int false "Items per page"
// @Success 200 {object} response.EquipmentListResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /equipment-types [get]
func (h *EquipementHandler) GetAll(c *gin.Context) {
	page, _ := strconv.Atoi((c.DefaultQuery("page", "1")))
	pageSize, _ := strconv.Atoi((c.DefaultQuery("page_size", "10")))

	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}

	preloadFields := utils.ParsePopulateQuery(c.Request.URL.Query())

	list, total, err := h.svc.GetAll(page, pageSize, preloadFields)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	ptrList := make([]*model.Equipment, len(list))
	for i := range list {
		ptrList[i] = &list[i]
	}

	res := dto.ToEquipmentResponseList(ptrList)

	response.Success(c, res, &response.Pagination{
		Page:     page,
		PageSize: pageSize,
		Total:    int(total),
	})
}

// Create godoc
// @Summary Create a new equipment type
// @Description Create an equipment type with given payload
// @Tags Equipment Types
// @Accept json
// @Produce json
// @Param body body model.Equipment true "Equipment data"
// @Success 201 {object} response.EquipmentResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /equipment-types [post]
func (h *EquipementHandler) Create(c *gin.Context) {
	var body model.Equipment
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := h.svc.Create(&body); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, body)
}

// Update godoc
// @Summary Update an equipment type (partial)
// @Description Partially update an equipment type by ID
// @Tags Equipment Types
// @Accept json
// @Produce json
// @Param id path int true "Equipment ID"
// @Param body body object true "Partial update fields"
// @Success 200 {object} response.EquipmentTypeResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /equipment-types/{id} [patch]
func (h *EquipementHandler) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var input map[string]interface{}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := h.svc.UpdatePartial(uint(id), input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	equipmentTypes, err := h.svc.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Cannot fetch updated Equipment"})
		return
	}

	c.JSON(http.StatusOK, equipmentTypes)
}

// Delete godoc
// @Summary Delete an equipment type
// @Description Delete an equipment type by ID
// @Tags Equipment Types
// @Param id path int true "Equipment ID"
// @Success 200 {object} model.Equipment
// @Failure 400 {object} response.ErrorResponse
// @Failure 404 {object} response.ErrorResponse
// @Router /equipment-types/{id} [delete]
func (h *EquipementHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid ID")
		return
	}

	deletedEquipmentType, err := h.svc.Delete(uint(id))
	if err != nil {
		response.Error(c, http.StatusNotFound, err.Error())
		return
	}

	response.Success(c, deletedEquipmentType, nil)
}

// ImportFromCSV godoc
// @Summary Import equipment from CSV file
// @Description Import multiple equipment records from a CSV file
// @Tags Equipment
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "CSV file to import"
// @Success 200 {object} response.SuccessResponse{data=service.ImportResult}
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /equipments/import [post]
func (h *EquipementHandler) ImportFromCSV(c *gin.Context) {
	// Get uploaded file
	file, err := c.FormFile("file")
	if err != nil {
		response.Error(c, http.StatusBadRequest, "No file uploaded or invalid file")
		return
	}

	// Validate file type
	if file.Header.Get("Content-Type") != "text/csv" && 
	   file.Header.Get("Content-Type") != "application/vnd.ms-excel" &&
	   file.Header.Get("Content-Type") != "application/octet-stream" {
		// Check file extension as fallback
		if !strings.HasSuffix(file.Filename, ".csv") {
			response.Error(c, http.StatusBadRequest, "File must be a CSV file")
			return
		}
	}

	// Open file
	src, err := file.Open()
	if err != nil {
		response.Error(c, http.StatusInternalServerError, fmt.Sprintf("Failed to open file: %s", err.Error()))
		return
	}
	defer src.Close()

	// Import CSV
	result, err := h.svc.ImportFromCSV(src)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, fmt.Sprintf("Failed to import CSV: %s", err.Error()))
		return
	}

	// Convert to response DTO
	importResponse := response.ToEquipmentImportResponse(
		result.Total,
		result.Success,
		result.Failed,
		result.Errors,
		result.Created,
	)

	response.Success(c, importResponse, nil)
}

// DownloadCSVTemplate godoc
// @Summary Download CSV template for equipment import
// @Description Download a CSV template file with headers for equipment import
// @Tags Equipment
// @Produce text/csv
// @Success 200 {file} file "CSV template file"
// @Router /equipments/template [get]
func (h *EquipementHandler) DownloadCSVTemplate(c *gin.Context) {
	// CSV headers based on EquipmentCSV struct
	headers := []string{
		"name_en", "name_zh", "name_vn", "code", "serial_number", "model", "manufacturer",
		"location", "purchase_date", "warranty_end_date", "installation_date", "status",
		"ip_address", "mac_address", "operating_system", "description", "notes",
		"last_maintenance_date", "next_maintenance_date", "department_id", "equipment_type_id",
		"responsible_user_id", "assigned_user_id",
	}

	// Sample data row
	sampleRow := []string{
		"Production Web Server 01", "生产网络服务器01", "Máy chủ Web Sản xuất 01",
		"SERVER-PROD-WEB-001", "SN1234567890ABC", "Dell PowerEdge R740", "Dell",
		"Server Room A, Rack 2, Unit 1-2", "2023-01-15T00:00:00Z", "2026-01-15T00:00:00Z",
		"2023-01-20T00:00:00Z", "active", "192.168.1.10", "00:1A:2B:3C:4D:5E",
		"Ubuntu Server 22.04 LTS", "Primary web server for production environment",
		"Configured with redundant power supplies", "2024-12-01T00:00:00Z", "2025-03-01T00:00:00Z",
		"1", "1", "1", "1",
	}

	// Set response headers
	c.Header("Content-Type", "text/csv")
	c.Header("Content-Disposition", "attachment; filename=equipment_template.csv")

	// Write CSV content
	writer := csv.NewWriter(c.Writer)
	defer writer.Flush()

	// Write headers
	if err := writer.Write(headers); err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to write CSV headers")
		return
	}

	// Write sample row
	if err := writer.Write(sampleRow); err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to write CSV sample data")
		return
	}
}
