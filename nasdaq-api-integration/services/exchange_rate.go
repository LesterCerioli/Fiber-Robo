package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"nasdaq-api-integration/config"
)

type ExchangeRateResponse struct {
	Dataset struct {
		Name        string        `json:"name"`
		ColumnNames []string      `json:"column_names"`
		Data        [][]interface{} `json:"data"`
	} `json:"dataset"`
}


func GetExchangeRate(currency string) (float64, error) {
	apiURL := fmt.Sprintf("https://data.nasdaq.com/api/v3/datasets/CUR/%s.json?api_key=%s", currency, config.ApiKey)
	resp, err := http.Get(apiURL)
	if err != nil {
		return 0, fmt.Errorf("error fetching exchange rate: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("error fetching exchange rate: status code %d", resp.StatusCode)
	}

	var exchangeRateResponse ExchangeRateResponse
	if err := json.NewDecoder(resp.Body).Decode(&exchangeRateResponse); err != nil {
		return 0, fmt.Errorf("error decoding exchange rate data: %v", err)
	}

	
	if len(exchangeRateResponse.Dataset.Data) == 0 || len(exchangeRateResponse.Dataset.Data[0]) == 0 {
		return 0, fmt.Errorf("no exchange rate data available for currency: %s", currency)
	}

	rate, ok := exchangeRateResponse.Dataset.Data[0][1].(float64)
	if !ok {
		return 0, fmt.Errorf("invalid exchange rate format for currency: %s", currency)
	}

	return rate, nil
}