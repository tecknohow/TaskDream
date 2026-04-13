package models

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

// GithubIntegration stores a user's GitHub integration configuration
type GithubIntegration struct {
	ID           int64     `xorm:"pk autoincr" json:"id"`
	UserID       int64     `xorm:"unique" json:"user_id"`
	AccessToken  string    `xorm:"text" json:"-"`         // never exposed to API
	Username     string    `xorm:"varchar(255)" json:"username"`
	AvatarURL    string    `xorm:"varchar(500)" json:"avatar_url"`
	Repos        RepoList  `xorm:"text" json:"repos"`     // tracked repositories
	SyncEnabled  bool      `xorm:"default true" json:"sync_enabled"`
	LastSyncedAt *time.Time `json:"last_synced_at"`
	Created      time.Time `xorm:"created" json:"created"`
	Updated      time.Time `xorm:"updated" json:"updated"`
}

func (GithubIntegration) TableName() string {
	return "github_integrations"
}

// RepoConfig represents a tracked repository
type RepoConfig struct {
	Owner       string `json:"owner"`
	Name        string `json:"name"`
	FullName    string `json:"full_name"`
	ProjectID   int64  `json:"project_id"` // link to a TaskDream project
	SyncIssues  bool   `json:"sync_issues"`
	SyncPRs     bool   `json:"sync_prs"`
}

type RepoList []RepoConfig

func (r RepoList) Value() (driver.Value, error) {
	return json.Marshal(r)
}

func (r *RepoList) Scan(value interface{}) error {
	if value == nil {
		*r = RepoList{}
		return nil
	}
	return json.Unmarshal(value.([]byte), r)
}

// GithubIssueSync tracks which GitHub issues are linked to tasks
type GithubIssueSync struct {
	ID            int64     `xorm:"pk autoincr" json:"id"`
	TaskID        int64     `xorm:"index" json:"task_id"`
	GithubIssueID int64     `json:"github_issue_id"`
	RepoFullName  string    `xorm:"varchar(255)" json:"repo_full_name"`
	IssueNumber   int       `json:"issue_number"`
	IssueTitle    string    `xorm:"varchar(500)" json:"issue_title"`
	IssueState    string    `xorm:"varchar(20)" json:"issue_state"` // open, closed
	IssueURL      string    `xorm:"varchar(500)" json:"issue_url"`
	IsPullRequest bool      `json:"is_pull_request"`
	LastSyncedAt  time.Time `json:"last_synced_at"`
	Created       time.Time `xorm:"created" json:"created"`
	Updated       time.Time `xorm:"updated" json:"updated"`
}

func (GithubIssueSync) TableName() string {
	return "github_issue_syncs"
}
