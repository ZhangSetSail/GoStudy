package main

import (
	"fmt"
	"github.com/ZhangSetSail/GoStudy/gin-demo/server"
	"os"
)

//入口函数
func main() {
	if err := server.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
