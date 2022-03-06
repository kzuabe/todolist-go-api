package entity

type Task struct {
	ID          string `json:"id"`      // UUID
	UserID      string `json:"user_id"` // ユーザID
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      int    `json:"status"` // 0: 未着手, 1: 完了
}

type TaskFetchParam struct {
	UserID string `form:"user_id"`
	Status *int   `form:"status"`
}
