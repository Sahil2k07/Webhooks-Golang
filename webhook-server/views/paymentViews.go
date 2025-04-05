package views

type (
	PaymentRequest struct {
		UserID     string  `json:"userID"`
		Amount     float32 `json:"amount"`
		Currency   string  `json:"currency"`
		WebhookURL string  `json:"webhookURL"`
	}

	PaymentResponse struct {
		TransactionID string `json:"transactionID"`
		Success       bool   `json:"success"`
	}
)
