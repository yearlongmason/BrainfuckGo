package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

type Token struct {
	token string
	row   int
	col   int
}

func getMatchingBrackets(tokens []Token) (map[int]int, error) {
	bracketPairs := make(map[int]int)
	testError := errors.New("Unmatched token!")

	return bracketPairs, testError
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
	codeLines := strings.Split(code, "\n")
	tokens := make([]Token, 0)

	// Loop through each line of brainfuck code
	for row, line := range codeLines {
		// Loop through each column (index) of each row
		for col := 0; col < len(line); col++ {
			// If the current character is a brainfuck character create a token and add it to tokens
			if strings.Contains("><+-.,[]", string(codeLines[row][col])) {
				tokens = append(tokens, Token{string(codeLines[row][col]), row + 1, col + 1})
			}
		}
	}

	return tokens
}

func main() {
	FILE_NAME := "./Examples/helloWorld.bf"
	code := getBFCode(FILE_NAME)
	fmt.Println(code)
}
