package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"

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

	links, err := link.Parse(resp.Body)
	if err != nil {
		fmt.Printf("There is an %s error", err)
	}
	for _, l := range links {
		fmt.Println(l)
	}

}
