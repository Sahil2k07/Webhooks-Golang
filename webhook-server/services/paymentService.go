package services

import (
	"net/http"

	"github.com/Sahil2k07/Webhooks-Golang/webhook-server/interfaces"
	"github.com/Sahil2k07/Webhooks-Golang/webhook-server/views"
	"github.com/labstack/echo"
)

type paymentService struct {
	webhookService interfaces.IWebhookService
	context        echo.Context
}

func NewPaymentService(ctx echo.Context) interfaces.IPaymentService {
	return &paymentService{
		context:        ctx,
		webhookService: NewWebHookService(),
	}
}

func (ps *paymentService) ProcessPayment() error {
	var (
		payload views.PaymentRequest

		trxID string = "randomID"
	)

	err := ps.context.Bind(&payload)
	if err != nil {
		return err
	}

	// Process some logic and after that:

	go ps.webhookService.PaymentSuccessWebhook(payload)

	response := views.PaymentResponse{
		TransactionID: trxID,
		Success:       true,
	}

	return ps.context.JSON(http.StatusOK, response)
}
