package model

type Task struct {
	ID          string `json:"id"`          // タスクID（自動で生成されるUUID）
	UserID      string `json:"user_id"`     // ユーザID
	Title       string `json:"title"`       // タスクタイトル
	Description string `json:"description"` // タスク説明文
	Status      int    `json:"status"`      // タスクのステータス（0: 未着手, 1: 完了）
}

type TaskFetchParam struct {
	UserID string `form:"user_id"`
	Status *int   `form:"status"`
}
