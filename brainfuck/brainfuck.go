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
	tokenStack := make([]int, 0)
	var lastIndex int

	// Loop through all tokens
	for i, token := range tokens {
		// If the current token is [ append it to the stack
		switch token.token {
		case "[":
			tokenStack = append(tokenStack, i)
		case "]":
			// If there is no bracket index to pop off, return an error
			if len(tokenStack) == 0 {
				return bracketPairs, errors.New("unmatched brackets")
			}
			// Pop token off the end
			tokenStack, lastIndex = tokenStack[:len(tokenStack)-1], tokenStack[len(tokenStack)-1]
			// Add bracket matches to the map
			bracketPairs[i] = lastIndex
			bracketPairs[lastIndex] = i
		}
	}

	// If we're done and we still have open brackets left, return an error
	if len(tokenStack) != 0 {
		return bracketPairs, errors.New("unmatched brackets")
	}

	return bracketPairs, nil
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
	FILE_NAME := "../Examples/helloWorld.bf"
	code := getBFCode(FILE_NAME)
	fmt.Println(code)
}
