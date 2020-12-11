package presenters

import (
	"os"

	"github.com/genki-sano/money-maneger-server/package/domains"
)

type totalResponse struct {
	Women uint32 `json:"women"`
	Men   uint32 `json:"men"`
}

// PaymentListResponse は支払情報一覧のレスポンスの構造体です
type PaymentListResponse struct {
	Items []*domains.Payment `json:"items"`
	Total totalResponse      `json:"total"`
}

// PaymentPresenter は支払情報関連のプレゼンターです
type PaymentPresenter interface {
	List([]*domains.Payment) PaymentListResponse
}

// PaymentFactory は支払情報関連のプレゼンターの構造体です
type PaymentFactory struct {
}

// NewPaymentPresenter はPaymentFactoryを返します
func NewPaymentPresenter() PaymentPresenter {
	return &PaymentFactory{}
}

// List は支払情報一覧を返します
func (p *PaymentFactory) List(payments []*domains.Payment) PaymentListResponse {
	var wp uint32
	var mp uint32
	for _, payment := range payments {
		if payment.Name == os.Getenv("WOMEN_NAME") {
			wp = wp + payment.Price
		}
		if payment.Name == os.Getenv("MEN_NAME") {
			mp = mp + payment.Price
		}
	}

	return PaymentListResponse{
		Items: payments,
		Total: totalResponse{
			Women: wp,
			Men:   mp,
		},
	}
}
