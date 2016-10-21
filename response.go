package gotrans

type APIResponse struct {
	StatusCode         string   `json:"status_code"`
	StatusMessage      string   `json:"status_message"`
	TransactionID      string   `json:"transaction_id,omitempty"`
	OrderID            string   `json:"order_id,omitempty"`
	GrossAmount        string   `json:"gross_amount,omitempty"`
	PaymentType        string   `json:"payment_type,omitempty"`
	TransactionTime    string   `json:"transaction_time,omitempty"`
	TransactionStatus  string   `json:"transaction_status,omitempty"`
	FraudStatus        string   `json:"fraid_status,omitempty"`
	MaskedCard         string   `json:"masked_card,omitempty"`
	Bank               string   `json:"bank,omitempty"`
	RedirectURL        string   `json:"redirect_url,omitempty"`
	ValidationMessages []string `json:"validation_messages,omitempty"`
	TokenID            string   `json:"token_id,omitempty"`
	ECI                string   `json:"eci,omitempty"`
}
