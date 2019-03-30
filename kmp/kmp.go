package kmp

//genNext 获取next
func genNext(str string) []int {
	if len(str) <= 2 {
		return []int{-1, 0}
	}
	next := make([]int, len(str))
	next[0] = -1
	next[1] = 0
	for j := 2; j < len(str); j++ {
		tmp := next[j-1]
		for str[j-1] != str[tmp] {
			tmp = next[tmp]
			break
		}
		next[j] = tmp + 1
	}
	return next
}

func StrContain(str1, str2 string) bool {
	next := genNext(str1)
	j := 0
	for i := 0; i < len(str1); i++ {
		if str1[i] == str2[j] {
			i++
			j++
		} else {
			j = next[j]
		}
		if j == -1 {
			i++
			j = 0
		}
		if j == len(str2) {
		}
		return true
	}
	return false
}
