package gotrans

type Transaction struct {
	PaymentType        string            `json:"payment_type,omitempty"`
	TransactionDetails TransactionDetail `json:"transaction_details"`
	BankTransfer       BankTransfer      `json:"bank_transfer,omitempty"`
	ItemDetails        []ItemDetail      `json:"item_detail,omitempty"`
	CustomerDetail     CustomerDetail    `json:"customer_detail,omitempty"`
	Echannel           Echannel          `json:"echannel,omitempty"`
	CustomField1       string            `json:"custom_field1,omitempty"` // max 255 char
	CustomField2       string            `json:"custom_field2,omitempty"` // max 255 char
	CustomField3       string            `json:"custom_field3,omitempty"` // max 255 char
	VtWeb              VtWeb             `json:"vtweb,omitempty"`
}

type TransactionDetail struct {
	OrderId     string `json:"order_id"` // limit of 50 chars
	GrossAmount int64  `json:"gross_amount"`
}

type BankTransfer struct {
	Bank string `json:"bank"`
}

type ItemDetail struct {
	ID           string  `json:"id"` // max 255 char
	Price        float64 `json:"price"`
	Quantity     float64 `json:"quantity"`
	Name         string  `json:"name"`          // max 50 char
	Brand        string  `json:"brand"`         // max 50 char
	Category     string  `json:"category"`      // max 50 char
	MerchantName string  `json:"merchant_name"` // max 50 char
}

type CustomerDetail struct {
	FirstName       string                 `json:"first_name"`          // max 20 char
	LastName        string                 `json:"last_name,omitempty"` // max 20 char
	Email           string                 `json:"email,omitempty"`     // max 45 char
	Phone           string                 `json:"phone,omitempty"`     // max 19 char
	BillingAddress  BillingShippingAddress `json:"billing_address,omitempty"`
	ShippingAddress BillingShippingAddress `json:"shipping_address,omitempty"`
}

type BillingShippingAddress struct {
	FirstName   string `json:"first_name"`            // max 20 char
	LastName    string `json:"last_name,omitempty"`   // max 20 char
	Email       string `json:"email,omitempty"`       // max 45 char
	Phone       string `json:"phone,omitempty"`       // max 19 char
	Address     string `json:"address"`               // max 200 char
	City        string `json:"city"`                  // max 100 char
	PostalCode  string `json:"postal_code,omitempty"` // max 10 char
	CountryCode string `json:"country_code"`          // 3 char
}

type Echannel struct {
	BillInfo1 string `json:"bill_info1"` // max 20 char
	BillInfo2 string `json:"bill_info2"` // max 20 char
	BillInfo3 string `json:"bill_info3"` // max 20 char
	BillInfo4 string `json:"bill_info4"` // max 20 char
	BillInfo5 string `json:"bill_info5"` // max 20 char
	BillInfo6 string `json:"bill_info6"` // max 20 char
	BillInfo7 string `json:"bill_info7"` // max 20 char
	BillInfo8 string `json:"bill_info8"` // max 20 char
}

type VtWeb struct {
	CreditCard3DSecure bool     `json:"credit_card_3d_secure,omitempty"`
	EnabledPayments    []string `json:"enabled_payments,omitempty"`
}
