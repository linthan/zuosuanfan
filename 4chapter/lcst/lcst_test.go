package lcst

import (
	"fmt"
	"testing"
)

func TestLCST(t *testing.T) {
	arr1 := "1AB2345CD"
	arr2 := "12345EF"
	a := lcst1(arr1, arr2)
	fmt.Println(a)
}
