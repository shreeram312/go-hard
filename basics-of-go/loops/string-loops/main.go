package main

import "fmt"

func main(){
// 	s := "Aह"

// // byte view
// for i := 0; i < len(s); i++ {
//     fmt.Println(s[i])
// }


str:="hello world 7🙂"

// for _, val := range str{
// 	fmt.Println(val)
// }

for idx:=0; idx<len(str); idx++ {
	fmt.Printf("%c",str[idx])
}
}
