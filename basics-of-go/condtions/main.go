package main

import "fmt"



func main(){
	// x:= uint8(8)
	// y:=int(5)
	// z:= int(x)==y
	// fmt.Println(z)

	// we have && || 
	abc:=10
	if abc > 100 {
		fmt.Println("greater")
	}else
	{
		fmt.Println("lesser")
	}


	a:=10
	// switch a{
	// case 10:
	// 	fmt.Println("a")
	// case 20:
	// 	fmt.Println("twenty")
	// default:
	// 	fmt.Println("not a number")
	// }



	//different type of switch case naked switch case


	//in go there is break statement by default so we dont need to write break statement in switch case
	switch {
	case a > 6:
		fmt.Println("a is greater than 10")
		fallthrough // it will execute the next case even if the condition is not met
    case a > 8:
	fmt.Println("a is greater than 8")

	default:
		fmt.Println("a is not greater than 10 or 8")
	}
	

	
}