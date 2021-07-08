package web

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func SendBTCRequest(baseURL string) (string, error) {
	resp, err := http.Get(baseURL + "/api/getCurrent")
	if err != nil {
		return "", fmt.Errorf("Failed to perform HTTP Request to the server: %s", err.Error())
	}
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Received non-200 status code: %d", resp.StatusCode)
	}

	bytes := new(bytes.Buffer)
	if _, err := bytes.ReadFrom(resp.Body); err != nil {
		return "", fmt.Errorf("Failed to send buffer to bytes: %s", err.Error())
	}

	respStruct := CurrentPriceResponse{}
	if err := json.Unmarshal(bytes.Bytes(), &respStruct); err != nil {
		return "", fmt.Errorf("Failed to unmarshall JSON: %s", err.Error())
	}

	return respStruct.Price, nil

}
