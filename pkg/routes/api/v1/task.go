package v1

import (
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/tecknohow/TaskDream/pkg/db"
	"github.com/tecknohow/TaskDream/pkg/models"
)

func ListTasks(c echo.Context) error {
	projectID := c.QueryParam("project_id")
	var tasks []models.Task

	query := db.Engine
	if projectID != "" {
		id, _ := strconv.ParseInt(projectID, 10, 64)
		query = query.Where("project_id = ?", id)
	}

	err := query.Find(&tasks)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "database error"})
	}

	if tasks == nil {
		tasks = []models.Task{}
	}

	return c.JSON(http.StatusOK, tasks)
}

func CreateTask(c echo.Context) error {
	userID := c.Get("user_id")
	if userID == nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "unauthorized"})
	}

	var req struct {
		Title       string    `json:"title" validate:"required"`
		Description string    `json:"description"`
		Priority    int       `json:"priority"`
		DueDate     *time.Time `json:"due_date"`
		ProjectID   int64     `json:"project_id" validate:"required"`
		BucketID    int64     `json:"bucket_id"`
		Labels      models.Labels `json:"labels"`
		StartDate   *time.Time `json:"start_date"`
		EndDate     *time.Time `json:"end_date"`
	}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	task := &models.Task{
		Title:       req.Title,
		Description: req.Description,
		Priority:    req.Priority,
		DueDate:     req.DueDate,
		ProjectID:   req.ProjectID,
		BucketID:    req.BucketID,
		Labels:      req.Labels,
		StartDate:   req.StartDate,
		EndDate:     req.EndDate,
		CreatedByID: userID.(int64),
		Done:        false,
	}

	_, err := db.Engine.Insert(task)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to create task"})
	}

	return c.JSON(http.StatusCreated, task)
}

func GetTask(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid task id"})
	}

	task := &models.Task{}
	has, err := db.Engine.ID(id).Get(task)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "database error"})
	}

	if !has {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "task not found"})
	}

	return c.JSON(http.StatusOK, task)
}

func UpdateTask(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid task id"})
	}

	var req struct {
		Title       string         `json:"title"`
		Description string         `json:"description"`
		Done        bool           `json:"done"`
		Priority    int            `json:"priority"`
		DueDate     *time.Time     `json:"due_date"`
		BucketID    int64          `json:"bucket_id"`
		Labels      models.Labels  `json:"labels"`
		Position    int            `json:"position"`
		PercentDone int            `json:"percent_done"`
		StartDate   *time.Time     `json:"start_date"`
		EndDate     *time.Time     `json:"end_date"`
	}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	task := &models.Task{
		ID:          id,
		Title:       req.Title,
		Description: req.Description,
		Done:        req.Done,
		Priority:    req.Priority,
		DueDate:     req.DueDate,
		BucketID:    req.BucketID,
		Labels:      req.Labels,
		Position:    req.Position,
		PercentDone: req.PercentDone,
		StartDate:   req.StartDate,
		EndDate:     req.EndDate,
	}

	_, err = db.Engine.ID(id).Update(task)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to update task"})
	}

	return c.JSON(http.StatusOK, task)
}

func DeleteTask(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid task id"})
	}

	_, err = db.Engine.ID(id).Delete(&models.Task{})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to delete task"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "task deleted"})
}
