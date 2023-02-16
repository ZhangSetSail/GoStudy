package main

import "fmt"

func main() {
	var test []*string
	for i, v := range test {
		fmt.Println(i, v)
	}
}
