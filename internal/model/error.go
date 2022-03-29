package model

type Error struct {
	StatusCode int    `json:"-"`       // HTTPステータスコード
	Message    string `json:"message"` // エラーメッセージ
}

func (e *Error) Error() string {
	return e.Message
}
