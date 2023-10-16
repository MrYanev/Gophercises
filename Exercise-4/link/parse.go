package link

import (
	"fmt"
	"io"

	"golang.org/x/net/html"
)

// Link represents a link in a HTML doc
type Link struct {
	Href string
	Text string
}

// The function will take in a HTML doc and it
// will return a slice of links parsed from it
func Parse(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	nodes := linkNodes(doc)
	for _, node := range nodes {
		fmt.Println(node)
	}
	//1. Find <a> nodes in doc
	//2. For each link node...
	//	2a Build a link
	//3. Return the links
	return nil, nil
}

func linkNodes(n *html.Node) []*html.Node {
	if n.Type == html.ElementNode && n.Data == "a" {
		return []*html.Node{n}
	}
	var ret []*html.Node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret = append(ret, linkNodes(c)...)
	}
	return ret
}
