package kmp

import (
	"fmt"
	"testing"
)

func TestKmp(t *testing.T) {
	a := genNext("abcabct")
	fmt.Println(a)
	fmt.Println("-----------", StrContain("aaasdfasdfabcabctasdfasdfa", "abcabct"))
}
