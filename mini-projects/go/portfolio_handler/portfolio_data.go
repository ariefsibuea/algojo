package main

var PortfolioData = []PortfolioResponse{
	{
		UserID: "1",
		Positions: []Position{
			{Coin: "BTC"},
			{Coin: "ETH"},
		},
	},
	{
		UserID: "2",
		Positions: []Position{
			{Coin: "BTC"},
			{Coin: "ETH"},
			{Coin: "BNB"},
		},
	},
	{
		UserID: "3",
		Positions: []Position{
			{Coin: "BTC"},
			{Coin: "BNB"},
		},
	},
	{
		UserID: "4",
		Positions: []Position{
			{Coin: "BTC"},
			{Coin: "USDT"},
			{Coin: "USDT"},
			{Coin: "USDC"},
		},
	},
}
