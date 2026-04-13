package v1

import (
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/tecknohow/TaskDream/pkg/db"
	"github.com/tecknohow/TaskDream/pkg/models"
)

func ListTimeTracking(c echo.Context) error {
	taskID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid task id"})
	}

	var trackings []models.TimeTracking
	err = db.Engine.Where("task_id = ?", taskID).Find(&trackings)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "database error"})
	}

	if trackings == nil {
		trackings = []models.TimeTracking{}
	}

	return c.JSON(http.StatusOK, trackings)
}

func CreateTimeTracking(c echo.Context) error {
	userID := c.Get("user_id")
	if userID == nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "unauthorized"})
	}

	taskID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid task id"})
	}

	var req struct {
		Start    time.Time `json:"start" validate:"required"`
		End      *time.Time `json:"end"`
		Duration int64     `json:"duration"`
		Comment  string    `json:"comment"`
	}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	// Calculate duration if not provided
	duration := req.Duration
	if duration == 0 && req.End != nil {
		duration = int64(req.End.Sub(req.Start).Seconds())
	}

	tracking := &models.TimeTracking{
		TaskID:   taskID,
		UserID:   userID.(int64),
		Start:    req.Start,
		End:      req.End,
		Duration: duration,
		Comment:  req.Comment,
	}

	_, err = db.Engine.Insert(tracking)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to create time tracking"})
	}

	return c.JSON(http.StatusCreated, tracking)
}

func UpdateTimeTracking(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid time tracking id"})
	}

	var req struct {
		Start    time.Time `json:"start"`
		End      *time.Time `json:"end"`
		Duration int64     `json:"duration"`
		Comment  string    `json:"comment"`
	}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	// Calculate duration if not provided
	duration := req.Duration
	if duration == 0 && req.End != nil {
		duration = int64(req.End.Sub(req.Start).Seconds())
	}

	tracking := &models.TimeTracking{
		ID:       id,
		Start:    req.Start,
		End:      req.End,
		Duration: duration,
		Comment:  req.Comment,
	}

	_, err = db.Engine.ID(id).Update(tracking)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to update time tracking"})
	}

	return c.JSON(http.StatusOK, tracking)
}

func DeleteTimeTracking(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid time tracking id"})
	}

	_, err = db.Engine.ID(id).Delete(&models.TimeTracking{})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to delete time tracking"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "time tracking deleted"})
}
