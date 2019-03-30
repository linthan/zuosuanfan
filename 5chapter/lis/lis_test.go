package lis

import (
	"fmt"
	"testing"
)

func TestLis(t *testing.T) {
	arr := []int{2, 1, 5, 3, 6, 4, 8, 9, 7}
	dp := geneDp2(arr)
	lisArr := generateLIS(arr, dp)
	fmt.Println(lisArr)
}
