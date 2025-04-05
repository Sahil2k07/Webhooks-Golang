package services

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"time"

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

	go ws.sendWebhookCall(request)
}

func (ws *webhookService) sendWebhookCall(request *http.Request) {
	retryCount, ticks := 0, 1

	for retryCount < 6 {
		retryCount++

		res, err := ws.client.Do(request)
		if err != nil {
			log.Println("Error making request:", err)
		} else {
			defer res.Body.Close()
		}

		if res.StatusCode == http.StatusOK {
			break
		}

		wait := time.Duration(min(20000, 1000*ticks)) * time.Millisecond
		time.Sleep(wait)

		ticks *= 2
	}
}

func min(val1, val2 int) int {
	if val1 > val2 {
		return val2
	}

	return val1
}
