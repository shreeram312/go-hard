package main

import "fmt"

func add(num1 int , num2 int) (int, string){
	return num1+ num2 , "hello world"
}



func callFunc( callable func(int) int) int {
	return callable(10)
}



func doubleNumber(num int) int{
	return num * 2

}


func tripleNumber(num int) int{
	return num * 3
}



func firstFunc(str string)  func(string) string {
	return func( str2 string ) string	 {
		return str + " " + str2
	}
}


//variadic parameters

func sum(nums ...int)(s int) { 
	for _, value := range nums {
		s+=value
	}

	// return s is not needed because s is already declared in the function signature 
	// named return values
	
	return 
}



func main(){
	 value, str:=add(10,200)
	 fmt.Println(value , str)


// callback function similar to javascript
	 value2 := callFunc(tripleNumber)
	 fmt.Println(value2)



	 // annonymous function
	 value3:= callFunc(func(x int) int{
		return x * 4
	 })

	 fmt.Println(value3)



	//  f1:= firstFunc("hello")
	//  f2:= f1("world")

	//  fmt.Println(f2)


	a:= sum(1,2,3,4,5)
	fmt.Println(a)


	// it works it sends slice one by one to the  but works only if it accpets varaiadic parameters
	b:= sum([]int{1,2,3,4,5}...)
	fmt.Println(b)


}