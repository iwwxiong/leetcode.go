package main

import "fmt"

func main() {
	// fmt.Println(11 / 2)
	// fmt.Println(11 % 2)
	// s := "1234567890"
	// fmt.Println(s[5:10])
	// fmt.Println(int32(-654) % 10)
	for _, ch := range "hello world" {
		fmt.Printf("%v\n", int(ch))
	}
	// var c int32
	// var d int = -91283472332
	// // var a int = 2147483647
	// // var b int = 2147483647
	// c = int32(d)
	// fmt.Printf("%v\n", c)
}
