package migration

import (
	"fmt"

	"github.com/tecknohow/TaskDream/pkg/models"
	"xorm.io/xorm"
)

func RunMigrations(engine *xorm.Engine) error {
	// Create tables if they don't exist
	tables := []interface{}{
		&models.User{},
		&models.Area{},
		&models.Project{},
		&models.Bucket{},
		&models.Label{},
		&models.Team{},
		&models.TeamMember{},
		&models.Task{},
		&models.TaskComment{},
		&models.TaskRelation{},
		&models.TimeTracking{},
		&models.Note{},
		&models.PomodoroSession{},
		&models.PomodoroSettings{},
		&models.GithubIntegration{},
		&models.GithubIssueSync{},
		&models.ActivityLog{},
		&models.DailySummary{},
	}

	for _, table := range tables {
		if err := engine.Sync2(table); err != nil {
			return fmt.Errorf("failed to sync table: %w", err)
		}
	}

	return nil
}
