package mincost

//minCost1 最小代价
func minCost1(str1, str2 string, ic, dc, rc int) int {
	row := len(str1) + 1
	col := len(str2) + 1
	dp := make([][]int, row)
	for idx := range dp {
		dp[idx] = make([]int, col)
	}
	for i := 1; i < row; i++ {
		dp[i][0] = dc * i
	}
	for j := 1; j < col; j++ {
		dp[0][j] = ic * j
	}
	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			if str1[i-1] == str2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = dp[i-1][j-1] + rc
			}
			if dp[i][j-1]+ic < dp[i][j] {
				dp[i][j] = dp[i][j-1] + ic
			}
			if dp[i-1][j]+dc < dp[i][j] {
				dp[i][j] = dp[i-1][j] + dc
			}

		}
	}
	return dp[row-1][col-1]
}
