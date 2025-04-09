package models

type Organisation struct {
	UserID             string   `bson:"user_id" json:"user_id"`
	JobsPosted         []string `bson:"jobs_posted" json:"jobs_posted"`
	CoursesPosted      []string `bson:"courses_posted" json:"course_posted"`
	JobApplications    []string `bson:"job_applications" json:"job_applications"`
	MentorshipRequests []string `bson:"mentorship_sessions" json:"mentorship_sessions"`
}
