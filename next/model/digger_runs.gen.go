// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameDiggerRun = "digger_runs"

// DiggerRun mapped from table <digger_runs>
type DiggerRun struct {
	ID                   string         `gorm:"column:id;primaryKey;default:gen_random_uuid()" json:"id"`
	CreatedAt            time.Time      `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt            time.Time      `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt            gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	Triggertype          string         `gorm:"column:triggertype;not null" json:"triggertype"`
	PrNumber             int64          `gorm:"column:pr_number" json:"pr_number"`
	Status               string         `gorm:"column:status;not null" json:"status"`
	CommitID             string         `gorm:"column:commit_id;not null" json:"commit_id"`
	DiggerConfig         string         `gorm:"column:digger_config" json:"digger_config"`
	GithubInstallationID int64          `gorm:"column:github_installation_id" json:"github_installation_id"`
	RepoID               int64          `gorm:"column:repo_id;not null" json:"repo_id"`
	RunType              string         `gorm:"column:run_type;not null" json:"run_type"`
	PlanStageID          string         `gorm:"column:plan_stage_id" json:"plan_stage_id"`
	ApplyStageID         string         `gorm:"column:apply_stage_id" json:"apply_stage_id"`
	ProjectName          string         `gorm:"column:project_name" json:"project_name"`
	IsApproved           bool           `gorm:"column:is_approved" json:"is_approved"`
	ApprovalAuthor       string         `gorm:"column:approval_author" json:"approval_author"`
	ApprovalDate         time.Time      `gorm:"column:approval_date" json:"approval_date"`
	ProjectID            string         `gorm:"column:project_id;not null" json:"project_id"`
	TerraformOutput      string         `gorm:"column:terraform_output" json:"terraform_output"`
	ApplyLogs            string         `gorm:"column:apply_logs" json:"apply_logs"`
	ApproverUserID       *string        `gorm:"column:approver_user_id" json:"approver_user_id"`
	TriggeredByUserID    *string        `gorm:"column:triggered_by_user_id" json:"triggered_by_user_id"`
}

// TableName DiggerRun's table name
func (*DiggerRun) TableName() string {
	return TableNameDiggerRun
}
