package max_profit_121

// MaxProfit
// @Description 买卖股票的最佳时机,只允许完成一笔交易,时间复杂度O(n),空间复杂度O(1)
// @Author Oberl-Fitzgerald 2024-07-11 16:39:19
// @Param  prices []int
// @Return int
func MaxProfit(prices []int) int {
	//  如果 prices 为空，则返回 0
	if len(prices) == 0 {
		return 0
	}

	minPrice := prices[0]
	maxProfit := 0

	//  遍历 prices，找到最小的价格和最大的利润
	for i := 1; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else if prices[i]-minPrice > maxProfit {
			maxProfit = prices[i] - minPrice
		}
	}

	return maxProfit
}
