package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	// http://go-colly.org/docs/introduction/install/
	c := colly.NewCollector()

	c.OnHTML("title", func(e *colly.HTMLElement) {
		fmt.Println(e.Text)
	})

	url := "https://example.com/"
	c.Visit(url)
}

