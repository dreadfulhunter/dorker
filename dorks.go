package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"regexp"
)

const (
	googleSearchURL = "https://www.google.com/search?q="
	bingSearchURL   = "https://www.bing.com/search?q="
	duckDuckGoURL   = "https://duckduckgo.com/?q="
	yandexSearchURL = "https://www.yandex.com/search/?text="
	yahooSearchURL  = "https://search.yahoo.com/search?p="
	baiduSearchURL  = "https://www.baidu.com/s?wd="
)

var searchEngines = map[string]string{
	"google": googleSearchURL,
	"bing":   bingSearchURL,
	"duck":   duckDuckGoURL,
	"yandex": yandexSearchURL,
	"yahoo":  yahooSearchURL,
	"baidu":  baiduSearchURL,
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage:", os.Args[0], "SEARCH_ENGINE QUERY")
		fmt.Println("Available search engines: google, bing, duck, yandex, yahoo, baidu")
		return
	}
	searchEngine := os.Args[1]
	query := os.Args[2]

	// Get the search URL for the selected engine
	searchURL, ok := searchEngines[searchEngine]
	if !ok {
		fmt.Println("Error: Invalid search engine.")
		return
	}

	// Build the search URL
	url := searchURL + url.QueryEscape(query)

	// Create a regular expression to match subdomains
	re := regexp.MustCompile(`(?i)https?://([a-z0-9-]+)\.[a-z-]+\.[a-z-]+`)

	// Start a new goroutine for each request
	go func(url string) {
		res, err := http.Get(url)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println("Response from", url, ":")
		fmt.Println("Matched subdomains:")
		matches := re.FindAllStringSubmatch(string(body), -1)
		for _, match := range matches {
			fmt.Println(match[1])
		}
	}(url)
}
