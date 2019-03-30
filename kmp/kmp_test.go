package kmp

import (
	"fmt"
	"testing"
)

func TestKmp(t *testing.T) {
	a := genNext("abcabc")
	fmt.Println(a)
	fmt.Println("-----------", StrContain("aaasdfasdfabcabctasdfasdfa", "abcabct"))
}
