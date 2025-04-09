package models

type ApplicationStatus string

const (
	Accepted ApplicationStatus = "accepted"
	Rejected ApplicationStatus = "rejected"
	Pending  ApplicationStatus = "pending"
)

type JobApplication struct {
	UserID string            `bson:"user_id" json:"user_id"`
	JobID  string            `bson:"job_id" json:"job_id"`
	Status ApplicationStatus `bson:"status" json:"status"`
}
