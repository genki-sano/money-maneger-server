package controllers

import (
	"net/http"

	"github.com/genki-sano/money-maneger-server/package/applications/usecases"
	"github.com/genki-sano/money-maneger-server/package/interfaces/handlers"
)

// PaymentController 支払情報のコントローラー
type PaymentController struct {
	u usecases.PaymentUseCase
}

// NewPaymentController コントローラーを生成
func NewPaymentController(u usecases.PaymentUseCase) *PaymentController {
	return &PaymentController{u: u}
}

// Index 支払情報を全件取得
func (controller *PaymentController) Index(ctx handlers.Context) {
	resp, err := controller.u.Payments()
	if err != nil {
		ctx.JSON(http.StatusForbidden, handlers.NewError(err))
		return
	}

	ctx.JSON(http.StatusOK, resp)
}
