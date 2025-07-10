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

type DepartmentHandler struct {
	svc *service.DepartmentService
}

func NewDepartmentHandler(s *service.DepartmentService) *DepartmentHandler {
	return &DepartmentHandler{svc: s}
}

// GetAll godoc
// @Summary Get all departments
// @Description Get list of departments with pagination
// @Tags Department
// @Accept json
// @Produce json
// @Param page query int false "Page number"
// @Param page_size query int false "Items per page"
// @Success 200 {object} response.DepartmentListResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /departments [get]
func (h *DepartmentHandler) GetAll(c *gin.Context) {
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
// @Summary Create a new department
// @Description Create a department with given payload
// @Tags Department
// @Accept json
// @Produce json
// @Param department body model.Department true "Department Data"
// @Success 201 {object} response.DepartmentResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /departments [post]
func (h *DepartmentHandler) Create(c *gin.Context) {
	var body model.Department
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
// @Summary Update a department (partial)
// @Description Update a department partially by ID
// @Tags Department
// @Accept json
// @Produce json
// @Param id path int true "Department ID"
// @Param body body object true "Partial update fields"
// @Success 200 {object} response.DepartmentResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /departments/{id} [patch]
func (h *DepartmentHandler) Update(c *gin.Context) {
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
	department, err := h.svc.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "cannot fetch updated department"})
		return
	}

	c.JSON(http.StatusOK, department)
}

// Delete godoc
// @Summary Delete a department
// @Description Delete a department by ID
// @Tags Department
// @Param id path int true "Department ID"
// @Success 200 {object} model.Department
// @Failure 400 {object} response.Body
// @Failure 404 {object} response.Body
// @Router /departments/{id} [delete]
func (h *DepartmentHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid ID")
		return
	}

	deletedDepartment, err := h.svc.Delete(uint(id))
	if err != nil {
		response.Error(c, http.StatusNotFound, err.Error())
		return
	}

	response.Success(c, deletedDepartment, nil)
}
