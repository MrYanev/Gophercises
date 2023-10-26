package main

import (
	"fmt"
	"strings"
)

func main() {
	var shift int
	var input, ans string
	fmt.Scanf("%s\n", &input)
	fmt.Scanf("%v\n", &shift)
	// for _, char := range input {
	// 	shifted := int(char) + shift
	// 	ans += fmt.Sprint(shifted)
	// }
	alphabetStr := "abcdefghijklmnopqrstuvwxyz"
	alphabet := []rune(alphabetStr)
	for _, ch := range input {
		if strings.IndexRune(alphabetStr, ch) >= 0 {
			ans += string(rotate(ch, shift, alphabet))
		} else {
			ans += string(ch)
		}
	}
	fmt.Println(ans)

}

func rotate(s rune, delta int, key []rune) rune {
	idx := strings.IndexRune(string(key), s)
	if idx < 0 {
		panic("idx < 0")
	}
	idx = (idx + delta) % len(key)
	return key[idx]
}
