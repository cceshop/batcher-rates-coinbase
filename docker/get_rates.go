package main

import (
	"encoding/json"
	"github.com/go-redis/redis"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type PriceCoinBase struct {
	Data struct {
		Base     string `json:"base"`
		Currency string `json:"currency"`
		Amount   string `json:"amount"`
	} `json:"data"`
}

func WriteToCache(price *PriceCoinBase) {
	client := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs:       []string{":26379", "16379"},
		MasterName:  "mymaster",
		DialTimeout: 3 * time.Second,
		Password:    "pjneUIy9RC",
		DB:          0, // use default DB
		MaxRetries:  6,
	})
	defer client.Close()

	ctx := client.Context()

	_, err := client.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}

	floatPrice, err := strconv.ParseFloat(price.Data.Amount, 64)
	if err != nil {
		panic(err)
	}

	err = client.Set(ctx, price.Data.Base, floatPrice, time.Duration(300*time.Second)).Err()
	if err != nil {
		panic(err)
	}
}

func GetExchangeRatesFromCoinbase() []string {
	urls := []string{"https://api.coinbase.com/v2/prices/BTC-CZK/buy",
		"https://api.coinbase.com/v2/prices/ETH-CZK/buy",
		"https://api.coinbase.com/v2/prices/LTC-CZK/buy"}
	var contents []byte
	var results []string

	for _, url := range urls {
		timeout := time.Duration(3 * time.Second)
		client := http.Client{
			Timeout: timeout,
		}

		resp, err := client.Get(url)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		contents, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}

		results = append(results, string(contents))
	}

	return results
}

func main() {
	var price PriceCoinBase
	var prices []PriceCoinBase
	var exchangeRates []string

	// get exchange rates
	exchangeRates = GetExchangeRatesFromCoinbase()
	for _, exchangeRate := range exchangeRates {
		json.Unmarshal([]byte(exchangeRate), &price)
		prices = append(prices, price)
	}

	// write exchange rates to a cache
	for _, p := range prices {
		WriteToCache(&p)
	}
}
