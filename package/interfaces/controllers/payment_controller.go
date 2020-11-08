package controllers

import (
	"net/http"

	"github.com/genki-sano/money-maneger-server/package/interfaces/repositories"
	"github.com/genki-sano/money-maneger-server/package/usecases"
)

// PaymentController 支払情報のコントローラー
type PaymentController struct {
	Interactor usecases.PaymentInteractor
}

// NewPaymentController コントローラーを生成
func NewPaymentController() *PaymentController {
	return &PaymentController{
		Interactor: usecases.PaymentInteractor{
			PaymentRepository: &repositories.PaymentRepository{},
		},
	}
}

// Index 支払情報を全件取得
func (controller *PaymentController) Index(c Context) {
	resp, err := controller.Interactor.Payments()
	if err != nil {
		c.JSON(http.StatusForbidden, NewError(err))
		return
	}

	c.JSON(http.StatusOK, resp)
}
