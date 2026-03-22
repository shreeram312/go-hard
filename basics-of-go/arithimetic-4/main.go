package main

import "fmt"


// byy default int is int 64
// by default unint is unint 64



func main(){
	x:=30
	y:=1000

	z:=float64(y)/float64(x)

	// always convert the type of variables to smaller type
	// among the above code x has uint8 and y has int so we convert x to int


	// if we convert into larger type it will have random values (overflow issue occurs)


	a:="hell"
	b:=65433
	c:=a+fmt.Sprint(b)  // just beacuse it was converting int o ascii  we used fmt.sprint to concatenate	

	// when we convert int to string it will convert into ascii value

	fmt.Println(c)

	fmt.Println(z)
	
}