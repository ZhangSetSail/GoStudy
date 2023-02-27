package main

import (
	"fmt"
	"strings"
)

func main() {
	a := "abc-def-hagh-jsk"
	fmt.Println(strings.SplitN(a, "-", 0))
	fmt.Println(strings.SplitAfterN(a, "-", 0))
}
