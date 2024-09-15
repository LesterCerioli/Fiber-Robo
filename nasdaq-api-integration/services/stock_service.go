package services

import (
	"encoding/json"
	"fmt"
	"nasdaq-api-integration/config"
	"net/http"
)

type StockData struct {
	Dataset struct {
		Ticker      string          `json:"dataset_code"`
		Name        string          `json:"name"`
		Data        [][]interface{} `json:"data"`
		ColumnNames []string        `json:"column_names"`
	} `json:"dataset"`
}

func FetchStockData(symbol string) (*StockData, error) {
	apiURL := fmt.Sprintf("https://data.nasdaq.com/api/v3/datasets/WIKI/%s.json?api_key=%s", symbol, config.ApiKey)
	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, fmt.Errorf("error fetching stock data: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error fetching stock data: status code %d", resp.StatusCode)
	}

	var stockData StockData
	if err := json.NewDecoder(resp.Body).Decode(&stockData); err != nil {
		return nil, fmt.Errorf("error decoding stock data: %v", err)
	}

	return &stockData, nil
}
