package api

import (
	"encoding/json"
	"io"
	"net/http"
	"pow-quotes-client/requests"
)

// getPoWData получает данные для PoW
func GetPoWData(url string) (*requests.PoWData, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var powData requests.PoWData
	err = json.Unmarshal(body, &powData)
	if err != nil {
		return nil, err
	}

	return &powData, nil
}
