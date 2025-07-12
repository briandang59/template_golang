package handler

import (
	"net/http"
	"strconv"

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
