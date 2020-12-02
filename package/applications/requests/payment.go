package requests

import (
	"time"
)

// PaymentInputData はテスト
type PaymentInputData struct {
	Date time.Time
}

// NewPaymentInputData はテスト
func NewPaymentInputData(date string) (req *PaymentInputData, err error) {
	t, err := time.ParseInLocation("20060102", date, time.Local)
	if err != nil {
		return nil, err
	}

	req = &PaymentInputData{
		Date: t,
	}
	return req, nil
}
