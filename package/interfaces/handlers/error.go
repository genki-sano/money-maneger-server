package handlers

// Error エラー発生時のレスポンス
type Error struct {
	Message string `json:"message"`
}

// NewError はErrorを返します
func NewError(err error) *Error {
	return &Error{
		Message: err.Error(),
	}
}
