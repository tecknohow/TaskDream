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
	"database/sql/driver"
	"encoding/json"
	"time"

	"code.vikunja.io/api/pkg/web"
)

// GithubIntegration stores a user's GitHub integration configuration.
type GithubIntegration struct {
	ID           int64      `xorm:"autoincr not null unique pk" json:"id"`
	UserID       int64      `xorm:"bigint not null unique" json:"user_id"`
	AccessToken  string     `xorm:"text" json:"-"`
	Username     string     `xorm:"varchar(255)" json:"username"`
	AvatarURL    string     `xorm:"varchar(500)" json:"avatar_url"`
	Repos        RepoList   `xorm:"text" json:"repos"`
	SyncEnabled  bool       `xorm:"default true" json:"sync_enabled"`
	LastSyncedAt *time.Time `json:"last_synced_at"`

	Created time.Time `xorm:"created not null" json:"created"`
	Updated time.Time `xorm:"updated not null" json:"updated"`

	web.CRUDable   `xorm:"-" json:"-"`
	web.Permissions `xorm:"-" json:"-"`
}

// TableName holds the table name for github integrations.
func (GithubIntegration) TableName() string {
	return "github_integrations"
}

// RepoConfig represents a tracked repository.
type RepoConfig struct {
	Owner      string `json:"owner"`
	Name       string `json:"name"`
	FullName   string `json:"full_name"`
	ProjectID  int64  `json:"project_id"`
	SyncIssues bool   `json:"sync_issues"`
	SyncPRs    bool   `json:"sync_prs"`
}

// RepoList is a JSON-serialised slice of RepoConfig stored in a single column.
type RepoList []RepoConfig

// Value implements the driver.Valuer interface for RepoList.
func (r RepoList) Value() (driver.Value, error) {
	return json.Marshal(r)
}

// Scan implements the sql.Scanner interface for RepoList.
func (r *RepoList) Scan(value interface{}) error {
	if value == nil {
		*r = RepoList{}
		return nil
	}
	return json.Unmarshal(value.([]byte), r)
}

// GithubIssueSync tracks which GitHub issues are linked to tasks.
type GithubIssueSync struct {
	ID            int64     `xorm:"autoincr not null unique pk" json:"id"`
	TaskID        int64     `xorm:"bigint not null index" json:"task_id" param:"task"`
	GithubIssueID int64     `xorm:"bigint not null" json:"github_issue_id"`
	RepoFullName  string    `xorm:"varchar(255) not null" json:"repo_full_name"`
	IssueNumber   int       `xorm:"int not null" json:"issue_number"`
	IssueTitle    string    `xorm:"varchar(500)" json:"issue_title"`
	IssueState    string    `xorm:"varchar(20)" json:"issue_state"`
	IssueURL      string    `xorm:"varchar(500)" json:"issue_url"`
	IsPullRequest bool      `json:"is_pull_request"`
	LastSyncedAt  time.Time `json:"last_synced_at"`

	Created time.Time `xorm:"created not null" json:"created"`
	Updated time.Time `xorm:"updated not null" json:"updated"`

	web.CRUDable   `xorm:"-" json:"-"`
	web.Permissions `xorm:"-" json:"-"`
}

// TableName holds the table name for github issue syncs.
func (GithubIssueSync) TableName() string {
	return "github_issue_syncs"
}
