package main

import "fmt"



func main(){

	// pointers : arr[1] --> 3
	// length --> numbers of elements in the slice
	// capacity --> number of 

	//   arr:=[5] int {1,2,3,4,5}
	//   sl:= arr[1:3]
	//   sl = sl[:4]
	//   fmt.Println(sl, arr , len(sl), cap(sl))

	// sl:= make([]int, 10, 10)


	sl:=[]string {"hello", "world", "how", "are", "you"}


	// for _, val:= range sl{
	// 	fmt.Println( val)
	// }

     test(sl)
	 fmt.Println(sl)

	 // in this case the slice is passed in fucntion so the slice can be modified which was not possible in arrays
	
}


func test(sl []string){
	sl[0]= "test"
}