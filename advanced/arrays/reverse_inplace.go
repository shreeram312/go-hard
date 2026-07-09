package main

import "fmt"

// ReverseInPlace solves the problem in O(n) time and O(1) space.

// {1,2,3,4,5} 0,3 --> 4,3,2,1,5
func ReverseInPlace(list []int, start, end int) {
	for start < end {
		temp := list[start]
		list[start] = list[end]
		list[end] = temp
		start++
		end--
	}
}

func main() {
	list := []int{1, 2, 3, 4, 5}
	ReverseInPlace(list, 0, 3)
	fmt.Println(list)
}