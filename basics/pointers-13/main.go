package main

import "fmt"


func main(){
	x:=0
	y:= &x

	*y=100
	fmt.Println(x,*y)
}