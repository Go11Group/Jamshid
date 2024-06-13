package model

type EnrollmentFilter struct {
	User_id        string
	Course_id      string
	EnrollmentDate string
	Limit, Offset  int
}
