package main

import (
	"fmt"
	"os"
)

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

func main() {
	FILE_NAME := "./Examples/helloWorld.bf"
	code := getBFCode(FILE_NAME)
	fmt.Println(code)
}
