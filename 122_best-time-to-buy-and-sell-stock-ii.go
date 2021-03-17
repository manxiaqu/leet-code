package leetcode

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	// 第一个买入点为，找到第一个a[i] < a[i+1]第值
	// 第一个卖出的点

	// a[i] > 买入点，a[i] >= a[i+1]买
	output := 0
	for buy, sell := 0, 0; buy < len(prices)-1 && sell < len(prices); {
		if prices[buy] >= prices[buy+1] {
			buy++
			continue
		}
		sell = buy + 1

		for sell < len(prices)-1 && prices[sell] < prices[sell+1] {
			sell++
		}
		output = output + prices[sell] - prices[buy]

		buy = sell + 1
	}

	return output
}

func maxProfit2(prices []int) int {
	output := 0

	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			output = output + prices[i] - prices[i-1]
		}
	}

	return output
}
