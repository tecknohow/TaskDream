package models

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

type Task struct {
	ID           int64          `xorm:"pk autoincr" json:"id"`
	Title        string         `xorm:"varchar(255)" json:"title"`
	Description  string         `xorm:"text" json:"description"`
	Done         bool           `xorm:"default false" json:"done"`
	Priority     int            `xorm:"default 0" json:"priority"` // 1-5
	DueDate      *time.Time     `json:"due_date"`
	ProjectID    int64          `xorm:"index" json:"project_id"`
	BucketID     int64          `xorm:"index" json:"bucket_id"`
	Labels       Labels         `xorm:"text" json:"labels"`
	Reminders    Reminders      `xorm:"text" json:"reminders"`
	RepeatAfter  string         `xorm:"varchar(255)" json:"repeat_after"`
	Attachments  Attachments    `xorm:"text" json:"attachments"`
	Position     int            `json:"position"`
	PercentDone  int            `xorm:"default 0" json:"percent_done"`
	StartDate    *time.Time     `json:"start_date"`
	EndDate      *time.Time     `json:"end_date"`
	CreatedByID  int64          `xorm:"index" json:"created_by_id"`
	Created      time.Time      `xorm:"created" json:"created"`
	Updated      time.Time      `xorm:"updated" json:"updated"`
}

func (Task) TableName() string {
	return "tasks"
}

// Labels is a JSON array of label IDs
type Labels []int64

func (l Labels) Value() (driver.Value, error) {
	return json.Marshal(l)
}

func (l *Labels) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), &l)
}

// Reminder represents a task reminder
type Reminder struct {
	RelativeTo string    `json:"relative_to"`
	Duration   int64     `json:"duration"` // in seconds
	Reminder   time.Time `json:"reminder"`
}

type Reminders []Reminder

func (r Reminders) Value() (driver.Value, error) {
	return json.Marshal(r)
}

func (r *Reminders) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), &r)
}

// Attachment represents a file attachment
type Attachment struct {
	ID        int64  `json:"id"`
	Filename  string `json:"filename"`
	FileSize  int64  `json:"file_size"`
	CreatedAt time.Time `json:"created_at"`
}

type Attachments []Attachment

func (a Attachments) Value() (driver.Value, error) {
	return json.Marshal(a)
}

func (a *Attachments) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), &a)
}
