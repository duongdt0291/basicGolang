package main

import "fmt"

func main() {
	//Cách sử dụng if else tương tự như các ngôn ngữ khác
	//Tuy nhiên trong Go, biểu thức điều kiện không cần trong ()
	//và đoạn code trong body bắt buộc phải trong {}, kể cả chỉ có 1 dòng lệnh
	num := 10
	if num < 9 {
		fmt.Println("Small number")
	} else {
		fmt.Println("Big number")
	}

	//Go còn cung cấp thêm cho if initialization statement ()
	//Cú pháp if init; cond {code body}
	if x := 11; x == 11 {
		fmt.Println("Thís ís eleven")
	}
	//Lưu ý, là biến khai báo x chỉ truy cập được trong if else body
	// fmt.Println(x) //error

	//Switch cũng có initialization statement tương tự
	switch food := "chocolate"; food {
	case "milk":
		fmt.Println("Milk")
	case "cheese":
		fmt.Println("Cheese")
	case "chocolate":
		fmt.Println("Chocolate")
	default:
		fmt.Println("Food not found")
	}
}
