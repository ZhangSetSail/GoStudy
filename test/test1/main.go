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
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			logrus.Errorf("Failed to request: %v\n", err)
			return
		}
		req.Header.Set("X-Request-Id", r.Header["X-Request-Id"][0])
		req.Header.Set("name", "zqh")
		resp, err := (&http.Client{}).Do(req)
		//resp, err := http.Get(serviceUrl + "/topic/query/false/lsj")
		if err != nil {
			logrus.Errorf("query topic failed: %v", err)
			return
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			logrus.Errorf("Failed to read response body: %v\n", err)
			return
		}
		fmt.Fprintln(w, "第一个：", string(body), "第二个：", r.Header["X-Request-Id"])
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
