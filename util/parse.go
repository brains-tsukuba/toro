package util

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

// Parse returns (`list of repo's path`, error), base on url
func Parse(url string) ([]string, error) {
	urls := []string{}

	doc, err := goquery.NewDocument(url)
	if err != nil {
		fmt.Print("error: url scrapping failed")
	}

	doc.Find(".Box-row .lh-condensed").Each(func(_ int, s *goquery.Selection) {
		s.Find("a").Each(func(_ int, a *goquery.Selection) {
			url, _ := a.Attr("href")
			urls = append(urls, url[1:])
		})
	})
	return urls, nil
}
