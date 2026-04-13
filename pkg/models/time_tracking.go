package models

import "time"

// TimeTracking represents time tracking entries (inspired by Super Productivity)
type TimeTracking struct {
	ID        int64      `xorm:"pk autoincr" json:"id"`
	TaskID    int64      `xorm:"index" json:"task_id"`
	UserID    int64      `xorm:"index" json:"user_id"`
	Start     time.Time  `json:"start"`
	End       *time.Time `json:"end"`
	Duration  int64      `json:"duration"` // in seconds
	Comment   string     `xorm:"text" json:"comment"`
	Created   time.Time  `xorm:"created" json:"created"`
	Updated   time.Time  `xorm:"updated" json:"updated"`
}

func (TimeTracking) TableName() string {
	return "time_tracking"
}
