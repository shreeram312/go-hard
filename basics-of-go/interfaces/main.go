package main

import "fmt"


type Shape interface {
	getPerimeter() uint


}


type Triangle struct {
	a uint
	b uint
	c uint
}


type Square struct {
	width uint

}


func (s Square ) getPerimeter() uint {
	return 4 * s.width
}

func(t Triangle) getPerimeter() uint {
	return t.a + t.b + t.c
}




func main(){
	var s Shape = Triangle{3, 4, 5}
	fmt.Println(s.getPerimeter())

	var s2 Shape = Square{width: 10}
	fmt.Println(s2.getPerimeter())


	var sl [] Shape = [] Shape{Triangle{1,3,4} , Square{width: 10}}

	perimeters:= uint(0)

	for _, shape:= range sl{
		perimeters = perimeters + shape.getPerimeter()
	}
}