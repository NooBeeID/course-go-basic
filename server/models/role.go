package models

type Role struct {
	ID   int
	Name string
}

func NewRole() *Role {
	return &Role{}
}
