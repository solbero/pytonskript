// main.go

package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/solbero/pytonskript/exec"
	"github.com/solbero/pytonskript/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	args := os.Args

	switch len(args) {
	case 1:
		fmt.Printf("Hei %s! Dette er programmeringsspråket Pyton!\n", user.Username)
		fmt.Printf("Her kan du skrive inn instruksjoner\n")
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
