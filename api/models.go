package api

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
	"github.com/arschaaa/pronunciation/util"
)

type PronResponse struct {
	Word          string   `json:"word"`
	Pronunciation string   `json:"ipa"`
	LettersPron   []string `json:"sbspronun"`
}

func (p *PronResponse) getPron(url string) error {
	client, err := util.SetProxy("localhost:9150")
	if err != nil {
		return err
	}

	res, err := client.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return err
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return err
	}

	doc.Find("div.sbs-section:nth-child(2)").Each(func(i int, s *goquery.Selection) {
		pron := s.Find("div.sbs-section:nth-child(2) > div:nth-child(1) > span:nth-child(1) > span:nth-child(2)").Text()
		p.Pronunciation = pron
		s.Find("li").Each(func(i int, s *goquery.Selection) {
			pron := fmt.Sprintf("%s as in %s\n", s.Find(".pron").Text(), s.Find(".word").Text())
			p.LettersPron = append(p.LettersPron, pron)
		})
	})

	return nil
}
