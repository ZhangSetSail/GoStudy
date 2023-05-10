package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	url := os.Getenv("URL")
	if url != "" {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("Failed to request: %v\n", err)
			return
		}
		logrus.Infof("header:%v", resp.Header)
	}
	fmt.Fprintln(w, r)
}

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Printf("http server failed, err:%v\n", err)
		return
	}
}
