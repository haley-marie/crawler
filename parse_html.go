package main

import (
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getH1fromHTML(html string) string {
	reader := strings.NewReader(html)
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return ""
	}

	selection := doc.Find("h1").First().Text()

	return strings.TrimSpace(selection)
}

func getFirstParagraphFromHTML(html string) string {
	reader := strings.NewReader(html)
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return ""
	}

	mainSelection := doc.Find("main")
	var p string
	if mainSelection.Length() > 0 {
		p = mainSelection.Find("p").First().Text()
	} else {
		p = doc.Find("p").First().Text()
	}

	return strings.TrimSpace(p)

}

func getURLsFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {
	return []string{}, nil
}

func getImagesFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {
	return []string{}, nil
}
