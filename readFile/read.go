package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	content, err := os.ReadFile("文件.txt")
	if err != nil {
		log.Fatal(err)
	}
	list := strings.Split(string(content), "\n")
	fmt.Println(list[1])
}
