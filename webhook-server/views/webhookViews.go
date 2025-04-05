package views

type (
	WebhookResponse struct {
		TransactionID string `json:"transactionID"`
		Success       bool   `json:"success"`
		UserID        string `json:"userID"`
	}
)
