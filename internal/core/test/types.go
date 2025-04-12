package test

type TestType string

type Test struct {
	Type       TestType `bson:"type" json:"type"`
	CourseID   string   `bson:"course_id" json:"course_id"`
	Skills     []string `bson:"skills" json:"skills"`
	Difficulty string   `bson:"difficulty" json:"difficulty"`
	Questions  []string `bson:"questions" json:"questions"`
}

type TestResult struct {
	TestID   string `bson:"test_id" json:"test_id"`
	UserID   string `bson:"user_id" json:"user_id"`
	Result   int    `bson:"result" json:"result"`
	FeedBack string `bson:"feedback" json:"feedback"`
}
