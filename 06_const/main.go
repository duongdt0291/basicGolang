package main

import "fmt"

func main() {
	//Const là 1 cách khai báo biến như var
	//Biến được khai báo bằng từ khóa const không thay đổi được giá trị dù giá trị có cùng kiểu
	const a = 10
	// a = 11 // error
	fmt.Printf("%T\n", a)

	b := a + 1.0
	fmt.Printf("%T\n", b)

	//The value of a constant should be known at compile time. 
	//Hence it cannot be assigned to a value returned by a function call 
	//since the function call takes place at run time.
	// c := sqrt(4) //hợp lệ
	// const d = sqrt(4) //Lỗi

	//Cần tìm hiểu về untyped const
	const c = 10
	const d = c + 1.1
	//với var sẽ bị lỗi
	// var c = 10
	// var d = c + 1.1
}
