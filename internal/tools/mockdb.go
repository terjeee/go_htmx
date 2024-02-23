package tools

import (
	"time"
)

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
	"alex": {
		AuthToken: "123ABC",
		Username:  "Alex",
	},
	"jason": {
		AuthToken: "456DEF",
		Username:  "Jason",
	},
	"marie": {
		AuthToken: "789GHI",
		Username:  "Marie",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"alex": {
		Coins:    100,
		Username: "Alex",
	},
	"jason": {
		Coins:    200,
		Username: "Jason",
	},
	"marie": {
		Coins:    300,
		Username: "Marie",
	},
}
