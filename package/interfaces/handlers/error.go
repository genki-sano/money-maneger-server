package handlers

// Error エラー発生時のレスポンス
type Error struct {
	Message string
}

// NewError エラーレスポンスを作成
func NewError(err error) *Error {
	return &Error{
		Message: err.Error(),
	}
}
