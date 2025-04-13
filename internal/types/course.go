package types

type CreateCourseRequest struct {
	UserID        string         `json:"user_id"`
	Banner        string         `json:"banner"`
	Title         string         `json:"title"`
	Description   string         `json:"description"`
	Difficulty    string         `json:"difficulty"`
	Duration      int            `json:"duration"`
	Skills        []string       `json:"skills"`
	Modules       []CreateModule `json:"modules"`
	Instructors   []string       `json:"instructors"`
	Prerequisites string         `json:"prerequisites"`
	Outcomes      string         `json:"outcomes"`
}

type CreateModule struct {
	Title   string         `json:"title"`
	OrderNo int            `json:"order_no"`
	Lessons []CreateLesson `json:"lessons"`
}

type CreateLesson struct {
	Title       string           `json:"title"`
	Description string           `json:"description"`
	ContentLink string           `json:"content_link"`
	ContentType string           `json:"content_type"`
	Resources   []CreateResource `json:"resources"`
	Duration    int              `json:"duration"`
	OrderNo     int              `json:"order_no"`
}

type CreateResource struct {
	Name string `json:"name"`
	Link string `json:"link"`
}

type UpdateCourseRequest struct {
	CourseID      string    `json:"course_id"`
	UserID        *string   `json:"user_id,omitempty"`
	Banner        *string   `json:"banner,omitempty"`
	Title         *string   `json:"title,omitempty"`
	Description   *string   `json:"description,omitempty"`
	Difficulty    *string   `json:"difficulty,omitempty"`
	Duration      *int      `json:"duration,omitempty"`
	Skills        *[]string `json:"skills,omitempty"`
	Prerequisites *string   `json:"prerequisites,omitempty"`
	Outcomes      *string   `json:"outcomes,omitempty"`
}

type UpdateModule struct {
	ModuleID string  `json:"module_id"`
	Title    *string `json:"title,omitempty"`
	OrderNo  *int    `json:"order_no,omitempty"`
}

type UpdateLesson struct {
	LessonID    string  `json:"lesson_id"`
	Title       *string `json:"title,omitempty"`
	Description *string `json:"description,omitempty"`
	ContentLink *string `json:"content_link,omitempty"`
	ContentType *string `json:"content_type,omitempty"`
	Duration    *int    `json:"duration,omitempty"`
	OrderNo     *int    `json:"order_no,omitempty"`
}

type UpdateResource struct {
	ResourceID string  `json:"resource_id"`
	Name       *string `json:"name,omitempty"`
	Link       *string `json:"link,omitempty"`
}

type AllCourseItem struct {
	ID          string   `json:"id" bson:"_id,omitempty"`
	Owner       string   `json:"owner" bson:"owner"`
	Banner      string   `json:"banner" bson:"banner"`
	Title       string   `json:"title" bson:"title"`
	Description string   `json:"description" bson:"description"`
	Difficulty  string   `json:"difficulty" bson:"difficulty"`
	Duration    int      `json:"duration" bson:"duration"`
	Skills      []string `json:"skills" bson:"skills"`
}

// FullCourse represents an expanded course with full nested details.
type FullCourse struct {
	ID            string          `json:"id" bson:"_id,omitempty"`
	Owner         BasicUserInfo   `json:"owner" bson:"owner"`
	Banner        string          `json:"banner" bson:"banner"`
	Title         string          `json:"title" bson:"title"`
	Description   string          `json:"description" bson:"description"`
	Difficulty    string          `json:"difficulty" bson:"difficulty"`
	Duration      int             `json:"duration" bson:"duration"`
	Skills        []string        `json:"skills" bson:"skills"`
	Modules       []FullModule    `json:"modules" bson:"modules"`
	Instructors   []BasicUserInfo `json:"instructors" bson:"instructors"`
	Discussions   []string        `json:"discussions" bson:"discussions"`
	Prerequisites string          `json:"prerequisites" bson:"prerequistes"`
	Outcomes      string          `json:"outcomes" bson:"outcome"`
}

// FullModule is an expanded module that contains full lesson data.
type FullModule struct {
	ID       string       `json:"id" bson:"_id,omitempty"`
	CourseID string       `json:"course_id" bson:"course_id"`
	Title    string       `json:"title" bson:"title"`
	OrderNo  int          `json:"order_no" bson:"order_no"`
	Lessons  []FullLesson `json:"lessons" bson:"lessons"`
}

// FullLesson is an expanded lesson containing full resource data.
type FullLesson struct {
	ID          string     `json:"id" bson:"_id,omitempty"`
	ModuleID    string     `json:"module_id" bson:"module_id"`
	Title       string     `json:"title" bson:"title"`
	Description string     `json:"description" bson:"description"`
	ContentLink string     `json:"content_link" bson:"content_link"`
	ContentType string     `json:"content_type" bson:"content_type"`
	Resources   []Resource `json:"resources" bson:"resources"`
	Duration    int        `json:"duration" bson:"duration"`
	OrderNo     int        `json:"order_no" bson:"order_no"`
}

// Resource represents a resource with additional metadata.
type Resource struct {
	ID       string `json:"id" bson:"_id,omitempty"`
	LessonID string `json:"lesson_id" bson:"lesson_id"`
	Name     string `json:"name" bson:"name"`
	Link     string `json:"link" bson:"link"`
}
