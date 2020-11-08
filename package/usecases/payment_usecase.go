package usecases

import (
	"github.com/genki-sano/money-maneger-server/package/domains"
)

// PaymentInteractor 支払情報のユースケース
type PaymentInteractor struct {
	PaymentRepository PaymentRepository
}

// Payments 支払情報を全件取得
func (interactor *PaymentInteractor) Payments() (payments domains.Payments, err error) {
	payments, err = interactor.PaymentRepository.FindAll()
	return
}
