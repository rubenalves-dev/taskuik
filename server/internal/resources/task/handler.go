package task

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rubenalves-dev/taskuik/internal/common"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Handle(r *gin.Engine) {
	rg := r.Group("/tasks")
	// Add routes without trailing slash to avoid 301 redirects that drop CORS headers
	rg.POST("", h.createTask)
	rg.GET("", h.listTasks)

	// Keep trailing-slash variants for compatibility
	rg.POST("/", h.createTask)
	rg.GET("/", h.listTasks)

	rg.GET("/:id", h.getTask)
	rg.PUT("/:id", h.updateTask)
	rg.DELETE("/:id", h.deleteTask)
}

func (h *Handler) createTask(c *gin.Context) {
	var task Task
	if err := c.ShouldBindJSON(&task); err != nil {
		common.ErrJson(c, http.StatusBadRequest, err)
		return
	}
	result := h.service.CreateTask(&task)
	if result.Err != nil {
		common.ErrJson(c, result.StatusCode, result.Err)
		return
	}
	common.OkJson(c, "task created", task)
}

func (h *Handler) listTasks(c *gin.Context) {
	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("page_size", "10")

	var p, ps int
	if _, err := fmt.Sscan(page, &p); err != nil || p < 1 {
		common.ErrJson(c, http.StatusBadRequest, fmt.Errorf("invalid page number"))
		return
	}
	if _, err := fmt.Sscan(pageSize, &ps); err != nil || ps < 1 {
		common.ErrJson(c, http.StatusBadRequest, fmt.Errorf("invalid page size"))
		return
	}
	result := h.service.ListTasks(p, ps)
	if result.Err != nil {
		common.ErrJson(c, result.StatusCode, result.Err)
		return
	}
	common.OkJson(c, "tasks listed", result)
}

func (h *Handler) getTask(c *gin.Context) {
	idParam := c.Param("id")
	var id uint
	if _, err := fmt.Sscan(idParam, &id); err != nil {
		common.ErrJson(c, http.StatusBadRequest, fmt.Errorf("invalid task ID"))
		return
	}
	result := h.service.GetTask(id)
	if result.Err != nil {
		common.ErrJson(c, result.StatusCode, result.Err)
		return
	}
	common.OkJson(c, "task retrieved", result.Data)
}

func (h *Handler) updateTask(c *gin.Context) {
	idParam := c.Param("id")
	var id uint
	if _, err := fmt.Sscan(idParam, &id); err != nil {
		common.ErrJson(c, http.StatusBadRequest, fmt.Errorf("invalid task ID"))
		return
	}
	var task Task
	if err := c.ShouldBindJSON(&task); err != nil {
		common.ErrJson(c, http.StatusBadRequest, err)
		return
	}
	task.ID = id
	result := h.service.UpdateTask(&task)
	if result.Err != nil {
		common.ErrJson(c, result.StatusCode, result.Err)
		return
	}
	common.OkJson(c, "task updated", task)
}

func (h *Handler) deleteTask(c *gin.Context) {
	idParam := c.Param("id")
	var id uint
	if _, err := fmt.Sscan(idParam, &id); err != nil {
		common.ErrJson(c, http.StatusBadRequest, fmt.Errorf("invalid task ID"))
		return
	}
	result := h.service.DeleteTask(id)
	if result.Err != nil {
		common.ErrJson(c, result.StatusCode, result.Err)
		return
	}
	c.Status(http.StatusNoContent)
}
