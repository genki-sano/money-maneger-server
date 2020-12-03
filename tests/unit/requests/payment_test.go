package requests_test

import (
	"testing"
	"time"

	"github.com/genki-sano/money-maneger-server/package/applications/requests"
	"github.com/stretchr/testify/assert"
)

func TestPaymentInputData_NewPaymentInputData(t *testing.T) {
	testCases := []struct {
		name     string
		date     string
		expected *requests.PaymentInputData
	}{
		{
			name: "正しい形式の場合",
			date: "20201201",
			expected: &requests.PaymentInputData{
				Date: time.Date(2020, 12, 1, 0, 0, 0, 0, time.Local),
			},
		},
		{
			name: "うるう年の場合",
			date: "20200229",
			expected: &requests.PaymentInputData{
				Date: time.Date(2020, 2, 29, 0, 0, 0, 0, time.Local),
			},
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual, err := requests.NewPaymentInputData(testCase.date)

			assert.NotNil(t, actual)
			assert.Exactly(t, actual.Date, testCase.expected.Date)

			assert.NoError(t, err)
		})
	}
}

func TestPaymentInputData_NewPaymentInputData_Failed(t *testing.T) {
	testCases := []struct {
		name string
		date string
		msg  string
	}{
		{
			name: "空文字の場合",
			date: "",
			msg:  "parsing time \"\" as \"20060102\": cannot parse \"\" as \"2006\"",
		},
		{
			name: "日付でない場合",
			date: "hoge",
			msg:  "parsing time \"hoge\" as \"20060102\": cannot parse \"hoge\" as \"2006\"",
		},
		{
			name: "形式が正しくない場合",
			date: "2020-12-01",
			msg:  "parsing time \"2020-12-01\" as \"20060102\": cannot parse \"-12-01\" as \"01\"",
		},
		{
			name: "範囲外の場合",
			date: "20200230",
			msg:  "parsing time \"20200230\": day out of range",
		},
		{
			name: "うるう年でない場合",
			date: "20190229",
			msg:  "parsing time \"20190229\": day out of range",
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual, err := requests.NewPaymentInputData(testCase.date)

			assert.Nil(t, actual)

			assert.Error(t, err)
			assert.EqualError(t, err, testCase.msg)
		})
	}
}