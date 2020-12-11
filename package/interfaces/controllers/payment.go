package controllers

import (
	"net/http"

	"github.com/genki-sano/money-maneger-server/package/applications/requests"
	"github.com/genki-sano/money-maneger-server/package/applications/usecases"
	"github.com/genki-sano/money-maneger-server/package/interfaces/handlers"
	"github.com/genki-sano/money-maneger-server/package/interfaces/presenters"
)

// PaymentController は支払情報関連のコントローラーの構造体です
type PaymentController struct {
	u usecases.PaymentUseCase
	p presenters.PaymentPresenter
}

// NewPaymentController はPaymentControllerを返します
func NewPaymentController(
	u usecases.PaymentUseCase,
	p presenters.PaymentPresenter,
) *PaymentController {
	return &PaymentController{
		u: u,
		p: p,
	}
}

// List は支払情報一覧を取得します
func (c *PaymentController) List(ctx handlers.Context) {
	req, err := requests.NewPaymentListInputData(ctx.Param("date"))
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, handlers.NewError(err))
		return
	}

	resp, err := c.u.List(req)
	if err != nil {
		ctx.JSON(http.StatusForbidden, handlers.NewError(err))
		return
	}

	ctx.JSON(http.StatusOK, c.p.List(resp))
}
