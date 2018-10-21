package main

import "fmt"

func main() {
	//Mỗi biến đang sử dụng đều được lưu tại 1 vị trí trong bộ nhớ (memory address)
	//Toán tử & trả về địa chỉ của biến trong bộ nhớ
	a := 10
	fmt.Println(&a)

	//Gán cho b giá trị địa chỉ của biến a
	//Lúc này ta gọi b là con trỏ của a
	//Con trỏ có type là *T
	//Giá trị zero value của con trỏ là nil
	//Reference
	//Lưu ý: Go không cho phép bất cứ phép toán số học nào thực hiện trên con trỏ
	var b *int
	b = &a
	fmt.Println("Address of variables a is:", b)
	// b ++ //error

	//Dereference a pointer - tức là việc lấy giá trị của biến ở ô địa chỉ được lưu trong con trỏ
	fmt.Println("Value of a:", *b) // 10
	//lấy giá trị của ô địa chỉ mà con trỏ b trỏ tới và gán bằng 11
	*b = 11
	//do vậy, giá trị của biến a cũng thay đổi theo
	fmt.Println("Value of a:", a) // 11

	//Ứng dụng lớn nhất của pointer là việc truyền nó vào trong function
	//Trong Go, khi truyền 1 giá trị vào trong func,
	//go sẽ thực hiện ngầm copy giá trị đó và gán vào biến, việc thực hiện thay đổi trong func không gây ảnh hưởng đến biến bên ngoài
	//Tuy nhiên khi gán 1 con trỏ vào func, bản chất là go vẫn copy giá trị
	//tuy nhiên do là con trỏ trỏ tới bộ nhớ, nên khi thay đổi giá trị của biến lưu tại ô địa chỉ mà con trỏ trỏ tới, biến đó sẽ bị thay đổi giá trị.
	increment1(a)
	fmt.Println("Value of a after execute func increment1:", a)

	increment2(&a)
	fmt.Println("Value of a after execute func increment2:", a)
}

func increment1(x int) {
	x++
}

func increment2(x *int) {
	*x++
}
