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
	status := c.QueryParam("status") // done, undone
	priority := c.QueryParam("priority")
	sortBy := c.QueryParam("sort")   // priority, due_date, created, position
	order := c.QueryParam("order")   // asc, desc
	filter := c.QueryParam("filter") // overdue, today, upcoming, no_date

	var tasks []models.Task

	query := db.Engine.Where("parent_id = 0 OR parent_id IS NULL") // exclude subtasks from top-level

	if projectID != "" {
		id, _ := strconv.ParseInt(projectID, 10, 64)
		query = query.And("project_id = ?", id)
	}

	if status == "done" {
		query = query.And("done = true")
	} else if status == "undone" {
		query = query.And("done = false")
	}

	if priority != "" {
		p, _ := strconv.Atoi(priority)
		query = query.And("priority = ?", p)
	}

	now := time.Now()
	startOfDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	endOfDay := startOfDay.Add(24 * time.Hour)
	endOfWeek := startOfDay.AddDate(0, 0, 7)

	switch filter {
	case "overdue":
		query = query.And("done = false AND due_date IS NOT NULL AND due_date < ?", now)
	case "today":
		query = query.And("due_date >= ? AND due_date < ?", startOfDay, endOfDay)
	case "upcoming":
		query = query.And("due_date >= ? AND due_date < ?", now, endOfWeek)
	case "no_date":
		query = query.And("due_date IS NULL AND done = false")
	}

	switch sortBy {
	case "priority":
		if order == "asc" {
			query = query.Asc("priority")
		} else {
			query = query.Desc("priority")
		}
	case "due_date":
		if order == "desc" {
			query = query.Desc("due_date")
		} else {
			query = query.Asc("due_date")
		}
	case "created":
		if order == "asc" {
			query = query.Asc("created")
		} else {
			query = query.Desc("created")
		}
	default:
		query = query.Asc("position").Desc("priority")
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
		Title         string       `json:"title" validate:"required"`
		Description   string       `json:"description"`
		Priority      int          `json:"priority"`
		Urgency       int          `json:"urgency"`
		Importance    int          `json:"importance"`
		DueDate       *time.Time   `json:"due_date"`
		ProjectID     int64        `json:"project_id"`
		BucketID      int64        `json:"bucket_id"`
		ParentID      int64        `json:"parent_id"`
		Labels        models.Labels `json:"labels"`
		StartDate     *time.Time   `json:"start_date"`
		EndDate       *time.Time   `json:"end_date"`
		EstimatedTime int64        `json:"estimated_time"`
		AssigneeID    int64        `json:"assignee_id"`
		RepeatAfter   string       `json:"repeat_after"`
		RepeatMode    string       `json:"repeat_mode"`
	}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	task := &models.Task{
		Title:         req.Title,
		Description:   req.Description,
		Priority:      req.Priority,
		Urgency:       req.Urgency,
		Importance:    req.Importance,
		DueDate:       req.DueDate,
		ProjectID:     req.ProjectID,
		BucketID:      req.BucketID,
		ParentID:      req.ParentID,
		Labels:        req.Labels,
		StartDate:     req.StartDate,
		EndDate:       req.EndDate,
		EstimatedTime: req.EstimatedTime,
		AssigneeID:    req.AssigneeID,
		RepeatAfter:   req.RepeatAfter,
		RepeatMode:    req.RepeatMode,
		CreatedByID:   userID.(int64),
		Done:          false,
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

	// Also fetch subtasks
	var subtasks []models.Task
	db.Engine.Where("parent_id = ?", id).Asc("position").Find(&subtasks)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"task":     task,
		"subtasks": subtasks,
	})
}

func UpdateTask(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid task id"})
	}

	var req struct {
		Title          string        `json:"title"`
		Description    string        `json:"description"`
		Done           bool          `json:"done"`
		Priority       int           `json:"priority"`
		Urgency        int           `json:"urgency"`
		Importance     int           `json:"importance"`
		DueDate        *time.Time    `json:"due_date"`
		BucketID       int64         `json:"bucket_id"`
		Labels         models.Labels `json:"labels"`
		Position       int           `json:"position"`
		PercentDone    int           `json:"percent_done"`
		StartDate      *time.Time    `json:"start_date"`
		EndDate        *time.Time    `json:"end_date"`
		EstimatedTime  int64         `json:"estimated_time"`
		TotalTimeSpent int64         `json:"total_time_spent"`
		AssigneeID     int64         `json:"assignee_id"`
		RepeatAfter    string        `json:"repeat_after"`
		RepeatMode     string        `json:"repeat_mode"`
	}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	task := &models.Task{
		ID:             id,
		Title:          req.Title,
		Description:    req.Description,
		Done:           req.Done,
		Priority:       req.Priority,
		Urgency:        req.Urgency,
		Importance:     req.Importance,
		DueDate:        req.DueDate,
		BucketID:       req.BucketID,
		Labels:         req.Labels,
		Position:       req.Position,
		PercentDone:    req.PercentDone,
		StartDate:      req.StartDate,
		EndDate:        req.EndDate,
		EstimatedTime:  req.EstimatedTime,
		TotalTimeSpent: req.TotalTimeSpent,
		AssigneeID:     req.AssigneeID,
		RepeatAfter:    req.RepeatAfter,
		RepeatMode:     req.RepeatMode,
	}

	_, err = db.Engine.ID(id).Update(task)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to update task"})
	}

	// Fetch updated task
	updated := &models.Task{}
	db.Engine.ID(id).Get(updated)

	return c.JSON(http.StatusOK, updated)
}

func DeleteTask(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid task id"})
	}

	// Also delete subtasks
	db.Engine.Where("parent_id = ?", id).Delete(&models.Task{})

	_, err = db.Engine.ID(id).Delete(&models.Task{})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to delete task"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "task deleted"})
}
