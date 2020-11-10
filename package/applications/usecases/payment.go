package usecases

import (
	"github.com/genki-sano/money-maneger-server/package/applications/repositories"
	"github.com/genki-sano/money-maneger-server/package/domains"
)

// PaymentUseCase ユースケース
type PaymentUseCase interface {
	Payments() (domains.Payments, error)
}

// PaymentInteractor 支払情報のユースケース
type PaymentInteractor struct {
	PaymentRepos repositories.PaymentRepository
}

// NewPaymentUsecase ユースケース
func NewPaymentUsecase(
	paymentRepos repositories.PaymentRepository,
) PaymentUseCase {
	return &PaymentInteractor{
		PaymentRepos: paymentRepos,
	}
}

// Payments 支払情報を全件取得
func (interactor *PaymentInteractor) Payments() (payments domains.Payments, err error) {
	payments, err = interactor.PaymentRepos.FindAll()
	return
}
