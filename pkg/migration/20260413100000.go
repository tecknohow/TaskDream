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

package migration

import (
	"src.techknowlogick.com/xormigrate"
	"xorm.io/xorm"
)

type pomodoroSession20260413100000 struct {
	ID            int64  `xorm:"autoincr not null unique pk"`
	TaskID        int64  `xorm:"bigint not null index"`
	UserID        int64  `xorm:"bigint not null index"`
	Duration      int    `xorm:"int not null default 1500"`
	BreakDuration int    `xorm:"int not null default 300"`
	Status        string `xorm:"varchar(20) not null default 'pending'"`
	StartedAt     int64  `xorm:"bigint null"`
	CompletedAt   int64  `xorm:"bigint null"`
	Created       int64  `xorm:"created not null"`
	Updated       int64  `xorm:"updated not null"`
}

func (pomodoroSession20260413100000) TableName() string {
	return "pomodoro_sessions"
}

type pomodoroSettings20260413100000 struct {
	ID                int64 `xorm:"autoincr not null unique pk"`
	UserID            int64 `xorm:"bigint not null unique"`
	WorkDuration      int   `xorm:"int not null default 1500"`
	ShortBreak        int   `xorm:"int not null default 300"`
	LongBreak         int   `xorm:"int not null default 900"`
	LongBreakInterval int   `xorm:"int not null default 4"`
	AutoStartBreaks   bool  `xorm:"default false"`
	AutoStartPomodoro bool  `xorm:"default false"`
	Created           int64 `xorm:"created not null"`
	Updated           int64 `xorm:"updated not null"`
}

func (pomodoroSettings20260413100000) TableName() string {
	return "pomodoro_settings"
}

type githubIntegration20260413100000 struct {
	ID           int64  `xorm:"autoincr not null unique pk"`
	UserID       int64  `xorm:"bigint not null unique"`
	AccessToken  string `xorm:"text"`
	Username     string `xorm:"varchar(255)"`
	AvatarURL    string `xorm:"varchar(500)"`
	Repos        string `xorm:"text"`
	SyncEnabled  bool   `xorm:"default true"`
	LastSyncedAt int64  `xorm:"bigint null"`
	Created      int64  `xorm:"created not null"`
	Updated      int64  `xorm:"updated not null"`
}

func (githubIntegration20260413100000) TableName() string {
	return "github_integrations"
}

type githubIssueSync20260413100000 struct {
	ID            int64  `xorm:"autoincr not null unique pk"`
	TaskID        int64  `xorm:"bigint not null index"`
	GithubIssueID int64  `xorm:"bigint not null"`
	RepoFullName  string `xorm:"varchar(255) not null"`
	IssueNumber   int    `xorm:"int not null"`
	IssueTitle    string `xorm:"varchar(500)"`
	IssueState    string `xorm:"varchar(20)"`
	IssueURL      string `xorm:"varchar(500)"`
	IsPullRequest bool   `xorm:"default false"`
	LastSyncedAt  int64  `xorm:"bigint null"`
	Created       int64  `xorm:"created not null"`
	Updated       int64  `xorm:"updated not null"`
}

func (githubIssueSync20260413100000) TableName() string {
	return "github_issue_syncs"
}

type activityLog20260413100000 struct {
	ID         int64  `xorm:"autoincr not null unique pk"`
	UserID     int64  `xorm:"bigint not null index"`
	EntityType string `xorm:"varchar(50) not null"`
	EntityID   int64  `xorm:"bigint not null"`
	Action     string `xorm:"varchar(50) not null"`
	Details    string `xorm:"text"`
	Created    int64  `xorm:"created not null"`
}

func (activityLog20260413100000) TableName() string {
	return "activity_logs"
}

type dailySummary20260413100000 struct {
	ID                 int64 `xorm:"autoincr not null unique pk"`
	UserID             int64 `xorm:"bigint not null index"`
	Date               int64 `xorm:"bigint not null"`
	TasksCompleted     int   `xorm:"int not null default 0"`
	TasksCreated       int   `xorm:"int not null default 0"`
	TimeTracked        int64 `xorm:"bigint not null default 0"`
	PomodorosCompleted int   `xorm:"int not null default 0"`
	EstimatedTime      int64 `xorm:"bigint not null default 0"`
	ActualTime         int64 `xorm:"bigint not null default 0"`
	Created            int64 `xorm:"created not null"`
}

func (dailySummary20260413100000) TableName() string {
	return "daily_summaries"
}

func init() {
	migrations = append(migrations, &xormigrate.Migration{
		ID:          "20260413100000",
		Description: "Create Super Productivity tables: pomodoro, github integration, and activity analytics",
		Migrate: func(tx *xorm.Engine) error {
			err := tx.Sync2(
				pomodoroSession20260413100000{},
				pomodoroSettings20260413100000{},
				githubIntegration20260413100000{},
				githubIssueSync20260413100000{},
				activityLog20260413100000{},
				dailySummary20260413100000{},
			)
			return err
		},
		Rollback: func(tx *xorm.Engine) error {
			return tx.DropTables(
				pomodoroSession20260413100000{},
				pomodoroSettings20260413100000{},
				githubIntegration20260413100000{},
				githubIssueSync20260413100000{},
				activityLog20260413100000{},
				dailySummary20260413100000{},
			)
		},
	})
}
