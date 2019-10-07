package util

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func Parse() ([]string, error) {
	doc, _ := goquery.NewDocument("http://favstar.fm/users/mattn_jp/")
	doc.Find(".fs-tweet").Each(func(_ int, s *goquery.Selection) {
		fmt.Println(s.Find(".fs-tweet-text").Text())

		s.Find("div[data-type='favs']").Each(func(_ int, s *goquery.Selection) {
			fmt.Printf("FAV(%s): %s ...\n",
				s.Find("li.fs-total").Text(),
				strings.Join(s.Find("a").Map(func(_ int, s *goquery.Selection) string {
					a, _ := s.Attr("title")
					return a
				}), ", "))
		})
		s.Find("div[data-type='retweets']").Each(func(_ int, s *goquery.Selection) {
			fmt.Printf("RT (%s): %s ...\n",
				s.Find("li.fs-total").Text(),
				strings.Join(s.Find("a").Map(func(_ int, s *goquery.Selection) string {
					a, _ := s.Attr("title")
					return a
				}), ", "))
		})

		fmt.Println()
	})
	return nil, nil
}
