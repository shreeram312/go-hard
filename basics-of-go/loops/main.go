package main

import "fmt"

func main() {
	//for loop
	// for idx:=0; idx <10; idx++ {
	// 	fmt.Println(idx)
	// }


	// while loop
	// there is no while loop in go but we can use for loop as while loop

	//while loop
	// a:=1
	// for a<=10{
	// 	fmt.Println(a)
	// 	a++
	// }

//string loops

const str string ="he🙂"

fmt.Println(len(str))

// ascii --> 1 byte 256 characters --> 0-255



// utf-8 --> 4 byte 1114112 characters --> 0-1114111

// genenrally all special characters like emois and hindi characters are represented in utf-8


// even if u access the index value it will give the ascii value of the character or utf bytes storage size for emojis


	for idx:=0; idx<len(str); idx++ {
	fmt.Println(string(str[idx]))
	
}
}