package modules

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"fmt"
)

func AutoRuParser() {
	doc, err := goquery.NewDocument("https://auto.ru/cars/used/sale/1034211938-0d152e/")
	if err != nil {
		log.Fatal(err)
	}
	m := make(map[string]string)
	prop := ""
	doc.Find(".card__info").Children().Each(func(i int, s *goquery.Selection) {
		if s.Is(".card__info-label") {
			prop = s.Text()
			m[prop]=""
		}
		if s.Is(".card__info-value") {
			m[prop] = s.Text()
		}
	})
	fmt.Print(m)
}
