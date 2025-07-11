package handler

import (
	"net/http"
	"strconv"

	"github.com/briandang59/be_scada/internal/http/response"
	"github.com/briandang59/be_scada/internal/model"
	"github.com/briandang59/be_scada/internal/service"
	"github.com/briandang59/be_scada/utils"
	"github.com/gin-gonic/gin"
)

type EquipementTypeHandler struct {
	svc *service.EquipmentTypeService
}

func NewEquipmentTypeHandler(s *service.EquipmentTypeService) *EquipementTypeHandler {
	return &EquipementTypeHandler{svc: s}
}

// GetAll godoc
// @Summary Get all equipment-types
// @Description Get list of equipment-types with pagination
// @Tags Equipment Types
// @Accept json
// @Produce json
// @Param page query int false "Page number"
// @Param page_size query int false "Items per page"
// @Success 200 {object} response.EquipmentTypeListResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /equipment-types [get]
func (h *EquipementTypeHandler) GetAll(c *gin.Context) {
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

	response.Success(c, list, &response.Pagination{
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
// @Param body body model.EquipmentType true "EquipmentType data"
// @Success 201 {object} response.EquipmentTypeResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /equipment-types [post]
func (h *EquipementTypeHandler) Create(c *gin.Context) {
	var body model.EquipmentType
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
// @Param id path int true "EquipmentType ID"
// @Param body body object true "Partial update fields"
// @Success 200 {object} response.EquipmentTypeResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /equipment-types/{id} [patch]
func (h *EquipementTypeHandler) Update(c *gin.Context) {
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
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Cannot fetch updated EquipmentType"})
		return
	}

	c.JSON(http.StatusOK, equipmentTypes)
}

// Delete godoc
// @Summary Delete an equipment type
// @Description Delete an equipment type by ID
// @Tags Equipment Types
// @Param id path int true "EquipmentType ID"
// @Success 200 {object} model.EquipmentType
// @Failure 400 {object} response.ErrorResponse
// @Failure 404 {object} response.ErrorResponse
// @Router /equipment-types/{id} [delete]
func (h *EquipementTypeHandler) Delete(c *gin.Context) {
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
