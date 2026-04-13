package models

import "time"

type Label struct {
	ID          int64     `xorm:"pk autoincr" json:"id"`
	Title       string    `xorm:"varchar(255)" json:"title"`
	Description string    `xorm:"text" json:"description"`
	HexColor    string    `xorm:"varchar(7)" json:"hex_color"`
	CreatedByID int64     `xorm:"index" json:"created_by_id"`
	Created     time.Time `xorm:"created" json:"created"`
	Updated     time.Time `xorm:"updated" json:"updated"`
}

func (Label) TableName() string {
	return "labels"
}
