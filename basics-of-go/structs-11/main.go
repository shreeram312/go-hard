package main

import "fmt"

// structs are just like classes in other languages
// structs is type defination of fields


type Person struct  {
	name string
	age uint
}


func (p Person ) getName()  string{
	return p.name
}

func (p Person ) setName(name string) {
	p.name = name
	fmt.Println(p.name)
}


func main() {
	var p1  Person = Person{ name: "John", age: 20,}

	// p1 is  sent as a copy of the struct so the name is not changed

	// so it cant be modified the original struct
	// ideally can be done using pointers
	p1.setName("Jane")
	fmt.Println(p1.getName())
}