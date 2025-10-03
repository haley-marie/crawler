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
	reader := strings.NewReader(htmlBody)
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return []string{}, err
	}

	var absURLs []string
	doc.Find("a[href]").Each(func(_ int, s *goquery.Selection) {
		raw, exists := s.Attr("href")

		if !exists || raw == "" {
			return
		}

		u, err := url.Parse(raw)
		if err != nil {
			return
		}
		abs := baseURL.ResolveReference(u)
		absURLs = append(absURLs, abs.String())
	})

	return absURLs, nil
}

func getImagesFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {
	reader := strings.NewReader(htmlBody)
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return []string{}, err
	}

	var absURLs []string
	doc.Find("img[src]").Each(func(_ int, s *goquery.Selection) {
		raw, exists := s.Attr("src")

		if !exists || raw == "" {
			return
		}

		u, err := url.Parse(raw)
		if err != nil {
			return
		}
		abs := baseURL.ResolveReference(u)
		absURLs = append(absURLs, abs.String())
	})

	return absURLs, nil
}
