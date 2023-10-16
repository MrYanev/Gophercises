package link

import "io"

//Link represents a link in a HTML doc
type Link struct {
	Href string
	Text string
}

func Parse(r io.Reader) ([]Link, error) {
	return nil, nil
}
