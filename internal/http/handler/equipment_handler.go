package handler

import (
	"net/http"
	"strconv"

	"github.com/briandang59/be_scada/internal/http/response"
	"github.com/briandang59/be_scada/internal/model"
	"github.com/briandang59/be_scada/internal/service"
	"github.com/gin-gonic/gin"
)

type EquipmentHandler struct {
	svc *service.EquipmentService
}

func NewEquipmentHandler(s *service.EquipmentService) *EquipmentHandler {
	return &EquipmentHandler{svc: s}
}

// GET /api/equipment
func (h *EquipmentHandler) GetAll(c *gin.Context) {
	// ── 1. Lấy query params, gán mặc định ─────────────────────
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}

	// ── 2. Gọi service lấy data + total ────────────────────────
	list, total, err := h.svc.GetAllPaginate(page, pageSize)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	// ── 3. Trả về format chuẩn ────────────────────────────────
	response.Success(c, list, &response.Pagination{
		Page:     page,
		PageSize: pageSize,
		Total:    int(total),
	})
}

// GET /api/equipment/:id
func (h *EquipmentHandler) Get(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	eq, err := h.svc.Get(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "not found"})
		return
	}
	c.JSON(http.StatusOK, eq)
}

// POST /api/equipment
func (h *EquipmentHandler) Create(c *gin.Context) {
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

// PUT /api/equipment/:id
func (h *EquipmentHandler) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var body model.Equipment
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	body.ID = uint(id)
	if err := h.svc.Update(&body); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, body)
}

// DELETE /api/equipment/:id
func (h *EquipmentHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.svc.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
