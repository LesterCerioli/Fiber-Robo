package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"

    "github.com/gofiber/fiber/v2"
)


type StockData struct {
    Dataset struct {
        Ticker   string    `json:"dataset_code"`
        Name     string    `json:"name"`
        Data     [][]interface{} `json:"data"`
        ColumnNames []string  `json:"column_names"`
    } `json:"dataset"`
}
const apiKey = "YOUR_API_KEY"
const apiURL = "https://data.nasdaq.com/api/v3/datasets/WIKI/%s.json?api_key=" + apiKey

func main() {
    
    app := fiber.New()

    
    app.Get("/stock/:symbol", func(c *fiber.Ctx) error {
        symbol := c.Params("symbol")

        
        stockData, err := fetchStockData(symbol)
        if err != nil {
            return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
        }

        return c.JSON(stockData)
    })

    
    log.Fatal(app.Listen(":3000"))
}


func fetchStockData(symbol string) (*StockData, error) {
    url := fmt.Sprintf(apiURL, symbol)

    
    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    
    var stockData StockData
    if err := json.NewDecoder(resp.Body).Decode(&stockData); err != nil {
        return nil, err
    }

    return &stockData, nil
}