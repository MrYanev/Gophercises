package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
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
	flag.Parse()

	fmt.Println(*urlFlag)
	resp, err := http.Get(*urlFlag)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)

	reqUrl := resp.Request.URL
	baseUrl := &url.URL{
		Scheme: reqUrl.Scheme,
		Host:   reqUrl.Host,
	}
	base := baseUrl.String()
	pages := hrefs(resp.Body, base)
	for _, page := range pages {
		fmt.Println(page)
	}
}

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
