package v1

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/tecknohow/TaskDream/pkg/db"
	"github.com/tecknohow/TaskDream/pkg/models"
)

func ListNotes(c echo.Context) error {
	projectID := c.QueryParam("project_id")
	var notes []models.Note

	query := db.Engine
	if projectID != "" {
		id, _ := strconv.ParseInt(projectID, 10, 64)
		query = query.Where("project_id = ?", id)
	}

	err := query.Find(&notes)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "database error"})
	}

	if notes == nil {
		notes = []models.Note{}
	}

	return c.JSON(http.StatusOK, notes)
}

func CreateNote(c echo.Context) error {
	userID := c.Get("user_id")
	if userID == nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "unauthorized"})
	}

	var req struct {
		Title     string `json:"title" validate:"required"`
		Content   string `json:"content"`
		ProjectID int64  `json:"project_id" validate:"required"`
	}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	note := &models.Note{
		Title:       req.Title,
		Content:     req.Content,
		ProjectID:   req.ProjectID,
		CreatedByID: userID.(int64),
	}

	_, err := db.Engine.Insert(note)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to create note"})
	}

	return c.JSON(http.StatusCreated, note)
}

func GetNote(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid note id"})
	}

	note := &models.Note{}
	has, err := db.Engine.ID(id).Get(note)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "database error"})
	}

	if !has {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "note not found"})
	}

	return c.JSON(http.StatusOK, note)
}

func UpdateNote(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid note id"})
	}

	var req struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	note := &models.Note{
		ID:      id,
		Title:   req.Title,
		Content: req.Content,
	}

	_, err = db.Engine.ID(id).Update(note)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to update note"})
	}

	return c.JSON(http.StatusOK, note)
}

func DeleteNote(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid note id"})
	}

	_, err = db.Engine.ID(id).Delete(&models.Note{})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to delete note"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "note deleted"})
}
