package v1

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/tecknohow/TaskDream/pkg/db"
	"github.com/tecknohow/TaskDream/pkg/models"
)

func ListLabels(c echo.Context) error {
	userID := c.Get("user_id")
	if userID == nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "unauthorized"})
	}

	var labels []models.Label
	err := db.Engine.Where("created_by_id = ?", userID).Find(&labels)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "database error"})
	}

	if labels == nil {
		labels = []models.Label{}
	}

	return c.JSON(http.StatusOK, labels)
}

func CreateLabel(c echo.Context) error {
	userID := c.Get("user_id")
	if userID == nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "unauthorized"})
	}

	var req struct {
		Title       string `json:"title" validate:"required"`
		Description string `json:"description"`
		HexColor    string `json:"hex_color"`
	}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	label := &models.Label{
		Title:       req.Title,
		Description: req.Description,
		HexColor:    req.HexColor,
		CreatedByID: userID.(int64),
	}

	_, err := db.Engine.Insert(label)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to create label"})
	}

	return c.JSON(http.StatusCreated, label)
}

func GetLabel(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid label id"})
	}

	label := &models.Label{}
	has, err := db.Engine.ID(id).Get(label)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "database error"})
	}

	if !has {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "label not found"})
	}

	return c.JSON(http.StatusOK, label)
}

func UpdateLabel(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid label id"})
	}

	var req struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		HexColor    string `json:"hex_color"`
	}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	label := &models.Label{
		ID:          id,
		Title:       req.Title,
		Description: req.Description,
		HexColor:    req.HexColor,
	}

	_, err = db.Engine.ID(id).Update(label)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to update label"})
	}

	return c.JSON(http.StatusOK, label)
}

func DeleteLabel(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid label id"})
	}

	_, err = db.Engine.ID(id).Delete(&models.Label{})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to delete label"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "label deleted"})
}
