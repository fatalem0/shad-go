//go:build !solution

package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

var (
	DeserializationError = errors.New("can't deserialize response's body")
	UrlValidationError   = errors.New("invalid URL")
)

func readResponseBody(response *http.Response) ([]byte, error) {
	body, err := io.ReadAll(response.Body)
	defer response.Body.Close()

	if err != nil {
		return []byte{}, DeserializationError
	}

	return body, nil
}

func validateUrl(urlForValidation string) (err error) {
	if _, err := url.ParseRequestURI(urlForValidation); err != nil {
		return UrlValidationError
	}

	return
}

func main() {
	urls := os.Args[1:]

	urlContent := make([]string, 0, len(urls))

	for _, url := range urls {
		if err := validateUrl(url); err != nil {
			log.Fatal(err)
		}

		res, err := http.Get(url)

		if err != nil {
			log.Fatal(err)
		}

		body, err := readResponseBody(res)

		if err != nil {
			log.Fatal(err)
		}

		urlContent = append(urlContent, string(body))
	}

	for _, v := range urlContent {
		fmt.Println(v)
	}
}
