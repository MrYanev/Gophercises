package exercise6

import "unicode"

func CamelCase(s string) int {
	var ans int
	for _, char := range s {
		if unicode.IsUpper(char) {
			ans++
		}
	}
	ans++
	return ans
}
