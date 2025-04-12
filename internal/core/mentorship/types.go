package mentorship

import "time"

type SessionType string

type MentorshipRequest struct {
	From     string      `bson:"from" json:"from"`
	To       string      `bson:"to" json:"to"`
	Date     time.Time   `bson:"date" json:"date"`
	Time     time.Time   `bson:"time" json:"time"`
	Duration int         `bson:"duration" json:"duration"`
	Type     SessionType `bson:"type" json:"type"`
	Note     string      `bson:"note" json:"note"`
}

type MentorShipSession struct {
	MentorID  string   `bson:"mentor_id" json:"mentor_id"`
	RequestID string   `bson:"request_id" json:"request_id"`
	Link      string   `bson:"session_link" json:"session_link"`
	Title     string   `bson:"title" json:"title"`
	Resources []string `bson:"resources" json:"resources"`
}
