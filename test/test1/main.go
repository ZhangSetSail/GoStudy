package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"os"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	url := os.Getenv("URL")
	if url != "" {
		resp, err := http.Get(url)
		if err != nil {
			logrus.Errorf("Failed to request: %v\n", err)
			return
		}

		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			logrus.Errorf("Failed to read response body: %v\n", err)
			return
		}

		fmt.Fprintln(w, "第一个：", string(body), "第二个：", r)
		return
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
