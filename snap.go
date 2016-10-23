package gotrans

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type SNAPTransaction struct {
	TransactionDetails *TransactionDetail `json:"transaction_details,omitempty"`
	EnabledPayments    []string           `json:"enabled_payments,omitempty"`
	CustomerDetails    *CustomerDetail    `json:"customer_details,omitempty"`
	Callbacks          *Callback          `json:"callbacks,omitempty"`
	CreditCard         *CreditCardDetail  `json:"credit_card,omitempty"`
	Expiry             *Expiry            `json:"expiry,omitempty"`
}

func (s *SNAPTransaction) String() string {
	data, _ := json.MarshalIndent(s, "", "    ")
	return string(data)
}

type SNAPResponse struct {
	ErrorMessages []string `json:"error_messages,omitempty"`
	Token         string   `json:"token"`
}

func (s SNAPResponse) String() string {
	data, _ := json.MarshalIndent(s, "", "    ")
	return string(data)
}

func (g *Gotrans) GetSNAPToken(transaction *SNAPTransaction) (SNAPResponse, error) {
	response := SNAPResponse{}
	targetURL := g.SNAPBaseURL + SNAPTransactionPath

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
