package gotrans

import "encoding/json"

type APIResponse struct {
	StatusCode           string      `json:"status_code"`
	StatusMessage        string      `json:"status_message"`
	TransactionID        string      `json:"transaction_id,omitempty"`
	OrderID              string      `json:"order_id,omitempty"`
	GrossAmount          string      `json:"gross_amount,omitempty"`
	PaymentType          string      `json:"payment_type,omitempty"`
	TransactionTime      string      `json:"transaction_time,omitempty"`
	TransactionStatus    string      `json:"transaction_status,omitempty"`
	FraudStatus          string      `json:"fraid_status,omitempty"`
	MaskedCard           string      `json:"masked_card,omitempty"`
	Bank                 string      `json:"bank,omitempty"`
	RedirectURL          string      `json:"redirect_url,omitempty"`
	ValidationMessages   []string    `json:"validation_messages,omitempty"`
	TokenID              string      `json:"token_id,omitempty"`
	ECI                  string      `json:"eci,omitempty"`
	VANumbers            []*VADetail `json:"va_numbers,omitempty"`
	BillerCode           string      `json:"biller_code,omitempty"`
	BillKey              string      `json:"bill_key,omitempty"`
	ApprovalCode         string      `json:"approval_code,omitempty"`
	SaveTokenID          string      `json:"save_token_id,omitempty"`
	SaveTokenIDExpiredAt string      `json:"save_token_id_expired_at,omitempty"`
	XLTunaiOrderID       string      `json:"xl_tunai_order_id,omitempty"`
	XLTunaiMerchantID    string      `json:"xl_tunai_merchant_id,omitempty"`
}

func (a APIResponse) String() string {
	data, _ := json.MarshalIndent(a, "", "    ")
	return string(data)
}

type VADetail struct {
	Bank     string `json:"bank,omitempty"`
	VANumber string `json:"va_number,omitempty"`
}
