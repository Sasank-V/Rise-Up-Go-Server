package job

import "time"

type WorkMode string
type JobType string
type Job struct {
	ID                 string    `bson:"_id" json:"id"`
	Owner              string    `bson:"owner" json:"owner"`
	Title              string    `bson:"title" json:"title"`
	Description        string    `bson:"description" json:"description"`
	SkillTags          []string  `bson:"skill_tags" json:"skill_tags"`
	WorkMode           WorkMode  `bson:"work_mode" json:"work_mode"`
	JobType            JobType   `bson:"job_type" json:"job_type"`
	Location           string    `bson:"location" json:"location"`
	SalaryRangeStart   int64     `bson:"salary_range_start" json:"salary_range_start"`
	SalaryRangeEnd     int64     `bson:"salary_range_end" json:"salary_range_end"`
	EvaluationCriteria string    `bson:"evaluation_criteria" json:"evaluation_criteria"`
	Active             bool      `bson:"active" json:"active"`
	Contact            string    `bson:"contact" json:"contact"`
	PostedAt           time.Time `bson:"posted_at" json:"posted_at"`
}

type ApplicationStatus string
type JobApplication struct {
	UserID          string            `bson:"user_id" json:"user_id"`
	JobID           string            `bson:"job_id" json:"job_id"`
	Status          ApplicationStatus `bson:"status" json:"status"`
	TestResult      string            `bson:"test_result" json:"test_result"`
	MatchPercentage int               `bson:"match_percentage" json:"match_percentage"`
}
