package gotrans

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

type Transaction struct {
	PaymentType        string             `json:"payment_type,omitempty"`
	TransactionDetails *TransactionDetail `json:"transaction_details,omitempty"`
	BankTransfer       *BankTransfer      `json:"bank_transfer,omitempty"`
	ItemDetails        []*ItemDetail      `json:"item_detail,omitempty"`
	CustomerDetails    *CustomerDetail    `json:"customer_details,omitempty"`
	Echannel           *Echannel          `json:"echannel,omitempty"`
	CustomField1       string             `json:"custom_field1,omitempty"` // max 255 char
	CustomField2       string             `json:"custom_field2,omitempty"` // max 255 char
	CustomField3       string             `json:"custom_field3,omitempty"` // max 255 char
	VtWeb              *VtWeb             `json:"vtweb,omitempty"`
	BCAKlikPay         *BCAKlikPay        `json:"bca_klikpay,omitempty"`
	BCAKlikBCA         *BCAKlikBCA        `json:"bca_klikbca,omitempty"`
	MandiriClickPay    *MandiriClickPay   `json:"mandiri_clickpay,omitempty"`
	MandiriECash       *MandiriECash      `json:"mandiri_ecash,omitempty"`
	CIMBClicks         *CIMBClicks        `json:"cimb_clicks,omitempty"`
	TelkomselCash      *TelkomselCash     `json:"telkomsel_cash,omitempty"`
	IndosatDompetku    *IndosatDompetku   `json:"indosat_dompetku,omitempty"`
	ConvenienceStore   *ConvenienceStore  `json:"cstore,omitempty"`
	CreditCard         *CreditCardDetail  `json:"credit_card,omitempty"`
	CustomExpiry       *Expiry            `json:"custom_expiry,omitempty"`
}

func (t *Transaction) String() string {
	data, _ := json.MarshalIndent(t, "", "    ")
	return string(data)
}

type Expiry struct {
	StartTime      string `json:"start_time,omitempty"` // 2017-04-13 18:11:08 +0700
	ExpiryDuration int    `json:"expiry_duration,omitempty"`
	Unit           string `json:"unit,omitempty"`       // minute || hour || day
	OrderTime      string `json:"order_time,omitempty"` // 2017-04-13 18:11:08 +0700
}

type CreditCardDetail struct {
	TokenID         string       `json:"token_id,omitempty"`
	Bank            string       `json:"bank,omitempty"`
	InstallmentTerm int          `json:"installment_term,omitempty"`
	Bins            []string     `json:"bins,omitempty"`
	Type            string       `json:"type,omitempty"`
	SaveTokenID     bool         `json:"save_token_id,omitempty"`
	Channel         string       `json:"channel,omitempty"`
	Installment     *Installment `json:"installment,omitempty"`
	WhitelistBins   []string     `json:"whitelist_bins,omitempty"`
	Secure          bool         `json:"secure,omitempty"`
	SaveCard        bool         `json:"save_card,omitempty"`
}

type Installment struct {
	Required bool  `json:"required,omitempty"`
	Terms    *Term `json:"terms,omitempty"`
}

type Term struct {
	BNI     []int `json:"bni,omitempty"`
	Mandiri []int `json:"mandiri,omitempty"`
	BCA     []int `json:"bca,omitempty"`
	BRI     []int `json:"bri,omitempty"`
	CIMB    []int `json:"cimb,omitempty"`
	Danamon []int `json:"danamon,omitempty"`
	Maybank []int `json:"maybank,omitempty"`
	Offline []int `json:"offline,omitempty"`
}

type BCAKlikPay struct {
	Type        int    `json:"type,omitempty"`
	Description string `json:"description,omitempty"` // max 60 char
}

type BCAKlikBCA struct {
	Description string `json:"description,omitempty"` // max 60 char
	UserID      string `json:"user_id,omitempty"`     // max 12 char
}

type CIMBClicks struct {
	Description string `json:"description,omitempty"` // max 60 char
}

type TelkomselCash struct {
	Customer   string `json:"customer,omitempty"` // max 10 char
	Promo      bool   `json:"promo,omitempty"`
	IsReversal int    `json:"is_reversal,omitempty"`
}

type IndosatDompetku struct {
	MSISDN string `json:"msisdn,omitempty"`
}

type MandiriClickPay struct {
	CardNumber string `json:"card_number,omitempty"`
	Input1     string `json:"input1,omitempty"`
	Input2     string `json:"input2,omitempty"`
	Input3     string `json:"input3,omitempty"`
	Token      string `json:"token,omitempty"`
}

type MandiriECash struct {
	Description string `json:"description,omitempty"`
}

// use json:"cstore"
type ConvenienceStore struct {
	Store   string `json:"store,omitempty"`   // max 20 char
	Message string `json:"message,omitempty"` // max 20 char
}

type TransactionDetail struct {
	OrderId       string `json:"order_id,omitempty"` // limit of 50 chars
	GrossAmount   int64  `json:"gross_amount,omitempty"`
	TransactionID string `json:"transaction_id,omitempty"`
}

type BankTransfer struct {
	Bank     string       `json:"bank,omitempty"`
	VANumber string       `json:"va_number,omitempty,omitempty"` // max 255 char
	FreeText *BCAFreeText `json:"free_text,omitempty,omitempty"`
}

type BCAFreeText struct {
	Inquiry []*InquiryPayment `json:"inquiry,omitempty"`
	Payment []*InquiryPayment `json:"payment,omitempty"`
}

type InquiryPayment struct {
	ID string `json:"id,omitempty"`
	EN string `json:"en,omitempty"`
}

type ItemDetail struct {
	ID           string  `json:"id,omitempty"` // max 255 char
	Price        float64 `json:"price,omitempty"`
	Quantity     float64 `json:"quantity,omitempty"`
	Name         string  `json:"name,omitempty"`          // max 50 char
	Brand        string  `json:"brand,omitempty"`         // max 50 char
	Category     string  `json:"category,omitempty"`      // max 50 char
	MerchantName string  `json:"merchant_name,omitempty"` // max 50 char
}

type CustomerDetail struct {
	FirstName       string                  `json:"first_name,omitempty"` // max 20 char
	LastName        string                  `json:"last_name,omitempty"`  // max 20 char
	Email           string                  `json:"email,omitempty"`      // max 45 char
	Phone           string                  `json:"phone,omitempty"`      // max 19 char
	BillingAddress  *BillingShippingAddress `json:"billing_address,omitempty"`
	ShippingAddress *BillingShippingAddress `json:"shipping_address,omitempty"`
}

type BillingShippingAddress struct {
	FirstName   string `json:"first_name,omitempty"`   // max 20 char
	LastName    string `json:"last_name,omitempty"`    // max 20 char
	Email       string `json:"email,omitempty"`        // max 45 char
	Phone       string `json:"phone,omitempty"`        // max 19 char
	Address     string `json:"address,omitempty"`      // max 200 char
	City        string `json:"city,omitempty"`         // max 100 char
	PostalCode  string `json:"postal_code,omitempty"`  // max 10 char
	CountryCode string `json:"country_code,omitempty"` // 3 char
}

type Echannel struct {
	BillInfo1 string `json:"bill_info1,omitempty"` // max 20 char
	BillInfo2 string `json:"bill_info2,omitempty"` // max 20 char
	BillInfo3 string `json:"bill_info3,omitempty"` // max 20 char
	BillInfo4 string `json:"bill_info4,omitempty"` // max 20 char
	BillInfo5 string `json:"bill_info5,omitempty"` // max 20 char
	BillInfo6 string `json:"bill_info6,omitempty"` // max 20 char
	BillInfo7 string `json:"bill_info7,omitempty"` // max 20 char
	BillInfo8 string `json:"bill_info8,omitempty"` // max 20 char
}

type VtWeb struct {
	CreditCard3DSecure bool     `json:"credit_card_3d_secure,omitempty"`
	EnabledPayments    []string `json:"enabled_payments,omitempty"`
}

type Callback struct {
	Finish string `json:"finish,omitempty"`
}

func (g *Gotrans) CaptureTransaction(transaction *TransactionDetail) (APIResponse, error) {
	response := APIResponse{}

	body, err := g.CreateBody(transaction)
	if err != nil {
		return response, err
	}

	req, err := http.NewRequest(http.MethodPost, g.BaseURL+CaptureTransactionPath, body)
	if err != nil {
		return response, err
	}
	req.Header.Set(ContentTypeHeader, ContentType)
	req.Header.Set(AcceptHeader, Accept)
	req.Header.Set(AuthorizationHeader, "Basic "+g.EncodeBase64(g.ServerKey+":"))

	resp, err := g.HttpClient.Do(req)
	if err != nil {
		return response, err
	}
	defer resp.Body.Close()

	if !g.CheckStatusCode(resp.StatusCode) {
		return response, errors.New("Failed to send POST request: " + resp.Status)
	}

	rawResp, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(rawResp, &response)
	if err != nil {
		return response, err
	}

	return response, nil
}

func (g *Gotrans) ApproveTransaction(orderOrTransactionID string) (APIResponse, error) {
	response := APIResponse{}
	targetURL := strings.Replace(g.BaseURL+ApproveTransactionPath, "{id}", orderOrTransactionID, 1)

	req, err := http.NewRequest(http.MethodPost, targetURL, nil)
	if err != nil {
		return response, err
	}
	req.Header.Set(ContentTypeHeader, ContentType)
	req.Header.Set(AcceptHeader, Accept)
	req.Header.Set(AuthorizationHeader, "Basic "+g.EncodeBase64(g.ServerKey+":"))

	resp, err := g.HttpClient.Do(req)
	if err != nil {
		return response, err
	}
	defer resp.Body.Close()

	if !g.CheckStatusCode(resp.StatusCode) {
		return response, errors.New("Failed to send POST request: " + resp.Status)
	}

	rawResp, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(rawResp, &response)
	if err != nil {
		return response, err
	}

	return response, nil
}

func (g *Gotrans) CancelTransaction(orderOrTransactionID string) (APIResponse, error) {
	response := APIResponse{}
	targetURL := strings.Replace(g.BaseURL+CancelTransactionPath, "{id}", orderOrTransactionID, 1)

	req, err := http.NewRequest(http.MethodPost, targetURL, nil)
	if err != nil {
		return response, err
	}
	req.Header.Set(ContentTypeHeader, ContentType)
	req.Header.Set(AcceptHeader, Accept)
	req.Header.Set(AuthorizationHeader, "Basic "+g.EncodeBase64(g.ServerKey+":"))

	resp, err := g.HttpClient.Do(req)
	if err != nil {
		return response, err
	}
	defer resp.Body.Close()

	if !g.CheckStatusCode(resp.StatusCode) {
		return response, errors.New("Failed to send POST request: " + resp.Status)
	}

	rawResp, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(rawResp, &response)
	if err != nil {
		return response, err
	}

	return response, nil
}

func (g *Gotrans) ExpireTransaction(orderOrTransactionID string) (APIResponse, error) {
	response := APIResponse{}
	targetURL := strings.Replace(g.BaseURL+ExpireTransactionPath, "{id}", orderOrTransactionID, 1)

	req, err := http.NewRequest(http.MethodPost, targetURL, nil)
	if err != nil {
		return response, err
	}
	req.Header.Set(ContentTypeHeader, ContentType)
	req.Header.Set(AcceptHeader, Accept)
	req.Header.Set(AuthorizationHeader, "Basic "+g.EncodeBase64(g.ServerKey+":"))

	resp, err := g.HttpClient.Do(req)
	if err != nil {
		return response, err
	}
	defer resp.Body.Close()

	if !g.CheckStatusCode(resp.StatusCode) {
		return response, errors.New("Failed to send POST request: " + resp.Status)
	}

	rawResp, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(rawResp, &response)
	if err != nil {
		return response, err
	}

	return response, nil
}
