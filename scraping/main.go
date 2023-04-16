package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

// type Title struct {
// 	ID int
// 	Title string
// }

// type Title string

func main() {
	allTitles := make([]string, 0)

	collector := colly.NewCollector(
		colly.AllowedDomains("anitrendz.com", "www.anitrendz.com"),
	)

	collector.OnHTML("body", func(e *colly.HTMLElement) {
		title := e.Text

		allTitles = append(allTitles, title)
	})

	collector.OnRequest(func(req *colly.Request) {
		fmt.Println("Visiting", req.URL.String())
	})

	url := "https://anitrendz.com/charts/top-anime"
	collector.Visit(url)

	fmt.Println(allTitles)
}
