package main


// Buy a stock on ith day and sell on other day
// Return the max profit
// Return 0 if no profit
func stockBuySell(prices []int) int {
	// Set a minimum price variable
	// Set a current price variable
	// If current price is smaller than minimun price , set current price to minimum
	// Else compare current profit with max profit and swap
	if len(prices) == 1 {
		return 0
	}
	minPrice := prices[0]
	maxProfit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else {
			currentProfit := prices[i] - minPrice
			if currentProfit > maxProfit {
				maxProfit = currentProfit
			}
		}
	}
	if minPrice == prices[len(prices)-1] {
		if maxProfit > 0 {
			return maxProfit
		}
		return 0
	}
	return maxProfit
}

// Return the maximum sum from a subarray
func maxSubArray(nums []int) int {
	maxSoFar := nums[0]
	maxEndHere := nums[0]
	for i := 1; i < len(nums); i++ {
		maxEndHere = max(nums[i], nums[i]+maxEndHere)
		maxSoFar = max(maxSoFar, maxEndHere)
	}
	return maxSoFar
}
