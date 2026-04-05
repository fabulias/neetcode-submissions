func maxProfit(prices []int) int {
	if len(prices) == 1 {
		return 0
	}

	buyIndex, sellIndex := 0, 1
	maxProfit := 0

	for sellIndex < len(prices) {
		if prices[buyIndex] < prices[sellIndex] {
			profit := prices[sellIndex] - prices[buyIndex]
			maxProfit = max(maxProfit, profit)
		} else {
			buyIndex = sellIndex
		}
		sellIndex++
	}

	return maxProfit
}
