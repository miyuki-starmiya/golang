package main

import "fmt"

func main() {
	// map
	m := make(map[string]int)
	m["key"] = 0
	fmt.Println(m)

	for k, v := range m {
		fmt.Println(k, v)
	}
}
