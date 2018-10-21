package main

import "fmt"

func main() {
	//array là tập hợp các phần tử có cùng kiểu dữ liệu
	//Array có kiểu dữ liệu là n[T], với n là số lượng phần tử trong array
	//T là kiểu dữ liệu các phần tử trong array
	//do n là phần trong kiểu dữ liệu của array nên chiều dài của array là ko thay đổi được
	//do đặc điểm này mà array rất ít được sử dụng
	//chủ yếu giữ nhiệm vụ là underlaying
	var arr [5]int
	fmt.Printf("%T", arr) //5[int]
	arr[0] = 1
	arr[1] = 2

	//Các cách khai báo khác
	var arr2 [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr2)

	//Khai báo dạng shorthand
	arr3 := [5]int{1, 2, 3, 4, 5}
	fmt.Println("arr3:", arr3)

	//Kiểu khai báo để Go tự nhận độ dài của mảng
	arr4 := [...]int{1, 2, 3, 4, 5}
	fmt.Println("arr4:", arr4)
	//để lấy chiều dài mảng dùng func len(arr)
	fmt.Println("Length of arr4:", len(arr4))

	//Array là 1 value type, tức là khi gán nó cho 1 biến mới
	//việc thay đổi trên biến mới không làm ảnh hưởng đến biến cũ
	arr5 := arr4
	arr5[0] = 100
	fmt.Println("arr5:", arr5)
	fmt.Println("arr4:", arr4)

	//Duyệt qua mảng bằng for
	for i := 0; i < len(arr5); i++ {
		fmt.Println(arr5[i])
	}

	//dùng for range
	for _, value := range arr5 {
		fmt.Println(value)
	}
}
