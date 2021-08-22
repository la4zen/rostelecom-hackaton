package models

type Event struct {
	User  User
	Event map[string]interface{}
}
