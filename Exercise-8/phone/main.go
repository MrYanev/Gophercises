package main

import (
	"bytes"

	_ "github.com/lib/pq"
)

func main() {

}

func normalize(phone string) string {
	var buf bytes.Buffer
	for _, i := range phone {
		if i >= '0' && i <= '9' {
			buf.WriteRune(i)
		}
	}
	return buf.String()
}
