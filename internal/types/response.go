package types

type MessageResponse struct {
	Message string `json:"message"`
}

type PaginatedCoursesResponse struct {
	Page       int64           `json:"page"`
	PageSize   int64           `json:"page_size"`
	TotalCount int64           `json:"total_count"`
	Courses    []AllCourseItem `json:"courses"`
}

type AllCoursesResponse struct {
	Page       int64           `json:"page"`
	PageSize   int64           `json:"page_size"`
	TotalCount int64           `json:"total_count"`
	Courses    []AllCourseItem `json:"courses"`
}
