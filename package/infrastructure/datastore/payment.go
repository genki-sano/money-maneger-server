package datastore

import (
	"os"

	"github.com/genki-sano/money-maneger-server/package/domains"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/jwt"
	"google.golang.org/api/sheets/v4"
)

// PaymentDatastore 支払情報リポジトリ（構造体）
type PaymentDatastore interface {
	GetAll() (domains.Payments, error)
}

// PaymentDatabase 支払情報リポジトリ（構造体）
type PaymentDatabase struct{}

// NewPaymentDatastore リポジトリを生成
func NewPaymentDatastore() PaymentDatastore {
	return &PaymentDatabase{}
}

// GetAll は全件取得します
func (d *PaymentDatabase) GetAll() (payments domains.Payments, err error) {
	email := os.Getenv("GOOGLE_SERVICE_ACCOUNT_EMAIL")
	key := os.Getenv("GOOGLE_SERVICE_ACCOUNT_PLIVATE_KEY")
	srv, err := newSheetService(email, key)
	if err != nil {
		return nil, err
	}

	spreadsheetID := os.Getenv("GOOGLE_SPREDSHEET_ID")
	valueRange := os.Getenv("GOOGLE_SPREDSHEET_RANGE")
	resp, err := srv.Spreadsheets.Values.Get(spreadsheetID, valueRange).Do()
	if err != nil {
		return nil, err
	}

	return generatePayment(resp)
}

func newSheetService(email string, key string) (*sheets.Service, error) {
	scopes := []string{
		"https://www.googleapis.com/auth/spreadsheets.readonly",
	}
	conf := &jwt.Config{
		Email:      email,
		PrivateKey: []byte(key),
		TokenURL:   google.JWTTokenURL,
		Scopes:     scopes,
	}
	client := conf.Client(oauth2.NoContext)

	return sheets.New(client)
}

func generatePayment(r *sheets.ValueRange) (payments domains.Payments, err error) {
	payments = make(domains.Payments, 0)
	if len(r.Values) == 0 {
		return payments, nil
	}

	for _, items := range r.Values {
		payments = append(payments, domains.Payment{
			ID:       items[0].(string),
			Name:     items[1].(string),
			Date:     items[2].(string),
			Price:    items[3].(string),
			Category: items[4].(string),
			Memo:     items[5].(string),
		})
	}
	return payments, nil
}