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
		switch token.token {
		// If the current token is [ append it to the stack
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
	var currentCharacter string

	// Loop through each line of brainfuck code
	for row, line := range codeLines {
		// Loop through each column (index) of each row
		for col := 0; col < len(line); col++ {
			currentCharacter = string(codeLines[row][col])
			// If the current character is a brainfuck character create a token and add it to tokens
			if strings.Contains("><+-.,[]", currentCharacter) {
				tokens = append(tokens, Token{currentCharacter, row + 1, col + 1})
			}
		}
	}

	return tokens
}

func interpret(tokens []Token) {
	// Keep track of the current instruction and the data pointer
	instructionPointer := 0
	dataPointer := 0

	// Get matching brackets and make sure syntax is valid
	matchingBrackets, err := getMatchingBrackets(tokens)
	if err != nil {
		println("Invalid syntax: %v", err)
		return
	}

	// Loop until we get past the last instruction (end of the program)
	for instructionPointer < len(tokens) {
		switch tokens[instructionPointer].token {
		case ">":
			// Increment the data pointer by one (to point to the next cell to the right)
		case "<":
			// Decrement the data pointer by one (to point to the next cell to the left)
		case "+":
			// Increment the byte at the data pointer by one
		case "-":
			// Decrement the byte at the data pointer by one
		case ".":
			// Output the byte at the data pointer
		case ",":
			// Accept one byte of input, storing its value in the byte at the data pointer
		case "[":
			// If the byte at the data pointer is 0, then jump the instruction pointer forward to the command after the matching ]
		case "]":
			// If the byte at the data pointer is not 0, then jump the instruction pointer back to the command after the matching [
		default:
			// Otherwise there was an unexpected token
			fmt.Println("ERROR: Unexpected token!")
			return
		}
	}

	// REMOVE AFTER finishing
	dataPointer += 0
	matchingBrackets[-1] = 0
}

func main() {
	FILE_NAME := "../Examples/helloWorld.bf"
	code := getBFCode(FILE_NAME)
	fmt.Println(code)
}
