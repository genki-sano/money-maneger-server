package usecases

import (
	"github.com/genki-sano/money-maneger-server/package/applications/repositories"
	"github.com/genki-sano/money-maneger-server/package/domains"
)

// PaymentUseCase 支払情報ユースケース
type PaymentUseCase interface {
	Payments() (domains.Payments, error)
}

// PaymentInteractor 支払情報ユースケース（構造体）
type PaymentInteractor struct {
	paymentRepos repositories.PaymentRepository
}

// NewPaymentUsecase ユースケースを生成
func NewPaymentUsecase(
	paymentRepos repositories.PaymentRepository,
) PaymentUseCase {
	return &PaymentInteractor{
		paymentRepos: paymentRepos,
	}
}

// Payments 支払情報を全件取得
func (i *PaymentInteractor) Payments() (payments domains.Payments, err error) {
	payments, err = i.paymentRepos.FindAll()
	return
}
