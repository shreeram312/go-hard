package main

import "fmt"

// import "fmt"

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

// func add(x int, y int, ch chan int) {
// 	fmt.Println("x + y is ", x+ y)
// 	ch <- x + y

// }

// func main(){
// 	ch :=  make(chan int)
// 	ch2 := make(chan int)

// 	go add(10,30,ch)
// 	go add(1332313130,12222, ch2)
// 	x:= <-ch
// 	y:=<-ch2

// 	fmt.Println(x,y)

// 	// by using sleect statement we can return whatver value first
// }

// go routines are deadlocked when the add function is not sending the value to the channel and main is waiting for the value to be sent to the channel.

// send and receive only channels
// send and receive only channels are used to send and receive values only to and from the channel
// send <- channel is used to send values to the channel
// receive <- channel is used to receive values from the channel
// chan<- send only channel
// <-chan receive only channel

// UnBuffered channels


func r(ch chan bool){
	<-ch
}


// func main(){
// 	ch := make(chan bool)


// 	go r(ch)
// 	go r(ch)

// 	ch <- true
// 	ch <-true
// 	fmt.Println("Done")
// }


// Buffered channels

func main(){
	ch:= make(chan bool,2)
	ch<- true
	ch<- true
	<-ch
	ch <- true
	fmt.Println("done")
}

// beacuse it has cpaacity of 2 only only when we remove from the queue e will be able to add another value in channel