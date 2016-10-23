package gotrans

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io"
	"math"
	"net/http"
	"time"
)

const (
	// Default timeout is 10s in ns
	DefaultTimeout = 10000000000

	DevURL  = "https://api.sandbox.veritrans.co.id/v2"
	ProdURL = "https://api.veritrans.co.id/v2"

	SNAPDevURL  = "https://app.sandbox.midtrans.com"
	SNAPProdURL = "https://app.midtrans.com"

	ChargePath               = "/charge"
	GetTransactionStatusPath = "/status"
	CaptureTransactionPath   = "/capture"
	ApproveTransactionPath   = "/v2/{id}/approve"
	CancelTransactionPath    = "/v2/{id}/cancel"
	ExpireTransactionPath    = "/v2/{id}/expire"

	SNAPTransactionPath = "/snap/v1/transactions"

	ContentTypeHeader   = "Content-Type"
	AcceptHeader        = "Accept"
	AuthorizationHeader = "Authorization"

	ContentType = "application/json"
	Accept      = "application/json"

	ErrorValidationCode             = 400
	ErrorInvalidClientServerKeyCode = 401
	ErrorMerchantNoAccessCode       = 402
	ErrorDuplicateOrderIDCode       = 406
	ErrorAccountInactiveCode        = 410
)

type Gotrans struct {
	HttpClient   *http.Client
	ServerKey    string
	ClientKey    string
	IsProduction bool
	IsSanitized  bool
	Is3Dsecure   bool
	BaseURL      string
	SNAPBaseURL  string
}

func New(serverKey, clientKey string, isProduction, isSanitized, is3Dsecure bool, timeout time.Duration) *Gotrans {
	if timeout == 0 {
		timeout = DefaultTimeout
	}
	t := &Gotrans{
		HttpClient: &http.Client{
			Timeout: timeout,
		},
		ServerKey:    serverKey,
		ClientKey:    clientKey,
		IsProduction: isProduction,
		IsSanitized:  isSanitized,
		Is3Dsecure:   is3Dsecure,
		BaseURL:      DevURL,
		SNAPBaseURL:  SNAPDevURL,
	}

	if isProduction {
		t.BaseURL = ProdURL
		t.SNAPBaseURL = SNAPProdURL
	}

	return t
}

func (g *Gotrans) CreateBody(body interface{}) (io.Reader, error) {
	data, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	return bytes.NewBuffer(data), nil
}

func (g *Gotrans) EncodeBase64(data string) string {
	return base64.StdEncoding.EncodeToString([]byte(data))
}

func (g *Gotrans) CheckStatusCode(code int) bool {
	return !(code != 200 && code != 201 && code != 202 && code != 407)
}

func (g *Gotrans) CalculateGrossAmount(transaction *Transaction) *Transaction {
	// set to zero first
	transaction.TransactionDetails.GrossAmount = 0
	// then iterate through the item details
	total := 0.0
	for _, item := range transaction.ItemDetails {
		total += (item.Price * item.Quantity)
	}
	// then round up the final price
	transaction.TransactionDetails.GrossAmount = int64(math.Ceil(total))

	return transaction
}
