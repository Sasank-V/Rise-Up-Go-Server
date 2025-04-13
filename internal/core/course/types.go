package course

type ContentType string
type Difficulty string

type Course struct {
	ID            string     `json:"id" bson:"_id,omitempty"`
	Owner         string     `bson:"owner" json:"owner"`
	Banner        string     `bson:"banner" json:"banner"`
	Title         string     `bson:"title" json:"title"`
	Description   string     `bson:"description" json:"description"`
	Difficulty    Difficulty `bson:"difficulty" json:"difficulty"`
	Duration      int        `bson:"duration" json:"duration"`
	Skills        []string   `bson:"skills" json:"skills"`
	Modules       []string   `bson:"modules" json:"modules"`
	Instructors   []string   `bson:"instructors" json:"instructors"`
	Discussions   []string   `bson:"discussions" json:"discussions"`
	Prerequisites string     `bson:"prerequistes" json:"prerequisites"`
	Outcomes      string     `bson:"outcome" json:"outcome"`
}

type Module struct {
	ID       string   `bson:"_id,omitempty" json:"id"`
	CourseID string   `bson:"course_id" json:"title_id"`
	Title    string   `bson:"title" json:"title"`
	OrderNo  int      `bson:"order_no" json:"order_no"`
	Lessons  []string `bson:"lessons" json:"lessons"`
}

type Lesson struct {
	ModuleID    string      `bson:"module_id" json:"module_id"`
	Title       string      `bson:"title" json:"title"`
	Description string      `bson:"description" json:"description"`
	ContentLink string      `bson:"content_link" json:"content_link"`
	ContentType ContentType `bson:"content_type" json:"content_type"`
	Resources   []string    `bson:"resources" json:"resources"`
	Duration    int         `bson:"duration" json:"duration"`
	OrderNo     int         `bson:"order_no" json:"order_no"`
}

type Resource struct {
	ID       string `bson:"_id,omitempty" json:"id"`
	LessonID string `bson:"lesson_id" json:"lesson_id"`
	Name     string `bson:"name" json:"name"`
	Link     string `bson:"link" json:"link"`
}

type CourseProgress struct {
	UserID           string   `bson:"user_id" json:"user_id"`
	CourseID         string   `bson:"course_id" json:"course_id"`
	LessonsCompleted []string `bson:"lessons_completed" json:"lessons_completed"`
	CourseCompleted  bool     `bson:"course_completed" json:"course_completed"`
}
