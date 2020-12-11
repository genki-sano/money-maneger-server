package usecases

import (
	"sort"

	"github.com/genki-sano/money-maneger-server/package/applications/repositories"
	"github.com/genki-sano/money-maneger-server/package/applications/requests"
	"github.com/genki-sano/money-maneger-server/package/domains"
)

// PaymentUseCase は支払情報に関するユースケースです
type PaymentUseCase interface {
	List(*requests.PaymentListInputData) ([]domains.Payment, error)
}

// PaymentInteractor は支払情報に関するユースケースの構造体です
type PaymentInteractor struct {
	paymentRepos repositories.PaymentRepository
}

// NewPaymentUsecase はPaymentInteractorを返します
func NewPaymentUsecase(
	paymentRepos repositories.PaymentRepository,
) PaymentUseCase {
	return &PaymentInteractor{
		paymentRepos: paymentRepos,
	}
}

// List は支払情報一覧を取得します
func (i *PaymentInteractor) List(req *requests.PaymentListInputData) (payments []domains.Payment, err error) {
	payments, err = i.paymentRepos.GetByDate(req.Date)

	sort.SliceStable(payments, func(i, j int) bool {
		return payments[i].Date > payments[j].Date
	})

	return payments, err
}
