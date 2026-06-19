package main

import (
	"bufio"
	"fmt"
	"os"
)

var hadError = false

func main() {
	args := os.Args[1:] // os.Args[0] is program's name

	if len(args) > 1 {
		fmt.Println("Usage golox [script]")
		os.Exit(64) // EX_USAGE exit code matching the book
	} else if len(args) == 1 {
		runFile(args[0])
	} else {
		runPrompt()
	}
}


func runFile(path string) {
	bytes, err := os.ReadFile(path)	// []byte is a slice
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(66)
	}
	run(string(bytes))

	// if err occurs during execution 
	if hadError {
		os.Exit(65) // dataErr
	}
}

func runPrompt() {
	reader := bufio.NewReader(os.Stdin) // Buffering: It reads data in larger chunks into a buffer, which minimizes the number of read operations.
	
	for {
		fmt.Println(">")
		line, err := reader.ReadString('\n') // read up to a new line 

		if err != nil {
			break // handles EOF (control-D)
		}
		 
		run(line)

		// Reset the error flag in interactive mode so one typo doesn't kill the REPL
		hadError = false
	}
}

// the core interpreter engine
func run(source string) {
	scanner := newScanner(source)

	tokens := scanner.ScanTokens()

	for _, token := range tokens {
		fmt.Println(token)
	}
}

func reportError(line int, message string) {
	report(line, "", message)
}

func report(line int, where string, message string) {
	fmt.Fprintf(os.Stderr, "[line %d] Error%s: %s\n", line, where, message)
	hadError = true
}

// stubs for now until chapter 4 revamps
type Token string // type aliasing

type Scanner struct {
	source string // why does scanner contain source here?
}

func newScanner(source string) *Scanner {
	return &Scanner{source: source}
}

func (s *Scanner) ScanTokens() []Token {
	if len(s.source) > 1 {
		return []Token{"DummyToken"}
	}
	return []Token{} // return a token slice
}
