package main

import "fmt"


// func main(){
// 	x:=0
// 	y:= &x

// 	*y=100
// 	fmt.Println(x,*y)
// }


// func change(num *int){
// 	*num = 100;
// }

// func main(){
// 	x:=10

// 	change(&x)
// 	fmt.Println(x)
// }


// type Book struct {
// 	id int
// 	title string
// }


// func (b *Book) setTitle(title string){
// 	b.title  = title

// }


// func main(){
// 	b:= Book{10, "Old"}
// 	b.setTitle("New")
// 	fmt.Println(b)
// }



func main() {
	a := 1
	b := 2

	values := []*int{&a, &b}

	fmt.Println("values =", values)

	for i, value := range values {
		fmt.Println("index =", i)
		fmt.Println("value =", value)   // address
		fmt.Println("*value =", *value) // actual number
		fmt.Println("-----")
	}
}