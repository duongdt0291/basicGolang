package main

import "fmt"

//Cú pháp:
//func func_name(param1 type_param1, param2 type_param2, ...) return_type {
// code block
//}
func greet(firstName string, lastName string) string {
	return firstName + " " + lastName
}

//Với các param cùng kiểu có thể viết ngắn gọn param1, param2 type
func greet1(firstName, lastName string) string {
	//fmt.Sprint trả về string được tạo từ các argument (các loại type, ko nhất thiết là string),
	//tự thêm khoảng trắng giữa các arg đó
	return fmt.Sprint(firstName, lastName)
}

//parameter và return type là không bắt buộc
func greet2() {
	fmt.Println("Welcome to function's world")
}

//named return value
func greet3(firstName string, lastName string) (fullName string) {
	fullName = fmt.Sprint(firstName, lastName)
	return //Go tự biết lấy giá trị của biến nào trả về
}

//1 điểm đặc biệt của Go so với các ngôn ngữ khác là có thể trả về 1 lúc nhiều giá trị
func rectProps(length, width float64) (float64, float64) {
	area := length * width
	perimeter := (length + width) * 2
	return area, perimeter
}

//1 func được gọi là variadic func khi param cuối cùng của func đó có kiểu là ...T
//lúc này, ngoài những param đã khai báo trước,
//những params còn lại có kiểu T sẽ được cho vào 1 slice có kiểu là []T
func tinhTong(nums ...int) int {
	sum := 0
	for _, value := range nums {
		sum += value
	}
	return sum
}

func tinhTongHaiSo(num1, num2 int) {
	fmt.Println(num1 + num2)
}

//Closure: tính đóng gói, xảy ra khi return 1 function
//func đó sẽ giữ tham chiếu đến các biến ở bên ngoài mà nó sử dụng
func testCl() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

//callback: 1 func đk gọi là cb khi nó được gán làm 1 arg cho 1 func khác
func testCb(sl []int, cb func(int)) {
	for _, v := range sl {
		cb(v)
	}
}

//example for defer
func hello() {
	fmt.Printf("hello ")
}

func world() {
	fmt.Printf("world")
}

func main() {
	greet2()                         // Welcome to function's world
	fmt.Println(greet("Duy", "Ram")) //Duy Ram

	//Lưu ý: không như 1 số ngôn ngữ khác,
	//trong Go không thể gán số lượng arg vượt quá hoặc ít hơn số  lượng params
	//fmt.Println(greet("Duy", "Ram", "ml")) //error
	//fmt.Println(greet("Duy")) //error

	fmt.Println(greet1("Long", "Yến")) //Long Yến
	fmt.Println(greet3("Su", "Do"))    //Su Do

	area, perimeter := rectProps(10, 5)
	fmt.Println("Area of rect is:", area)           //50
	fmt.Println("Perimeter of rect is:", perimeter) //30
	fmt.Println(tinhTong(1, 2, 3, 4, 5))            //15

	//variadic arguments
	//chỉ dùng cho variadic func
	sl := []int{1, 2, 3, 4, 5}
	fmt.Println(tinhTong(sl...))

	//Các cách khai báo func bên trên đk gọi là func declaration
	//Ko thể khai báo 1 func bằng cách trên trong 1 func khác
	//Go cung cấp cách khai báo khác trong trường hợp này là func expression
	//Cú pháp tên_biến := func (params type) return_type {}
	fc := func() {
		fmt.Println("I'm from a function expression")
	}
	fc()

	//callback
	testCb([]int{1, 2, 3, 4}, func(num int) {
		fmt.Println(num) //màn hình sẽ in ra từng dòng từ 1 đến 4
	})

	//closure
	increment1 := testCl()
	fmt.Println("increment1:", increment1()) //1
	fmt.Println("increment1:", increment1()) //2
	increment2 := testCl()
	fmt.Println("increment2:", increment2()) // 2

	//Lưu ý: Mọi thứ trong Go đều pass by value
	//nếu gán 1 biến x làm arg của 1 func
	//thì Go sẽ lấy giá tri của biến x để gán cho 1 biến mới y
	//nếu x ko phải là con trỏ hoặc reference type như slice, map, channel
	//thì việc thay đổi giá trị biến y không làm thay đổi đến x
	//còn nếu là các loại dữ liệu trên, thì việc thay đổi y sẽ dẫn đến thay đổi x

	//Anonymous self executing func
	func() {
		fmt.Println("Im self executing func")
	}()

	//defer: function sau keyword defer sẽ chạy trước khi func bao quanh nó kết thúc
	//có thể hiểu là trước khi return
	//đoạn code dưới cho ra kết quả là hello world
	defer world()
	hello()
}
