package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/parikshith078/interpreter/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! this is the monkey \n", user.Username)
	fmt.Printf("Fell free to type in commands")
	repl.Start(os.Stdin, os.Stdout)
}
