package model

type LessonFilter struct {
	CourseId      string
	Title         string
	Content       string
	Limit, Offset int
}
