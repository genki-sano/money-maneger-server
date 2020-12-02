package controllers

import (
	"net/http"

	"github.com/genki-sano/money-maneger-server/package/applications/requests"

	"github.com/genki-sano/money-maneger-server/package/applications/usecases"
	"github.com/genki-sano/money-maneger-server/package/interfaces/handlers"
)

// PaymentController 支払情報のコントローラー（構造体）
type PaymentController struct {
	u usecases.PaymentUseCase
}

// NewPaymentController コントローラーを生成
func NewPaymentController(u usecases.PaymentUseCase) *PaymentController {
	return &PaymentController{u: u}
}

// Index 支払情報を全件取得
func (c *PaymentController) Index(ctx handlers.Context) {
	resp, err := c.u.Payments()
	if err != nil {
		ctx.JSON(http.StatusForbidden, handlers.NewError(err))
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

// Show 支払情報を全件取得
func (c *PaymentController) Show(ctx handlers.Context) {
	req, err := requests.NewPaymentInputData(ctx.Param("date"))
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, handlers.NewError(err))
		return
	}

	resp, err := c.u.PaymentsByDate(req)
	if err != nil {
		ctx.JSON(http.StatusForbidden, handlers.NewError(err))
		return
	}

	ctx.JSON(http.StatusOK, resp)
}
