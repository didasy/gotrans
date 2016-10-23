package gotrans

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

func (g *Gotrans) Charge(targetURL, serverKey string, transaction *Transaction) (APIResponse, error) {
	response := APIResponse{}

	body, err := g.CreateBody(transaction)
	if err != nil {
		return response, err
	}

	req, err := http.NewRequest(http.MethodPost, targetURL, body)
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

func (g *Gotrans) ChargeCreditCard(transaction *Transaction) (APIResponse, error) {
	transaction.PaymentType = "credit_card"
	if transaction.CreditCard == nil {
		return APIResponse{}, errors.New("CreditCard property must be exists")
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

func (g *Gotrans) ChargeConvenienceStore(transaction *Transaction) (APIResponse, error) {
	transaction.PaymentType = "cstore"
	if transaction.ConvenienceStore == nil {
		return APIResponse{}, errors.New("ConvenienceStore property must be exists")
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

func (g *Gotrans) ChargeEWallet(transaction *Transaction) (APIResponse, error) {
	response := APIResponse{}

	switch transaction.PaymentType {
	case "telkomsel_cash":
		if transaction.TelkomselCash == nil {
			return APIResponse{}, errors.New("TelkomselCash property must be exists")
		}
		break
	case "xl_tunai":
		break
	case "indosat_dompetku":
		if transaction.IndosatDompetku == nil {
			return APIResponse{}, errors.New("IndosatDompetku property must be exists")
		}
		break
	case "mandiri_ecash":
		if transaction.MandiriECash == nil {
			return APIResponse{}, errors.New("MandiriECash property must be exists")
		}
		break
	default:
		return response, errors.New("Invalid payment type: must be telkomsel_cash, xl_tunai, indosat_dompetku, or mandiri_ecash")
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

func (g *Gotrans) ChargeInternetBanking(transaction *Transaction) (APIResponse, error) {
	response := APIResponse{}

	switch transaction.PaymentType {
	case "bca_klikpay":
		if transaction.BCAKlikPay == nil {
			return APIResponse{}, errors.New("BCAKlikPay property must be exists")
		}
		break
	case "bca_klikbca":
		if transaction.BCAKlikBCA == nil {
			return APIResponse{}, errors.New("BCAKlikBCA property must be exists")
		}
		break
	case "mandiri_clickpay":
		if transaction.MandiriClickPay == nil {
			return APIResponse{}, errors.New("MandiriClickPay property must be exists")
		}
		break
	case "bri_epay":
		break
	case "cimb_clicks":
		if transaction.CIMBClicks == nil {
			return APIResponse{}, errors.New("CIMBClicks property must be exists")
		}
		break
	default:
		return response, errors.New("Invalid payment type: must be bca_klikpay, bca_klikpay, mandiri_clickpay, bri_epay, or cimb_clicks")
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

func (g *Gotrans) ChargeBankTransfer(transaction *Transaction) (APIResponse, error) {
	switch transaction.PaymentType {
	case "bank_transfer":
		if transaction.BankTransfer == nil {
			return APIResponse{}, errors.New("BankTransfer property must be exists")
		}
		break
	case "echannel":
		if transaction.Echannel == nil {
			return APIResponse{}, errors.New("Echannel property must be exists")
		}
		break
	default:
		return APIResponse{}, errors.New("Invalid payment type: must be bank_transfer or echannel")
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
