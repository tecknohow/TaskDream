package v1

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/tecknohow/TaskDream/pkg/db"
	"github.com/tecknohow/TaskDream/pkg/models"
)

func GetCurrentUser(c echo.Context) error {
	userID := c.Get("user_id")
	if userID == nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "unauthorized"})
	}

	user := &models.User{}
	has, err := db.Engine.ID(userID).Get(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "database error"})
	}

	if !has {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "user not found"})
	}

	// Don't return password hash
	user.Password = ""

	return c.JSON(http.StatusOK, user)
}

func UpdateCurrentUser(c echo.Context) error {
	userID := c.Get("user_id")
	if userID == nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "unauthorized"})
	}

	var req struct {
		Name            string `json:"name"`
		AvatarProvider  string `json:"avatar_provider"`
		AvatarFileID    int64  `json:"avatar_file_id"`
	}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	user := &models.User{}
	_, err := db.Engine.ID(userID).Update(&models.User{
		Name:           req.Name,
		AvatarProvider: req.AvatarProvider,
		AvatarFileID:   req.AvatarFileID,
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to update user"})
	}

	has, err := db.Engine.ID(userID).Get(user)
	if err != nil || !has {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to fetch user"})
	}

	user.Password = ""
	return c.JSON(http.StatusOK, user)
}
