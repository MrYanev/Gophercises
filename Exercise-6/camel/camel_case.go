package main

import (
	"fmt"
	"unicode"
)

func CamelCase() {
	var ans int
	var input string
	fmt.Scanf("%s\n", &input)
	for _, char := range input {
		if unicode.IsUpper(char) {
			ans++
		}
	}
	ans++
	fmt.Println(ans)
}
