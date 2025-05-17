package main

import (
	"fmt"
	"os"
)

type Token struct {
	token string
	row   int
	col   int
}

func getBFCode(filename string) string {
	// Read file
	fileBytes, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}

	// Cast as string and return
	codeString := string(fileBytes)

	return codeString
}

func tokenize(code string) []Token {
	return make([]Token, 0)
}

func main() {
	FILE_NAME := "./Examples/helloWorld.bf"
	code := getBFCode(FILE_NAME)
	fmt.Println(code)
}
