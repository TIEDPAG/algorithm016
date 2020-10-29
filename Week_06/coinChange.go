func coinChange(coins []int, amount int) int {
    max := amount + 1
    l := len(coins)
    dp := make([]int, max)
    for i := 1; i < max; i++ {
        dp[i] = max
    }
    for i := 1; i < max; i++ {
        for j := 0; j < l; j++ {
            if coins[j] <= i {
                dp[i] = min(dp[i], dp[i - coins[j]] + 1)
            }
        }
    }
    if dp[amount] > amount {
        return -1
    }
    return dp[amount]
}

func min(a, b int) int {
    if a > b {
        return b
    }
    return a
}