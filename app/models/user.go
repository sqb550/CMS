package models

type User struct {
	ID       uint   `json:"user_id"`
	Username string `json:"-"`
	UserType int    `json:"user_type"`
	Password string `json:"-"`
	Name     string `json:"-"`
}
