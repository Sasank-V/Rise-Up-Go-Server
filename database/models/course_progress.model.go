package models

type CourseProgress struct {
	UserID           string   `bson:"user_id" json:"user_id"`
	CourseID         string   `bson:"course_id" json:"course_id"`
	LessonsCompleted []string `bson:"lessons_completed" json:"lessons_completed"`
	CourseCompleted  bool     `bson:"course_completed" json:"course_completed"`
}
