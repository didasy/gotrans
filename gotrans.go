package gotrans

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"math"
	"net/http"
	"time"
)

const (
	// Default timeout is 10s in ns
	DefaultTimeout      = 10000000000
	DevURL              = "https://api.sandbox.veritrans.co.id/v2"
	ProdURL             = "https://api.veritrans.co.id/v2"
	ChargePath          = "/charge"
	ContentTypeHeader   = "Content-Type"
	AcceptHeader        = "Accept"
	AuthorizationHeader = "Authorization"
	ContentType         = "application/json"
	Accept              = "application/json"
)

type Gotrans struct {
	HttpClient   *http.Client
	ServerKey    string
	ClientKey    string
	IsProduction bool
	IsSanitized  bool
	Is3Dsecure   bool
	BaseURL      string
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
	}

	if isProduction {
		t.BaseURL = ProdURL
	}

	return t
}

func (g *Gotrans) CreateBody(transaction Transaction) (io.Reader, error) {
	data, err := json.Marshal(transaction)
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

func (g *Gotrans) Post(targetUrl, serverKey string, transaction Transaction) (string, error) {
	body, err := g.CreateBody(transaction)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest(http.MethodPost, targetUrl, body)
	if err != nil {
		return "", err
	}
	req.Header.Set(ContentTypeHeader, ContentType)
	req.Header.Set(AcceptHeader, Accept)
	req.Header.Set(AuthorizationHeader, "Basic "+g.EncodeBase64(serverKey+":"))

	resp, err := g.HttpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if !g.CheckStatusCode(resp.StatusCode) {
		return "", errors.New("Failed to send POST request: " + resp.Status)
	}

	rawResp, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	response := APIResponse{}
	err = json.Unmarshal(rawResp, &response)
	if err != nil {
		return "", err
	}

	return response.RedirectURL, nil
}

func (g *Gotrans) GetVtWebRedirectionUrl(transaction Transaction) (string, error) {
	transaction.PaymentType = "vtweb"
	transaction.VtWeb = VtWeb{
		CreditCard3DSecure: true,
	}

	// calculate gross_amount if there is ItemDetails
	if len(transaction.ItemDetails) > 0 {
		// set to zero first
		transaction.TransactionDetails.GrossAmount = 0
		// then iterate through the item details
		total := 0.0
		for _, item := range transaction.ItemDetails {
			total += (item.Price * item.Quantity)
		}
		// then round up the final price
		transaction.TransactionDetails.GrossAmount = int64(math.Ceil(total))
	}

	// sanitize
	if g.IsSanitized {
		// to be created
	}

	return g.Post(g.BaseURL+ChargePath, g.ServerKey, transaction)
}
