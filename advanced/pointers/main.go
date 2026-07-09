package main

import "log"



func main(){
	truck:=42
	anotherTruck:=&truck

	log.Print(&truck)
	log.Print(anotherTruck)


	truck = 0

	log.Print(*anotherTruck)
}



