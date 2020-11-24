package repositories

import (
	"time"

	"github.com/genki-sano/money-maneger-server/package/domains"
)

// PaymentRepository 支払情報のリポジトリ
type PaymentRepository interface {
	FindAll() (domains.Payments, error)
	GetByDate(time.Time) (domains.Payments, error)
}
