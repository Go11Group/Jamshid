package handlers

import "net/http"

type Handler struct {
	client *http.Client
}

//func NewHandler(user *postgres.UserRepository, course *postgres.CourseRepository, lesson *postgres.LessonRepository, enrollment *postgres.EnrollmentRepository) *Handler {
//	return &Handler{UserRepo: user, CourseRepo: course, LessonRepo: lesson, EnrollmentRepo: enrollment}
//}
