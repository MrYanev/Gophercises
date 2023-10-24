package main

import (
	"flag"
	"fmt"
)

func main() {
	url := flag.String("url", "https://gophercises.com", "the URL to buld the site map for")
	flag.Parse()

	fmt.Println(url)
}
