package lcst

//getdp 获取动态规划
func getdp(arr1, arr2 string) [][]int {
	ret := make([][]int, len(arr1))
	for idx := range ret {
		ret[idx] = make([]int, len(arr2))
	}
	for i := 0; i < len(arr1); i++ {
		if arr1[i] == arr2[0] {
			ret[i][0] = 1
		}
	}
	for j := 1; j < len(arr2); j++ {
		if arr2[j] == arr1[0] {
			ret[0][j] = 1
		}
	}
	for i := 1; i < len(arr1); i++ {
		for j := 1; j < len(arr2); j++ {
			if arr1[i] == arr2[j] {
				ret[i][j] = ret[i-1][j-1] + 1
			}
		}
	}
	return ret
}

func lcst1(arr1, arr2 string) string {
	dp := getdp(arr1, arr2)
	end := 0
	max := 0
	for i := 0; i < len(arr1); i++ {
		for j := 0; j < len(arr2); j++ {
			if dp[i][j] > max {
				end = i
				max = dp[i][j]
			}
		}
	}
	return arr1[end-max+1 : end+1]
}

//lcst2 获取最长公共子串
func lcst2(arr1, arr2 string) string {
	row := 0
	col := len(arr2) - 1
	max := 0
	end := 0
	for row < len(arr1) {
		i := row
		j := col
		leng := 0
		for i < len(arr1) && j < len(arr2) {
			if arr1[i] == arr2[j] {
				leng++
			} else {
				leng = 0
			}

			if leng > max {
				max = leng
				end = i
			}
			i++
			j++
		}
		if col > 0 {
			col--
		} else {
			row++
		}
	}
	return arr1[end-max+1 : end+1]
}
