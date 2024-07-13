//go:build !solution

package main

import (
	"fmt"
	"os"
	s "strings"
)

func main() {
	filenames := os.Args[1:]
	occurrences := make(map[string]int)

	for _, filename := range filenames {
		content, err := os.ReadFile(filename)

		if err != nil {
			fmt.Printf("Deserialization error: %s\n", err)
			return
		}

		stringContent := s.Split(string(content), "\n")

		for _, str := range stringContent {
			if count, ok := occurrences[str]; ok {
				occurrences[str] = count + 1
			} else {
				occurrences[str] = 1
			}
		}
	}

	for k, v := range occurrences {
		if occurrences[k] > 1 {
			fmt.Printf("%d\t%s\n", v, k)
		}
	}
}
