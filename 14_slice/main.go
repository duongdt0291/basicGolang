package main

import "fmt"

func main() {
	//Slice là 1 tập hợp (danh sách) các phần tử có cùng kiểu, ko cố định độ dài
	//Slice bản chất ko lưu trữ dữ liệu,
	//mà nó chỉ là tham chiếu (reference type) đến 1 mảng nền (underlying array)

	//Tạo slice từ mảng arr với các phần tử từ vị trí 1 đến 4 (ko bao gồm 4) của mảng arr
	arr := [5]int{1, 2, 3, 4, 5}
	sl1 := arr[1:4]
	fmt.Println(sl1) //[2, 3, 4]

	//Tạo slice từ slice
	sl2 := sl1[:]
	fmt.Println(sl2)

	//Do slice là reference type, nên khi thay đổi giá trị vị trí nào đó trong sl2,
	//sẽ dẫn đến việc thay đổi của sl1, arr
	sl2[0] = 6
	fmt.Println("sl2 after changed:", sl2)
	fmt.Println("sl1 after sl2 changed:", sl1)
	fmt.Println("arr after sl2 changed", arr)

	//1 cách khai báo slice khác
	//bản chất cách khai báo dưới là go sẽ tạo 1 underlying arr có chứa các phần tử khai báo
	//và gán cho sl3 là tham chiếu của mảng đó
	sl3 := []int{1, 2, 3, 4, 5}
	fmt.Println("sl3:", sl3)

	//len và cap
	//func len cho biết độ dài của slice
	//cap cho biết sức chứa của underlying array,
	//tính từ vị trí phần tử mà slice bắt đầu tham chiếu đến phần tử cuối cùng của mảng
	fmt.Println("len of sl1:", len(sl1)) //3
	fmt.Println("len of sl1:", cap(sl1)) //4

	//cách khai báo bằng func make([]T, len, cap)
	//func make tạo 1 slice có kiểu các phần tử là T, với số phần tử là len, và sức chứa của underlying arr là cap
	//cap ko nhất thiết phải có khi gọi func, khi đó go sẽ mặc định cap bằng len
	sl4 := make([]int, 5, 10)
	fmt.Println("sl4:", sl4) //[0 0 0 0 0] - do 0 là zero value của kiểu int
	sl4[0] = 1
	sl4[1] = 2
	fmt.Println("sl4:", sl4) //[1 2 0 0 0]

	//Zero value của slice là nil, len và cap lúc này là 0
	var sl5 []int
	if sl5 == nil {
		fmt.Println("sl5's value is nil")
	}

	//Thêm phần tử vào slice
	//func append(s []T, x ...T) []T

	sl5 = append(sl5, 1, 2, 3, 4, 5) //do sl5 có giá trị ban đầu là 5, lúc này sl5 có len là 5, cap là 6
	fmt.Println("sl5:", sl5)
	fmt.Println("len of sl5:", len(sl5))
	fmt.Println("cap of sl5:", cap(sl5))

	//Khi thêm 1 phần tử mới bằng func append, nếu vượt quá cap
	//go sẽ ngầm copy toàn bộ phần tử của underlying arr hiện tại vào 1 array mới
	//sl sẽ tham chiếu đến array mới và array mới sẽ có cap gấp đôi array cũ
	sl6 := make([]int, 3)
	fmt.Println("sl6:", sl6)
	fmt.Println("len of sl6", len(sl6))
	fmt.Println("cap of sl6", cap(sl6))
	sl6 = append(sl6, 4)
	fmt.Println("cap of sl6 after append func:", cap(sl6))

	fmt.Println("sl6:")
	for _, v := range sl6 {
		fmt.Println(v)
	}

	//pass slice vào function
	//go sẽ pass value của sl6 vào param sl của func changeSl
	//do sl6 là reference type, nên khi thay đổi sl sẽ dẫn đến thay đổi sl6
	changeSl(sl6)
	fmt.Println("sl6:", sl6)
}

func changeSl(sl []int) {
	sl[0] = 6
}
