package recurse

import (
	"fmt"
	"testing"
)

func TestQueue(t *testing.T) {
	a := "1!2!3!4!#!#!5!#!#!#!#!"
	head := reconByLevelString(a)
	b := serialByLevel(head)
	fmt.Println("---------", b)
}
