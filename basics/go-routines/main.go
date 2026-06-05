package main

import (
	"fmt"
	"time"
)



func run1(){
	time.Sleep(2 * time.Second)
	fmt.Println("run 1")
}

func run2(){
	time.Sleep(2 * time.Second)
	fmt.Println("run 2")
}


func main(){
	go run1()
	go run2()
	time.Sleep(3 * time.Second)
	
	fmt.Println("done")
}