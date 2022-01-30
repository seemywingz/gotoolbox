package coinbase

import "fmt"

const apiURL = "https://api.coinbase.com"
const sandboxURL = "https://api-public.sandbox.exchange.coinbase.com"

var baseURL string

func GetAccount(sandbox bool) {
	setBaseURL(sandbox)

	fmt.Println("Using Base URL:", baseURL)

}

func setBaseURL(sandbox bool) {
	if sandbox {
		baseURL = sandboxURL
	} else {
		baseURL = apiURL
	}
}