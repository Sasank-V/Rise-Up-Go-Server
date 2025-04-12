package user

import "time"

type Role string

type User struct {
	ID           string `bson:"_id" json:"_id"`
	Name         string `bson:"name" json:"name"`
	Email        string `bson:"email" json:"email"`
	Picture      string `bson:"picture" json:"picture"`
	Bio          string `bson:"bio" json:"bio"`
	Location     string `bson:"location" json:"location"`
	Role         string `bson:"role" json:"role"`
	RoleID       string `bson:"role_id" json:"role_id"`
	AccessToken  string `bson:"access_token" json:"access_token"`
	RefreshToken string `bson:"refresh_token" json:"refresh_token"`
	ExpiresAt    string `bson:"expires_at" json:"expires_at"`
}

type Learner struct {
	UserID             string   `bson:"user_id" json:"user_id"`
	Skills             []string `bson:"skills" json:"skills"`
	Interests          []string `bson:"interests" json:"interests"`
	JobPreferences     []string `bson:"job_preferences" json:"job_preferences"` //Later
	LanguagePreferred  string   `bson:"language_preferred" json:"language_preferred"`
	Education          []string `bson:"education" json:"education"`
	ProfileCompletion  int      `bson:"profile_completion" json:"profile_completion"`
	EnrolledCourses    []string `bson:"enrolled_courses" json:"enrolled_courses"`
	AppliedJobs        []string `bson:"applied_jobs" json:"applied_jobs"`
	TestsTaken         []string `bson:"tests_taken" json:"tests_taken"`
	MentorshipRequests []string `bson:"mentorship_requests" json:"mentorship_requests"`
	MentorshipSessions []string `bson:"mentorship_sessions" json:"mentorship_sessions"`
	Reviews            []string `bson:"reviews" json:"reviews"`
}

type Mentor struct {
	UserID             string   `bson:"user_id" json:"user_id"`
	Skills             []string `bson:"skills" json:"skills"`
	Experience         []string `bson:"experience" json:"experience"`
	RegisteredCourses  []string `bson:"registered_courses" json:"registered_courses"`
	MentorshipRequests []string `bson:"mentorship_requests" json:"mentorship_requests"`
	MentorShipSessions []string `bson:"mentorship_sessions" json:"mentorship_sessions"`
	TestsTaken         []string `bson:"tests_taken" json:"tests_taken"`
	Reviews            []string `bson:"reviews" json:"reviews"`
	Available          bool     `bson:"available" json:"available"`
}

type Organisation struct {
	UserID             string   `bson:"user_id" json:"user_id"`
	OrganisationName   string   `bson:"organisation_name" json:"organisation_name"`
	About              string   `bson:"about" json:"about"`
	Website            string   `bson:"website" json:"website"`
	JobsPosted         []string `bson:"jobs_posted" json:"jobs_posted"`
	CoursesPosted      []string `bson:"courses_posted" json:"course_posted"`
	JobApplications    []string `bson:"job_applications" json:"job_applications"`
	MentorshipRequests []string `bson:"mentorship_sessions" json:"mentorship_sessions"`
}

type Education struct {
	UserID    string    `bson:"user_id" json:"user_id"`
	Institute string    `bson:"institute" json:"institute"`
	Degree    string    `bson:"degree" json:"degree"`
	StartDate time.Time `bson:"start_date" json:"start_date"`
	EndDate   time.Time `bson:"end_date" json:"end_date"`
}

type Experience struct {
	UserID      string    `bson:"user_id" json:"user_id"`
	Company     string    `bson:"company" json:"company"`
	Position    string    `bson:"position" json:"position"`
	StartDate   time.Time `bson:"start_date" json:"start_date"`
	EndDate     time.Time `bson:"end_date" json:"end_date"`
	Description string    `bson:"description" json:"description"`
}

type Review struct {
	From   string `bson:"from" json:"from"`
	To     string `bson:"to" json:"to"`
	Rating int    `bson:"rating" json:"rating"`
	Body   string `bson:"body" json:"body"`
}
