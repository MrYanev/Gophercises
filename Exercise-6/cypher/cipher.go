package main

import "fmt"

func main() {
	var shift int
	var input, ans string
	fmt.Scanf("%s\n", &input)
	fmt.Scanf("%v\n", &shift)
	for _, char := range input {
		shifted := int(char) + shift
		ans += fmt.Sprint(shifted)
	}
	fmt.Println(ans)
}
