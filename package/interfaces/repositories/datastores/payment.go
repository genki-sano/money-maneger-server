package datastores

import (
	"github.com/genki-sano/money-maneger-server/package/domains"
)

// PaymentDatastore 支払情報リポジトリ（構造体）
type PaymentDatastore interface {
	GetAll() (domains.Payments, error)
}
