package han

import "fmt"

//Han 汉诺塔递归
func Han(n int, x, y, z string) {
	if n == 0 {
		return
	}
	Han(n-1, x, z, y)
	fmt.Printf("%s=>%s \n", x, z)
	Han(n-1, y, x, z)
}
