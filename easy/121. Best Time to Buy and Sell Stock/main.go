package main

func maxProfit(prices []int) int {
	curSell := 0
	profit := 0

	for i := len(prices) - 1; i >= 0; i-- {
		if prices[i] > curSell {
			curSell = prices[i]
		} else {
			tmpProfit := curSell - prices[i]
			if tmpProfit > profit {
				profit = tmpProfit
			}
		}
	}

	return profit
}

func main() {

}
