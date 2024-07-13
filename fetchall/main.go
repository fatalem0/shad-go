//go:build !solution

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type Response struct {
	TimeSpent     float64
	ContentLength int
	URL           string
	Error         error
}

func main() {
	startTimeMain := time.Now()
	urls := os.Args[1:]
	responsesCh := make(chan Response)

	for _, url := range urls {
		go func(url string) {
			startTime := time.Now()
			resp, err := http.Get(url)
			timeSpent := time.Since(startTime)

			if err != nil {
				responsesCh <- Response{Error: err}
				return
			}

			body, err := io.ReadAll(resp.Body)
			defer resp.Body.Close()

			if err != nil {
				responsesCh <- Response{Error: err}
			} else {
				responsesCh <- Response{timeSpent.Seconds(), len(body), url, nil}
			}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		res := <-responsesCh

		if res.Error != nil {
			fmt.Println(res.Error)
		} else {
			fmt.Printf("%.2fs    %d  %s\n", res.TimeSpent, res.ContentLength, res.URL)
		}
	}

	endTimeMain := time.Since(startTimeMain)
	fmt.Printf("%.2fs elapsed", endTimeMain.Seconds())
}
