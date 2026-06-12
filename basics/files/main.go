package main

import (
	
	"os"
)

func main(){
	// f, err := os.Open("basics/files/example.txt")
	// if err!=nil{
	// 	panic(err)
	// }

	// fileInfo,err:=f.Stat()
	// if err!=nil{
	// 	panic(err)
	// }
	// fmt.Println(fileInfo.Name())
	// fmt.Println(fileInfo.IsDir())
	// fmt.Println(fileInfo.ModTime())


	// read file

	f, err := os.Open("basics/files/example.txt")

	if err!= nil {
		panic(err)
	
	}

	defer f.Close()

	buf:=make([] byte,30)


	d, err:= f.Read(buf)

	if err!=nil{
		panic(err)
	}


	for i:=0;i<len(buf);i++{
		println(string(buf[i]),d)
	}





}