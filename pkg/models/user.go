package models

import "time"

type User struct {
	ID              int64     `xorm:"pk autoincr" json:"id"`
	Username        string    `xorm:"varchar(255) unique" json:"username"`
	Email           string    `xorm:"varchar(255) unique" json:"email"`
	Password        string    `xorm:"varchar(255)" json:"-"`
	Name            string    `xorm:"varchar(255)" json:"name"`
	AvatarProvider  string    `xorm:"varchar(50)" json:"avatar_provider"`
	AvatarFileID    int64     `json:"avatar_file_id"`
	IsActive        bool      `xorm:"default true" json:"is_active"`
	Created         time.Time `xorm:"created" json:"created"`
	Updated         time.Time `xorm:"updated" json:"updated"`
}

func (User) TableName() string {
	return "users"
}
