package v1

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tecknohow/TaskDream/pkg/db"
	"github.com/tecknohow/TaskDream/pkg/models"
)

func SearchTasks(c echo.Context) error {
	userID := c.Get("user_id")
	if userID == nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "unauthorized"})
	}

	query := c.QueryParam("q")
	if query == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "search query is required"})
	}

	var tasks []models.Task
	searchPattern := "%" + query + "%"
	err := db.Engine.Where("(title LIKE ? OR description LIKE ?) AND created_by_id = ?",
		searchPattern, searchPattern, userID).
		Limit(50).
		Desc("updated").
		Find(&tasks)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "database error"})
	}

	if tasks == nil {
		tasks = []models.Task{}
	}

	// Also search projects
	var projects []models.Project
	err = db.Engine.Where("(title LIKE ? OR description LIKE ?) AND owner_id = ?",
		searchPattern, searchPattern, userID).
		Limit(20).
		Desc("updated").
		Find(&projects)
	if err != nil {
		projects = []models.Project{}
	}

	// Also search notes
	var notes []models.Note
	err = db.Engine.Where("(title LIKE ? OR content LIKE ?) AND created_by_id = ?",
		searchPattern, searchPattern, userID).
		Limit(20).
		Desc("updated").
		Find(&notes)
	if err != nil {
		notes = []models.Note{}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"tasks":    tasks,
		"projects": projects,
		"notes":    notes,
	})
}
