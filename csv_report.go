package main

import (
	"encoding/csv"
	"log"
	"os"
	"sort"
	"strings"
)

func writeCSVReport(pages map[string]PageData, filename string) error {
	report, err := os.Create(filename)
	if err != nil {
		log.Println("Error creating csv: ", err)
		return err
	}

	w := csv.NewWriter(report)

	err = w.Write([]string{
		"page_url",
		"h1",
		"first_paragraph",
		"outgoing_link_urls",
		"image_urls",
	})
	if err != nil {
		log.Println("Error writing csv header: ", err)
		return err
	}

	keys := make([]string, 0, len(pages))
	for k := range pages {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, normalizedURL := range keys {
		page := pages[normalizedURL]
		err = w.Write([]string{
			page.URL,
			page.H1,
			page.FirstParagraph,
			strings.Join(page.OutgoingLinks, ";"),
			strings.Join(page.ImageURLs, ";"),
		})
		if err != nil {
			log.Printf(
				"Error writing page data for %s: %v\n",
				page.URL,
				err,
			)
			return err
		}
	}

	return nil
}
