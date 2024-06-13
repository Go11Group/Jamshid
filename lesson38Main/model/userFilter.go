package model

type UserFilter struct {
	Name          string
	Email         string
	Birthday      string
	Password      string
	Limit, Offset int
}
