package main

func findLongestSameString(a, b string) string {
	m := 0
	maxLen := 0
	dp := make([][]int, len(a)+1)
	for i, _ := range dp {
		dp[i] = make([]int, len(b)+1)
	}
	for i, x := range a {
		for j, y := range b {
			if x == y {
				dp[i+1][j+1] = dp[i][j] + 1
				if dp[i+1][j+1] > maxLen {
					m = i
					maxLen = dp[i+1][j+1]
				}
			} else {
				dp[i+1][j+1] = 0
			}
		}
	}
	return a[m-maxLen+1 : m+1]
}
