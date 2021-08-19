package models

type User struct {
	ID       uint   `gorm:"primary"`
	Login    string `json:"login,omitempty"`
	Password string `json:"password,omitempty"`
}
