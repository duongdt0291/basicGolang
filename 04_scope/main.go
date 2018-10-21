package main

import (
	"fmt"
)

//Biến được khai báo ở package scope, có thể truy cập được từ mọi vị trí trong package
//Kể cả trường hợp biến được khai báo dưới func sử dụng nó.
//Có thể hiểu như cơ chế hoisting với biến var trong js.
var a int = 1

func main() {
	//Trong Go, scope nhỏ nhất cho 1 biến là block scope.
	//Biến được khai báo trong 1 block {}, chỉ có thể truy cập trong block đó
	b := "Im only accessible in main function"

	fmt.Println("In func main, a =", a)
	fmt.Println("b =", b)

	test()
}

func test() {
	fmt.Println("In func test, a =", a)

	// Biến b được khai báo trong func main, do vậy func test không truy cập được
	// fmt.Println("In func test, b =", b) //Error, undefined b
}

//Ngoài ra trong go còn khái niệm file scope
//Là khi import 1 package khác, thì chỉ có file nào được import package đó sử dụng được, 
//còn các file khác dù trong cùng package nếu không import thì không sử dụng được.

//Universe scope
//Những biến mà có thể dùng ở mọi vị trí, được go qui định