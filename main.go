package main

import (
	"fmt"
	"interpreter/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Welcome %s, this is the monkey programming language.\n", user.Username)
	fmt.Printf("Feel free to type commands.\n")
	repl.Start(os.Stdin, os.Stdout)
}
