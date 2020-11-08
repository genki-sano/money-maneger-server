package repositories

import (
	"os"

	"github.com/genki-sano/money-maneger-server/package/domains"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/jwt"
	"google.golang.org/api/sheets/v4"
)

// PaymentRepository 支払情報のリポジトリ
type PaymentRepository struct{}

// FindAll 全件取得
func (repo *PaymentRepository) FindAll() (payments domains.Payments, err error) {
	email := os.Getenv("GOOGLE_SERVICE_ACCOUNT_EMAIL")
	key := os.Getenv("GOOGLE_SERVICE_ACCOUNT_PLIVATE_KEY")

	conf := &jwt.Config{
		Email:      email,
		PrivateKey: []byte(key),
		TokenURL:   google.JWTTokenURL,
		Scopes: []string{
			"https://www.googleapis.com/auth/spreadsheets.readonly",
		},
	}

	client := conf.Client(oauth2.NoContext)
	srv, err := sheets.New(client)
	if err != nil {
		return nil, err
	}

	spreadsheetID := os.Getenv("GOOGLE_SPREDSHEET_ID")
	valueRange := os.Getenv("GOOGLE_SPREDSHEET_RANGE")
	resp, err := srv.Spreadsheets.Values.Get(spreadsheetID, valueRange).Do()
	if err != nil {
		return nil, err
	}

	payments = make(domains.Payments, 0)
	if len(resp.Values) == 0 {
		return payments, nil
	}

	for _, items := range resp.Values {
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
