package gotrans

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func (g *Gotrans) HandleNotification(notificationJSONBody []byte) (APIResponse, error) {
	resp := APIResponse{}
	err := json.Unmarshal(notificationJSONBody, resp)
	if err != nil {
		return resp, err
	}

	return g.GetTransactionStatus(g.BaseURL+GetTransactionStatusPath, resp.TransactionID, g.ServerKey)
}

func (g *Gotrans) GetTransactionStatus(targetURL, transactionID, serverKey string) (APIResponse, error) {
	response := APIResponse{}

	req, err := http.NewRequest(http.MethodGet, targetURL, nil)
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
