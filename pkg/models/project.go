package models

import "time"

type Project struct {
	ID                int64     `xorm:"pk autoincr" json:"id"`
	Title             string    `xorm:"varchar(255)" json:"title"`
	Description       string    `xorm:"text" json:"description"`
	Color             string    `xorm:"varchar(7)" json:"color"`
	IsArchived        bool      `xorm:"default false" json:"is_archived"`
	BackgroundFileID  int64     `json:"background_file_id"`
	Position          int       `json:"position"`
	AreaID            int64     `xorm:"index" json:"area_id"`
	OwnerID           int64     `xorm:"index" json:"owner_id"`
	Created           time.Time `xorm:"created" json:"created"`
	Updated           time.Time `xorm:"updated" json:"updated"`
}

func (Project) TableName() string {
	return "projects"
}
