package main

import (
	"errors"
	"fmt"
	"log"
)

type Truck struct{
	id string

}


func (t *Truck) LoadCargo() error{
	return nil

}




var (
	ErrNotImplemented  = errors.New("Not Implemented ")
	ErrTruckNotFound = errors.New("Truck not found")
)

func processTruck(truck Truck)error {

	fmt.Println("Processing truck",truck)



	return ErrNotImplemented
	
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
		if err:=processTruck(truck[i]); err!=nil{
			// if errors.Is(err,ErrNotImplemented){
			// 	// we do this 
			// }

			// if errors.Is(err,ErrTruckNotFound){
			// 	// we do this 
			// }

			// or use switch
			log.Fatalf("Error processing %s",err)
		}
		fmt.Println("ok")
	}
}