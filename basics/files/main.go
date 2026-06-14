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

	// f, err := os.Open("basics/files/example.txt")

	// if err!= nil {
	// 	panic(err)
	
	// }

	// defer f.Close()

	// buf:=make([] byte,30)


	// d, err:= f.Read(buf)

	// if err!=nil{
	// 	panic(err)
	// }


	// for i:=0;i<len(buf);i++{
	// 	println(string(buf[i]),d)
	// }


	// f, err := os.ReadFile("basics/files/example.txt") // not to use Readfile generally coz it loads all the data once in the memory 
	// if err!=nil{
	// 	panic(err)
	// }

	// fmt.Println(string(f))


	// // read folder

	// dir, err := os.Open("../")

	// if err!=nil{
	// 	panic(err)
	// }

	// defer dir.Close()

	// fileInfo ,err:=dir.ReadDir(2)

	// for _, fi :=range fileInfo {
	// 	fmt.Println(fi.Name())
	// }


	// create a file

	f,err:= os.Create("example2.txt")

	if err!=nil{
		panic(err)
	}

	defer f.Close()

	// f.WriteString("hi go bro")
	// f.WriteString(" Hello")

	bytes:= []byte("Go lang")

	f.Write(bytes)
	


	// read and write file in streaming fashnion
	// using reader and writer from buffio


	// delete a file 

	os.Remove("example2.txt")
}