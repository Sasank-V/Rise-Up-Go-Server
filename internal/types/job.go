package types

import "time"

type CreateJobRequest struct {
	Owner              string   `json:"owner" `
	Title              string   `json:"title" `
	Description        string   `json:"description"`
	SkillTags          []string `json:"skill_tags"`
	WorkMode           string   `json:"work_mode"`
	JobType            string   `json:"job_type"`
	Location           string   `json:"location"`
	SalaryRangeStart   int64    `json:"salary_range_start"`
	SalaryRangeEnd     int64    `json:"salary_range_end"`
	Contact            string   `json:"contact"`
	EvaluationCriteria string   `json:"evaluation_criteria"`
	Active             bool     `json:"active"`
}

type UpdateJobRequest struct {
	UserId             string    `json:"user_id"`
	JobID              string    `json:"job_id"`
	Title              *string   `json:"title,omitempty"`
	Description        *string   `json:"description,omitempty"`
	SkillTags          *[]string `json:"skill_tags,omitempty"`
	WorkMode           *string   `json:"work_mode,omitempty"`
	JobType            *string   `json:"job_type,omitempty"`
	Location           *string   `json:"location,omitempty"`
	SalaryRangeStart   *int64    `json:"salary_range_start,omitempty"`
	SalaryRangeEnd     *int64    `json:"salary_range_end,omitempty"`
	EvaluationCriteria *string   `json:"evaluation_criteria,omitempty"`
	Active             *bool     `json:"active,omitempty"`
	Contact            *string   `json:"contact,omitempty"`
}

type AllJobItem struct {
	ID           string `json:"id"`
	OwnerName    string `json:"owner_name"`
	OwnerPicture string `json:"owner_picture"`
	Title        string `json:"title"`
	JobType      string `json:"job_type"`
	WorkMode     string `json:"work_mode"`
	
}

type FullJob struct {
	ID                 string    `json:"id"`
	Owner              string    `json:"owner_id"`
	OwnerName          string    `json:"owner_name"`
	OwnerPicture       string    `json:"owner_picture"`
	Title              string    `json:"title"`
	Description        string    `json:"description"`
	SkillTags          []string  `json:"skill_tags"`
	WorkMode           string    `json:"work_mode"`
	JobType            string    `json:"job_type"`
	Location           string    `json:"location"`
	SalaryRangeStart   int64     `json:"salary_range_start"`
	SalaryRangeEnd     int64     `json:"salary_range_end"`
	EvaluationCriteria string    `json:"evaluation_criteria"`
	Contact            string    `json:"contact"`
	PostedAt           time.Time `json:"posted_at"`
	Active             bool      `json:"active"`
}
