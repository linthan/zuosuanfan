package nqueen

import (
	"math"
)

func num1(n int) int {
	if n < 1 {
		return 0
	}
	record := make([]int, n)
	return process1(0, record, n)
}

func process1(i int, record []int, n int) int {
	if i == n {
		return 1
	}
	res := 0
	for j := 0; j < n; j++ {
		if isValid(record, i, j) {
			record[i] = j
			res += process1(i+1, record, n)
		}
	}
	return res
}

func isValid(record []int, i, j int) bool {
	for k := 0; k < i; k++ {
		for j == record[k] || math.Abs(float64(record[k]-j)) == math.Abs(float64(i-k)) {
			return false
		}
	}
	return true
}

func num2(n int) int {
	if n < 1 || n > 32 {
		return 0
	}
	upperLim := 0
	if n == 32 {
		upperLim = -1
	} else {
		upperLim = 1<<uint32(n) - 1
	}
	return process2(upperLim, 0, 0, 0)
}

func process2(upperLim, colLim, leftDiaLim, rightDiaLim int) int {
	if colLim == upperLim {
		return 1
	}
	pos := 0
	mostRightOne := 0
	pos = upperLim & (^(colLim | leftDiaLim | rightDiaLim))
	res := 0
	for pos != 0 {
		mostRightOne = pos & (^pos + 1)
		pos = pos - mostRightOne
		res += process2(upperLim, colLim|mostRightOne, (leftDiaLim|mostRightOne)<<1, (rightDiaLim|mostRightOne)>>1)
	}
	return res
}
