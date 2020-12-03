package requests

import (
	"time"
)

// PaymentInputData 支払情報のリクエスト
type PaymentInputData struct {
	Date time.Time
}

// NewPaymentInputData リクエストを作成
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
