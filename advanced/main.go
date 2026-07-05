package main

import (
	"errors"
	"fmt"
	"log"
)

type Truck struct{
	id string
}



func processTruck(truck *Truck)error {
	truck.id ="23"
	fmt.Println("Processing truck",truck)


	return errors.New("Some error")
	
}


func main(){
	truck:=[]Truck{
		{id: "10"},
		{id:"20"},
		{id:"30"},
	}



	for i := range truck{
		
		// err:=processTruck(&truck[i])

		// if err!=nil{
		// 	log.Fatalf("Error processing %s",err)
		// }


		// another way
		if err:=processTruck(&truck[i]); err!=nil{
			log.Fatalf("Error processing %s",err)
		}
		fmt.Println("ok")
	}
}