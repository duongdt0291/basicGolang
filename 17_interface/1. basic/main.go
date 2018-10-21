package main

import (
	"fmt"
	"math"
)

//Trong Go, 1 interface A là tập hợp các method
//Khi 1 type B có tất cả các method trong 1 interface A
//thì ta coi type B implement interface A
//Lưu ý: interface chỉ xác định những method nào 1 type nên có
//còn chính type đó sẽ tự quyết định việc thực hiện method thế nào.

//Ở ví dụ dưới, chta sẽ có 2 struct Square và Circle đều có method Area, Perimeter
//có 1 interface Shape chứa 2 func Area, Perimeter
//Lúc này Go sẽ tự động coi interface Shape có giá trị nên là Square và Circle
//Mọi function có params là kiểu Shape đều có thể gán Square và Circle vào làm arguments

type Square struct {
	size float64
}

func (sqr Square) Area() float64 {
	return sqr.size * sqr.size
}

func (sqr Square) Perimeter() float64 {
	return sqr.size * 4
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

type Shape interface {
	Area() float64
	Perimeter() float64
}

func getInfo(s Shape) (area float64, perimeter float64) {
	area = s.Area()
	perimeter = s.Perimeter()
	return
}

func calcAllArea(shapes ...Shape) float64 {
	sum := 0.0
	for _, s := range shapes {
		sum += s.Area()
	}
	return sum
}

//empty interface
//vì Empty không có method nào, nên mọi type đều implement Empty interface
type empty interface{}

func displaySelf(e empty) {
	fmt.Println(e)
}

func main() {
	sqr := Square{10}
	area, perimeter := getInfo(sqr)
	fmt.Printf("Square info: area = %f, perimeter = %f\n", area, perimeter)

	cir := Circle{5}
	area, perimeter = getInfo(cir)
	fmt.Printf("Circle info: area = %f, perimeter = %f\n", area, perimeter)

	//Ở ví dụ trên, tính ứng dụng của interface có thể hơi mập mờ, chưa rõ ràng
	//giờ chúng ta sẽ thử nghĩ đến bài toán yêu cầu tính tổng diện tích các hình
	//đang có (vuông, tròn , chữ nhật ...)
	//vì chúng khác struct, nên nếu ko dùng interface thì chúng ta viết như sau
	sum := 0.0
	sum += sqr.Area() + cir.Area()
	fmt.Println("Sum area all shape above is", sum)
	//Còn nếu dùng interface
	fmt.Println("Sum area all shape above is", calcAllArea(sqr, cir))
	//nhìn chung, chúng ta cần sử dụng interface khi cần tương tác giữa nhiều loại type với nhau
	//còn nếu chỉ 1 type, tốt hơn là dùng method

	//Empty interface
	displaySelf(sqr)

	//Type assertion
	var n interface{} = 24 //ở đây hiểu rằng interface n có underlying value là 24, và concrete type là int
	fmt.Println("n:", n)
	assert(n)
	var s interface{} = "Hello interface"
	fmt.Println("s:", s)
	assert(s)

	//Type switch
	findType(24)
	findType("Hello interface")
	findType([]int{})

}

//Type assertion dùng để  lấy thông tin của giá trị nền interface (underlying value)
//Cú pháp: i.(T)
func assert(i interface{}) {
	v, ok := i.(int)
	fmt.Println(v, ok)
}

//type switch
//Cú pháp lấy concrete type của interface: i.(type)
//lưu ý là chỉ được dùng i.(type) trong câu lệnh switch như trong ví dụ dưới
func findType(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Println("String:", i.(string))
	case int:
		fmt.Println("Integer:", i.(int))
	default:
		fmt.Println("Unknown type")
	}
}
