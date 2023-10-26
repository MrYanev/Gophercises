package main

import (
	"fmt"
	"strings"
)

func CamelCase() {
	var ans int
	var input string
	fmt.Scanf("%s\n", &input)
	for _, char := range input {
		str := string(char)
		if strings.ToUpper(str) == str {
			ans++
		}
	}
	ans++
	fmt.Println(ans)
}
