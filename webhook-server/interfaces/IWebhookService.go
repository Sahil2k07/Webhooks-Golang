package interfaces

import "github.com/Sahil2k07/Webhooks-Golang/webhook-server/views"

type IWebhookService interface {
	PaymentSuccessWebhook(payload views.PaymentRequest)
}
