package main


import "fmt"


func main() {
	var name string
	fmt.Print("Enter your name:")

	fmt.Print("Enter your age:")
	var age int
	fmt.Scanln(&name, &age)
	fmt.Println("Hello", name, "you are", age, "years old")
}