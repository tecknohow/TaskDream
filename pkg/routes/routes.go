package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/tecknohow/TaskDream/pkg/routes/api/v1"
)

func SetupRoutes(e *echo.Echo) {
	// Health check
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
	})

	// API v1 routes
	api := e.Group("/api/v1")

	// Auth routes (public)
	api.POST("/auth/login", v1.Login)
	api.POST("/auth/register", v1.Register)
	api.POST("/auth/refresh", v1.RefreshToken)

	// Protected routes
	api.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte("your-secret-key"),
	}))

	// User routes
	api.GET("/users/me", v1.GetCurrentUser)
	api.PUT("/users/me", v1.UpdateCurrentUser)

	// Project routes
	api.GET("/projects", v1.ListProjects)
	api.POST("/projects", v1.CreateProject)
	api.GET("/projects/:id", v1.GetProject)
	api.PUT("/projects/:id", v1.UpdateProject)
	api.DELETE("/projects/:id", v1.DeleteProject)

	// Task routes
	api.GET("/tasks", v1.ListTasks)
	api.POST("/tasks", v1.CreateTask)
	api.GET("/tasks/:id", v1.GetTask)
	api.PUT("/tasks/:id", v1.UpdateTask)
	api.DELETE("/tasks/:id", v1.DeleteTask)

	// Subtask routes
	api.GET("/tasks/:id/subtasks", v1.ListSubtasks)
	api.POST("/tasks/:id/subtasks", v1.CreateSubtask)

	// Task comment routes
	api.GET("/tasks/:id/comments", v1.ListTaskComments)
	api.POST("/tasks/:id/comments", v1.CreateTaskComment)
	api.DELETE("/comments/:id", v1.DeleteTaskComment)

	// Label routes
	api.GET("/labels", v1.ListLabels)
	api.POST("/labels", v1.CreateLabel)
	api.GET("/labels/:id", v1.GetLabel)
	api.PUT("/labels/:id", v1.UpdateLabel)
	api.DELETE("/labels/:id", v1.DeleteLabel)

	// Bucket routes
	api.GET("/buckets", v1.ListBuckets)
	api.POST("/buckets", v1.CreateBucket)
	api.PUT("/buckets/:id", v1.UpdateBucket)
	api.DELETE("/buckets/:id", v1.DeleteBucket)

	// Area routes
	api.GET("/areas", v1.ListAreas)
	api.POST("/areas", v1.CreateArea)
	api.GET("/areas/:id", v1.GetArea)
	api.PUT("/areas/:id", v1.UpdateArea)
	api.DELETE("/areas/:id", v1.DeleteArea)

	// Time tracking routes
	api.GET("/tasks/:id/time-tracking", v1.ListTimeTracking)
	api.POST("/tasks/:id/time-tracking", v1.CreateTimeTracking)
	api.PUT("/time-tracking/:id", v1.UpdateTimeTracking)
	api.DELETE("/time-tracking/:id", v1.DeleteTimeTracking)

	// Note routes
	api.GET("/notes", v1.ListNotes)
	api.POST("/notes", v1.CreateNote)
	api.GET("/notes/:id", v1.GetNote)
	api.PUT("/notes/:id", v1.UpdateNote)
	api.DELETE("/notes/:id", v1.DeleteNote)

	// Pomodoro routes
	api.GET("/pomodoro/sessions", v1.ListPomodoroSessions)
	api.POST("/pomodoro/start", v1.StartPomodoro)
	api.POST("/pomodoro/:id/complete", v1.CompletePomodoro)
	api.POST("/pomodoro/:id/cancel", v1.CancelPomodoro)
	api.GET("/pomodoro/settings", v1.GetPomodoroSettings)
	api.PUT("/pomodoro/settings", v1.UpdatePomodoroSettings)
	api.GET("/pomodoro/stats", v1.GetPomodoroStats)

	// Search routes
	api.GET("/search", v1.SearchTasks)

	// Analytics routes
	api.GET("/analytics/dashboard", v1.GetDashboardStats)
	api.GET("/analytics/trend", v1.GetProductivityTrend)

	// GitHub integration routes
	api.GET("/integrations/github", v1.GetGithubIntegration)
	api.POST("/integrations/github", v1.SaveGithubIntegration)
	api.DELETE("/integrations/github", v1.DeleteGithubIntegration)
	api.PUT("/integrations/github/repos", v1.UpdateGithubRepos)
	api.GET("/integrations/github/issues", v1.ListGithubIssueSyncs)
}
