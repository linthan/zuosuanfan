package kmp

//genNext 获取next
func genNext(str string) []int {
	if len(str) <= 2 {
		return []int{-1, 0}
	}
	next := make([]int, len(str)+1)
	next[0] = -1
	next[1] = 0
	for j := 2; j < len(str)+1; j++ {
		tmp := next[j-1]
		for str[j-1] != str[tmp] {
			tmp = next[tmp]
			break
		}
		next[j] = tmp + 1
	}
	return next
}

//StrContain  获取字符串
func StrContain(str1, str2 string) bool {
	next := genNext(str1)
	j := 0
	i := 0
	for i < len(str1) && j < len(str2) {
		if str1[i] == str2[j] {
			i++
			j++
		} else if next[j] == -1 {
			i++
		} else {
			j = next[j]
		}
		if j == len(str2) {
			return true
		}
	}
	return false
}

//StrContain1  获取字符串
func StrContain1(str1, str2 string) int {
	next := genNext(str1)
	j := 0
	i := 0
	for i < len(str1) && j < len(str2) {
		if str1[i] == str2[j] {
			i++
			j++

		} else if next[j] == -1 {
			i++
		} else {
			j = next[j]
		}
	}
	if j == len(str2) {
		return i - j
	}
	return -1
}
