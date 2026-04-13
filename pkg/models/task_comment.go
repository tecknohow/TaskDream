package models

import "time"

type TaskComment struct {
	ID        int64     `xorm:"pk autoincr" json:"id"`
	TaskID    int64     `xorm:"index" json:"task_id"`
	UserID    int64     `xorm:"index" json:"user_id"`
	Comment   string    `xorm:"text" json:"comment"`
	Created   time.Time `xorm:"created" json:"created"`
	Updated   time.Time `xorm:"updated" json:"updated"`
}

func (TaskComment) TableName() string {
	return "task_comments"
}
