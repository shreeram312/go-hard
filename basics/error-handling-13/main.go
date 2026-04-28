package main

import "fmt"


func divide(a int, b int) int {
	
	return a / b
}


func deferredFunc() {
	fmt.Println("deferred function")
	r:= recover()
	fmt.Println(r)
}


// if we add recover() in the deferred function then the program will not stop and will print the message

func main(){
	// fmt.Println(divide(2,0))
// panic is used to stop the program and print the message

defer deferredFunc()

panic("hello world")
}