package main

import (
	"fmt"
	"time"
	"sync"
)


type Counter struct {
	value int
	lock sync.Mutex
}


func count(counter *Counter){
	counter.lock.Lock()
	defer counter.lock.Unlock()
	counter.value++
	fmt.Println(counter.value)
}


func main(){
	counter := Counter{}


	for i:=0; i<100; i++{
		go count(&counter)
	}

	time.Sleep(2 * time.Second)

}