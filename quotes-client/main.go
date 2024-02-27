package main

import (
	"fmt"
	"pow-quotes-client/api"
	"pow-quotes-client/cmd"
	"pow-quotes-client/pow"
)

func main() {
	config, err := cmd.ParseFlags()
	if err != nil {
		fmt.Println(err)
		return
	}

	powUrl := fmt.Sprintf("%s/api/v1/pow/", config.BaseURL)

	// getting PoW data
	powData, err := api.GetPoWData(powUrl)
	if err != nil {
		fmt.Println("Error getting PoW data:", err)
		return
	}

	// Searching for a nonce
	nonce := pow.FindNonce(powData.Data, powData.LeadingZerosCount)
	fmt.Printf("Found nonce: %d\n", nonce)

	// Getting quotes
	err = api.GetQuotes(config.BaseURL, powData.Token, nonce)
	if err != nil {
		fmt.Println("Error getting quotes:", err)
		return
	}
}
