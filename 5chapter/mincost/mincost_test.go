package mincost

import (
	"fmt"
	"testing"
)

func TestCost(t *testing.T) {
	a := minCost1("abc", "adc", 5, 3, 2)
	fmt.Println(a)
}
