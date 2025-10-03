package main

import (
	"net/url"
)

type PageData struct {
	URL            string
	H1             string
	FirstParagraph string
	OutgoingLinks  []string
	ImageURLs      []string
}

func extractPageData(html, pageURL string) PageData {
	dat := PageData{
		URL:            pageURL,
		H1:             getH1fromHTML(html),
		FirstParagraph: getFirstParagraphFromHTML(html),
	}

	u, err := url.Parse(pageURL)
	if err != nil {
		return dat
	}

	if outgoing, err := getURLsFromHTML(html, u); err == nil {
		dat.OutgoingLinks = outgoing
	}

	if images, err := getImagesFromHTML(html, u); err == nil {
		dat.ImageURLs = images
	}

	return dat
}
