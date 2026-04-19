package main

import "fmt"


func divide(a int, b int) int {
	
	return a / b
}


func deferredFunc() {
	fmt.Println("deferred function")
}

func main(){
	// fmt.Println(divide(2,0))
// panic is used to stop the program and print the message

defer deferredFunc()

panic("hello world")
}