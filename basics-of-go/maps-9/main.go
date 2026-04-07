package main

import "fmt"

func main() {

	// var mp map[string] int  = map[string]int{"a":1}   verbose map syntax

	// shorhand map syntax
	mp := map[string]int{"a": 3} // or curly {} at the end 
	fmt.Println(mp)


	// another way to create a map
	// mp2 := map[string] [] int {"a":{1,2,3}}
	// // add a new key value pair
	// mp2["b"] =[]int {3,4,5}
	// // delete a key value pair
	// delete(mp2, "b")


	// //check if a key exists

	// value , ok := mp2["b"]

	// fmt.Println(value, ok)


	mp3:= map[uint]uint{}
	n:=100

	for i:=uint(0); i<uint(n); i++ {
		for d:=uint(1); d<=5; d++ {
			if i % d==0 {
				mp3[d]++
			}
		}
	}

	fmt.Println(mp3)


}