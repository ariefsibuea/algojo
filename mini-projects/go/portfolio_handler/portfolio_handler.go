package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Position struct {
	Coin   string  `json:"coin"`
	Amount float64 `json:"amount"`
}

type PortfolioResponse struct {
	UserID    string     `json:"userId"`
	Positions []Position `json:"positions"`
}

type portfolioHandler struct {
	data []PortfolioResponse
}

func PortfolioListHandler(data []PortfolioResponse) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !strings.HasPrefix(r.URL.Path, ExactPortfolio) {
			WriteJSON(w, http.StatusNotFound, ErrorResponse{Error: "endpoint not found"})
			return
		}

		// NOTE: Currently only support pattern /portfolio?userId={userID}
		userID := r.URL.Query().Get("userId")
		if userID == "" {
			WriteJSON(w, http.StatusBadRequest, ErrorResponse{Error: "missing user id"})
			return
		}

		h := &portfolioHandler{data: data}

		switch r.Method {
		case http.MethodGet:
			h.getPortfolioByUserID(w, userID)
		default:
			WriteJSON(w, http.StatusMethodNotAllowed, ErrorResponse{Error: "method not allowed"})
		}
	}
}

func (h *portfolioHandler) getPortfolioByUserID(w http.ResponseWriter, userID string) {
	var result PortfolioResponse

	for _, v := range h.data {
		if v.UserID == userID {
			result = v
		}
	}

	if result.UserID == "" {
		WriteJSON(w, http.StatusNotFound, ErrorResponse{Error: fmt.Sprintf("user %s not found", userID)})
		return
	}

	for i, v := range result.Positions {
		coinPrice, err := getCoinPrice(fmt.Sprintf("%s-USD", v.Coin))
		if err != nil {
			WriteJSON(w, http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
			return
		}

		amount, _ := strconv.ParseFloat(coinPrice.Data.Amount, 64)
		result.Positions[i].Amount = amount
	}

	WriteJSON(w, http.StatusOK, result)
}

type CoinPrice struct {
	Data struct {
		Amount   string `json:"amount"`
		Currency string `json:"currency"`
	} `json:"data"`
}

func getCoinPrice(currencyPair string) (CoinPrice, error) {
	log.Printf("handle request get coin price for currency pair: %s ...\n", currencyPair)

	var coinPrice CoinPrice

	requestUrl := fmt.Sprintf("https://api.coinbase.com/v2/prices/%s/spot", currencyPair)

	req, err := http.NewRequest(http.MethodGet, requestUrl, nil)
	if err != nil {
		return coinPrice, err
	}

	transport := &http.Transport{
		TLSHandshakeTimeout: 3 * time.Second,
	}

	client := &http.Client{
		Timeout:   5 * time.Second,
		Transport: transport,
	}

	res, err := client.Do(req)
	if err != nil {
		return coinPrice, err
	}
	defer res.Body.Close()

	if res.StatusCode < 200 || res.StatusCode >= 400 {
		resBody, errRead := io.ReadAll(res.Body)
		if errRead != nil {
			return coinPrice, fmt.Errorf("request failed with status code %d and error reading response body: %v", res.StatusCode, errRead)
		}
		return coinPrice, fmt.Errorf("request failed with status code %d: %s", res.StatusCode, string(resBody))
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return coinPrice, err
	}

	if err := json.Unmarshal(resBody, &coinPrice); err != nil {
		return coinPrice, fmt.Errorf("failed to parse response body: %v", err)
	}

	return coinPrice, nil
}
