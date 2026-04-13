package v1

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/tecknohow/TaskDream/pkg/db"
	"github.com/tecknohow/TaskDream/pkg/models"
)

func GetDashboardStats(c echo.Context) error {
	userID := c.Get("user_id")
	if userID == nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "unauthorized"})
	}

	now := time.Now()
	startOfDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	endOfDay := startOfDay.Add(24 * time.Hour)
	startOfWeek := startOfDay.AddDate(0, 0, -int(startOfDay.Weekday()))

	// Total tasks
	totalTasks, _ := db.Engine.Where("created_by_id = ?", userID).Count(&models.Task{})

	// Open tasks (not done)
	openTasks, _ := db.Engine.Where("created_by_id = ? AND done = false", userID).Count(&models.Task{})

	// Completed tasks
	completedTasks, _ := db.Engine.Where("created_by_id = ? AND done = true", userID).Count(&models.Task{})

	// Tasks completed today
	completedToday, _ := db.Engine.Where("created_by_id = ? AND done = true AND updated >= ? AND updated < ?",
		userID, startOfDay, endOfDay).Count(&models.Task{})

	// Tasks completed this week
	completedThisWeek, _ := db.Engine.Where("created_by_id = ? AND done = true AND updated >= ?",
		userID, startOfWeek).Count(&models.Task{})

	// Overdue tasks
	overdueTasks, _ := db.Engine.Where("created_by_id = ? AND done = false AND due_date IS NOT NULL AND due_date < ?",
		userID, now).Count(&models.Task{})

	// Tasks due today
	dueTodayTasks, _ := db.Engine.Where("created_by_id = ? AND done = false AND due_date >= ? AND due_date < ?",
		userID, startOfDay, endOfDay).Count(&models.Task{})

	// Time tracked today
	var todayTrackings []models.TimeTracking
	db.Engine.Where("user_id = ? AND start >= ? AND start < ?", userID, startOfDay, endOfDay).Find(&todayTrackings)
	var todayTimeTracked int64
	for _, t := range todayTrackings {
		todayTimeTracked += t.Duration
	}

	// Time tracked this week
	var weekTrackings []models.TimeTracking
	db.Engine.Where("user_id = ? AND start >= ?", userID, startOfWeek).Find(&weekTrackings)
	var weekTimeTracked int64
	for _, t := range weekTrackings {
		weekTimeTracked += t.Duration
	}

	// Pomodoros completed today
	pomodorosToday, _ := db.Engine.Where("user_id = ? AND status = 'completed' AND started_at >= ? AND started_at < ?",
		userID, startOfDay, endOfDay).Count(&models.PomodoroSession{})

	// Active projects
	activeProjects, _ := db.Engine.Where("owner_id = ? AND is_archived = false", userID).Count(&models.Project{})

	// Estimation accuracy (for completed tasks that had estimates)
	var completedWithEstimates []models.Task
	db.Engine.Where("created_by_id = ? AND done = true AND estimated_time > 0 AND total_time_spent > 0",
		userID).Find(&completedWithEstimates)

	var totalEstimated, totalActual int64
	for _, t := range completedWithEstimates {
		totalEstimated += t.EstimatedTime
		totalActual += t.TotalTimeSpent
	}

	estimationAccuracy := float64(0)
	if totalEstimated > 0 && len(completedWithEstimates) > 0 {
		estimationAccuracy = float64(totalActual) / float64(totalEstimated) * 100
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"total_tasks":          totalTasks,
		"open_tasks":           openTasks,
		"completed_tasks":      completedTasks,
		"completed_today":      completedToday,
		"completed_this_week":  completedThisWeek,
		"overdue_tasks":        overdueTasks,
		"due_today":            dueTodayTasks,
		"today_time_tracked":   todayTimeTracked,
		"week_time_tracked":    weekTimeTracked,
		"pomodoros_today":      pomodorosToday,
		"active_projects":      activeProjects,
		"estimation_accuracy":  estimationAccuracy,
		"tasks_with_estimates": len(completedWithEstimates),
	})
}

func GetProductivityTrend(c echo.Context) error {
	userID := c.Get("user_id")
	if userID == nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "unauthorized"})
	}

	daysStr := c.QueryParam("days")
	days := 7
	if daysStr != "" {
		if d, err := time.ParseDuration(daysStr); err == nil {
			days = int(d.Hours() / 24)
		}
	}
	if days > 90 {
		days = 90
	}
	if days < 1 {
		days = 7
	}

	now := time.Now()
	type DayData struct {
		Date           string `json:"date"`
		TasksCompleted int    `json:"tasks_completed"`
		TimeTracked    int64  `json:"time_tracked"`
		Pomodoros      int    `json:"pomodoros"`
	}

	trend := make([]DayData, days)
	for i := 0; i < days; i++ {
		day := now.AddDate(0, 0, -(days - 1 - i))
		startOfDay := time.Date(day.Year(), day.Month(), day.Day(), 0, 0, 0, 0, day.Location())
		endOfDay := startOfDay.Add(24 * time.Hour)

		completed, _ := db.Engine.Where("created_by_id = ? AND done = true AND updated >= ? AND updated < ?",
			userID, startOfDay, endOfDay).Count(&models.Task{})

		var trackings []models.TimeTracking
		db.Engine.Where("user_id = ? AND start >= ? AND start < ?", userID, startOfDay, endOfDay).Find(&trackings)
		var timeTracked int64
		for _, t := range trackings {
			timeTracked += t.Duration
		}

		pomodoros, _ := db.Engine.Where("user_id = ? AND status = 'completed' AND started_at >= ? AND started_at < ?",
			userID, startOfDay, endOfDay).Count(&models.PomodoroSession{})

		trend[i] = DayData{
			Date:           startOfDay.Format("2006-01-02"),
			TasksCompleted: int(completed),
			TimeTracked:    timeTracked,
			Pomodoros:      int(pomodoros),
		}
	}

	return c.JSON(http.StatusOK, trend)
}
