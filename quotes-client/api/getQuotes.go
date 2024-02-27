package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ResposeFromQuotes struct {
	Author string `json:"author"`
	Quote  string `json:"quote"`
}

// getQuotes отправляет запрос на получение цитат
func GetQuotes(baseUrl string, token string, nonce int64) error {
	url := fmt.Sprintf("%s/api/v1/quotes/?token=%s&nonce=%d", baseUrl, token, nonce)
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	quote := ResposeFromQuotes{}
	err = json.Unmarshal(body, &quote)
	fmt.Println("Received quote!")
	fmt.Printf("Author: %s\nQuote: %s\n", quote.Author, quote.Quote)
	fmt.Println()
	return nil
}
