package main

import (
	"net/http"
)

func main() {
	url := "http://www.example.com"
	// HTMLリソースを取得
	resp, err := http.Get(url)

}
