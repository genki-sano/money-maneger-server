package domains

// Payment は支払情報の構造体です
type Payment struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Date     string `json:"date"`
	Price    string `json:"price"`
	Category string `json:"category"`
	Memo     string `json:"memo"`
}

// NewPayment は支払情報を作成します
func NewPayment(id string, name string, date string, price string, category string, memo string) *Payment {
	return &Payment{
		ID:       id,
		Name:     name,
		Date:     date,
		Price:    price,
		Category: category,
		Memo:     memo,
	}
}
