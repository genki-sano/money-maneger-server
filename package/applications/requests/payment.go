package requests

import (
	"time"
)

// PaymentListInputData は支払情報一覧を取得するリクエストの構造体です
type PaymentListInputData struct {
	Date time.Time
}

// NewPaymentListInputData はPaymentListInputDataを返します
func NewPaymentListInputData(date string) (req *PaymentListInputData, err error) {
	t, err := time.ParseInLocation("20060102", date, time.Local)
	if err != nil {
		return nil, err
	}

	req = &PaymentListInputData{
		Date: t,
	}
	return req, nil
}
