package maxprofit

import (
	"math"
)

func MaxProfit(prices []int) int {
	minDayFound := math.Inf(1)
	maxProfitFound := 0

	for _, price := range prices {
		if float64(price) < minDayFound {
			minDayFound = float64(price)
			continue
		}

		profit := price - int(minDayFound)

		if profit > maxProfitFound {
			maxProfitFound = profit
		}
	}

	return maxProfitFound
}
