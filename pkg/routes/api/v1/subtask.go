package v1

import (
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/tecknohow/TaskDream/pkg/db"
	"github.com/tecknohow/TaskDream/pkg/models"
)

func ListSubtasks(c echo.Context) error {
	parentID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid task id"})
	}

	var subtasks []models.Task
	err = db.Engine.Where("parent_id = ?", parentID).Asc("position").Find(&subtasks)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "database error"})
	}

	if subtasks == nil {
		subtasks = []models.Task{}
	}

	return c.JSON(http.StatusOK, subtasks)
}

func CreateSubtask(c echo.Context) error {
	userID := c.Get("user_id")
	if userID == nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "unauthorized"})
	}

	parentID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid task id"})
	}

	// Check parent task exists
	parent := &models.Task{}
	has, _ := db.Engine.ID(parentID).Get(parent)
	if !has {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "parent task not found"})
	}

	var req struct {
		Title       string     `json:"title" validate:"required"`
		Description string     `json:"description"`
		Priority    int        `json:"priority"`
		DueDate     *time.Time `json:"due_date"`
	}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	subtask := &models.Task{
		Title:       req.Title,
		Description: req.Description,
		Priority:    req.Priority,
		DueDate:     req.DueDate,
		ProjectID:   parent.ProjectID,
		ParentID:    parentID,
		CreatedByID: userID.(int64),
		Done:        false,
	}

	_, err = db.Engine.Insert(subtask)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to create subtask"})
	}

	return c.JSON(http.StatusCreated, subtask)
}
