package repositories

import (
	"time"

	"github.com/genki-sano/money-maneger-server/package/domains"
)

// PaymentRepository は支払情報関連のリポジトリです
type PaymentRepository interface {
	GetByDate(time.Time) ([]*domains.Payment, error)
}
