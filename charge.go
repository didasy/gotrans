package gotrans

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

func (g *Gotrans) Charge(targetUrl, serverKey string, transaction *Transaction) (APIResponse, error) {
	response := APIResponse{}

	body, err := g.CreateBody(transaction)
	if err != nil {
		return response, err
	}

	req, err := http.NewRequest(http.MethodPost, targetUrl, body)
	if err != nil {
		return response, err
	}
	req.Header.Set(ContentTypeHeader, ContentType)
	req.Header.Set(AcceptHeader, Accept)
	req.Header.Set(AuthorizationHeader, "Basic "+g.EncodeBase64(serverKey+":"))

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

func (g *Gotrans) ChargeBankTransfer(transaction *Transaction) (APIResponse, error) {
	transaction.PaymentType = "bank_transfer"

	// recalculate gross_amount if there is ItemDetails
	if len(transaction.ItemDetails) > 0 {
		transaction = g.CalculateGrossAmount(transaction)
	}

	// sanitize
	if g.IsSanitized {
		// to be created
	}

	return g.Charge(g.BaseURL+ChargePath, g.ServerKey, transaction)
}

func (g *Gotrans) ChargeVtWeb(transaction *Transaction) (APIResponse, error) {
	transaction.PaymentType = "vtweb"
	transaction.VtWeb = &VtWeb{
		CreditCard3DSecure: true,
	}

	// recalculate gross_amount if there is ItemDetails
	if len(transaction.ItemDetails) > 0 {
		transaction = g.CalculateGrossAmount(transaction)
	}

	// sanitize
	if g.IsSanitized {
		// to be created
	}

	return g.Charge(g.BaseURL+ChargePath, g.ServerKey, transaction)
}
