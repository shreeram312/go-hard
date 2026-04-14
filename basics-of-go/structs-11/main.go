package main

import "fmt"

// structs are just like classes in other languages
// structs is type defination of fields


type Person struct  {

	//setting first letter of the field name to capital to make it public  and can be accessed outside the package

	Name string
	Age uint
}


func (p Person ) GetName()  string{
	return p.Name
}

func (p Person ) SetName(name string) {
	p.Name = name
	fmt.Println(p.Name)
}


func main() {
	var p1  Person = Person{ Name: "John", Age: 20,}

	// p1 is  sent as a copy of the struct so the name is not changed

	// so it cant be modified the original struct
	// ideally can be done using pointers
	p1.SetName("Jane")
	fmt.Println(p1.GetName())
}