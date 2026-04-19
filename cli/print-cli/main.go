package main

import (
	"fmt"
	"os"
)


func main() {
	args:= os.Args


	if len(args) < 2 {
		fmt.Println("Please provide a message to print")
		return
	}


	message:= args[1]


	switch message {
	case "hello":
		fmt.Println("hellow ow")

	case "bye":
		fmt.Println("bye bye")

	default:
		fmt.Println("Invalid message")
	}
}