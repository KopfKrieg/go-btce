package btce

import (
	"encoding/json"
	"net/http"
	"strconv"
)

// PublicAPI provides access to such information as tickers of currency pairs, active orders on different pairs, the latest trades for each pair etc.
type PublicAPI struct{}

const API_URL = "https://btc-e.com/api/3/"

// Info provides all the information about currently active pairs, such as the maximum number of digits after the decimal point, the minimum price, the maximum price, the minimum transaction size, whether the pair is hidden, the commission for each pair.
func (api *PublicAPI) Info() (Info, error) {

	url := API_URL + "info"
	r, err := http.Get(url)

	if err == nil {
		data := Info{}
		json.NewDecoder(r.Body).Decode(&data)
		return data, nil
	}
	return Info{}, err
}

// Ticker provides all the information about currently active pairs, such as: the maximum price, the minimum price, average price, trade volume, trade volume in currency, the last trade, Buy and Sell price.
// All information is provided over the past 24 hours.
func (api *PublicAPI) Ticker(currency []string) (Ticker, error) {

	url := API_URL + "ticker/"
	for _, c := range currency {
		url = url + c + "-"
	}
	r, err := http.Get(url)

	if err == nil {
		data := make(Ticker, len(currency))
		json.NewDecoder(r.Body).Decode(&data)
		return data, nil
	}

	return nil, err
}

// Depth provides the information about active orders on the pair.
func (api *PublicAPI) Depth(currency []string, limit int) (Depth, error) {

	url := API_URL + "depth/"
	for _, c := range currency {
		url = url + c + "-"
	}
	if limit > 0 {
		url = url + "?limit=" + strconv.Itoa(limit)
	}
	r, err := http.Get(url)

	if err == nil {
		data := make(Depth, len(currency))
		json.NewDecoder(r.Body).Decode(&data)
		return data, nil
	}

	return nil, err

}

// Trades provides the information about the last trades.
func (api *PublicAPI) Trades(currency []string, limit int) (Trades, error) {

	url := API_URL + "trades/"
	for _, c := range currency {
		url = url + c + "-"
	}
	if limit > 0 {
		url = url + "?limit=" + strconv.Itoa(limit)
	}
	r, err := http.Get(url)

	if err == nil {
		data := make(Trades, len(currency))
		json.NewDecoder(r.Body).Decode(&data)
		return data, nil
	}

	return nil, err
}