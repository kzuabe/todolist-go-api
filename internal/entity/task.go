package entity

type Task struct {
	ID          string `json:"id"` // UUID
	UserID      string `json:"-"`  // NOTE: 自明のためリクエスト・レスポンスボディに含めない
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      int    `json:"status"` // 0: 未着手, 1: 完了
}

type TaskFetchParam struct {
	UserID string `form:"-"`
	Status *int   `form:"status"`
}
