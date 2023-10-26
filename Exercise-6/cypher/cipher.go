package main

import "fmt"

func main() {
	var shift int
	var input, ans string
	fmt.Scan("%s\n", &input)
	fmt.Scan("%v\n", &shift)
	for _, char := range input {
		shifted := int(char) + shift
		ans += fmt.Sprint(shifted)
	}
	fmt.Println(ans)
}
