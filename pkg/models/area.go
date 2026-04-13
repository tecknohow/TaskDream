package models

import "time"

// Area represents a collection of projects (inspired by Tududi)
type Area struct {
	ID          int64     `xorm:"pk autoincr" json:"id"`
	Title       string    `xorm:"varchar(255)" json:"title"`
	Description string    `xorm:"text" json:"description"`
	Color       string    `xorm:"varchar(7)" json:"color"` // Hex color code
	CreatedByID int64     `xorm:"index" json:"created_by_id"`
	Created     time.Time `xorm:"created" json:"created"`
	Updated     time.Time `xorm:"updated" json:"updated"`
}

func (Area) TableName() string {
	return "areas"
}
