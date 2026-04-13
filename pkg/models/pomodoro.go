package models

import "time"

// PomodoroSession represents a pomodoro timer session
type PomodoroSession struct {
	ID            int64      `xorm:"pk autoincr" json:"id"`
	TaskID        int64      `xorm:"index" json:"task_id"`
	UserID        int64      `xorm:"index" json:"user_id"`
	Duration      int        `xorm:"default 1500" json:"duration"`      // in seconds (default 25 min)
	BreakDuration int        `xorm:"default 300" json:"break_duration"` // in seconds (default 5 min)
	Status        string     `xorm:"varchar(20) default 'pending'" json:"status"` // pending, running, completed, cancelled
	StartedAt     *time.Time `json:"started_at"`
	CompletedAt   *time.Time `json:"completed_at"`
	Created       time.Time  `xorm:"created" json:"created"`
	Updated       time.Time  `xorm:"updated" json:"updated"`
}

func (PomodoroSession) TableName() string {
	return "pomodoro_sessions"
}

// PomodoroSettings represents user-specific pomodoro settings
type PomodoroSettings struct {
	ID                int64 `xorm:"pk autoincr" json:"id"`
	UserID            int64 `xorm:"unique" json:"user_id"`
	WorkDuration      int   `xorm:"default 1500" json:"work_duration"`       // 25 min default
	ShortBreak        int   `xorm:"default 300" json:"short_break"`          // 5 min default
	LongBreak         int   `xorm:"default 900" json:"long_break"`           // 15 min default
	LongBreakInterval int   `xorm:"default 4" json:"long_break_interval"`    // every 4 sessions
	AutoStartBreaks   bool  `xorm:"default false" json:"auto_start_breaks"`
	AutoStartPomodoro bool  `xorm:"default false" json:"auto_start_pomodoro"`
	Created           time.Time `xorm:"created" json:"created"`
	Updated           time.Time `xorm:"updated" json:"updated"`
}

func (PomodoroSettings) TableName() string {
	return "pomodoro_settings"
}
