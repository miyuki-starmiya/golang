package main

import (
	"fmt"
	"regexp"
)

func main() {
	result := []string{}
	input := "abcde"
	pattern := regexp.MustCompile(`cd`)

	matched := pattern.MatchString(input)
	if matched {
		result = append(result, input)
	}
	fmt.Println(result)
}
