package han

import (
	"github.com/linthan/zuosuanfan/1chapter/minstatck"

	"fmt"
)

//Han 汉诺塔递归
func Han(n int, x, y, z string) {
	if n == 0 {
		return
	}
	Han(n-1, x, z, y)
	fmt.Printf("%s=>%s \n", x, z)
	Han(n-1, y, x, z)
}

//Han1 汉诺塔递归
func Han1(n int, x, y, z *minstatck.Stack) {
	if n == 0 {
		return
	}
	Han1(n-1, x, z, y)
	fmt.Println("-------------")
	Han1(n-1, y, x, z)
}
