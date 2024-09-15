# Nasdaq API Integration with Go and Fiber

This project is a Go-based backend that integrates with the Nasdaq API to fetch stock data and exchange rates. It uses the Fiber web framework for building RESTful APIs, providing endpoints to retrieve stock market information and exchange rates for currencies like USD and EUR.

## Features

- **Fetch Stock Data**: Retrieve stock data from the Nasdaq API for any given symbol.
- **Get Exchange Rates**: Fetch exchange rates (USD, EUR, etc.) from the Nasdaq API.
- **Fiber Web Framework**: A lightweight, fast web server built using Fiber, inspired by Express.js.
- **Error Handling**: Graceful error handling for failed requests and API communication.
- **Environment Configuration**: Load API keys and other sensitive data from environment variables.

## Project Structure

```bash
/nasdaq-api-integration
  /config
    config.go                # Configuration file to load API keys
  /services
    stock_service.go         # Service for fetching stock data
    exchange_rate.go         # Service for fetching exchange rates
  main.go                    # Application entry point
  go.mod                     # Go module dependencies
  go.sum                     # Go module dependency checksums

## Requirements:

- ** Go 1.17+
- ** Nasdaq API key (Get API Key)
- ** Fiber framework
- ** HTTP client