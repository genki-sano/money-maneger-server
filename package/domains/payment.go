package domains

// Payments 支払情報
type Payments []Payment

// Payment 支払情報
type Payment struct {
	ID       string
	Name     string
	Date     string
	Price    string
	Category string
	Memo     string
}
