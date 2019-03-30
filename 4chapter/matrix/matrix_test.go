package matrix

import (
	"fmt"
	"testing"
)

func TestQueue(t *testing.T) {
	data := [][]int{{1, 3, 5, 9}, {8, 1, 3, 4}, {5, 0, 6, 1}, {8, 8, 4, 0}}
	a := minPathSum1(data)
	fmt.Println("------", a)
}
