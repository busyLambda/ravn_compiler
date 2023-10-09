package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/busylambda/raven/parser"
)

func main() {
  file, err := os.ReadFile("main.rvn")
  if err != nil {
    fmt.Println("Error reading file: ", err)
    panic(err)
  }

  scanner := parser.NewScanner(string(file))
  for {
    tokenKind, literal := scanner.Scan()
    if tokenKind == parser.EOF {
      break
    }
    
    token := parser.NewToken(tokenKind, literal)

    fmt.Println(token.String())
  }

  repl()

  fmt.Println("Pull the trigger!")
}

func repl() {
  fmt.Print("RAVN 0.1 LEXER REPL >>> ")
}

func unknownCommand(text string) {
  fmt.Println("Unknown command -> ", text)
}

func clearScreen() {
    cmd := exec.Command("clear")
    cmd.Stdout = os.Stdout
    cmd.Run()
}

