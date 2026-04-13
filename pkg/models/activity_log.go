// Vikunja is a to-do list application to facilitate your life.
// Copyright 2018-present Vikunja and contributors. All rights reserved.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package models

import (
	"time"

	"code.vikunja.io/api/pkg/web"
)

// ActivityLog tracks user activity for analytics.
type ActivityLog struct {
	ID         int64  `xorm:"autoincr not null unique pk" json:"id"`
	UserID     int64  `xorm:"bigint not null index" json:"user_id"`
	EntityType string `xorm:"varchar(50) not null" json:"entity_type"`
	EntityID   int64  `xorm:"bigint not null" json:"entity_id"`
	Action     string `xorm:"varchar(50) not null" json:"action"`
	Details    string `xorm:"text" json:"details"`

	Created time.Time `xorm:"created not null" json:"created"`

	web.CRUDable   `xorm:"-" json:"-"`
	web.Permissions `xorm:"-" json:"-"`
}

// TableName holds the table name for activity logs.
func (ActivityLog) TableName() string {
	return "activity_logs"
}

// DailySummary stores precomputed daily productivity summaries.
type DailySummary struct {
	ID                 int64     `xorm:"autoincr not null unique pk" json:"id"`
	UserID             int64     `xorm:"bigint not null index" json:"user_id"`
	Date               time.Time `xorm:"not null" json:"date"`
	TasksCompleted     int       `xorm:"int not null default 0" json:"tasks_completed"`
	TasksCreated       int       `xorm:"int not null default 0" json:"tasks_created"`
	TimeTracked        int64     `xorm:"bigint not null default 0" json:"time_tracked"`
	PomodorosCompleted int       `xorm:"int not null default 0" json:"pomodoros_completed"`
	EstimatedTime      int64     `xorm:"bigint not null default 0" json:"estimated_time"`
	ActualTime         int64     `xorm:"bigint not null default 0" json:"actual_time"`

	Created time.Time `xorm:"created not null" json:"created"`

	web.CRUDable   `xorm:"-" json:"-"`
	web.Permissions `xorm:"-" json:"-"`
}

// TableName holds the table name for daily summaries.
func (DailySummary) TableName() string {
	return "daily_summaries"
}
