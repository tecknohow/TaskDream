package models

import "time"

// ActivityLog tracks user activity for analytics
type ActivityLog struct {
	ID         int64     `xorm:"pk autoincr" json:"id"`
	UserID     int64     `xorm:"index" json:"user_id"`
	EntityType string    `xorm:"varchar(50)" json:"entity_type"` // task, project, note, etc.
	EntityID   int64     `json:"entity_id"`
	Action     string    `xorm:"varchar(50)" json:"action"` // created, updated, completed, deleted
	Details    string    `xorm:"text" json:"details"`
	Created    time.Time `xorm:"created" json:"created"`
}

func (ActivityLog) TableName() string {
	return "activity_logs"
}

// DailySummary stores precomputed daily productivity summaries
type DailySummary struct {
	ID               int64     `xorm:"pk autoincr" json:"id"`
	UserID           int64     `xorm:"index" json:"user_id"`
	Date             time.Time `json:"date"`
	TasksCompleted   int       `json:"tasks_completed"`
	TasksCreated     int       `json:"tasks_created"`
	TimeTracked      int64     `json:"time_tracked"`      // in seconds
	PomodorosCompleted int     `json:"pomodoros_completed"`
	EstimatedTime    int64     `json:"estimated_time"`     // sum of estimates for completed tasks
	ActualTime       int64     `json:"actual_time"`        // sum of actual time for completed tasks
	Created          time.Time `xorm:"created" json:"created"`
}

func (DailySummary) TableName() string {
	return "daily_summaries"
}
