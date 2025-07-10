package handler

import (
	"net/http"
	"strconv"

	"github.com/briandang59/be_scada/internal/http/response"
	"github.com/briandang59/be_scada/internal/model"
	"github.com/briandang59/be_scada/internal/service"
	"github.com/gin-gonic/gin"
)

type FactoryHandler struct {
	svc *service.FactoryService
}

func NewFactoryHandler(s *service.FactoryService) *FactoryHandler {
	return &FactoryHandler{svc: s}
}

// GetAll godoc
// @Summary Get all factories
// @Description Get list of factories with pagination
// @Tags Factory
// @Accept json
// @Produce json
// @Param page query int false "Page number"
// @Param page_size query int false "Items per page"
// @Success 200 {object} response.FactoryListResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /factories [get]
func (h *FactoryHandler) GetAll(c *gin.Context) {
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
	list, total, err := h.svc.GetAll(page, pageSize)
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

// Create godoc
// @Summary Create a new factory
// @Description Create a factory with given payload
// @Tags Factory
// @Accept json
// @Produce json
// @Param factory body model.Factory true "Factory Data"
// @Success 201 {object} response.FactoryResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /factories [post]
func (h *FactoryHandler) Create(c *gin.Context) {
	var body model.Factory
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
// @Summary Update a factory (partial)
// @Description Update a factory partially by ID
// @Tags Factory
// @Accept json
// @Produce json
// @Param id path int true "Factory ID"
// @Param body body object true "Partial update fields"
// @Success 200 {object} response.FactoryResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /factories/{id} [patch]
func (h *FactoryHandler) Update(c *gin.Context) {
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

	// 🔁 Lấy lại bản ghi sau khi update
	factory, err := h.svc.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "cannot fetch updated factory"})
		return
	}

	c.JSON(http.StatusOK, factory)
}

// Delete godoc
// @Summary Delete a factory
// @Description Delete a factory by ID
// @Tags Factory
// @Param id path int true "Factory ID"
// @Success 200 {object} model.Factory
// @Failure 400 {object} response.Body
// @Failure 404 {object} response.Body
// @Router /factories/{id} [delete]
func (h *FactoryHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid ID")
		return
	}

	deletedFactory, err := h.svc.Delete(uint(id))
	if err != nil {
		response.Error(c, http.StatusNotFound, err.Error())
		return
	}

	response.Success(c, deletedFactory, nil)
}
