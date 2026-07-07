package main

import (
	"fmt"
	"log"
)


type Truck interface {
	LoadCargo() error
	UnloadCargo() error
}

type NormalTruck struct{
	id string
	cargo int
}


type ElectricTruck struct{
	id string
	cargo int
	battery float64
}


func (e *ElectricTruck) LoadCargo() error{
	e.cargo+=1
	return nil
}


func (e *ElectricTruck) UnloadCargo() error{
	e.cargo=0
	return nil
}



func (t *NormalTruck) LoadCargo() error{
	t.cargo+=10
	return nil
}

func (t *NormalTruck) UnloadCargo() error{
	t.cargo=0
	return nil
}







// var (
// 	ErrNotImplemented  = errors.New("Not Implemented ")
// 	ErrTruckNotFound = errors.New("Truck not found")
// )

func processTruck(truck  Truck) error {
	fmt.Printf("Processing truck %+v \n",truck)	

	err:= truck.LoadCargo()
	if err!=nil{
		return fmt.Errorf("Error Loading")
	}
	return nil
}


func main(){
	truck:=&NormalTruck{
	id: "10",
	cargo: 20,
	}

	electricTruck:=&ElectricTruck{
		id: "12",
		cargo: 10,
		battery: 22.33,
	}


	processTruck(truck)

	processTruck(electricTruck)
	

	log.Println(truck.cargo)
	log.Println(electricTruck.cargo)





}