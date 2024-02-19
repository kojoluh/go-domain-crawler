package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	var domain string
	flag.StringVar(&domain, "domain", "", "The target domain to collect js files from <urlLink>")
	docType := flag.String("docType", "js", "The documentType e.g. js | css | a ")
	flag.Parse()

	client := &http.Client{}

	req, err := http.NewRequest("GET", domain, nil)
	if err != nil {
		fmt.Println("Error while retrieving site", err)
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "Golang_Spider_Crawler_Bot/3.0")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error while retrieving site", err)
		log.Fatal(err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Eroor while reading response body", err)
		log.Fatal(err)
	}

	urls := []string{}
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(body)))
	if err != nil {
		fmt.Println("Error while parsing response into goquery")
		log.Fatal(err)
	}

	if *docType == "js" {
		doc.Find("script").Each(func(i int, e *goquery.Selection) {
			src, ok := e.Attr("src")
			if ok {
				urls = append(urls, src)
			}
		})
	}
	if *docType == "css" {
		doc.Find("link").Each(func(i int, e *goquery.Selection) {
			src, ok := e.Attr("href")
			if ok {
				urls = append(urls, src)
			}
		})
	}

	if *docType == "a" {
		doc.Find("a").Each(func(i int, e *goquery.Selection) {
			src, ok := e.Attr("href")
			if ok {
				urls = append(urls, src)
			}
		})
	}

	fmt.Printf("\nFound %v urls matching %v (script/css link/hreftag): \n", len(urls), *docType)
	for _, url := range urls {
		fmt.Println(url)
	}
}
