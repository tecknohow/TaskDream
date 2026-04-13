package v1

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/tecknohow/TaskDream/pkg/db"
	"github.com/tecknohow/TaskDream/pkg/models"
)

func ListProjects(c echo.Context) error {
	userID := c.Get("user_id")
	if userID == nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "unauthorized"})
	}

	var projects []models.Project
	err := db.Engine.Where("owner_id = ?", userID).Find(&projects)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "database error"})
	}

	if projects == nil {
		projects = []models.Project{}
	}

	return c.JSON(http.StatusOK, projects)
}

func CreateProject(c echo.Context) error {
	userID := c.Get("user_id")
	if userID == nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "unauthorized"})
	}

	var req struct {
		Title             string `json:"title" validate:"required"`
		Description       string `json:"description"`
		Color             string `json:"color"`
		BackgroundFileID  int64  `json:"background_file_id"`
		AreaID            int64  `json:"area_id"`
	}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	project := &models.Project{
		Title:            req.Title,
		Description:      req.Description,
		Color:            req.Color,
		BackgroundFileID: req.BackgroundFileID,
		AreaID:           req.AreaID,
		OwnerID:          userID.(int64),
		IsArchived:       false,
	}

	_, err := db.Engine.Insert(project)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to create project"})
	}

	return c.JSON(http.StatusCreated, project)
}

func GetProject(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid project id"})
	}

	project := &models.Project{}
	has, err := db.Engine.ID(id).Get(project)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "database error"})
	}

	if !has {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "project not found"})
	}

	return c.JSON(http.StatusOK, project)
}

func UpdateProject(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid project id"})
	}

	var req struct {
		Title             string `json:"title"`
		Description       string `json:"description"`
		Color             string `json:"color"`
		IsArchived        bool   `json:"is_archived"`
		BackgroundFileID  int64  `json:"background_file_id"`
		Position          int    `json:"position"`
		AreaID            int64  `json:"area_id"`
	}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	project := &models.Project{
		ID:               id,
		Title:            req.Title,
		Description:      req.Description,
		Color:            req.Color,
		IsArchived:       req.IsArchived,
		BackgroundFileID: req.BackgroundFileID,
		Position:         req.Position,
		AreaID:           req.AreaID,
	}

	_, err = db.Engine.ID(id).Update(project)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to update project"})
	}

	return c.JSON(http.StatusOK, project)
}

func DeleteProject(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid project id"})
	}

	_, err = db.Engine.ID(id).Delete(&models.Project{})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to delete project"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "project deleted"})
}
