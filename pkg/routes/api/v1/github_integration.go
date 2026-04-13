package v1

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/tecknohow/TaskDream/pkg/db"
	"github.com/tecknohow/TaskDream/pkg/models"
)

func GetGithubIntegration(c echo.Context) error {
	userID := c.Get("user_id")
	if userID == nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "unauthorized"})
	}

	integration := &models.GithubIntegration{}
	has, err := db.Engine.Where("user_id = ?", userID).Get(integration)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "database error"})
	}

	if !has {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"connected": false,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"connected":      true,
		"username":       integration.Username,
		"avatar_url":     integration.AvatarURL,
		"repos":          integration.Repos,
		"sync_enabled":   integration.SyncEnabled,
		"last_synced_at": integration.LastSyncedAt,
	})
}

func SaveGithubIntegration(c echo.Context) error {
	userID := c.Get("user_id")
	if userID == nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "unauthorized"})
	}

	var req struct {
		AccessToken string `json:"access_token" validate:"required"`
		Username    string `json:"username"`
		AvatarURL   string `json:"avatar_url"`
	}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	// Upsert
	existing := &models.GithubIntegration{}
	has, _ := db.Engine.Where("user_id = ?", userID).Get(existing)

	integration := &models.GithubIntegration{
		UserID:      userID.(int64),
		AccessToken: req.AccessToken,
		Username:    req.Username,
		AvatarURL:   req.AvatarURL,
		SyncEnabled: true,
	}

	if has {
		_, err := db.Engine.Where("user_id = ?", userID).Update(integration)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to update integration"})
		}
	} else {
		_, err := db.Engine.Insert(integration)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to create integration"})
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"connected":    true,
		"username":     integration.Username,
		"sync_enabled": integration.SyncEnabled,
	})
}

func DeleteGithubIntegration(c echo.Context) error {
	userID := c.Get("user_id")
	if userID == nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "unauthorized"})
	}

	_, err := db.Engine.Where("user_id = ?", userID).Delete(&models.GithubIntegration{})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to delete integration"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "integration removed"})
}

func UpdateGithubRepos(c echo.Context) error {
	userID := c.Get("user_id")
	if userID == nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "unauthorized"})
	}

	var req struct {
		Repos models.RepoList `json:"repos"`
	}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	integration := &models.GithubIntegration{
		Repos: req.Repos,
	}

	_, err := db.Engine.Where("user_id = ?", userID).Cols("repos").Update(integration)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to update repos"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "repos updated"})
}

func ListGithubIssueSyncs(c echo.Context) error {
	taskIDStr := c.QueryParam("task_id")
	repoName := c.QueryParam("repo")

	var syncs []models.GithubIssueSync
	query := db.Engine.NewSession()

	if taskIDStr != "" {
		taskID, _ := strconv.ParseInt(taskIDStr, 10, 64)
		query = query.Where("task_id = ?", taskID)
	}

	if repoName != "" {
		query = query.Where("repo_full_name = ?", repoName)
	}

	err := query.Find(&syncs)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "database error"})
	}

	if syncs == nil {
		syncs = []models.GithubIssueSync{}
	}

	return c.JSON(http.StatusOK, syncs)
}
