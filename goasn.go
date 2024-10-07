package main

import (
	"bufio"
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// retry requests and handle errors
func fetchWithRetries(url string, retries int, verbose bool) (*http.Response, error) {
	var resp *http.Response
	var err error

	for i := 0; i < retries; i++ {
		resp, err = http.Get(url)
		if err == nil && resp.StatusCode == 200 {
			return resp, nil
		}
		if verbose {
			fmt.Printf("Retry %d: Failed to fetch %s, error: %v\n", i+1, url, err)
		}
		time.Sleep(2 * time.Second) // wait before retrying
	}

	return nil, fmt.Errorf("after %d retries, failed to fetch %s", retries, url)
}

func searchASN(query string, verbose bool) {
	// regex to detect if the input is a domain
	isDomain := regexp.MustCompile(`\.[a-zA-Z]{2,}$`)

	var searchUrl string
	if isDomain.MatchString(query) {
		// search by domain
		searchUrl = "https://bgp.he.net/dns/" + strings.TrimSpace(query)
	} else {
		// search by organization name
		searchUrl = "https://bgp.he.net/search?search%5Bsearch%5D=" + url.QueryEscape(strings.TrimSpace(query)) + "&commit=Search"
	}

	// page with retries
	resp, err := fetchWithRetries(searchUrl, 3, verbose)
	if err != nil {
		if verbose {
			fmt.Printf("Error fetching the page: %v\n", err)
		}
		return
	}
	defer resp.Body.Close()

	// use goquery to parse the results
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		if verbose {
			fmt.Println("Error reading the page:", err)
		}
		return
	}

	// scrape ASN numbers
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		asn := s.Text()
		if strings.HasPrefix(asn, "AS") {
			if verbose {
				fmt.Printf("ASN for %s: %s\n", query, asn)
			} else {
				fmt.Println(asn)
			}
		}
	})
}

func main() {
	var file string
	var verbose bool
	flag.StringVar(&file, "f", "", "Input file containing domains or organization names")
	flag.BoolVar(&verbose, "v", false, "Enable verbose mode")
	flag.Parse()

	if file != "" {
		f, err := os.Open(file)
		if err != nil {
			fmt.Println("Error opening file:", err)
			return
		}
		defer f.Close()

		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			query := scanner.Text()
			searchASN(query, verbose)
		}

		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading file:", err)
		}
	} else {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			query := scanner.Text()
			searchASN(query, verbose)
		}
	}
}
