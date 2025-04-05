package views

type (
	AccountUpgradeRequest struct {
		TransactionID string `json:"transactionID"`
		Success       bool   `json:"success"`
		UserID        string `json:"userID"`
	}

	PaymentWebhookRequest struct {
		UserID     string  `json:"userID"`
		Amount     float32 `json:"amount"`
		Currency   string  `json:"currency"`
		WebhookURL string  `json:"webhookURL"`
	}

	PaymentProcessResponse struct {
		TransactionID string `json:"transactionID"`
		Success       bool   `json:"success"`
	}
)
