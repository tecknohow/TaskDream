package v1

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/tecknohow/TaskDream/pkg/db"
	"github.com/tecknohow/TaskDream/pkg/models"
)

func ListBuckets(c echo.Context) error {
	projectID := c.QueryParam("project_id")
	var buckets []models.Bucket

	query := db.Engine
	if projectID != "" {
		id, _ := strconv.ParseInt(projectID, 10, 64)
		query = query.Where("project_id = ?", id)
	}

	err := query.Order("position").Find(&buckets)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "database error"})
	}

	if buckets == nil {
		buckets = []models.Bucket{}
	}

	return c.JSON(http.StatusOK, buckets)
}

func CreateBucket(c echo.Context) error {
	var req struct {
		Title     string `json:"title" validate:"required"`
		ProjectID int64  `json:"project_id" validate:"required"`
		Limit     int    `json:"limit"`
		Position  int    `json:"position"`
	}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	bucket := &models.Bucket{
		Title:     req.Title,
		ProjectID: req.ProjectID,
		Limit:     req.Limit,
		Position:  req.Position,
	}

	_, err := db.Engine.Insert(bucket)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to create bucket"})
	}

	return c.JSON(http.StatusCreated, bucket)
}

func UpdateBucket(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid bucket id"})
	}

	var req struct {
		Title    string `json:"title"`
		Limit    int    `json:"limit"`
		Position int    `json:"position"`
	}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	bucket := &models.Bucket{
		ID:       id,
		Title:    req.Title,
		Limit:    req.Limit,
		Position: req.Position,
	}

	_, err = db.Engine.ID(id).Update(bucket)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to update bucket"})
	}

	return c.JSON(http.StatusOK, bucket)
}

func DeleteBucket(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid bucket id"})
	}

	_, err = db.Engine.ID(id).Delete(&models.Bucket{})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to delete bucket"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "bucket deleted"})
}
