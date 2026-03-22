package main

import "fmt"

func main(){
	x:=3
	// prints them adding spaces with new line /n by default
	fmt.Println("hello",x,2)


	y:=false

	// %T prints the types of the variable 
	fmt.Printf("%T %T",y,y)


	z:=345
	//represents %b the binary format of variable
	fmt.Printf("%b", z)
	fmt.Println("")


	a:=10.5435353453
	// %e returns the scientific value
	// %f returns the smallest floating number 
	// %.2f retuns upto 2 decimals 
	// double %% for percentage at the end
	fmt.Printf("%.5f %%\n",a )



	b:="hello"
	// %s is used for string formatting
	fmt.Printf("%s",b)

}