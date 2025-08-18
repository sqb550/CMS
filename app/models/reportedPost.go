package models

type ReportedPost struct {
	ID       uint   `json:"report_id"`
	UserID   int    `json:"user_id"`
	PostID   int    `json:"post_id"`
	Content  string `json:"content"`
	Reason   string `json:"reason"`
	Status   int    `json:"status"`
	Username string `json:"username"`
}
