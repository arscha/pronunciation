package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/proxy"
)

func main() {
	socksProxy, err := proxy.SOCKS5("tcp", "localhost:9150", nil, proxy.Direct)
	if err != nil {
		log.Fatalln(err)
	}
	transport := &http.Transport{Dial: socksProxy.Dial}

	client := &http.Client{
		Transport: transport,
	}

	res, err := client.Get("https://dictionary.cambridge.org/pronunciation/english/island")
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
		fmt.Println(s.Find("div.sbs-section:nth-child(2) > div:nth-child(1) > span:nth-child(1) > span:nth-child(2)").Text())
		s.Find("li").Each(func(i int, s *goquery.Selection) {
			text := fmt.Sprintf("%s as in %s\n", s.Find(".pron").Text(), s.Find(".word").Text())
			fmt.Print(text)
		})
	})
}
