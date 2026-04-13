package models

import "time"

// TaskRelation represents relationships between tasks (parent/child/related/blocking)
type TaskRelation struct {
	ID        int64     `xorm:"pk autoincr" json:"id"`
	TaskID    int64     `xorm:"index" json:"task_id"`
	RelatedID int64     `xorm:"index" json:"related_id"`
	Kind      string    `xorm:"varchar(50)" json:"kind"` // parent, child, related, blocking, blocked_by, duplicates, duplicated_by
	Created   time.Time `xorm:"created" json:"created"`
	Updated   time.Time `xorm:"updated" json:"updated"`
}

func (TaskRelation) TableName() string {
	return "task_relations"
}
