package main

import (
	"fmt"
	"os"
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

func match(fileName string) (string, error) {
	return "", nil
}
