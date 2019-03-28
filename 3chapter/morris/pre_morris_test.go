package morris

import (
	"fmt"
	"testing"

	"github.com/linthan/zuosuanfan/3chapter/recurse"
)

/**
		1
	2	   3
4      # #    5
**/

func TestQueue(t *testing.T) {
	a := "1!2!3!4!#!#!5!#!#!#!#!"
	head := recurse.ReconByLevelString(a)
	Post(head)
	fmt.Println()
	fmt.Println("---------------")
}
