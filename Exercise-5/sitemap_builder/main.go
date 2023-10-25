package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/MrYanev/Gophercises/Exercise-4/link"
)

func main() {
	/*
		1. Get the webpage
		2. Parse all the links on the page
		3. Build proper urls with the links
		4.Filter out any links w/ a diff domain
		5. Find all pages
		6. Print out the XML
	*/
	urlFlag := flag.String("url", "https://gophercises.com", "the URL to buld the site map for")
	maxDepth := flag.Int("deoth", 5, "The maximum number of links depth to traverse")
	flag.Parse()

	pages := bfs(*urlFlag, *maxDepth)
	for _, page := range pages {
		fmt.Println(page)
	}

}

//	struct{}{} -> means type empty struct{}

func bfs(urlStr string, maxDepth int) []string {
	seen := make(map[string]struct{})
	var q map[string]struct{}
	nq := map[string]struct{}{
		urlStr: struct{}{},
	}
	for i := 0; i <= maxDepth; i++ {
		q, nq = nq, make(map[string]struct{})
		for url, _ := range q {
			if _, ok := seen[url]; ok {
				continue
			}
			seen[url] = struct{}{}
			for _, link := range get(url) {
				nq[link] = struct{}{}
			}
		}
	}
	ret := make([]string, 0, len(seen))
	for url, _ := range seen {
		ret = append(ret, url)
	}
	return ret
}

// A function to get the URLs
func get(UurlStr string) []string {
	resp, err := http.Get(UurlStr)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	reqUrl := resp.Request.URL
	baseUrl := &url.URL{
		Scheme: reqUrl.Scheme,
		Host:   reqUrl.Host,
	}
	base := baseUrl.String()
	return filter(hrefs(resp.Body, base), withPrefix(base))
}

// A function to parse the links with the extentions
// Or use the full http. request
func hrefs(r io.Reader, base string) []string {
	links, err := link.Parse(r)
	if err != nil {
		fmt.Printf("There is an %s error", err)
	}

	var ret []string
	for _, l := range links {
		switch {
		case strings.HasPrefix(l.Href, "/"):
			ret = append(ret, base+l.Href)
		case strings.HasPrefix(l.Href, "http"):
			ret = append(ret, l.Href)
		}
	}
	return ret
}

// A function to filter out repetative links
func filter(links []string, keepFn func(string) bool) []string {
	var ret []string
	for _, link := range links {
		if keepFn(link) {
			ret = append(ret, link)
		}
	}
	return ret
}

// A func to keep the links with prefix
func withPrefix(pfx string) func(string) bool {
	return func(link string) bool {
		return strings.HasPrefix(link, pfx)
	}
}
