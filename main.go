package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	res, err := http.Get("https://dictionary.cambridge.org/pronunciation/english/uncle")
	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("div.sbs-section:nth-child(2)").Each(func(i int, s *goquery.Selection) {
		s.Find("li").Each(func(i int, s *goquery.Selection) {
			text := fmt.Sprintf("%s as in %s\n", s.Find(".pron").Text(), s.Find(".word").Text())
			fmt.Println(text)
		})
	})
}
