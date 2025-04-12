package types

type CreateCourseRequest struct {
	UserID      string         `json:"user_id"`
	Banner      string         `json:"banner"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Difficulty  string         `json:"difficulty"`
	Duration    int            `json:"duration"`
	Skills      []string       `json:"skills"`
	Modules     []CreateModule `json:"modules"`
	Instructors []string       `json:"instructors"`
}

type CreateModule struct {
	Title   string         `json:"title"`
	Lessons []CreateLesson `json:"lessons"`
}

type CreateLesson struct {
	Title       string           `json:"title"`
	Description string           `json:"description"`
	ContentLink string           `json:"content_link"`
	ContentType string           `json:"content_type"`
	Resources   []CreateResource `json:"resources"`
	Duration    int              `json:"duration"`
}

type CreateResource struct {
	Name string `json:"name"`
	Link string `json:"link"`
}
