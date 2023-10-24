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
	*/
}
