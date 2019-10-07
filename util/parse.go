package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func Parse() ([]string, error) {
	urls := []string{}

	doc, err := goquery.NewDocument("https://github.com/trending")
	if err != nil {
		fmt.Print("error: url scrapping failed")
	}

	doc.Find(".Box-row .lh-condensed").Each(func(_ int, s *goquery.Selection) {
		s.Find("a").Each(func(_ int, a *goquery.Selection) {
			url, _ := a.Attr("href")
			urls = append(urls, url)
		})
	})
	return urls, nil
}

func main() {
	urls, _ := Parse()
	fmt.Println(urls)
}
