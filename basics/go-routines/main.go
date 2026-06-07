package main

import "fmt"

// "fmt"
// "time"

// func run1(){
// 	time.Sleep(2 * time.Second)
// 	fmt.Println("run 1")
// }

// func run2(){
// 	time.Sleep(2 * time.Second)
// 	fmt.Println("run 2")
// }

// func main(){
// 	go run1()
// 	go run2()
// 	time.Sleep(3 * time.Second)

// 	fmt.Println("done")
// }


func add(x int, y int, ch chan int) {
	fmt.Println("x + y is ", x+ y)
	ch <- x + y


}


func main(){
	ch :=  make(chan int)
	go add(10,20,ch)
	go add(30,40,ch)
	go add(40,50,ch)
	go add(40,3,ch)

	x:= <-ch
	x= <-ch
	x= <-ch
	x= <-ch

	fmt.Println(x)
}


// go routines are deadlocked when the add function is not sending the value to the channel and main is waiting for the value to be sent to the channel.