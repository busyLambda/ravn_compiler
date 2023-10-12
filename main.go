package main

import (
	"fmt"
	"os"

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

	for {
		repl()
		var command string
		fmt.Scanln(&command)

		switch command {
		case "clear":
			clearScreen()
		case "exit":
			os.Exit(0)
		case "help":
			fmt.Println("Commands: clear, exit, help")
		default:
			scanner := parser.NewScanner(command)

			for {
				tokenKind, literal := scanner.Scan()
				if tokenKind == parser.EOF {
					break
				}

				token := parser.NewToken(tokenKind, literal)

				fmt.Println(token.String())
			}
		}
	}
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
