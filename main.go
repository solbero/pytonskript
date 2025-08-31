// main.go

package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/solbero/pyton/exec"
	"github.com/solbero/pyton/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	args := os.Args

	switch len(args) {
	case 1:
		fmt.Printf("Hello %s! This is the Monkey programming language!\n", user.Username)
		fmt.Printf("Feel free to type in commands\n")
		repl.Start(os.Stdin, os.Stdout)
	case 2:
		file, err := os.Open(args[1])
		if err != nil {
			panic(err)
		}
		defer file.Close()
		exec.Start(file, os.Stdout)
	default:
		fmt.Fprintf(os.Stderr, "%q: incorrect usage: Usage: `monkey [filePath]`\n", os.Args[0])
	}

}
