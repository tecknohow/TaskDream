package v1

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/tecknohow/TaskDream/pkg/db"
	"github.com/tecknohow/TaskDream/pkg/models"
)

func ListAreas(c echo.Context) error {
	userID := c.Get("user_id")
	if userID == nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "unauthorized"})
	}

	var areas []models.Area
	err := db.Engine.Where("created_by_id = ?", userID).Find(&areas)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "database error"})
	}

	if areas == nil {
		areas = []models.Area{}
	}

	return c.JSON(http.StatusOK, areas)
}

func CreateArea(c echo.Context) error {
	userID := c.Get("user_id")
	if userID == nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "unauthorized"})
	}

	var req struct {
		Title       string `json:"title" validate:"required"`
		Description string `json:"description"`
		Color       string `json:"color"`
	}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	area := &models.Area{
		Title:       req.Title,
		Description: req.Description,
		Color:       req.Color,
		CreatedByID: userID.(int64),
	}

	_, err := db.Engine.Insert(area)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to create area"})
	}

	return c.JSON(http.StatusCreated, area)
}

func GetArea(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid area id"})
	}

	area := &models.Area{}
	has, err := db.Engine.ID(id).Get(area)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "database error"})
	}

	if !has {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "area not found"})
	}

	return c.JSON(http.StatusOK, area)
}

func UpdateArea(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid area id"})
	}

	var req struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Color       string `json:"color"`
	}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	area := &models.Area{
		ID:          id,
		Title:       req.Title,
		Description: req.Description,
		Color:       req.Color,
	}

	_, err = db.Engine.ID(id).Update(area)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to update area"})
	}

	return c.JSON(http.StatusOK, area)
}

func DeleteArea(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid area id"})
	}

	_, err = db.Engine.ID(id).Delete(&models.Area{})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to delete area"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "area deleted"})
}
