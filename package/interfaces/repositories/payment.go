package repositories

import (
	"time"

	repos "github.com/genki-sano/money-maneger-server/package/applications/repositories"
	"github.com/genki-sano/money-maneger-server/package/domains"
	"github.com/genki-sano/money-maneger-server/package/infrastructure/datastore"
)

// PaymentDataAccess は支払情報関連のリポジトリの構造体です
type PaymentDataAccess struct {
	store datastore.PaymentDatastore
}

// NewPaymentRepository はPaymentDataAccessを返します
func NewPaymentRepository(store datastore.PaymentDatastore) repos.PaymentRepository {
	return &PaymentDataAccess{store: store}
}

// GetByDate は特定月の支払情報を取得します
func (d *PaymentDataAccess) GetByDate(t time.Time) (payments []*domains.Payment, err error) {
	all, err := d.store.GetAll()
	if err != nil {
		return nil, err
	}
	if len(all) == 0 {
		return all, nil
	}

	first := time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, time.Local)
	last := first.AddDate(0, 1, -1)

	payments = make([]*domains.Payment, 0)
	for _, items := range all {
		if date, _ := time.ParseInLocation("2006/01/02", items.Date, time.Local); date.Before(first) || date.After(last) {
			continue
		}
		payments = append(payments, items)
	}

	return payments, nil
}
