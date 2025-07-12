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

type PersonnelHandler struct {
	svc *service.PersonnelService
}

func NewPersonnelHandler(s *service.PersonnelService) *PersonnelHandler {
	return &PersonnelHandler{svc: s}
}

// GetAll godoc
// @Summary Get all personnels
// @Description Get list of personnels with pagination
// @Tags Personnel
// @Accept json
// @Produce json
// @Param page query int false "Page number"
// @Param page_size query int false "Items per page"
// @Success 200 {object} response.PersonnelListResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /personnels [get]
func (h *PersonnelHandler) GetAll(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
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
// @Summary Create a new personnel
// @Description Create a personnel with given payload
// @Tags Personnel
// @Accept json
// @Produce json
// @Param personnel body model.Personnel true "personnel Data"
// @Success 201 {object} response.PersonnelResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /personnels [post]
func (h *PersonnelHandler) Create(c *gin.Context) {
	var body model.Personnel
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
// @Summary Update a personnel (partial)
// @Description Update a personnel partially by ID
// @Tags Personnel
// @Accept json
// @Produce json
// @Param id path int true "personnel ID"
// @Param body body object true "Partial update fields"
// @Success 200 {object} response.PersonnelResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /personnels/{id} [patch]
func (h *PersonnelHandler) Update(c *gin.Context) {
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
	personnel, err := h.svc.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "cannot fetch updated personnel"})
		return
	}

	c.JSON(http.StatusOK, personnel)
}
