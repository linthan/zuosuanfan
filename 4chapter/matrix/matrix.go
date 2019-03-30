package matrix

//minPathSum1 最小路径和
func minPathSum1(m [][]int) int {
	if m == nil || len(m) == 0 || len(m[0]) == 0 {
		return 0
	}
	row := len(m)
	col := len(m[0])
	dp := make([][]int, row)
	for idx := range dp {
		dp[idx] = make([]int, col)
	}
	dp[0][0] = m[0][0]
	for i := 1; i < row; i++ {
		dp[i][0] = dp[i-1][0] + m[i][0]
	}

	for j := 1; j < col; j++ {
		dp[0][j] = dp[0][j-1] + m[0][j]
	}
	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			if dp[i][j-1] > dp[i-1][j] {
				dp[i][j] = dp[i-1][j] + m[i][j]
			} else {
				dp[i][j] = dp[i][j-1] + m[i][j]
			}

		}
	}
	return dp[row-1][col-1]
}
