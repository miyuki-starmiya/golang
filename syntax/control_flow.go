package main

import "fmt"

func main() {
	for i := 0; i < 3; i++ {
		if i == 1 {
			continue
		}
		fmt.Println(i)
	}
}
