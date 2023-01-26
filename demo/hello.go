package main

import (
	"fmt"
)

type Human struct {
	Age  int
	Name string
}

func main() {
	iterable := []int{1, 2, 3}
	for _, e := range iterable {
		fmt.Println(e)
	}
}
