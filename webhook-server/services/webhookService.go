package services

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/Sahil2k07/Webhooks-Golang/webhook-server/interfaces"
	"github.com/Sahil2k07/Webhooks-Golang/webhook-server/views"
)

type webhookService struct {
	client http.Client
}

func NewWebHookService() interfaces.IWebhookService {
	return &webhookService{client: http.Client{}}
}

func (ws *webhookService) PaymentSuccessWebhook(payload views.PaymentRequest) {
	data := views.WebhookResponse{
		TransactionID: "random_trx_id",
		Success:       true,
		UserID:        payload.UserID,
	}

	response, err := json.Marshal(data)
	if err != nil {
		log.Println("Error while Marshal:", err)
		return
	}

	request, err := http.NewRequest(http.MethodPut, payload.WebhookURL, bytes.NewReader(response))
	if err != nil {
		log.Println("Error while forming request:", err)
		return
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("X-API-Token", "some_random_token")

	res, err := ws.client.Do(request)
	if err != nil {
		log.Println("Error making request:", err)
	}
	defer res.Body.Close()
}
