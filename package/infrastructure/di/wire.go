// +build wireinject

package di

import (
	"github.com/genki-sano/money-maneger-server/package/applications/usecases"
	"github.com/genki-sano/money-maneger-server/package/interfaces/controllers"
	"github.com/genki-sano/money-maneger-server/package/interfaces/repositories"
	"github.com/google/wire"
)

// InitializePayment DI
func InitializePayment() *controllers.PaymentController {
	wire.Build(
		controllers.NewPaymentController,
		usecases.NewPaymentUsecase,
		repositories.NewPaymentRepository,
	)
	return nil
}
