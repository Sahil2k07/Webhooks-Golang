package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Sahil2k07/Webhooks-Golang/main-server/interfaces"
	"github.com/Sahil2k07/Webhooks-Golang/main-server/views"
	"github.com/labstack/echo"
)

type accountService struct {
	ctx    echo.Context
	client http.Client
}

func NewAccountService(c echo.Context) interfaces.IAccountService {
	return &accountService{
		ctx:    c,
		client: http.Client{},
	}
}

func (as *accountService) UpgradeUserAccount() error {
	var request views.AccountUpgradeRequest

	err := as.ctx.Bind(&request)
	if err != nil {
		return err
	}

	// Perform some logic

	fmt.Println("Request data:", request)

	return nil
}

func (as *accountService) MakeUserPayment() error {
	reqPayload := views.PaymentWebhookRequest{
		UserID:     "random-id",
		Amount:     10,
		Currency:   "INR",
		WebhookURL: "http://" + as.ctx.Request().Host + "/api/account/upgrade/webhook",
	}

	reqBody, err := json.Marshal(reqPayload)
	if err != nil {
		return err
	}

	paymentExternalAPI := "http://localhost:4000/api/payment" // Hard coded for now

	request, err := http.NewRequest(http.MethodPost, paymentExternalAPI, bytes.NewReader(reqBody))
	if err != nil {
		return err
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("X-API-Token", "xxxxxxxx")

	r, err := as.client.Do(request)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	var response views.PaymentProcessResponse

	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bodyBytes, &response)
	if err != nil {
		return err
	}

	return as.ctx.JSON(http.StatusOK, response)
}
