package lcse

import (
	"fmt"
	"testing"
)

func TestLCST(t *testing.T) {
	arr1 := "1A2C3D4B56"
	arr2 := "B1D23CA45B6A"
	a := lcse(arr1, arr2)
	fmt.Println(a)
}
