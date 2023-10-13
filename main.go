package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/busylambda/raven/parser"
)

func main() {
	// file, err := os.ReadFile("main.rvn")
	// if err != nil {
	// 	fmt.Println("Error reading file: ", err)
	// 	panic(err)
	// }

	// scanner := parser.NewScanner(string(file))
	// for {
	// 	tokenKind, literal := scanner.Scan()
	// 	if tokenKind == parser.EOF {
	// 		break
	// 	}

	// 	token := parser.NewToken(tokenKind, literal)

	// 	fmt.Println(token.String())
	// }
	// Write a repl for the lexer
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("[Welcome to the RAVN 0.1 LEXER/PARSER REPL]\n\nType .help for help duh...\n\n")
	repl()
	for scanner.Scan() {
		command := cleanInput(scanner.Text())

		switch command {
		case ".clear":
			clearScreen()
		case ".exit":
			os.Exit(0)
		case ".help":
			fmt.Println("# Commands\n - .clear\n - .exit\n - .help")
		default:
			parser := parser.NewParser(command)

			parser.NextNode()
		}

		repl()
	}
	fmt.Println()
}

// cleanInput preprocesses input to the db repl
func cleanInput(text string) string {
	output := strings.TrimSpace(text)
	output = strings.ToLower(output)
	return output
}

func repl() {
	fmt.Print("RAVN 0.1 LEXER REPL >>> ")
}

func unknownCommand(text string) {
	fmt.Println("Unknown command -> ", text)
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}
