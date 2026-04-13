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

	"code.vikunja.io/api/pkg/user"
	"code.vikunja.io/api/pkg/web"
)

// PomodoroSession represents a pomodoro timer session linked to a task.
type PomodoroSession struct {
	ID            int64      `xorm:"autoincr not null unique pk" json:"id"`
	TaskID        int64      `xorm:"bigint not null index" json:"task_id" param:"task"`
	UserID        int64      `xorm:"bigint not null index" json:"-"`
	Duration      int        `xorm:"int not null default 1500" json:"duration"`
	BreakDuration int        `xorm:"int not null default 300" json:"break_duration"`
	Status        string     `xorm:"varchar(20) not null default 'pending'" json:"status"`
	StartedAt     *time.Time `json:"started_at"`
	CompletedAt   *time.Time `json:"completed_at"`

	// Read-only, resolved from UserID
	User *user.User `xorm:"-" json:"user"`

	Created time.Time `xorm:"created not null" json:"created"`
	Updated time.Time `xorm:"updated not null" json:"updated"`

	web.CRUDable   `xorm:"-" json:"-"`
	web.Permissions `xorm:"-" json:"-"`
}

// TableName holds the table name for pomodoro sessions.
func (PomodoroSession) TableName() string {
	return "pomodoro_sessions"
}

// PomodoroSettings represents user-specific pomodoro configuration.
type PomodoroSettings struct {
	ID                int64 `xorm:"autoincr not null unique pk" json:"id"`
	UserID            int64 `xorm:"bigint not null unique" json:"user_id"`
	WorkDuration      int   `xorm:"int not null default 1500" json:"work_duration"`
	ShortBreak        int   `xorm:"int not null default 300" json:"short_break"`
	LongBreak         int   `xorm:"int not null default 900" json:"long_break"`
	LongBreakInterval int   `xorm:"int not null default 4" json:"long_break_interval"`
	AutoStartBreaks   bool  `xorm:"default false" json:"auto_start_breaks"`
	AutoStartPomodoro bool  `xorm:"default false" json:"auto_start_pomodoro"`

	Created time.Time `xorm:"created not null" json:"created"`
	Updated time.Time `xorm:"updated not null" json:"updated"`

	web.CRUDable   `xorm:"-" json:"-"`
	web.Permissions `xorm:"-" json:"-"`
}

// TableName holds the table name for pomodoro settings.
func (PomodoroSettings) TableName() string {
	return "pomodoro_settings"
}
