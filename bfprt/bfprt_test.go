package bfprt

import (
	"fmt"
	"testing"
)

//bfprt
func TestBfprt(t *testing.T) {
	a := []float64{8, 9, 5, 7, 6, 4, 3, 10, 11, 2, 1}
	c := bfptrFloat64(a, 0, 10, 5)
	fmt.Println(c, a)
}
