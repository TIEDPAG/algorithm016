package t122

func maxProfit(prices []int) int {
	count, l := 0, len(prices)
	for i := 1; i < l; i++ {
		if prices[i] > prices[i-1] {
			count += prices[i] - prices[i-1]
		}
	}
	return count
}