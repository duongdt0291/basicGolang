package main

import (
	"fmt"
	"math"
)

//method đơn giản là 1 func đặc biệt có receiver type được khai báo ở giữa
//từ khóa func và tên function
//receiver có thể là struct hoặc nonstruct type
//receiver có thể truy cập trong function
//cú pháp:
//func (t Type) methodName(parameter list) {
//}

type Person struct {
	firstName, lastName string
	age                 int
	job                 string
}

//func displayFullName có receiver type là Person
func (p Person) displayFullName() {
	fmt.Println(p.firstName, p.lastName)
}

//Tại sao lại cần method khi đã có function ?
//Vì không thể đặt tên function giống nhau, trong khi có nhiều thực thể lại có các hành động giống nhau
//còn các method có tên giống nhau, nhưng có receiver type khác nhau,thì đều được chấp nhận
type Square struct {
	size float64
}

type Circle struct {
	radius float64
}

func (s Square) Area() float64 {
	return s.size * s.size
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

//Pointer receiver vs Value receiver
//ở những ví dụ trên, chta mới chỉ sử dụng value receiver, vậy với pointer thì sẽ thế nào?
//
//method có pointer receiver
func (p *Person) changeAge(newAge int) {
	p.age = newAge
	//chính xác thì phải viết dòng trên là (*p).age = newAge
	//tuy nhiên Go sẽ tự động hiểu nên cách viết p.age giúp code dễ đọc hơn
}

//method dùng value receiver
func (p Person) changeJob(newJob string) {
	p.job = newJob
}

//Method của anonymous (embedded) field
//Ở ví dụ dưới, có Address là 1 anonymous field của struct Worker
type Worker struct {
	Name string
	Age  int
	Address
}

type Address struct {
	City, Country string
}

func (a Address) showAddress() {
	fmt.Printf("City %s, country %s\n", a.City, a.Country)
}

func main() {
	p1 := Person{"Duy", "Ram", 23, "developer"}
	//Dùng method: v.Method()
	p1.displayFullName() //Duy Ram
	//Ngoài ra, Go còn cung cấp 1 kiểu gọi khác cho method: (T).Method(v)
	Person.displayFullName(p1) //Duy Ram

	//Method vs Function
	s1 := Square{10}
	fmt.Printf("Area of s1: %f\n", s1.Area())
	c1 := Circle{5}
	fmt.Printf("Area of c1: %f\n", c1.Area())
	//Còn với function chúng ta sẽ phải đặt các hàm như AreaSquare, AreaCircle,
	//trong khi vẫn chỉ mục đích là tính Area

	//Value receiver vs pointer receiver
	fmt.Println("Job p1 before change with method use value pointer:", p1.job)
	p1.changeJob("director")
	fmt.Println("Job p1 after change with method use value pointer:", p1.job)

	fmt.Println("Age p1 before change with method use receiver pointer:", p1.age)
	p1.changeAge(35)
	//tương tự như trên đã nói, bản chất câu lệnh trên là: (&p1).changeAge(35)
	fmt.Println("Age p1 after change with method use value pointer:", p1.age)
	//Từ 2 ví dụ trên, thấy rằng, khi nào cần thay đổi gì đến giá trị các trường
	//trong struct, thì chúng ta sẽ dùng pointer receiver, còn lại có thể dùng value receiver
	//ngoài ra khi dữ liệu quá lớn, quá phức tạp, việc dùng pointer sẽ giúp tiết kiệm hơn,
	//vì Go chỉ cần copy con trỏ(địa chỉ) mà ko phải thực hiện copy toàn bộ dữ liệu

	//Method của anonymous field
	//tương tự như promoted fields
	w1 := Worker{
		Name: "Duy Ram",
		Age:  25,
		Address: Address{
			City:    "Ha Noi",
			Country: "Viet Nam",
		},
	}
	w1.showAddress()

	//Value receiver method vs value argument function
	//Với 1 function có value arg, thì nó chỉ nhận giá trị truyền vào là value
	//Với 1 method value receiver, nó nhận cả value và pointer receiver
	//Tương tự, 1 function có pointer arg, thì nó chỉ nhận giá trị truyền vào là pointer
	//còn method thì nhận cả 2

	//Lưu ý: Để định nghĩa 1 method, thì recevier và method cần được định nghĩa hết ở trong 1 package
	//Từ qui tắc trên, chta thấy rằng ko thể  tạo 1 method cho những type có sẵn của Go như int, string,...
	//Cách giải quyết ở đây là tạo 1 type là alias của những type đó
	//Ví dụ: type myInt int - tạo type myInt là alias của type int, cách dùng tương tự.
	//User defined type

}
