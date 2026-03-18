package main

import "fmt"

//uint --> unsigned integer --> means just positive integer
// int --> signed integer --> means just positive and negative integer
//float --> means just floating point number
//bool --> means just true or false
// string --> means just text
//rune --> means just character (rune(8 bits) == int32)
// byte --> means just byte of memory  (byte(8 bits) == uint8)

// unint8 --> means just 8 bits of memory
// uint16 --> means just 16 bits of memory

func main(){
	var x string ="hello world bhai";
	fmt.Println(x)

	const y uint8 = 244
	fmt.Println(y)

	z:= 100000000
	fmt.Println(z)
	
}