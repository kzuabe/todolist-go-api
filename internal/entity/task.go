package entity

type Task struct {
	ID          string `json:"id"` // UUID
	UserID      string `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      int    `json:"status"` // 0: 未着手, 1: 進行中, 2: 完了
}
