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
