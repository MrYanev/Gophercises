package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/MrYanev/Gophercises/cyoa"
)

func main() {
	port := flag.Int("port", 3000, "The port to strat the CYOA application on")
	filename := flag.String("file", "gopher.json", "the JSON file with the CYOA")
	flag.Parse()
	fmt.Printf("Using the story in %s.\n", *filename)

	f, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}

	story, err := cyoa.JsonStory(f)
	if err != nil {
		panic(err)
	}
	h := cyoa.NewHandler(story)
	fmt.Printf("Starting the server on port: %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))
}