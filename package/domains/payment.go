package domains

// Payments 支払情報
type Payments []Payment

// Payment 支払情報
type Payment struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Date     string `json:"date"`
	Price    string `json:"price"`
	Category string `json:"category"`
	Memo     string `json:"memo"`
}
