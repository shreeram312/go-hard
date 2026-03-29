package main

import "fmt"

func main(){
	// var arr [3] int
	// fmt.Println(arr)

	// implicit assignment of array
	// arr:=[3] int{1,2,4}
	// fmt.Println(arr)

	// explicit assignment of array
	// arr2:=[3] int{1:2,2:3}
	// fmt.Println(arr2)


	array:= [3][3] int {{1,2,3},{4,5,6},{7,8,9}}
	array[0] = [3] int {4,5,6}
	fmt.Println(array)
	fmt.Println(len(array))


	// for idx:=0 ; idx<len(array); idx++ {
	// 	for jdx:=0; jdx<len(array[idx]); jdx++ {
	// 		fmt.Println(array[idx][jdx])
	// 	}
	// }


	for _, val := range array{
		for _, val2 := range val{
			fmt.Println(val2,val)
		}
	}
}