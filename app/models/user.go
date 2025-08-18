package models

type User struct {
	ID       uint   `json:"user_id"`
	Username string `json:"username"`
	UserType int    `json:"user_type"`
	Password string `json:"password"`
	Name     string `json:"name"`
}
