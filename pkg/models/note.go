package models

import "time"

// Note represents a note (inspired by Tududi)
type Note struct {
	ID          int64     `xorm:"pk autoincr" json:"id"`
	Title       string    `xorm:"varchar(255)" json:"title"`
	Content     string    `xorm:"text" json:"content"`
	ProjectID   int64     `xorm:"index" json:"project_id"`
	CreatedByID int64     `xorm:"index" json:"created_by_id"`
	Created     time.Time `xorm:"created" json:"created"`
	Updated     time.Time `xorm:"updated" json:"updated"`
}

func (Note) TableName() string {
	return "notes"
}
