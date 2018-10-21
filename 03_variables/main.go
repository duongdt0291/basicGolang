package main

import "fmt"

func main() {
	//Co 2 dang khai bao bien co ban trong Go
	//1. var , cu phap: var name type = value
	//Khi dung var, neu khong khoi tao ngay gia tri ban dau
	//go se tu dong khoi tao ngam gia tri cho bien, tuy theo kieu cua bien do
	//Gia tri nay duoc goi la zero value
	//Khi khai bao co gia tri ban dau luon, thi co the khong can khai bao kieu, Go se tu dong
	//nhan kieu gia tri tu gia tri ban dau
	var a int = 10
	fmt.Println("a =", a)

	var b string
	fmt.Println("b =", b)

	var d = false
	fmt.Printf("Type of d is %T\n", d)

	//2. short hand, cu phap: name := value
	//Dang nay bat buoc phai co gia tri ban dau, go se tu dua vao gia tri ban dau
	//de xac dinh type cua bien
	//Luu y: Chi co khai bao kieu var moi duoc dung ngoai function
	c := 10
	fmt.Println("c =", c)

	//Khai bao nhieu bien 1 luc
	// var e, f ,g int = 1, 2, 3;
	// var e, f, g = 1, true, "Hello";

	//Lưu ý: Khi đã khai báo 1 biến cho 1 kiểu dữ liệu, không thể gán biến đó cho 1 giá trị kiểu khác
	//var e = 10 //biến có kiểu là int
	//e = "Hello" //Error, biến có kiểu int, ko thể dùng giá trị string

}
