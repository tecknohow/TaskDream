package v1

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/tecknohow/TaskDream/pkg/db"
	"github.com/tecknohow/TaskDream/pkg/models"
)

func ListTaskComments(c echo.Context) error {
	taskID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid task id"})
	}

	var comments []models.TaskComment
	err = db.Engine.Where("task_id = ?", taskID).Order("created desc").Find(&comments)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "database error"})
	}

	if comments == nil {
		comments = []models.TaskComment{}
	}

	return c.JSON(http.StatusOK, comments)
}

func CreateTaskComment(c echo.Context) error {
	userID := c.Get("user_id")
	if userID == nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "unauthorized"})
	}

	taskID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid task id"})
	}

	var req struct {
		Comment string `json:"comment" validate:"required"`
	}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	comment := &models.TaskComment{
		TaskID:  taskID,
		UserID:  userID.(int64),
		Comment: req.Comment,
	}

	_, err = db.Engine.Insert(comment)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to create comment"})
	}

	return c.JSON(http.StatusCreated, comment)
}

func DeleteTaskComment(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid comment id"})
	}

	_, err = db.Engine.ID(id).Delete(&models.TaskComment{})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to delete comment"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "comment deleted"})
}
