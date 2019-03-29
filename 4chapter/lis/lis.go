package lis

import "math"

//生成dp数组
func geneDp1(arr []int) []int {
	length := len(arr)
	ret := make([]int, length)
	length = len(arr)
	for idx := 0; idx < length; idx++ {
		ret[idx] = 1
		for j := 0; j < idx; j++ {
			if arr[idx] > arr[j] {
				ret[idx] = int(math.Max(float64(ret[idx]), float64(ret[j]+1)))
			}
		}
	}
	return ret
}

//geneDp2 生成dp2 非常的nice
func geneDp2(arr []int) []int {
	dp := make([]int, len(arr))
	ends := make([]int, len(arr))
	dp[0] = 1
	ends[0] = arr[0]
	right := 0
	l := 0
	r := 0
	m := 0
	for i := 1; i < len(arr); i++ {
		l = 0
		r = right
		for l <= r {
			m = (l + r) / 2
			if ends[m] >= arr[i] {
				r = m - 1
			} else {
				l = m + 1
			}
		}
		if l > right {
			right = l
		}
		ends[l] = arr[i]
		dp[i] = l + 1
	}
	return dp
}

//generateLIS 生成最长子序列
func generateLIS(arr []int, dp []int) []int {
	lenght := 0
	idx := 0
	for i := 0; i < len(dp); i++ {
		if dp[i] > lenght {
			lenght = dp[i]
			idx = i
		}
	}
	lisArr := make([]int, lenght)
	lenght--
	lisArr[lenght] = arr[idx]
	for i := idx; i >= 0; i-- {
		if arr[i] < arr[idx] && dp[i] == dp[idx]-1 {
			lenght--
			lisArr[lenght] = arr[i]
			idx = i
		}
	}
	return lisArr
}
