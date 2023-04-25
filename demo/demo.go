package main

import "fmt"

type Character int

func main() {
	const (
		C Character = iota
		CPP
		Python
	)

	fmt.Println(C, CPP, Python)
}


