package main

import (
	"flag"
	"fmt"
)

func main() {
	url := flag.String("url", "https://gophercises.com", "the URL to buld the site map for")
	flag.Parse()

	fmt.Println(url)

	/*
		1. Get the webpage
		2. Parse all the links on the page
		3. Build proper urls with the links
		4.Filter out any links w/ a diff domain
		5. Find all pages
		6. Print out the XML
	*/
}
