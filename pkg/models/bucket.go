package models

import "time"

// Bucket represents a Kanban column
type Bucket struct {
	ID        int64     `xorm:"pk autoincr" json:"id"`
	Title     string    `xorm:"varchar(255)" json:"title"`
	ProjectID int64     `xorm:"index" json:"project_id"`
	Limit     int       `json:"limit"`
	Position  int       `json:"position"`
	Created   time.Time `xorm:"created" json:"created"`
	Updated   time.Time `xorm:"updated" json:"updated"`
}

func (Bucket) TableName() string {
	return "buckets"
}
