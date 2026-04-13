package models

import "time"

type Team struct {
	ID          int64     `xorm:"pk autoincr" json:"id"`
	Name        string    `xorm:"varchar(255)" json:"name"`
	Description string    `xorm:"text" json:"description"`
	CreatedByID int64     `xorm:"index" json:"created_by_id"`
	Created     time.Time `xorm:"created" json:"created"`
	Updated     time.Time `xorm:"updated" json:"updated"`
}

func (Team) TableName() string {
	return "teams"
}

type TeamMember struct {
	ID     int64 `xorm:"pk autoincr" json:"id"`
	TeamID int64 `xorm:"index" json:"team_id"`
	UserID int64 `xorm:"index" json:"user_id"`
	Admin  bool  `xorm:"default false" json:"admin"`
}

func (TeamMember) TableName() string {
	return "team_members"
}
