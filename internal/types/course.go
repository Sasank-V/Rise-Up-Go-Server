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
