package main

import "fmt"



func main(){

	// pointers : arr[1] --> 3
	// length --> numbers of elements in the slice
	// capacity --> number of 

	  arr:=[5] int {1,2,3,4,5}
	  sl:= arr[1:3]
	  sl = sl[:4]
	  fmt.Println(sl, arr , len(sl), cap(sl))
}