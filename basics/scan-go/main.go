package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Type something and press Enter:")

	input, _ := reader.ReadString('\n')

	fmt.Println("You typed:", input)
}