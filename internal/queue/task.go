package queue

type Task struct {
	ChatID int64  `json:"chat_id"`
	Email  string `json:"email"`
}
