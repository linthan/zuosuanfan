package lcse

import "strings"

//getDp 获取动态规划
func getDp(arr1, arr2 string) [][]int {
	dp := make([][]int, len(arr1))
	for idx := range dp {
		dp[idx] = make([]int, len(arr2))
	}
	for i := 0; i < len(arr1); i++ {
		if arr1[i] == arr2[0] {
			dp[i][0] = 1
		}
	}
	for j := 0; j < len(arr2); j++ {
		if arr2[j] == arr1[0] {
			dp[0][j] = 1
		}
	}
	for i := 1; i < len(arr1); i++ {
		for j := 1; j < len(arr2); j++ {
			if dp[i][j-1] > dp[i-1][j] {
				dp[i][j] = dp[i][j-1]
			} else {
				dp[i][j] = dp[i-1][j]
			}
			if arr1[i] == arr2[j] {
				if dp[i-1][j-1]+1 > dp[i][j] {
					dp[i][j] = dp[i-1][j-1] + 1
				}
			}

		}
	}
	return dp
}

func lcse(arr1, arr2 string) string {
	m := len(arr1) - 1
	n := len(arr2) - 1
	dp := getDp(arr1, arr2)
	ret := make([]string, dp[m][n])
	idx := len(ret) - 1
	for idx >= 0 {
		if n > 0 && dp[m][n] == dp[m][n-1] {
			n--
		} else if m > 0 && dp[m][n] == dp[m-1][n] {
			m--
		} else {
			ret[idx] = string(arr1[m])
			idx--
			m--
			n--
		}
	}
	return strings.Join(ret, "")
}
