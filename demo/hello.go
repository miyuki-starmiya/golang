package main

import (
	"fmt"
)

type Human struct {
	Age  int
	Name string
}

func main() {
	h := Human{14, "hitoe"}
	fmt.Println(Human{14, "hitoe"})
	fmt.Println(h.Age)
}
