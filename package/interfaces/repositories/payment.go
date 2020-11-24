package repositories

import (
	"time"

	repos "github.com/genki-sano/money-maneger-server/package/applications/repositories"
	"github.com/genki-sano/money-maneger-server/package/domains"
	"github.com/genki-sano/money-maneger-server/package/interfaces/repositories/datastores"
)

// PaymentDataAccess 支払情報リポジトリ（構造体）
type PaymentDataAccess struct {
	store datastores.PaymentDatastore
}

// NewPaymentRepository リポジトリを生成
func NewPaymentRepository(store datastores.PaymentDatastore) repos.PaymentRepository {
	return &PaymentDataAccess{store: store}
}

// FindAll 全件取得
func (d *PaymentDataAccess) FindAll() (payments domains.Payments, err error) {
	return d.store.GetAll()
}

// GetByDate 特定の月のみ
func (d *PaymentDataAccess) GetByDate(t time.Time) (payments domains.Payments, err error) {
	all, err := d.store.GetAll()
	if err != nil {
		return nil, err
	}
	if len(all) == 0 {
		return all, nil
	}

	first := time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, time.UTC)
	last := first.AddDate(0, 1, -1)

	payments = make(domains.Payments, 0)
	for _, items := range all {
		if date, _ := time.Parse("2006/01/02", items.Date); date.Before(first) || date.After(last) {
			continue
		}
		payments = append(payments, items)
	}
	return payments, nil
}
