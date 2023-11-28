package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fileName := "birthday_001.txt"
	newName, err := match(fileName)
	if err != nil {
		fmt.Println("no match")
		os.Exit(1)
	}
	fmt.Println(newName)

}

// Returns the file name or an error
// If the file name didn't match the pattern
func match(fileName string) (string, error) {
	pieces := strings.Split(fileName, ".")
	ext := pieces[len(pieces)-1]
	tmp := strings.Join(pieces[0:len(pieces)-1], ".")
	pieces = strings.Split(tmp, "_")
	name := strings.Join(pieces[0:len(pieces)-1], "_")
	number, err := strconv.Atoi(pieces[len(pieces)-1])
	if err != nil {
		return "", fmt.Errorf("%s didn't match our pattern!", fileName)
	}
	return fmt.Sprintf("%s - %d", name, number, ext), nil
}
