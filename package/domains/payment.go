package domains

import "strconv"

// Payment は支払情報の構造体です
type Payment struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Date     string `json:"date"`
	Price    uint32 `json:"price"`
	Category string `json:"category"`
	Memo     string `json:"memo"`
}

// NewPayment は支払情報を作成します
func NewPayment(id, name, date, price, category, memo string) *Payment {
	num, _ := strconv.Atoi(price)
	return &Payment{
		ID:       id,
		Name:     name,
		Date:     date,
		Price:    uint32(num),
		Category: category,
		Memo:     memo,
	}
}
