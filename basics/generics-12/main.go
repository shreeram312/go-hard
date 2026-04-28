package main

import (

	"fmt"
)





func add[T int | float64](x T, y T) T{
	return x + y
}



func getValues[K comparable, V any](mp map[K]V) []V {
	values:= []V{}

	for  _, value := range mp{
		values = append(values, value)
	}
	return values
}

func main() {
	value:= add(1, 2)
	value2:= add(1.0, 2.0)
	value3:= add(1.0, 2.2)
	value4:=  add(7,2.3)
	fmt.Println("RUNNING THIS FILE NOW")
	fmt.Println(value,value2,value3,value4)




	mp:=getValues(map[string]int{"a":1, "b":2, "c":3})
	mp2:=getValues(map[int]float64{1:1.1, 2:2.2})
	fmt.Println(mp)
	fmt.Println(mp2)
}