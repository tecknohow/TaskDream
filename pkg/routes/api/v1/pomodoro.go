package v1

import (
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/tecknohow/TaskDream/pkg/db"
	"github.com/tecknohow/TaskDream/pkg/models"
)

func ListPomodoroSessions(c echo.Context) error {
	userID := c.Get("user_id")
	if userID == nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "unauthorized"})
	}

	taskIDStr := c.QueryParam("task_id")
	dateStr := c.QueryParam("date")

	var sessions []models.PomodoroSession
	query := db.Engine.Where("user_id = ?", userID)

	if taskIDStr != "" {
		taskID, _ := strconv.ParseInt(taskIDStr, 10, 64)
		query = query.And("task_id = ?", taskID)
	}

	if dateStr != "" {
		date, err := time.Parse("2006-01-02", dateStr)
		if err == nil {
			nextDay := date.Add(24 * time.Hour)
			query = query.And("started_at >= ? AND started_at < ?", date, nextDay)
		}
	}

	err := query.Desc("created").Find(&sessions)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "database error"})
	}

	if sessions == nil {
		sessions = []models.PomodoroSession{}
	}

	return c.JSON(http.StatusOK, sessions)
}

func StartPomodoro(c echo.Context) error {
	userID := c.Get("user_id")
	if userID == nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "unauthorized"})
	}

	var req struct {
		TaskID   int64 `json:"task_id"`
		Duration int   `json:"duration"` // in seconds, 0 means use settings default
	}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	// Get user's pomodoro settings
	settings := &models.PomodoroSettings{}
	has, _ := db.Engine.Where("user_id = ?", userID).Get(settings)
	if !has {
		settings.WorkDuration = 1500 // 25 min default
		settings.ShortBreak = 300
	}

	duration := req.Duration
	if duration == 0 {
		duration = settings.WorkDuration
	}

	now := time.Now()
	session := &models.PomodoroSession{
		TaskID:        req.TaskID,
		UserID:        userID.(int64),
		Duration:      duration,
		BreakDuration: settings.ShortBreak,
		Status:        "running",
		StartedAt:     &now,
	}

	_, err := db.Engine.Insert(session)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to start pomodoro"})
	}

	return c.JSON(http.StatusCreated, session)
}

func CompletePomodoro(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid session id"})
	}

	now := time.Now()
	session := &models.PomodoroSession{
		Status:      "completed",
		CompletedAt: &now,
	}

	_, err = db.Engine.ID(id).Cols("status", "completed_at").Update(session)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to complete pomodoro"})
	}

	// Fetch the updated session
	updated := &models.PomodoroSession{}
	db.Engine.ID(id).Get(updated)

	return c.JSON(http.StatusOK, updated)
}

func CancelPomodoro(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid session id"})
	}

	session := &models.PomodoroSession{
		Status: "cancelled",
	}

	_, err = db.Engine.ID(id).Cols("status").Update(session)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to cancel pomodoro"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "pomodoro cancelled"})
}

func GetPomodoroSettings(c echo.Context) error {
	userID := c.Get("user_id")
	if userID == nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "unauthorized"})
	}

	settings := &models.PomodoroSettings{}
	has, err := db.Engine.Where("user_id = ?", userID).Get(settings)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "database error"})
	}

	if !has {
		// Return defaults
		settings = &models.PomodoroSettings{
			UserID:            userID.(int64),
			WorkDuration:      1500,
			ShortBreak:        300,
			LongBreak:         900,
			LongBreakInterval: 4,
		}
	}

	return c.JSON(http.StatusOK, settings)
}

func UpdatePomodoroSettings(c echo.Context) error {
	userID := c.Get("user_id")
	if userID == nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "unauthorized"})
	}

	var req struct {
		WorkDuration      int  `json:"work_duration"`
		ShortBreak        int  `json:"short_break"`
		LongBreak         int  `json:"long_break"`
		LongBreakInterval int  `json:"long_break_interval"`
		AutoStartBreaks   bool `json:"auto_start_breaks"`
		AutoStartPomodoro bool `json:"auto_start_pomodoro"`
	}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	// Upsert
	existing := &models.PomodoroSettings{}
	has, _ := db.Engine.Where("user_id = ?", userID).Get(existing)

	settings := &models.PomodoroSettings{
		UserID:            userID.(int64),
		WorkDuration:      req.WorkDuration,
		ShortBreak:        req.ShortBreak,
		LongBreak:         req.LongBreak,
		LongBreakInterval: req.LongBreakInterval,
		AutoStartBreaks:   req.AutoStartBreaks,
		AutoStartPomodoro: req.AutoStartPomodoro,
	}

	if has {
		_, err := db.Engine.Where("user_id = ?", userID).Update(settings)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to update settings"})
		}
	} else {
		_, err := db.Engine.Insert(settings)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to create settings"})
		}
	}

	return c.JSON(http.StatusOK, settings)
}

func GetPomodoroStats(c echo.Context) error {
	userID := c.Get("user_id")
	if userID == nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "unauthorized"})
	}

	dateStr := c.QueryParam("date")
	date := time.Now()
	if dateStr != "" {
		parsed, err := time.Parse("2006-01-02", dateStr)
		if err == nil {
			date = parsed
		}
	}

	startOfDay := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
	endOfDay := startOfDay.Add(24 * time.Hour)

	// Count today's completed pomodoros
	todayCompleted, _ := db.Engine.Where("user_id = ? AND status = 'completed' AND started_at >= ? AND started_at < ?",
		userID, startOfDay, endOfDay).Count(&models.PomodoroSession{})

	// Count total time in completed pomodoros today
	var sessions []models.PomodoroSession
	db.Engine.Where("user_id = ? AND status = 'completed' AND started_at >= ? AND started_at < ?",
		userID, startOfDay, endOfDay).Find(&sessions)

	totalSeconds := 0
	for _, s := range sessions {
		totalSeconds += s.Duration
	}

	// Weekly count
	startOfWeek := startOfDay.AddDate(0, 0, -int(startOfDay.Weekday()))
	weeklyCompleted, _ := db.Engine.Where("user_id = ? AND status = 'completed' AND started_at >= ?",
		userID, startOfWeek).Count(&models.PomodoroSession{})

	return c.JSON(http.StatusOK, map[string]interface{}{
		"today_completed":  todayCompleted,
		"today_total_time": totalSeconds,
		"weekly_completed": weeklyCompleted,
	})
}
