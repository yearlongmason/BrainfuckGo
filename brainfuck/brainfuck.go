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

func incByte(num int) int {
	if num >= 255 {
		return 0
	}
	return num + 1
}

func decByte(num int) int {
	if num <= 0 {
		return 255
	}
	return num - 1
}

func getInput() int {
	var input int
	fmt.Scan(&input)

	// Validate input
	if input < 0 {
		return 0
	} else if input > 255 {
		return 255
	}
	return input
}

func interpret(tokens []Token) {
	// Keep track of the current instruction and the data pointer
	instructionPointer := 0
	dataPointer := 0
	data := make(map[int]int)

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
			dataPointer += 1
		case "<":
			// Decrement the data pointer by one (to point to the next cell to the left)
			dataPointer -= 1
		case "+":
			// Increment the byte at the data pointer by one
			data[dataPointer] = incByte(data[dataPointer])
		case "-":
			// Decrement the byte at the data pointer by one
			data[dataPointer] = decByte(data[dataPointer])
		case ".":
			// Output the byte at the data pointer
			fmt.Printf("%c", data[dataPointer])
		case ",":
			// Accept one byte of input, storing its value in the byte at the data pointer
			data[dataPointer] = getInput()
		case "[":
			// If the byte at the data pointer is 0, then jump the instruction pointer forward to the command after the matching ]
			if data[dataPointer] == 0 {
				instructionPointer = matchingBrackets[instructionPointer]
			}
		case "]":
			// If the byte at the data pointer is not 0, then jump the instruction pointer back to the command after the matching [
			if data[dataPointer] != 0 {
				instructionPointer = matchingBrackets[instructionPointer]
			}
		default:
			// Otherwise there was an unexpected token
			fmt.Println("ERROR: Unexpected token!")
			return
		}

		// Incriment instruction pointer to move to the next token
		instructionPointer += 1
	}

	// REMOVE AFTER finishing
	dataPointer += 0
	matchingBrackets[-1] = 0
}

func main() {
	FILE_NAME := "../Examples/mandelbrot.bf"
	code := getBFCode(FILE_NAME)
	tokens := tokenize(code)
	interpret(tokens)
}
